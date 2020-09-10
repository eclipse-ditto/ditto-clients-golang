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
	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func (client *Client) defaultMessageHandler(mqttClient MQTT.Client, message MQTT.Message) {
	DEBUG.Printf("unexpected message received: %v", message)
}

func (client *Client) honoMessageHandler(mqttClient MQTT.Client, message MQTT.Message) {
	DEBUG.Printf("received Hono message for client subscription: %v", message)
	dittoMsg, err := getMessage(message.Payload())
	if err != nil {
		ERROR.Printf("error getting Ditto message: %v", err)
		return
	}
	requestID := extractHonoRequestId(message.Topic())
	if requestID == "" {
		DEBUG.Printf("no Hono request ID is available in the received message with topic: %s", message.Topic())
	} else {
		DEBUG.Printf("received a command from Hono with request ID: %s", requestID)
	}
	client.handlersLock.RLock()
	defer client.handlersLock.RUnlock()
	if client.handler != nil {
		client.handler(requestID, dittoMsg)
	}
}
