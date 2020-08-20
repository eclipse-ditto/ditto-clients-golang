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

func (client *Client) clientConnectHandler(pahoClient MQTT.Client) {
	if token := client.pahoClient.Subscribe(honoMQTTTopicSubscribeCommands, 1, client.honoMessageHandler); token.Wait() && token.Error() != nil {
		ERROR.Printf("error subscribing to root Hono topic %s : %v", honoMQTTTopicSubscribeCommands, token.Error())
	}
	client.notifyClientConnected()
}

func (client *Client) notifyClientConnected() {
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

func (client *Client) clientConnectionLostHandler(pahoClient MQTT.Client, err error) {
	client.notifyClientConnectionLost(err)
}

func (client *Client) notifyClientConnectionLost(err error) {
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

func (client *Client) publish(topic string, message *protocol.Message, qos byte, retained bool) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if token := client.pahoClient.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
