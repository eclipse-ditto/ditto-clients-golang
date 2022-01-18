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

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestNewConfiguration(t *testing.T) {
	want := &Configuration{
		keepAlive:          defaultKeepAlive,
		disconnectTimeout:  defaultDisconnectTimeout,
		connectTimeout:     defaultConnectTimeout,
		acknowledgeTimeout: defaultAcknowledgeTimeout,
		subscribeTimeout:   defaultSubscribeTimeout,
		unsubscribeTimeout: defaultUnsubscribeTimeout,
	}

	got := NewConfiguration()
	internal.AssertEqual(t, want, got)
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
			got := testCase.testConfiguration.Broker()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}
func TestConnectTimeout(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_connect_timeout": {
			testConfiguration: NewConfiguration(),
			want:              defaultConnectTimeout,
		},
		"test_any_connect_timeout": {
			testConfiguration: &Configuration{
				connectTimeout: 30,
			},
			want: 30,
		},
		"test_empty_connect_timeout": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.ConnectTimeout()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestAcknowledgeTimeout(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_acknowledge_timeout": {
			testConfiguration: NewConfiguration(),
			want:              defaultAcknowledgeTimeout,
		},
		"test_any_acknowledge_timeout": {
			testConfiguration: &Configuration{
				acknowledgeTimeout: 30,
			},
			want: 30,
		},
		"test_empty_acknowledge_timeout": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.AcknowledgeTimeout()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestSubscribeTimeout(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_subscribe_timeout": {
			testConfiguration: NewConfiguration(),
			want:              defaultSubscribeTimeout,
		},
		"test_any_subscribe_timeout": {
			testConfiguration: &Configuration{
				subscribeTimeout: 30,
			},
			want: 30,
		},
		"test_empty_subscribe_timeout": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.SubscribeTimeout()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestUnsubscribeTimeout(t *testing.T) {
	tests := map[string]struct {
		testConfiguration *Configuration
		want              time.Duration
	}{
		"test_default_unsubscribe_timeout": {
			testConfiguration: NewConfiguration(),
			want:              defaultUnsubscribeTimeout,
		},
		"test_any_unsubscribe_timeout": {
			testConfiguration: &Configuration{
				unsubscribeTimeout: 30,
			},
			want: 30,
		},
		"test_empty_unsubscribe_timeout": {
			testConfiguration: &Configuration{},
			want:              0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testConfiguration.UnsubscribeTimeout()
			internal.AssertEqual(t, testCase.want, got)
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
			got := testCase.testConfiguration.KeepAlive()
			internal.AssertEqual(t, testCase.want, got)
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
			got := testCase.testConfiguration.DisconnectTimeout()
			internal.AssertEqual(t, testCase.want, got)
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
			internal.AssertEqual(t, testCase.want, got)
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
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestWithBroker(t *testing.T) {
	arg := "test.broker"

	testConfiguration := &Configuration{}

	want := &Configuration{
		broker: arg,
	}

	got := testConfiguration.WithBroker(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithConnectTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		connectTimeout: arg,
	}

	got := testConfiguration.WithConnectTimeout(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithAcknowledgeTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		acknowledgeTimeout: arg,
	}

	got := testConfiguration.WithAcknowledgeTimeout(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithSubscribeTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		subscribeTimeout: arg,
	}

	got := testConfiguration.WithSubscribeTimeout(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithUnsubscribeTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		unsubscribeTimeout: arg,
	}

	got := testConfiguration.WithUnsubscribeTimeout(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithKeepAlive(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		keepAlive: arg,
	}

	got := testConfiguration.WithKeepAlive(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithDisconnectTimeout(t *testing.T) {
	arg := time.Second

	testConfiguration := &Configuration{}

	want := &Configuration{
		disconnectTimeout: arg,
	}

	got := testConfiguration.WithDisconnectTimeout(arg)
	internal.AssertEqual(t, want, got)
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

	got := testConfiguration.WithCredentials(arg)
	internal.AssertEqual(t, want, got)
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

	got := testConfiguration.WithTLSConfig(arg)
	internal.AssertEqual(t, want, got)
}
