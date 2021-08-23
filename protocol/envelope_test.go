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

package protocol

import (
	"reflect"
	"testing"
)

func TestEnvelopeWithTopic(t *testing.T) {
	t.Run("TestEnvelopeWithTopic", func(t *testing.T) {
		arg := &Topic{
			Namespace: "anamespace",
			EntityID:  "aentityid",
			Group:     GroupThings,
			Channel:   ChannelTwin,
			Criterion: CriterionMessages,
			Action:    ActionSubscribe,
		}
		msg := &Envelope{}
		if got := msg.WithTopic(arg); !reflect.DeepEqual(got.Topic, arg) {
			t.Errorf("Envelope.WithTopic() = %v, want %v", got.Topic, arg)
		}
	})
}

func TestEnvelopeWithHeaders(t *testing.T) {
	t.Run("TestEnvelopeWithHeaders", func(t *testing.T) {
		arg := NewHeaders(WithChannel("something"))
		msg := &Envelope{}
		if got := msg.WithHeaders(arg); !reflect.DeepEqual(got.Headers, arg) {
			t.Errorf("Envelope.WithHeaders() = %v, want %v", got.Headers, arg)
		}
	})
}

func TestEnvelopeWithPath(t *testing.T) {
	t.Run("TestEnvelopeWithPath", func(t *testing.T) {
		arg := "somePath"
		msg := &Envelope{}
		if got := msg.WithPath(arg); !reflect.DeepEqual(got.Path, arg) {
			t.Errorf("Envelope.WithPath() = %v, want %v", got.Path, arg)
		}
	})
}

func TestEnvelopeWithValue(t *testing.T) {
	t.Run("TestEnvelopeWithValue", func(t *testing.T) {
		arg := "someValue"
		msg := &Envelope{}
		if got := msg.WithValue(arg); !reflect.DeepEqual(got.Value, arg) {
			t.Errorf("Envelope.WithValue() = %v, want %v", got.Value, arg)
		}
	})
}

func TestEnvelopeWithFields(t *testing.T) {
	t.Run("TestEnvelopeWithFields", func(t *testing.T) {
		arg := "someFields"
		msg := &Envelope{}
		if got := msg.WithFields(arg); !reflect.DeepEqual(got.Fields, arg) {
			t.Errorf("Envelope.WithFields() = %v, want %v", got.Fields, arg)
		}
	})
}

func TestEnvelopeWithExtra(t *testing.T) {
	t.Run("TestEnvelopeWithExtra", func(t *testing.T) {
		arg := "Extra"
		msg := &Envelope{}
		if got := msg.WithExtra(arg); !reflect.DeepEqual(got.Extra, arg) {
			t.Errorf("Envelope.WithExtra() = %v, want %v", got.Extra, arg)
		}
	})
}

func TestEnvelopeWithStatus(t *testing.T) {
	t.Run("TestEnvelopeWithStatus", func(t *testing.T) {
		arg := 202
		msg := &Envelope{}
		if got := msg.WithStatus(arg); !reflect.DeepEqual(got.Status, arg) {
			t.Errorf("Envelope.WithStatus() = %v, want %v", got.Status, arg)
		}
	})
}

func TestEnvelopeWithRevision(t *testing.T) {
	t.Run("TestEnvelopeWithRevision", func(t *testing.T) {
		arg := int64(1001)
		msg := &Envelope{}
		if got := msg.WithRevision(arg); !reflect.DeepEqual(got.Revision, arg) {
			t.Errorf("Envelope.WithRevision() = %v, want %v", got.Revision, arg)
		}
	})
}

func TestEnvelopeWithTimestamp(t *testing.T) {
	t.Run("TestEnvelopeWithTimestamp", func(t *testing.T) {
		arg := "10"
		msg := &Envelope{}
		if got := msg.WithTimestamp(arg); !reflect.DeepEqual(got.Timestamp, arg) {
			t.Errorf("Envelope.WithTimestamp() = %v, want %v", got.Timestamp, arg)
		}
	})
}
