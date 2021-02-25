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

package things

import (
	"fmt"
	"github.com/eclipse/ditto-clients-golang/model"
	"github.com/eclipse/ditto-clients-golang/protocol"
)

const (
	inbox              = "inbox"
	outbox             = "outbox"
	pathMessagesFormat = "%s/%s/messages/%s"
)

// Message represents a message entity defined by the Ditto protocol for the Things group that defines an instant communication with the underlying device/implementation.
// This is a special Message that is always bound to a specific Thing instance, it's always exchanged vie the
// Live communication channel and it provides the capabilities to configure:
// - the type of the communication - Inbox, Outbox
// - the entity that was affected - the whole Thing (the default) or a single Feature of the Thing (Feature).
// Note: Only one communication type can be configured to the live message - if using the methods for configuring it - only the last one applies.
// Note: Only one entity that the message targts can be configured to the live message - if using the methods for configuring it - only the last one applies.
type Message struct {
	Topic                *protocol.Topic
	Subject              string
	Mailbox              string
	AddressedPartOfThing string
	Payload              interface{}
}

// NewMessage creates a new Message instance for the defined by the provided NamespacedID Thing.
func NewMessage(thingID *model.NamespacedID) *Message {
	return &Message{
		Topic: (&protocol.Topic{}).
			WithNamespace(thingID.Namespace).
			WithEntityID(thingID.Name).
			WithGroup(protocol.GroupThings).
			WithChannel(protocol.ChannelLive).
			WithCriterion(protocol.CriterionMessages),
		AddressedPartOfThing: "",
	}
}

// Inbox configures the live Message to be sent to the inbox of the target entity, i.e. it defines an incoming communication.
// The Message is configured to serve only one subject - the one provided.
func (msg *Message) Inbox(subject string) *Message {
	msg.Topic.WithAction(protocol.TopicAction(subject))
	msg.Subject = subject
	msg.Mailbox = inbox
	return msg
}

// Outbox configures the live Message to be sent to the outbox of the target entity, i.e. it defines an outgoing communication.
// The Message is configured to serve only one subject - the one provided.
func (msg *Message) Outbox(subject string) *Message {
	msg.Topic.WithAction(protocol.TopicAction(subject))
	msg.Subject = subject
	msg.Mailbox = outbox
	return msg
}

// WithPayload sets the data to be sent in the message, i.e. its content.
func (msg *Message) WithPayload(payload interface{}) *Message {
	msg.Payload = payload
	return msg
}

// Feature configures the Message's target to be the specified by the featureID Thing's Feature.
func (msg *Message) Feature(featureID string) *Message {
	msg.AddressedPartOfThing = fmt.Sprintf(pathThingFeatureFormat, featureID)
	return msg
}

// Envelope generates the Ditto message applying all configurations and optionally all Headers provided.
func (msg *Message) Envelope(headerOpts ...protocol.HeaderOpt) *protocol.Envelope {
	res := &protocol.Envelope{
		Topic: msg.Topic,
		Path:  fmt.Sprintf(pathMessagesFormat, msg.AddressedPartOfThing, msg.Mailbox, msg.Subject),
		Value: msg.Payload,
	}
	if headerOpts != nil {
		res.Headers = protocol.NewHeaders(headerOpts...)
	}
	return res
}
