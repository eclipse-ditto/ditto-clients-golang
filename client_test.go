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
	mockMqttClient *mock.MockClient
	mockToken      *mock.MockToken
)

func setup(controller *gomock.Controller) {
	mockMqttClient = mock.NewMockClient(controller)
	mockToken = mock.NewMockToken(controller)
}

func TestNewClient(t *testing.T) {
	config := &Configuration{}

	want := &Client{
		cfg:      config,
		handlers: map[string]Handler{},
	}

	if got := NewClient(config); !reflect.DeepEqual(got, want) {
		t.Errorf("NewClient()= %v, want %v", got, want)
	}
}

type mockExecNewClientMqtt func(mockMqttClient *mock.MockClient, config *Configuration, message string) (*Client, error)

func TestNewClientMqtt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	setup(mockCtrl)

	tests := map[string]struct {
		arg           *Configuration
		mockExecution mockExecNewClientMqtt
		errorMassage  string
	}{
		"test_connected_client": {
			arg:           &Configuration{},
			mockExecution: mockExecNewClientMqttNoErrors,
		},
		"test_not_connected_client": {
			arg:           &Configuration{},
			mockExecution: mockExecNewClientMqttNotConnectedError,
			errorMassage:  "MQTT client is not connected",
		},
		"test_configuration_nil": {
			arg:           nil,
			mockExecution: mockExecNewClientMqttNoErrors,
		},
		"test_configuration_broker_error": {
			arg: &Configuration{
				broker: "nil",
			},
			mockExecution: mockExecNewClientMqttConfigurationError,
			errorMassage:  "broker is not expected when using external MQTT client",
		},
		"test_configuration_credentials_error": {
			arg: &Configuration{
				credentials: &Credentials{},
			},
			mockExecution: mockExecNewClientMqttConfigurationError,
			errorMassage:  "credentials are not expected when using external MQTT client",
		},
		"test_configuration_disconnect_timeout_error": {
			arg: &Configuration{
				disconnectTimeout: 50,
			},
			mockExecution: mockExecNewClientMqttConfigurationError,
			errorMassage:  "disconnectTimeout is not expected when using external MQTT client",
		},
		"test_configuration_keep_alive_error": {
			arg: &Configuration{
				keepAlive: 50,
			},
			mockExecution: mockExecNewClientMqttConfigurationError,
			errorMassage:  "keepAlive is not expected when using external MQTT client",
		},
		"test_configuration_TLS_configuration_error": {
			arg: &Configuration{
				tlsConfig: &tls.Config{},
			},
			mockExecution: mockExecNewClientMqttConfigurationError,
			errorMassage:  "TLS configuration is not expected when using external MQTT client",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			expectedClient, expectedError := testCase.mockExecution(mockMqttClient, testCase.arg, testCase.errorMassage)
			actualClient, actualError := NewClientMqtt(mockMqttClient, testCase.arg)

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
				pahoClient:         mockMqttClient,
				externalMqttClient: true,
			},
			mockExec: mockExecConnectNoError,
		},
		"test_external_mqtt_client_error": {
			client: &Client{
				pahoClient:         mockMqttClient,
				externalMqttClient: true,
			},
			mockExec: mockExecConnectError,
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
		pahoClient:         mockMqttClient,
		externalMqttClient: false,
	}

	mockMqttClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().Wait().Return(false)
	mockMqttClient.EXPECT().Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds())).Times(1)

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
				pahoClient:         mockMqttClient,
				externalMqttClient: true,
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
				pahoClient:         mockMqttClient,
				externalMqttClient: true,
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
		pahoClient: mockMqttClient,
	}

	tests := map[string]struct {
		arg           string
		arg2          *protocol.Envelope
		mockExecution mockExecPublish
	}{
		"test_send_without_error": {
			arg:           "testRequestID",
			arg2:          &protocol.Envelope{},
			mockExecution: mockExecPublishNoErrors,
		},
		"test_send_token_error": {
			arg:           "testRequestID",
			arg2:          &protocol.Envelope{},
			mockExecution: mockExecPublishErrors,
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
		pahoClient: mockMqttClient,
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
			internal.AssertEqual(t, len(handlers), len(testCase.want))
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
			internal.AssertEqual(t, len(handlers), len(testCase.want))
			for key, element := range testCase.want {
				got := handlers[key]
				if got == nil || reflect.ValueOf(got).Pointer() != reflect.ValueOf(element).Pointer() {
					t.Errorf("Client.Unsubscribe()= %v, want %v", got, element)
				}
			}
		})
	}

}

// Mock executions -------------------------------------------------------------
// NewClientMqtt -------------------------------------------------------------
func mockExecNewClientMqttNoErrors(mockMqttClient *mock.MockClient, config *Configuration, _ string) (*Client, error) {
	mockMqttClient.EXPECT().IsConnected().Return(true)
	return &Client{
		cfg:                config,
		pahoClient:         mockMqttClient,
		externalMqttClient: true,
	}, nil
}

func mockExecNewClientMqttNotConnectedError(mockMqttClient *mock.MockClient, config *Configuration, message string) (*Client, error) {
	mockMqttClient.EXPECT().IsConnected().Return(false)
	err := errors.New(message)
	return nil, err
}

func mockExecNewClientMqttConfigurationError(mockMqttClient *mock.MockClient, config *Configuration, message string) (*Client, error) {
	mockMqttClient.EXPECT().IsConnected().Return(true)
	err := errors.New(message)
	return nil, err
}

// MqttClientPublish -------------------------------------------------------------
func mockExecPublishNoErrors(topic string, payload interface{}) error {
	mockToken.EXPECT().Wait().Return(false)
	mockMqttClient.EXPECT().Publish(topic, byte(1), false, payload).Return(mockToken)
	return nil
}

func mockExecPublishErrors(topic string, payload interface{}) error {
	err := MQTT.ErrNotConnected
	mockToken.EXPECT().Wait().Return(true)
	mockToken.EXPECT().Error().AnyTimes().Return(err)
	mockMqttClient.EXPECT().Publish(topic, byte(1), false, payload).Return(mockToken)
	return err
}

// MqttClientDisconnect -------------------------------------------------------------
func mockExecUnsubscribeNoError(client *Client, _ error) {
	mockMqttClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().Wait().Return(false)
	mockMqttClient.EXPECT().Disconnect(uint(client.cfg.disconnectTimeout.Milliseconds())).Times(0)
}

func mockExecUnsubscribeError(client *Client, err error) {
	mockMqttClient.EXPECT().Unsubscribe(honoMQTTTopicSubscribeCommands).Return(mockToken)
	mockToken.EXPECT().Wait().Return(true)
	mockToken.EXPECT().Error().AnyTimes().Return(err)
}

// MqttClientConnect -------------------------------------------------------------
func mockExecConnectNoError(testWg *sync.WaitGroup) error {
	testWg.Add(1)
	var qos byte = 1
	mockMqttClient.EXPECT().Subscribe(honoMQTTTopicSubscribeCommands, qos, gomock.Any()).Return(mockToken)
	mockToken.EXPECT().Wait().Return(false)
	return nil
}

func mockExecConnectError(testWg *sync.WaitGroup) error {
	mockMqttClient.EXPECT().Subscribe(honoMQTTTopicSubscribeCommands, byte(1), gomock.Any()).Return(mockToken)
	mockToken.EXPECT().Wait().Return(true)
	mockToken.EXPECT().Error().AnyTimes().Return(MQTT.ErrNotConnected)
	return MQTT.ErrNotConnected
}
