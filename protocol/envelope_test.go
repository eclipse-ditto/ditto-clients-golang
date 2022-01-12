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
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestEnvelopeWithTopic(t *testing.T) {
	t.Run("TestEnvelopeWithTopic", func(t *testing.T) {
		arg := &Topic{
			Namespace:  "namespace",
			EntityName: "entity_name",
			Group:      GroupThings,
			Channel:    ChannelTwin,
			Criterion:  CriterionMessages,
			Action:     ActionSubscribe,
		}
		msg := &Envelope{}

		got := msg.WithTopic(arg)
		internal.AssertEqual(t, arg, got.Topic)
	})
}

func TestEnvelopeWithHeaders(t *testing.T) {
	t.Run("TestEnvelopeWithHeaders", func(t *testing.T) {
		arg := NewHeaders(WithChannel("something"))
		msg := &Envelope{}

		got := msg.WithHeaders(arg)
		internal.AssertEqual(t, arg, got.Headers)
	})
}

func TestEnvelopeWithPath(t *testing.T) {
	t.Run("TestEnvelopeWithPath", func(t *testing.T) {
		arg := "somePath"
		msg := &Envelope{}

		got := msg.WithPath(arg)
		internal.AssertEqual(t, arg, got.Path)
	})
}

func TestEnvelopeWithValue(t *testing.T) {
	t.Run("TestEnvelopeWithValue", func(t *testing.T) {
		arg := "someValue"
		msg := &Envelope{}

		got := msg.WithValue(arg)
		internal.AssertEqual(t, arg, got.Value)
	})
}

func TestEnvelopeWithFields(t *testing.T) {
	t.Run("TestEnvelopeWithFields", func(t *testing.T) {
		arg := "someFields"
		msg := &Envelope{}

		got := msg.WithFields(arg)
		internal.AssertEqual(t, arg, got.Fields)
	})
}

func TestEnvelopeWithExtra(t *testing.T) {
	t.Run("TestEnvelopeWithExtra", func(t *testing.T) {
		arg := "Extra"
		msg := &Envelope{}

		got := msg.WithExtra(arg)
		internal.AssertEqual(t, arg, got.Extra)
	})
}

func TestEnvelopeWithStatus(t *testing.T) {
	t.Run("TestEnvelopeWithStatus", func(t *testing.T) {
		arg := 202
		msg := &Envelope{}

		got := msg.WithStatus(arg)
		internal.AssertEqual(t, arg, got.Status)
	})
}

func TestEnvelopeWithRevision(t *testing.T) {
	t.Run("TestEnvelopeWithRevision", func(t *testing.T) {
		arg := int64(1001)
		msg := &Envelope{}

		got := msg.WithRevision(arg)
		internal.AssertEqual(t, arg, got.Revision)
	})
}

func TestEnvelopeWithTimestamp(t *testing.T) {
	t.Run("TestEnvelopeWithTimestamp", func(t *testing.T) {
		arg := "10"
		msg := &Envelope{}

		got := msg.WithTimestamp(arg)
		internal.AssertEqual(t, arg, got.Timestamp)
	})
}
