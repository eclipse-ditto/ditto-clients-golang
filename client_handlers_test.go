// Copyright (c) 2021 Contributors to the Eclipse Foundation
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
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/eclipse/ditto-clients-golang/protocol"
)

func TestMessageHandlingSuccess(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	unitUnderTest := NewClient(&Configuration{})
	validMessage := []byte("{\"test\": 15}")
	requestId := "expected"
	topic := createTopic(requestId)

	mqttMessage := newMockMQTTMessage(validMessage, topic)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handler := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	unitUnderTest.Subscribe(handler)
	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertEqual(t, mqttMessage.payloadCalls, 1)

	assertPassed(t, &wg, 5)
}

func TestInvalidMessageHandling(t *testing.T) {
	unitUnderTest := NewClient(&Configuration{})
	invalidJson := []byte("{\"t\"}")
	requestId := "expected"
	topic := createTopic(requestId)

	mqttMessage := newMockMQTTMessage(invalidJson, topic)

	handler := func(requestID string, message *protocol.Envelope) {
		t.Errorf("handler should not be called")
		t.Fail()
	}

	unitUnderTest.Subscribe(handler)
	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertEqual(t, mqttMessage.payloadCalls, 1)
}

func TestWithoutHandlersDoesNotPanic(t *testing.T) {
	mqttMessage := newMockMQTTMessage(nil, "")
	unitUnderTest := NewClient(&Configuration{})

	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertEqual(t, mqttMessage.payloadCalls, 0)
}

func TestMultipleHandlers(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	unitUnderTest := NewClient(&Configuration{})
	validMessage := []byte("{\"test\": 15}")
	requestId := "expected"
	topic := createTopic(requestId)
	mqttMessage := newMockMQTTMessage(validMessage, topic)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	unitUnderTest.Subscribe(handlerOne)
	unitUnderTest.Subscribe(handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertPassed(t, &wg, 5)
	assertEqual(t, mqttMessage.payloadCalls, 1)
	assertEqual(t, mqttMessage.topicCalls, 1)
}

func TestAddMultipleHandlers(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	unitUnderTest := NewClient(&Configuration{})

	validMessage := []byte("{\"test\": 15}")
	requestId := "expected"
	topic := createTopic(requestId)
	mqttMessage := newMockMQTTMessage(validMessage, topic)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertPassed(t, &wg, 5)
	assertEqual(t, mqttMessage.payloadCalls, 1)
	assertEqual(t, mqttMessage.topicCalls, 1)
}

func TestRemoveAllHandlers(t *testing.T) {
	mqttMessage := newMockMQTTMessage(nil, "")
	unitUnderTest := NewClient(&Configuration{})

	handlerOne := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	// We already know this works from another test
	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.Unsubscribe()

	unitUnderTest.honoMessageHandler(nil, mqttMessage)
	assertEqual(t, mqttMessage.payloadCalls, 0)
	assertEqual(t, mqttMessage.topicCalls, 0)

}

func TestRemoveSingleHandler(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	unitUnderTest := NewClient(&Configuration{})

	validMessage := []byte("{\"test\": 15}")
	requestId := "expected"
	topic := createTopic(requestId)
	expectedEnvelope, _ := getEnvelope(validMessage)
	mqttMessage := newMockMQTTMessage(validMessage, topic)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		assertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.Unsubscribe(handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mqttMessage)

	assertPassed(t, &wg, 5)
	assertEqual(t, mqttMessage.payloadCalls, 1)
	assertEqual(t, mqttMessage.topicCalls, 1)

}

func testHandler(requestID string, message *protocol.Envelope) {}

func TestGetHandlerName(t *testing.T) {
	expectedName := "github.com/eclipse/ditto-clients-golang.testHandler"

	actualName := getHandlerName(testHandler)

	assertEqual(t, expectedName, actualName)
}

func createTopic(requestId string) string {
	return fmt.Sprintf("command///req/%s/dosomething", requestId)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v , got %v", expected, actual)
		t.Fail()
	}
}

func assertPassed(t *testing.T, wg *sync.WaitGroup, timeout time.Duration) {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return // completed normally
	case <-time.After(timeout * time.Second):
		t.Fatal("timed out")
	}
}

// mockMQTTMessage is a mock of MQTTMessage interface
type mockMQTTMessage struct {
	payload      []byte
	payloadCalls int
	payloadMutex sync.Mutex
	topic        string
	topicCalls   int
	topicMutex   sync.Mutex
}

func newMockMQTTMessage(payload []byte, topic string) *mockMQTTMessage {
	return &mockMQTTMessage{
		payload:      payload,
		payloadCalls: 0,
		payloadMutex: sync.Mutex{},

		topic:      topic,
		topicCalls: 0,
		topicMutex: sync.Mutex{},
	}
}

// Duplicate mocks base method
func (m *mockMQTTMessage) Duplicate() bool { return false }

// Qos mocks base method
func (m *mockMQTTMessage) Qos() byte { return byte(0) }

// Retained mocks base method
func (m *mockMQTTMessage) Retained() bool { return false }

// Topic mocks base method
func (m *mockMQTTMessage) Topic() string {
	m.topicMutex.Lock()
	defer m.topicMutex.Unlock()

	m.topicCalls++

	return m.topic
}

// MessageID mocks base method
func (m *mockMQTTMessage) MessageID() uint16 { return uint16(1) }

// Payload mocks base method
func (m *mockMQTTMessage) Payload() []byte {
	m.payloadMutex.Lock()
	defer m.payloadMutex.Unlock()

	m.payloadCalls++

	return m.payload
}

// Ack mocks base method
func (m *mockMQTTMessage) Ack() {}
