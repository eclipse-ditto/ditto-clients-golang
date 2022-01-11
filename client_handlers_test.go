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
	"sync"
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
	"github.com/eclipse/ditto-clients-golang/internal/mock"
	"github.com/eclipse/ditto-clients-golang/protocol"
	"github.com/golang/mock/gomock"
)

func testHandler(requestID string, message *protocol.Envelope) {}

func TestHonoMessageHandlingSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	wg := sync.WaitGroup{}
	wg.Add(1)

	unitUnderTest := NewClient(&Configuration{})
	validMessage := []byte("{\"test\": 15}")
	requestID := "expected"
	topic := createTopic(requestID)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handler := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	mockMQTTMessage.EXPECT().Payload().Return(validMessage)
	mockMQTTMessage.EXPECT().Topic().Return(topic)

	unitUnderTest.Subscribe(handler)
	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)

	internal.AssertWithTimeout(t, &wg, 5)
}

func TestHonoInvalidMesssageHandling(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	unitUnderTest := NewClient(&Configuration{})
	invalidJSON := []byte("{\"t\"}")

	handler := func(requestID string, message *protocol.Envelope) {
		t.Errorf("handler should not be called")
		t.Fail()
	}

	mockMQTTMessage.EXPECT().Payload().Return(invalidJSON)

	unitUnderTest.Subscribe(handler)
	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)
}

func TestHonoWithoutHandlersDoesNotPanic(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	unitUnderTest := NewClient(&Configuration{})

	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)
}

func TestHonoMultipleHanlders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	wg := sync.WaitGroup{}
	wg.Add(2)

	unitUnderTest := NewClient(&Configuration{})
	validMessage := []byte("{\"test\": 15}")
	requestID := "expected"
	topic := createTopic(requestID)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	mockMQTTMessage.EXPECT().Payload().Return(validMessage)
	mockMQTTMessage.EXPECT().Topic().Return(topic)

	unitUnderTest.Subscribe(handlerOne)
	unitUnderTest.Subscribe(handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)

	internal.AssertWithTimeout(t, &wg, 5)
}

func TestHonoAddMultipleHanlders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	wg := sync.WaitGroup{}
	wg.Add(2)

	unitUnderTest := NewClient(&Configuration{})

	validMessage := []byte("{\"test\": 15}")
	requestID := "expected"
	topic := createTopic(requestID)

	expectedEnvelope, _ := getEnvelope(validMessage)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	mockMQTTMessage.EXPECT().Payload().Return(validMessage)
	mockMQTTMessage.EXPECT().Topic().Return(topic)

	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)

	internal.AssertWithTimeout(t, &wg, 5)
}

func TestRemoveAllHanlders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	unitUnderTest := NewClient(&Configuration{})

	handlerOne := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	mockMQTTMessage.EXPECT().Payload().Times(0)
	mockMQTTMessage.EXPECT().Topic().Times(0)

	// We already know this works from another test
	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.Unsubscribe()

	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)

}

func TestRemoveSingleHanlder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMQTTMessage := mock.NewMockMessage(mockCtrl)

	wg := sync.WaitGroup{}
	wg.Add(1)

	unitUnderTest := NewClient(&Configuration{})

	validMessage := []byte("{\"test\": 15}")
	requestID := "expected"
	topic := createTopic(requestID)
	expectedEnvelope, _ := getEnvelope(validMessage)

	handlerOne := func(requestID string, message *protocol.Envelope) {
		internal.AssertEqual(t, expectedEnvelope, message)
		wg.Done()
	}

	handlerTwo := func(requestID string, message *protocol.Envelope) {
		t.Errorf("should not be called")
		t.Fail()
	}

	mockMQTTMessage.EXPECT().Payload().Return(validMessage)
	mockMQTTMessage.EXPECT().Topic().Return(topic)

	unitUnderTest.Subscribe(handlerOne, handlerTwo)

	unitUnderTest.Unsubscribe(handlerTwo)

	unitUnderTest.honoMessageHandler(nil, mockMQTTMessage)
	internal.AssertWithTimeout(t, &wg, 5)
}

func TestGetHandlerName(t *testing.T) {
	expectedName := "github.com/eclipse/ditto-clients-golang.testHandler"

	actualName := getHandlerName(testHandler)

	internal.AssertEqual(t, expectedName, actualName)
}

func createTopic(requestID string) string {
	return fmt.Sprintf("command///req/%s/dosomething", requestID)
}
