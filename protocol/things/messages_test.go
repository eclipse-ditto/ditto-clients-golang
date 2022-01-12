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
package things

import (
	"fmt"
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
	"github.com/eclipse/ditto-clients-golang/model"
	"github.com/eclipse/ditto-clients-golang/protocol"
)

func TestNewMessage(t *testing.T) {
	want := &Message{
		Topic: &protocol.Topic{
			Namespace:  testNamespaceID.Namespace,
			EntityName: testNamespaceID.Name,
			Group:      protocol.GroupThings,
			Channel:    protocol.ChannelLive,
			Criterion:  protocol.CriterionMessages,
		},
		AddressedPartOfThing: "",
	}

	got := NewMessage(testNamespaceID)
	internal.AssertEqual(t, want, got)
}

func TestInbox(t *testing.T) {
	arg := "testSubject"

	testMessage := &Message{
		Topic: &protocol.Topic{},
	}

	want := &Message{
		Topic: &protocol.Topic{
			Action: protocol.TopicAction(arg),
		},
		Subject: arg,
		Mailbox: inbox,
	}

	got := testMessage.Inbox(arg)
	internal.AssertEqual(t, want, got)
}

func TestOutbox(t *testing.T) {
	arg := "testSubject"

	testMessage := &Message{
		Topic: &protocol.Topic{},
	}

	want := &Message{
		Topic: &protocol.Topic{
			Action: protocol.TopicAction(arg),
		},
		Subject: arg,
		Mailbox: outbox,
	}

	got := testMessage.Outbox(arg)
	internal.AssertEqual(t, want, got)
}

func TestWithPayload(t *testing.T) {
	arg := &model.Thing{}

	testMessage := &Message{}

	want := &Message{
		Payload: arg,
	}

	got := testMessage.WithPayload(arg)
	internal.AssertEqual(t, want, got)
}

func TestMessageFeature(t *testing.T) {
	testMessage := &Message{}

	want := &Message{
		AddressedPartOfThing: fmt.Sprintf(pathThingFeatureFormat, testFeatureID),
	}

	got := testMessage.Feature(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestMessageEnvelope(t *testing.T) {
	msg := NewMessage(testNamespaceID)

	tests := map[string]struct {
		arg  []protocol.HeaderOpt
		want *protocol.Envelope
	}{
		"test_without_header": {
			arg: nil,
			want: &protocol.Envelope{
				Topic: msg.Topic,
				Path:  fmt.Sprintf(pathMessagesFormat, msg.AddressedPartOfThing, msg.Mailbox, msg.Subject),
				Value: msg.Payload,
			},
		},
		"test_with_any_headers": {
			arg: []protocol.HeaderOpt{
				protocol.WithChannel("testChannel"),
			},
			want: &protocol.Envelope{
				Topic: msg.Topic,
				Path:  fmt.Sprintf(pathMessagesFormat, msg.AddressedPartOfThing, msg.Mailbox, msg.Subject),
				Value: msg.Payload,
				Headers: &protocol.Headers{
					Values: map[string]interface{}{
						protocol.HeaderChannel: "testChannel",
					},
				},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := msg.Envelope(testCase.arg...)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}
