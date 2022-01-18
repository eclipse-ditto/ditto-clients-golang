// Copyright (c) 2020 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// http://www.eclipse.org/legal/epl-2.0
//
// SPDX-License-Identifier: EPL-2.0

package ditto

import (
	"errors"
	"github.com/eclipse/ditto-clients-golang/protocol"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"sync"
)

var (
	// ErrAcknowledgeTimeout is an error that acknowledgement is not received within the timeout.
	ErrAcknowledgeTimeout = errors.New("acknowledge timeout")
	// ErrSubscribeTimeout is an error that subscription confirmation is not received within the timeout.
	ErrSubscribeTimeout = errors.New("subscribe timeout")
	// ErrUnsubscribeTimeout is an error that unsubscription confirmation is not received within the timeout.
	ErrUnsubscribeTimeout = errors.New("unsubscribe timeout")
)

// Handler represents a callback handler that is called on each received message.
// If the underlying transport (e.g. Hono) provides a special requestID related to the Envelope,
// it's also provided to the handler so that chained responses to the ID can be later sent properly.
type Handler func(requestID string, message *protocol.Envelope)

// Client is the Ditto's library actual client's implementation.
// It provides the connect/disconnect capabilities along with the options to subscribe/unsubscribe
// for receiving all Ditto messages being exchanged using the underlying transport (MQTT/WS).
type Client struct {
	cfg                *Configuration
	pahoClient         MQTT.Client
	handlers           map[string]Handler
	handlersLock       sync.RWMutex
	externalMQTTClient bool
	wgConnectHandler   sync.WaitGroup
}

// NewClient creates a new Client instance with the provided Configuration.
func NewClient(cfg *Configuration) *Client {
	client := &Client{
		cfg:      cfg,
		handlers: map[string]Handler{},
	}
	return client
}

// NewClientMQTT creates a new Client instance with the Configuration, if such is provided, that is going
// to use the external MQTT client.
//
// It is expected that the provided MQTT client is already connected. So this Client must be controlled
// from outside and its Connect/Disconnect methods must be invoked accordingly.
//
// If a Configuration is provided it may include ConnectHandler and ConnectionLostHandler, as well as acknowledge,
// subscribe and unsubscribe timeout. As an external MQTT client is used, other fields are not needed and
// regarded as invalid ones.
//
// Returns an error if the provided MQTT client is not connected or the Configuration contains invalid fields.
func NewClientMQTT(mqttClient MQTT.Client, cfg *Configuration) (*Client, error) {
	if !mqttClient.IsConnected() {
		return nil, errors.New("MQTT client is not connected")
	}

	if err := validateConfiguration(cfg); err != nil {
		return nil, err
	}

	client := &Client{
		cfg:                cfg,
		pahoClient:         mqttClient,
		externalMQTTClient: true,
	}
	return client, nil
}

// Connect connects the client to the configured Ditto endpoint provided via the Client's Configuration at creation time.
// If any error occurs during the connection's initiation - it's returned here.
// An actual connection status is callbacked to the provided ConnectHandler
// as soon as the connection is established and all Client's internal preparations are performed.
// If the connection gets lost during runtime - the ConnectionLostHandler is notified to handle the case.
//
// When the client is created using an external MQTT client, only internal preparations are performed.
// The Client will be functional once this method returns without error. However, for consistency, if
// there is a provided ConnectHandler, it will be notified.
// In the case of an external MQTT client, if any error occurs during the internal preparations - it's returned here.
func (client *Client) Connect() error {
	if client.externalMQTTClient {
		client.wgConnectHandler.Add(1)

		token := client.pahoClient.Subscribe(honoMQTTTopicSubscribeCommands, 1, client.honoMessageHandler)
		if !token.WaitTimeout(client.cfg.subscribeTimeout) || token.Error() != nil {
			client.wgConnectHandler.Done()
			if err := token.Error(); err != nil {
				return err
			}
			return ErrSubscribeTimeout
		}

		go client.notifyClientConnected()
		return nil
	}

	pahoOpts := MQTT.NewClientOptions().
		AddBroker(client.cfg.broker).
		SetClientID(uuid.New().String()).
		SetDefaultPublishHandler(client.defaultMessageHandler).
		SetKeepAlive(client.cfg.keepAlive).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetOnConnectHandler(client.clientConnectHandler).
		SetConnectionLostHandler(client.clientConnectionLostHandler).
		SetTLSConfig(client.cfg.tlsConfig).
		SetConnectTimeout(client.cfg.connectTimeout)

	if client.cfg.credentials != nil {
		pahoOpts = pahoOpts.SetCredentialsProvider(func() (username string, password string) {
			return client.cfg.credentials.Username, client.cfg.credentials.Password
		})
	}

	//create and start a client using the created ClientOptions
	client.pahoClient = MQTT.NewClient(pahoOpts)

	if token := client.pahoClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Disconnect in the case of an external MQTT client, only undoes internal preparations, otherwise - it also disconnects
// the client from the configured Ditto endpoint. A call to Disconnect will cause a ConnectionLostHandler to be notified
// only if an external MQTT client is used.
func (client *Client) Disconnect() {
	var err error
	token := client.pahoClient.Unsubscribe(honoMQTTTopicSubscribeCommands)
	if token.WaitTimeout(client.cfg.unsubscribeTimeout) {
		err = token.Error()
		if client.externalMQTTClient && err == MQTT.ErrNotConnected {
			go client.notifyClientConnectionLost(err) // expected: external MQTT client has already been disconnected
			return
		}
	} else {
		err = ErrUnsubscribeTimeout
	}

	if err != nil {
		ERROR.Printf("error while disconnecting client: %v", err)
	}

	if client.externalMQTTClient { // do not disconnect when external MQTT client, the connection should be managed only externally
		go client.notifyClientConnectionLost(nil)
	} else {
		client.pahoClient.Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds()))
	}
}

// Reply is an auxiliary method to send replies for specific requestIDs if such has been provided along with the incoming protocol.Envelope.
// The requestID must be the same as the one provided with the request protocol.Envelope.
// An error is returned if the reply could not be sent for some reason.
func (client *Client) Reply(requestID string, message *protocol.Envelope) error {
	if err := client.publish(generateHonoResponseTopic(requestID, message.Status), message, 1, false); err != nil {
		return err
	}
	return nil
}

// Send sends a protocol.Envelope to the Client's configured Ditto endpoint.
func (client *Client) Send(message *protocol.Envelope) error {
	if err := client.publish(honoMQTTTopicPublishEvents, message, 1, false); err != nil {
		return err
	}
	return nil
}

// Subscribe ensures that all incoming Ditto messages will be transferred to the provided Handlers.
// As subscribing in Ditto is transport-specific - this is a lightweight version of a default subscription that is applicable in the MQTT use case.
func (client *Client) Subscribe(handlers ...Handler) {
	client.handlersLock.Lock()
	defer client.handlersLock.Unlock()

	if client.handlers == nil {
		client.handlers = make(map[string]Handler)
	}

	for _, handler := range handlers {
		client.handlers[getHandlerName(handler)] = handler
	}
}

// Unsubscribe cancels sending incoming Ditto messages from the client to the provided Handlers
// and removes them from the subscriptions list of the client.
// If Unsubscribe is called without arguments, it will cancel and remove all currently subscribed Handlers.
func (client *Client) Unsubscribe(handlers ...Handler) {
	client.handlersLock.Lock()
	defer client.handlersLock.Unlock()

	if len(handlers) == 0 {
		client.handlers = make(map[string]Handler)
	} else {
		for _, handler := range handlers {
			delete(client.handlers, getHandlerName(handler))
		}
	}
}
