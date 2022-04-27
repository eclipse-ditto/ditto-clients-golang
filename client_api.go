// Copyright (c) 2022 Contributors to the Eclipse Foundation
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
	"github.com/eclipse/ditto-clients-golang/protocol"
)

// Handler represents a callback handler that is called on each received message.
// If the underlying transport provides a special requestID related to the Envelope,
// it's also provided to the handler so that chained responses to the ID can be later sent properly.
type Handler func(requestID string, message *protocol.Envelope)

// Client is the Ditto's library main interface definition. The interface is intended to abstract multiple implementations
// over different transports. Client has connect/disconnect capabilities along with the options to subscribe/unsubscribe
// for receiving all Ditto messages being exchanged using the underlying transport.
type Client interface {

	// Connect connects the client to the configured Ditto endpoint provided via the Client's Configuration at creation time.
	// If any error occurs during the connection's initiation - it's returned here.
	// An actual connection status is callbacked to the provided ConnectHandler
	// as soon as the connection is established and all Client's internal preparations are performed.
	// If the connection gets lost during runtime - the ConnectionLostHandler is notified to handle the case.
	Connect() error

	// Disconnect disconnects the client from the configured Ditto endpoint.
	Disconnect()

	// Reply is an auxiliary method to send replies for specific requestIDs if such has been provided along with the incoming protocol.Envelope.
	// The requestID must be the same as the one provided with the request protocol.Envelope.
	// An error is returned if the reply could not be sent for some reason.
	Reply(requestID string, message *protocol.Envelope) error

	// Send sends a protocol.Envelope to the Client's configured Ditto endpoint.
	// An error is returned if the envelope could not be sent for some reason.
	Send(message *protocol.Envelope) error

	// Subscribe ensures that all incoming Ditto messages will be transferred to the provided Handlers.
	Subscribe(handlers ...Handler)

	// Unsubscribe cancels sending incoming Ditto messages from the client to the provided Handlers
	// and removes them from the subscriptions list of the client.
	// If Unsubscribe is called without arguments, it will cancel and remove all currently subscribed Handlers.
	Unsubscribe(handlers ...Handler)
}
