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
	"reflect"
	"testing"
	"time"
)

func TestNewConfiguration(t *testing.T) {
	want := &Configuration{
		keepAlive:         defaultKeepAlive,
		disconnectTimeout: defaultDisconnectTimeout,
	}

	if got := NewConfiguration(); !reflect.DeepEqual(got, want) {
		t.Errorf("NewConfiguration() = %v, want %v", got, want)
	}
}

func TestBroker(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              string
	}{
		"test_empty_broker": {
			testConfiguration: NewConfiguration(),
			want:              "",
		},
		"test_any_broker": {
			testConfiguration: &Configuration{
				broker: "test.broker",
			},
			want: "test.broker",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testConfiguration.Broker(); got != testCase.want {
				t.Errorf("Broker() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestKeepAlive(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_keep_alive": {
			testConfiguration: NewConfiguration(),
			want:              defaultKeepAlive,
		},
		"test_any_keep_alive": {
			testConfiguration: &Configuration{
				keepAlive: 30,
			},
			want: 30,
		},
		"test_empty_keep_alive": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testConfiguration.KeepAlive(); got != testCase.want {
				t.Errorf("KeepAlive() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestDisconnectTimeout(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_disconnect_timeout": {
			testConfiguration: NewConfiguration(),
			want:              defaultDisconnectTimeout,
		},
		"test_any_disconnect_timeout": {
			testConfiguration: &Configuration{
				disconnectTimeout: 30,
			},
			want: 30,
		},
		"test_empty_disconnect_timeout": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testConfiguration.DisconnectTimeout(); got != testCase.want {
				t.Errorf("DisconnectTimeout() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestCredentials(t *testing.T) {
	var (
		emptyCredentials = &Credentials{}
		emptyUsername    = &Credentials{
			Password: "test.password",
		}
		emptyPassword = &Credentials{
			Username: "test.username",
		}
		anyCredentials = &Credentials{
			Username: "test.username",
			Password: "test.password",
		}
	)

	tests := map[string]struct {
		testConfiguration *Configuration
		want              *Credentials
	}{
		"test_nil_credentials": {
			testConfiguration: &Configuration{},
			want:              nil,
		},
		"test_empty_credentials": {
			testConfiguration: &Configuration{
				credentials: emptyCredentials,
			},
			want: emptyCredentials,
		},
		"test_empty_username": {
			testConfiguration: &Configuration{
				credentials: emptyUsername,
			},
			want: emptyUsername,
		},
		"test_empty_password": {
			testConfiguration: &Configuration{
				credentials: emptyPassword,
			},
			want: emptyPassword,
		},
		"test_any_credentials": {
			testConfiguration: &Configuration{
				credentials: anyCredentials,
			},
			want: anyCredentials,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.Credentials()
			if got != testCase.want {
				t.Error("Credentials objects are not the same")
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Credentials() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestConnectHandler(t *testing.T) {
	var mockFunction = func(client *Client) {}

	tests := map[string]struct {
		testConfiguration *Configuration
		want              ConnectHandler
	}{
		"test_nil_connection_handler": {
			testConfiguration: &Configuration{},
			want:              nil,
		},
		"test_any_connection_handler": {
			testConfiguration: &Configuration{
				connectHandler: mockFunction,
			},
			want: mockFunction,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testConfiguration.ConnectHandler(); reflect.ValueOf(got).Pointer() != reflect.ValueOf(testCase.want).Pointer() {
				t.Errorf("ConnectHandler() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestConnectionLostHandler(t *testing.T) {
	var mockFunction = func(client *Client, err error) {}

	tests := map[string]struct {
		testConfiguration *Configuration
		want              ConnectionLostHandler
	}{
		"test_nil_connection_lost_handler": {
			testConfiguration: &Configuration{},
			want:              nil,
		},
		"test_any_connection_lost_handler": {
			testConfiguration: &Configuration{
				connectionLostHandler: mockFunction,
			},
			want: mockFunction,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testConfiguration.ConnectionLostHandler(); reflect.ValueOf(got).Pointer() != reflect.ValueOf(testCase.want).Pointer() {
				t.Errorf("ConnectionLostHandler() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestTLSConfig(t *testing.T) {
	var (
		emptyTLSConfig = &tls.Config{}
		anyTLSConfig   = &tls.Config{
			Certificates: []tls.Certificate{},
		}
	)

	tests := map[string]struct {
		testConfiguration *Configuration
		want              *tls.Config
	}{
		"test_empty_tls_config": {
			testConfiguration: &Configuration{
				tlsConfig: emptyTLSConfig,
			},
			want: emptyTLSConfig,
		},
		"test_nil_tls_config": {
			testConfiguration: &Configuration{},
			want:              nil,
		},
		"test_any_tls_config": {
			testConfiguration: &Configuration{
				tlsConfig: anyTLSConfig,
			},
			want: anyTLSConfig,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.TLSConfig()
			if got != testCase.want {
				t.Error("tls.Config objects are not the same")
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("TLSConfig() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestWithBroker(t *testing.T) {
	arg := "test.broker"

	testConfiguration := &Configuration{}

	want := &Configuration{
		broker: arg,
	}

	if got := testConfiguration.WithBroker(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("WithBroker() = %v, want %v", got, want)
	}
}

func TestWithKeepAlive(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		keepAlive: arg,
	}

	if got := testConfiguration.WithKeepAlive(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("WithKeepAlive() = %v, want %v", got, want)
	}
}

func TestWithDisconnectTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		disconnectTimeout: arg,
	}

	if got := testConfiguration.WithDisconnectTimeout(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("WithDisconnectTimeout() = %v, want %v", got, want)
	}
}

func TestWitWithCredentials(t *testing.T) {
	arg := &Credentials{
		Username: "test.username",
		Password: "test.password",
	}

	testConfiguration := &Configuration{}

	want := &Configuration{
		credentials: arg,
	}

	if got := testConfiguration.WithCredentials(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("WithCredentials() = %v, want %v", got, want)
	}
}

func TestWithConnectHandler(t *testing.T) {
	arg := func(clint *Client) {}

	testConfiguration := &Configuration{}

	want := &Configuration{
		connectHandler: arg,
	}

	if got := testConfiguration.WithConnectHandler(arg); reflect.ValueOf(got.connectHandler).Pointer() != reflect.ValueOf(arg).Pointer() {
		t.Errorf("WithConnectHandler() = %v, want %v", got, want)
	}
}

func TestWithConnectionLostHandler(t *testing.T) {
	arg := func(client *Client, err error) {}

	testConfiguration := &Configuration{}

	want := &Configuration{
		connectionLostHandler: arg,
	}

	if got := testConfiguration.WithConnectionLostHandler(arg); reflect.ValueOf(got.connectionLostHandler).Pointer() != reflect.ValueOf(arg).Pointer() {
		t.Errorf("WithConnectionLostHandler() = %v, want %v", got, want)
	}
}

func TestWithTLSConfig(t *testing.T) {
	arg := &tls.Config{}

	testConfiguration := &Configuration{}

	want := &Configuration{
		tlsConfig: arg,
	}

	if got := testConfiguration.WithTLSConfig(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("WithTLSConfig() = %v, want %v", got, want)
	}
}
