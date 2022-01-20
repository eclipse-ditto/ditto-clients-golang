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
	"crypto/tls"
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/eclipse/ditto-clients-golang/internal"
	"github.com/eclipse/ditto-clients-golang/internal/mock"
	"github.com/eclipse/ditto-clients-golang/protocol"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/mock/gomock"
)

var (
	mockMQTTClient *mock.MockClient
	mockToken      *mock.MockToken
)

func setup(controller *gomock.Controller) {
	mockMQTTClient = mock.NewMockClient(controller)
	mockToken = mock.NewMockToken(controller)
}

func TestNewClient(t *testing.T) {
	config := &Configuration{}

	want := &Client{
		cfg:      config,
		handlers: map[string]Handler{},
	}

	got := NewClient(config)
	internal.AssertEqual(t, want, got)
}

type mockExecNewClientMQTT func(mockMQTTClient *mock.MockClient, config *Configuration, message string) (*Client, error)

func TestNewClientMQTT(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	tests := map[string]struct {
		arg           *Configuration
		mockExecution mockExecNewClientMQTT
		errorMassage  string
	}{
		"test_connected_client": {
			arg:           &Configuration{},
			mockExecution: mockExecNewClientMQTTNoErrors,
		},
		"test_not_connected_client": {
			arg:           &Configuration{},
			mockExecution: mockExecNewClientMQTTNotConnectedError,
			errorMassage:  "MQTT client is not connected",
		},
		"test_configuration_nil": {
			arg:           nil,
			mockExecution: mockExecNewClientMQTTNoErrors,
		},
		"test_configuration_broker_error": {
			arg: &Configuration{
				broker: "nil",
			},
			mockExecution: mockExecNewClientMQTTConfigurationError,
			errorMassage:  "broker is not expected when using external MQTT client",
		},
		"test_configuration_credentials_error": {
			arg: &Configuration{
				credentials: &Credentials{},
			},
			mockExecution: mockExecNewClientMQTTConfigurationError,
			errorMassage:  "credentials are not expected when using external MQTT client",
		},
		"test_configuration_disconnect_timeout_error": {
			arg: &Configuration{
				disconnectTimeout: 50,
			},
			mockExecution: mockExecNewClientMQTTConfigurationError,
			errorMassage:  "disconnectTimeout is not expected when using external MQTT client",
		},
		"test_configuration_keep_alive_error": {
			arg: &Configuration{
				keepAlive: 50,
			},
			mockExecution: mockExecNewClientMQTTConfigurationError,
			errorMassage:  "keepAlive is not expected when using external MQTT client",
		},
		"test_configuration_TLS_configuration_error": {
			arg: &Configuration{
				tlsConfig: &tls.Config{},
			},
			mockExecution: mockExecNewClientMQTTConfigurationError,
			errorMassage:  "TLS configuration is not expected when using external MQTT client",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			expectedClient, expectedError := testCase.mockExecution(mockMQTTClient, testCase.arg, testCase.errorMassage)
			actualClient, actualError := NewClientMQTT(mockMQTTClient, testCase.arg)

			internal.AssertEqual(t, expectedClient, actualClient)
			internal.AssertError(t, expectedError, actualError)
		})
	}
}

type mockExecConnect func(testWg *sync.WaitGroup) error

func TestConnect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	testWg := &sync.WaitGroup{}

	tests := map[string]struct {
		client   *Client
		mockExec mockExecConnect
	}{
		"test_external_mqtt_client_no_error": {
			client: &Client{
				cfg: &Configuration{
					connectHandler: func(client *Client) {
						testWg.Done()
					},
				},
				pahoClient:         mockMQTTClient,
				externalMQTTClient: true,
			},
			mockExec: mockExecConnectNoError,
		},
		"test_external_mqtt_client_error": {
			client: &Client{
				cfg:                &Configuration{},
				pahoClient:         mockMQTTClient,
				externalMQTTClient: true,
			},
			mockExec: mockExecConnectError,
		},
		"test_external_mqtt_client_timeout_error": {
			client: &Client{
				cfg:                &Configuration{},
				pahoClient:         mockMQTTClient,
				externalMQTTClient: true,
			},
			mockExec: mockExecConnectTimeoutError,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			expectedError := testCase.mockExec(testWg)
			actualError := testCase.client.Connect()
			internal.AssertWithTimeout(t, testWg, 5*time.Second)
			internal.AssertError(t, expectedError, actualError)
		})
	}
}

func TestDisconnectInternalClient(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	client := &Client{
		cfg: &Configuration{
			disconnectTimeout: defaultDisconnectTimeout,
		},
		pahoClient:         mockMQTTClient,
		externalMQTTClient: false,
	}

	mockMQTTClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(time.Duration(0)).Return(false)
	mockMQTTClient.EXPECT().Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds())).Times(1)

	client.Disconnect()
}

type mockExecUnsubscribe func(client *Client, err error)

func TestDisconnectExternalClient(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	testWg := &sync.WaitGroup{}

	tests := map[string]struct {
		client   *Client
		err      error
		mockExec mockExecUnsubscribe
	}{
		"test_disconnect_with_error": {
			client: &Client{
				cfg: &Configuration{
					connectionLostHandler: func(client *Client, err error) {
						testWg.Done()
					},
				},
				pahoClient:         mockMQTTClient,
				externalMQTTClient: true,
			},
			err:      MQTT.ErrNotConnected,
			mockExec: mockExecUnsubscribeError,
		},
		"test_disconnect_without_error": {
			client: &Client{
				cfg: &Configuration{
					connectionLostHandler: func(client *Client, err error) {
						testWg.Done()
					},
				},
				pahoClient:         mockMQTTClient,
				externalMQTTClient: true,
			},
			mockExec: mockExecUnsubscribeNoError,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			testWg.Add(1)
			testCase.mockExec(testCase.client, testCase.err)
			testCase.client.Disconnect()
			internal.AssertWithTimeout(t, testWg, 5*time.Second)
		})
	}
}

type mockExecPublish func(topic string, payload interface{}) error

func TestReply(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	client := &Client{
		cfg:        &Configuration{},
		pahoClient: mockMQTTClient,
	}

	tests := map[string]struct {
		arg           string
		arg2          *protocol.Envelope
		mockExecution mockExecPublish
	}{
		"test_reply_without_error": {
			arg:           "testRequestID",
			arg2:          &protocol.Envelope{},
			mockExecution: mockExecPublishNoErrors,
		},
		"test_reply_token_error": {
			arg:           "testRequestID",
			arg2:          &protocol.Envelope{},
			mockExecution: mockExecPublishErrors,
		},
		"test_reply_timeout_error": {
			arg:           "testRequestID",
			arg2:          &protocol.Envelope{},
			mockExecution: mockExecPublishTimeoutErrors,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			topic := generateHonoResponseTopic(testCase.arg, testCase.arg2.Status)
			payload, _ := json.Marshal(testCase.arg2)
			expectedError := testCase.mockExecution(topic, payload)
			actualError := client.Reply(testCase.arg, testCase.arg2)
			internal.AssertError(t, expectedError, actualError)
		})
	}
}

func TestSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	client := &Client{
		cfg:        &Configuration{},
		pahoClient: mockMQTTClient,
	}

	tests := map[string]struct {
		arg           *protocol.Envelope
		mockExecution mockExecPublish
	}{
		"test_send_without_error": {
			arg:           &protocol.Envelope{},
			mockExecution: mockExecPublishNoErrors,
		},
		"test_send_token_error": {
			arg:           &protocol.Envelope{},
			mockExecution: mockExecPublishErrors,
		},
		"test_send_timeout_error": {
			arg:           &protocol.Envelope{},
			mockExecution: mockExecPublishTimeoutErrors,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			payload, _ := json.Marshal(testCase.arg)
			expectedError := testCase.mockExecution(honoMQTTTopicPublishEvents, payload)
			actualError := client.Send(testCase.arg)

			internal.AssertError(t, expectedError, actualError)
		})
	}
}

func TestSubscribe(t *testing.T) {
	handler := func(requestID string, message *protocol.Envelope) {}
	secondHandler := func(requestID string, message *protocol.Envelope) {}

	tests := map[string]struct {
		arg        []Handler
		testClient *Client
		want       map[string]Handler
	}{
		"test_client_handlers_nil": {
			arg: nil,
			testClient: &Client{
				handlers: nil,
			},
			want: make(map[string]Handler),
		},
		"test_client_handlers_empty": {
			arg: []Handler{
				handler,
			},
			testClient: &Client{},
			want: map[string]Handler{
				getHandlerName(handler): handler,
			},
		},
		"test_client_not_empty_handler": {
			arg: []Handler{
				secondHandler,
			},
			testClient: &Client{
				handlers: map[string]Handler{
					getHandlerName(handler): handler,
				},
			},
			want: map[string]Handler{
				getHandlerName(handler):       handler,
				getHandlerName(secondHandler): secondHandler,
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			testCase.testClient.Subscribe(testCase.arg...)
			handlers := testCase.testClient.handlers
			internal.AssertEqual(t, len(testCase.want), len(handlers))
			for key, element := range testCase.want {
				got := handlers[key]
				if got == nil || reflect.ValueOf(got).Pointer() != reflect.ValueOf(element).Pointer() {
					t.Errorf("Client.Subscribe()= %v, want %v", got, element)
				}
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	handler := func(requestID string, message *protocol.Envelope) {}
	secondHandler := func(requestID string, message *protocol.Envelope) {}

	tests := map[string]struct {
		arg        []Handler
		testClient *Client
		want       map[string]Handler
	}{
		"test_remove_all_handlers": {
			arg: []Handler{},
			testClient: &Client{
				handlers: map[string]Handler{
					getHandlerName(handler):       handler,
					getHandlerName(secondHandler): secondHandler,
				},
			},
			want: make(map[string]Handler),
		},
		"test_remove_nil_argument": {
			arg: nil,
			testClient: &Client{
				handlers: map[string]Handler{
					getHandlerName(handler):       handler,
					getHandlerName(secondHandler): secondHandler,
				},
			},
			want: make(map[string]Handler),
		},
		"test_remove_arg_handler": {
			arg: []Handler{
				handler,
			},
			testClient: &Client{
				handlers: map[string]Handler{
					getHandlerName(handler):       handler,
					getHandlerName(secondHandler): secondHandler,
				},
			}, want: map[string]Handler{
				getHandlerName(secondHandler): secondHandler,
			},
		},
		"test_remove_not_existing_handler": {
			arg: []Handler{
				handler,
			},
			testClient: &Client{
				handlers: map[string]Handler{
					getHandlerName(secondHandler): secondHandler,
				},
			}, want: map[string]Handler{
				getHandlerName(secondHandler): secondHandler,
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			testCase.testClient.Unsubscribe(testCase.arg...)
			handlers := testCase.testClient.handlers
			internal.AssertEqual(t, len(testCase.want), len(handlers))
			for key, element := range testCase.want {
				got := handlers[key]
				internal.AssertNotNil(t, got)
				if reflect.ValueOf(got).Pointer() != reflect.ValueOf(element).Pointer() {
					t.Errorf("Client.Unsubscribe()= %v, want %v", got, element)
				}
			}
		})
	}

}

// Mock executions -------------------------------------------------------------
// NewClientMQTT -------------------------------------------------------------
func mockExecNewClientMQTTNoErrors(mockMQTTClient *mock.MockClient, config *Configuration, _ string) (*Client, error) {
	mockMQTTClient.EXPECT().IsConnected().Return(true)
	return &Client{
		cfg:                config,
		pahoClient:         mockMQTTClient,
		externalMQTTClient: true,
	}, nil
}

func mockExecNewClientMQTTNotConnectedError(mockMQTTClient *mock.MockClient, config *Configuration, message string) (*Client, error) {
	mockMQTTClient.EXPECT().IsConnected().Return(false)
	err := errors.New(message)
	return nil, err
}

func mockExecNewClientMQTTConfigurationError(mockMQTTClient *mock.MockClient, config *Configuration, message string) (*Client, error) {
	mockMQTTClient.EXPECT().IsConnected().Return(true)
	err := errors.New(message)
	return nil, err
}

// MQTTClientPublish -------------------------------------------------------------
func mockExecPublishNoErrors(topic string, payload interface{}) error {
	mockMQTTClient.EXPECT().Publish(topic, byte(1), false, payload).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Return(nil)
	return nil
}

func mockExecPublishErrors(topic string, payload interface{}) error {
	err := MQTT.ErrNotConnected
	mockMQTTClient.EXPECT().Publish(topic, byte(1), false, payload).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Return(err)
	return err
}

func mockExecPublishTimeoutErrors(topic string, payload interface{}) error {
	mockMQTTClient.EXPECT().Publish(topic, byte(1), false, payload).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(false)
	return ErrAcknowledgeTimeout
}

// MQTTClientDisconnect -------------------------------------------------------------
func mockExecUnsubscribeNoError(client *Client, _ error) {
	mockMQTTClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Return(nil)
	mockMQTTClient.EXPECT().Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds())).Times(0)
}

func mockExecUnsubscribeError(client *Client, err error) {
	mockMQTTClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Return(err)
}

// MQTTClientConnect -------------------------------------------------------------
func mockExecConnectNoError(testWg *sync.WaitGroup) error {
	testWg.Add(1)
	var qos byte = 1
	mockMQTTClient.EXPECT().Subscribe(honoMQTTTopicSubscribeCommands, qos, gomock.Any()).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Return(nil)
	return nil
}

func mockExecConnectError(testWg *sync.WaitGroup) error {
	mockMQTTClient.EXPECT().Subscribe(honoMQTTTopicSubscribeCommands, byte(1), gomock.Any()).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(true)
	mockToken.EXPECT().Error().Times(2).Return(MQTT.ErrNotConnected)
	return MQTT.ErrNotConnected
}

func mockExecConnectTimeoutError(testWg *sync.WaitGroup) error {
	mockMQTTClient.EXPECT().Subscribe(honoMQTTTopicSubscribeCommands, byte(1), gomock.Any()).Return(mockToken)
	mockToken.EXPECT().WaitTimeout(gomock.Any()).Return(false)
	mockToken.EXPECT().Error().Return(nil)
	return ErrSubscribeTimeout
}
