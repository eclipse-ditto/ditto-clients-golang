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
	"crypto/tls"
	"time"
)

const (
	defaultDisconnectTimeout = 250 * time.Millisecond
	defaultKeepAlive         = 30 * time.Second
)

// ConnectHandler is called when a successful connection to the configured Ditto endpoint is established and
// all Client's internal preparations are done.
type ConnectHandler func(client *Client)

// ConnectionLostHandler is called is the connection is lost during runtime.
type ConnectionLostHandler func(client *Client, err error)

// Credentials represents a user credentials for authentication used by the underlying connection (e.g. MQTT)
type Credentials struct {
	Username string
	Password string
}

// Configuration provides the Client's configuration.
type Configuration struct {
	broker                string
	keepAlive             time.Duration
	disconnectTimeout     time.Duration
	connectHandler        ConnectHandler
	connectionLostHandler ConnectionLostHandler
	tlsConfig             *tls.Config
	credentials           *Credentials
}

// NewConfiguration creates a new Configuration instance.
func NewConfiguration() *Configuration {
	return &Configuration{
		keepAlive:         defaultKeepAlive,
		disconnectTimeout: defaultDisconnectTimeout,
	}
}

// Broker provides the current MQTT broker the client is to connect to.
func (cfg *Configuration) Broker() string {
	return cfg.broker
}

// KeepAlive provides the keep alive connection's period
// The default is 30 seconds.
func (cfg *Configuration) KeepAlive() time.Duration {
	return cfg.keepAlive
}

// DisconnectTimeout provides the timeout for disconnecting the client.
// The default is 250 milliseconds.
func (cfg *Configuration) DisconnectTimeout() time.Duration {
	return cfg.disconnectTimeout
}

// Credentials provides the currently configured authentication credentials used for the underlying connection.
func (cfg *Configuration) Credentials() *Credentials {
	return cfg.credentials
}

// ConnectHandler provides the currently configured ConnectHandler.
func (cfg *Configuration) ConnectHandler() ConnectHandler {
	return cfg.connectHandler
}

// ConnectionLostHandler provides the currently configured ConnectionLostHandler.
func (cfg *Configuration) ConnectionLostHandler() ConnectionLostHandler {
	return cfg.connectionLostHandler
}

// TLSConfig provides the current TLS configuration for the underlying connection.
func (cfg *Configuration) TLSConfig() *tls.Config {
	return cfg.tlsConfig
}

// WithBroker configures the MQTT's broker the Client to connect to.
func (cfg *Configuration) WithBroker(broker string) *Configuration {
	cfg.broker = broker
	return cfg
}

// WithKeepAlive configures the keep alive time period for the underlying Client's connection.
func (cfg *Configuration) WithKeepAlive(keepAlive time.Duration) *Configuration {
	cfg.keepAlive = keepAlive
	return cfg
}

// WithDisconnectTimeout configures the timeout for disconnection of the Client.
func (cfg *Configuration) WithDisconnectTimeout(disconnectTimeout time.Duration) *Configuration {
	cfg.disconnectTimeout = disconnectTimeout
	return cfg
}

// WithCredentials configures the credentials to be used for authentication by the underlying connection of the Client.
func (cfg *Configuration) WithCredentials(credentials *Credentials) *Configuration {
	cfg.credentials = credentials
	return cfg
}

// WithConnectHandler configures the connectHandler to be notified when the Client's connection is established.
func (cfg *Configuration) WithConnectHandler(connectHandler ConnectHandler) *Configuration {
	cfg.connectHandler = connectHandler
	return cfg
}

// WithConnectionLostHandler configures the connectionLostHandler to be notified is the Client's connection gets lost during rutnime.
func (cfg *Configuration) WithConnectionLostHandler(connectionLostHandler ConnectionLostHandler) *Configuration {
	cfg.connectionLostHandler = connectionLostHandler
	return cfg
}

// WithTLSConfig sets the TLS configuration to be used by the Client's underlying connection.
func (cfg *Configuration) WithTLSConfig(tlsConfig *tls.Config) *Configuration {
	cfg.tlsConfig = tlsConfig
	return cfg
}
