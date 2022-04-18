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
	"encoding/json"
	"errors"
	"github.com/eclipse/ditto-clients-golang/protocol"
	"sync"
	"time"

	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	honoMQTTTopicSubscribeCommands = "command///req/#"
	honoMQTTTopicPublishTelemetry  = "t"
	honoMQTTTopicPublishEvents     = "e"
)

func (client *client) clientConnectHandler(pahoClient MQTT.Client) {
	client.wgConnectHandler.Add(1)
	token := client.pahoClient.Subscribe(honoMQTTTopicSubscribeCommands, 1, client.honoMessageHandler)

	var err error
	if token.WaitTimeout(client.cfg.subscribeTimeout) {
		err = token.Error()
	} else {
		err = ErrSubscribeTimeout
	}

	if err != nil {
		ERROR.Printf("error subscribing to root Hono topic %s : %v", honoMQTTTopicSubscribeCommands, err)
	}
	client.notifyClientConnected()
}

func (client *client) notifyClientConnected() {
	defer client.wgConnectHandler.Done()
	if client.cfg == nil {
		return
	}

	notifyChan := make(chan error, 1)
	var notifyOnce sync.Once
	go func() {
		notifyOnce.Do(func() {
			if client.cfg.connectHandler != nil {
				client.cfg.connectHandler(client)
			}
		})
		notifyChan <- nil
	}()

	select {
	case <-notifyChan:
		DEBUG.Println("notified for client initialization successfully")
	case <-time.After(60 * time.Second):
		ERROR.Printf("%v", errors.New("timed out waiting for initialization notification to be handled"))
	}
}

func (client *client) clientConnectionLostHandler(pahoClient MQTT.Client, err error) {
	client.notifyClientConnectionLost(err)
}

func (client *client) notifyClientConnectionLost(err error) {
	if client.cfg == nil {
		return
	}

	notifyChan := make(chan error, 1)
	var notifyOnce sync.Once
	go func() {
		notifyOnce.Do(func() {
			if client.cfg.connectionLostHandler != nil {
				client.cfg.connectionLostHandler(client, err)
			}
		})
		notifyChan <- nil
	}()

	select {
	case <-notifyChan:
		DEBUG.Println("notified for client connection lost successfully")
	case <-time.After(60 * time.Second):
		ERROR.Printf("%v", errors.New("timed out waiting for connection lost notification to be handled"))
	}
}

func (client *client) publish(topic string, message *protocol.Envelope, qos byte, retained bool) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	token := client.pahoClient.Publish(topic, qos, retained, payload)
	if !token.WaitTimeout(client.cfg.acknowledgeTimeout) {
		return ErrAcknowledgeTimeout
	}
	return token.Error()
}
