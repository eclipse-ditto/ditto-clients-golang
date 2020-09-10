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

package protocol

// Message represents the Ditto's Envelope specification. As a Ditto's message consists of an envelope along with a Ditto-compliant
// payload, the structure is to be used as a ready to use Ditto message.
type Message struct {
	Topic     *Topic      `json:"topic"`
	Headers   *Headers    `json:"headers,omitempty"`
	Path      string      `json:"path"`
	Value     interface{} `json:"value,omitempty"`
	Fields    string      `json:"fields,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Status    int         `json:"status,omitempty"`
	Revision  int64       `json:"revision,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
}

// WithTopic sets the topic of the Message.
func (msg *Message) WithTopic(topic *Topic) *Message {
	msg.Topic = topic
	return msg
}

// WithHeaders sets the Headers of the Message.
func (msg *Message) WithHeaders(headers *Headers) *Message {
	msg.Headers = headers
	return msg
}

// WithPath sets the Ditto path of the Message.
func (msg *Message) WithPath(path string) *Message {
	msg.Path = path
	return msg
}

// WithValue sets the Ditto value of the Message.
func (msg *Message) WithValue(value interface{}) *Message {
	msg.Value = value
	return msg
}

// WithFields sets the fields of the Message as defined by the Ditto protocol specification.
func (msg *Message) WithFields(fields string) *Message {
	msg.Fields = fields
	return msg
}

// WithExtra sets any extra Message configurations as defined by the Ditto protocol specification.
func (msg *Message) WithExtra(extra interface{}) *Message {
	msg.Extra = extra
	return msg
}

// WithStatus sets the Message's status based on the HTTP codes available.
func (msg *Message) WithStatus(status int) *Message {
	msg.Status = status
	return msg
}

// WithRevision sets the current revision number of an entity this Message refers to.
func (msg *Message) WithRevision(revision int64) *Message {
	msg.Revision = revision
	return msg
}

// WithTimestamp sets the timestamp of the Message.
func (msg *Message) WithTimestamp(timestamp string) *Message {
	msg.Timestamp = timestamp
	return msg
}
