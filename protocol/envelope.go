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

// Envelope represents the Ditto's Envelope specification. As a Ditto's message consists of an envelope along with a Ditto-compliant
// payload, the structure is to be used as a ready to use Ditto message.
type Envelope struct {
	Topic     *Topic      `json:"topic"`
	Headers   Headers     `json:"headers,omitempty"`
	Path      string      `json:"path"`
	Value     interface{} `json:"value,omitempty"`
	Fields    string      `json:"fields,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Status    int         `json:"status,omitempty"`
	Revision  int64       `json:"revision,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
}

// WithTopic sets the topic of the Envelope.
func (msg *Envelope) WithTopic(topic *Topic) *Envelope {
	msg.Topic = topic
	return msg
}

// WithHeaders sets the Headers of the Envelope.
func (msg *Envelope) WithHeaders(headers Headers) *Envelope {
	msg.Headers = headers
	return msg
}

// WithPath sets the Ditto path of the Envelope.
func (msg *Envelope) WithPath(path string) *Envelope {
	msg.Path = path
	return msg
}

// WithValue sets the Ditto value of the Envelope.
func (msg *Envelope) WithValue(value interface{}) *Envelope {
	msg.Value = value
	return msg
}

// WithFields sets the fields of the Envelope as defined by the Ditto protocol specification.
func (msg *Envelope) WithFields(fields string) *Envelope {
	msg.Fields = fields
	return msg
}

// WithExtra sets any extra Envelope configurations as defined by the Ditto protocol specification.
func (msg *Envelope) WithExtra(extra interface{}) *Envelope {
	msg.Extra = extra
	return msg
}

// WithStatus sets the Envelope's status based on the HTTP codes available.
func (msg *Envelope) WithStatus(status int) *Envelope {
	msg.Status = status
	return msg
}

// WithRevision sets the current revision number of an entity this Envelope refers to.
func (msg *Envelope) WithRevision(revision int64) *Envelope {
	msg.Revision = revision
	return msg
}

// WithTimestamp sets the timestamp of the Envelope.
func (msg *Envelope) WithTimestamp(timestamp string) *Envelope {
	msg.Timestamp = timestamp
	return msg
}
