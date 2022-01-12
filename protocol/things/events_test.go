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

func TestNewEvent(t *testing.T) {
	want := &Event{
		Topic: &protocol.Topic{
			Namespace:  testNamespaceID.Namespace,
			EntityName: testNamespaceID.Name,
			Group:      protocol.GroupThings,
			Channel:    protocol.ChannelTwin,
			Criterion:  protocol.CriterionEvents,
		},
		Path: pathThing,
	}

	got := NewEvent(testNamespaceID)
	internal.AssertEqual(t, want, got)
}

func TestCreated(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Action: protocol.ActionCreated,
		},
		Payload: &model.Thing{},
	}

	got := testEvent.Created(&model.Thing{})
	internal.AssertEqual(t, want, got)
}

func TestModified(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Action: protocol.ActionModified,
		},
		Payload: &model.Feature{},
	}

	got := testEvent.Modified(&model.Feature{})
	internal.AssertEqual(t, want, got)
}

func TestMerged(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Action: protocol.ActionMerged,
		},
		Payload: &model.Feature{},
	}

	got := testEvent.Merged(&model.Feature{})
	internal.AssertEqual(t, want, got)
}

func TestDeleted(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Action: protocol.ActionDeleted,
		},
	}

	got := testEvent.Deleted()
	internal.AssertEqual(t, want, got)
}

func TestEventPolicyID(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingPolicyID,
	}

	got := testEmptyEvent.PolicyID()
	internal.AssertEqual(t, want, got)
}

func TestEventDefinition(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingDefinition,
	}

	got := testEmptyEvent.Definition()
	internal.AssertEqual(t, want, got)
}

func TestEventAttributes(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingAttributes,
	}

	got := testEmptyEvent.Attributes()
	internal.AssertEqual(t, want, got)
}

func TestEventAttribute(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingAttributeFormat, testAttributeID),
	}

	got := testEmptyEvent.Attribute(testAttributeID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatures(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingFeatures,
	}

	got := testEmptyEvent.Features()
	internal.AssertEqual(t, want, got)
}

func TestEventFeature(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureFormat, testFeatureID),
	}

	got := testEmptyEvent.Feature(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatureDefinition(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDefinitionFormat, testFeatureID),
	}

	got := testEmptyEvent.FeatureDefinition(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatureProperties(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeaturePropertiesFormat, testFeatureID),
	}

	got := testEmptyEvent.FeatureProperties(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatureProperty(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeaturePropertyFormat, testFeatureID, testPropertyID),
	}

	got := testEmptyEvent.FeatureProperty(testFeatureID, testPropertyID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatureDesiredProperties(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, testFeatureID),
	}

	got := testEmptyEvent.FeatureDesiredProperties(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestEventFeatureDesiredProperty(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, testFeatureID, testPropertyPath),
	}

	got := testEmptyEvent.FeatureDesiredProperty(testFeatureID, testPropertyPath)
	internal.AssertEqual(t, want, got)
}

func TestEventLive(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Channel: protocol.ChannelLive,
		},
	}

	got := testEvent.Live()
	internal.AssertEqual(t, want, got)
}

func TestEventTwin(t *testing.T) {
	testEvent := &Event{
		Topic: &protocol.Topic{},
	}

	want := &Event{
		Topic: &protocol.Topic{
			Channel: protocol.ChannelTwin,
		},
	}

	got := testEvent.Twin()
	internal.AssertEqual(t, want, got)
}

func TestEventEnvelope(t *testing.T) {
	event := NewEvent(testNamespaceID)

	tests := map[string]struct {
		arg  []protocol.HeaderOpt
		want *protocol.Envelope
	}{
		"test_without_header": {
			arg: nil,
			want: &protocol.Envelope{
				Topic: event.Topic,
				Path:  event.Path,
				Value: event.Payload,
			},
		},
		"test_with_any_headers": {
			arg: []protocol.HeaderOpt{
				protocol.WithChannel("testChannel"),
			},
			want: &protocol.Envelope{
				Topic: event.Topic,
				Path:  event.Path,
				Value: event.Payload,
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
			got := event.Envelope(testCase.arg...)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}
