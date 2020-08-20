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
	"github.com/eclipse/ditto-clients-golang/protocol"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"sync"
	"time"
)

const (
	defaultDisconnectTimeout = 250 * time.Millisecond
	defaultKeepAlive         = 30 * time.Second
)

// Handler represents a callback handler that is called on each received message.
// If the underlying transport (e.g. Hono) provides a special requestID related to the Message,
// it's also provided to the handler so that chained responses to the ID can be later sent properly.
type Handler func(requestID string, message *protocol.Message)

// Client is the Ditto's library actual client's imple,entation.
// It provides the connect/disconnect capabilities along with the options to subscribe/unsubscribe
// for receiving all Ditto messages being exchanged using the underlying transport (MQTT/WS).
type Client struct {
	cfg          *Configuration
	pahoClient   MQTT.Client
	handler      Handler
	handlersLock sync.RWMutex
}

// NewClient creates a new Client instance with the provided Configuration.
func NewClient(cfg *Configuration) *Client {
	client := &Client{
		cfg: cfg,
	}
	return client
}

// Connect connects the client to the configured Ditto endpoint provided via the Client's Configuration at creation time.
// If any error occurs during the connection's initiation - it's returned here.
// An actual connection status is callbacked to the provided ConnectHandler
// as soon as the connection is established and all Client's internal preparations are performed.
// If the connection gets lost during runtime - the ConnectionLostHandler is notified to handle the case.
func (client *Client) Connect() error {

	pahoOpts := MQTT.NewClientOptions().
		AddBroker(client.cfg.broker).
		SetClientID(uuid.New().String()).
		SetDefaultPublishHandler(client.defaultMessageHandler).
		SetKeepAlive(client.cfg.keepAlive).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetOnConnectHandler(client.clientConnectHandler).
		SetConnectionLostHandler(client.clientConnectionLostHandler).
		SetTLSConfig(client.cfg.tlsConfig)

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

// Disconnect disconnects the client from the configured Ditto endpoint.
// A call to Disconnect will not cause a ConnectionLostHandler to be called.
func (client *Client) Disconnect() {
	if token := client.pahoClient.Unsubscribe(honoMQTTTopicSubscribeCommands); token.Wait() && token.Error() != nil {
		ERROR.Printf("error while disconnecting client: %v", token.Error())
	}
	client.pahoClient.Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds()))
}

// Reply is an auxiliary method to send replies for specific requestIDs if such has been provided along with the incoming protocol.Message.
// The requestID must be the same as the one provided with the request protocol.Message.
// An error is returned if the reply could not be sent for some reason.
func (client *Client) Reply(requestID string, message *protocol.Message) error {
	if err := client.publish(generateHonoResponseTopic(requestID, message.Status), message, 1, false); err != nil {
		return err
	}
	return nil
}

// Send sends a protocol.Message to the Client's configured Ditto endpoint.
func (client *Client) Send(message *protocol.Message) error {
	if err := client.publish(honoMQTTTopicPublishEvents, message, 1, false); err != nil {
		return err
	}
	return nil
}

// Subscribe ensures that all incoming Ditto messages will be transferred to the provided Handler.
// As subscribing in Ditto is transport-specific - this is a lighweight version of a default subsciption that is applicable in the MQTT use case.
func (client *Client) Subscribe(handler Handler) {
	client.handlersLock.Lock()
	defer client.handlersLock.Unlock()
	client.handler = handler
}

// Unsubscribe ensures that protocol.Message-s will no longer be forwarded to the provided Handler.
func (client *Client) Unsubscribe() {
	client.handlersLock.Lock()
	defer client.handlersLock.Unlock()
	client.handler = nil
}
