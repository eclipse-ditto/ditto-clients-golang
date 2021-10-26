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
	"reflect"
	"testing"

	"github.com/eclipse/ditto-clients-golang/model"
	"github.com/eclipse/ditto-clients-golang/protocol"
)

func TestNewEvent(t *testing.T) {
	want := &Event{
		Topic: &protocol.Topic{
			Namespace: testNamespaceID.Namespace,
			EntityID:  testNamespaceID.Name,
			Group:     protocol.GroupThings,
			Channel:   protocol.ChannelTwin,
			Criterion: protocol.CriterionEvents,
		},
		Path: pathThing,
	}

	if got := NewEvent(testNamespaceID); !reflect.DeepEqual(got, want) {
		t.Errorf("NewEvent() = %v want: %v", got, want)
	}
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

	if got := testEvent.Created(&model.Thing{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Created() = %v want: %v", got, want)
	}
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

	if got := testEvent.Modified(&model.Feature{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Modified() = %v want: %v", got, want)
	}
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

	if got := testEvent.Deleted(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Deleted() = %v want: %v", got, want)
	}
}

func TestEventPolicyID(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingPolicyID,
	}

	if got := testEmptyEvent.PolicyID(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.PolicyID() = %v want: %v", got, want)
	}
}

func TestEventDefinition(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingDefinition,
	}

	if got := testEmptyEvent.Definition(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Definition() = %v want: %v", got, want)
	}
}

func TestEventAttributes(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingAttributes,
	}

	if got := testEmptyEvent.Attributes(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Attributes() = %v want: %v", got, want)
	}
}

func TestEventAttribute(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingAttributeFormat, testAttributeID),
	}

	if got := testEmptyEvent.Attribute(testAttributeID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Attribute() = %v want: %v", got, want)
	}
}

func TestEventFeatures(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: pathThingFeatures,
	}

	if got := testEmptyEvent.Features(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Features() = %v want: %v", got, want)
	}
}

func TestEventFeature(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureFormat, testFeatureID),
	}

	if got := testEmptyEvent.Feature(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Feature() = %v want: %v", got, want)
	}
}

func TestEventFeatureDefinition(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDefinitionFormat, testFeatureID),
	}

	if got := testEmptyEvent.FeatureDefinition(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.FeatureDefinition() = %v want: %v", got, want)
	}
}

func TestEventFeatureProperties(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeaturePropertiesFormat, testFeatureID),
	}

	if got := testEmptyEvent.FeatureProperties(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.FeatureProperties() = %v want: %v", got, want)
	}
}

func TestEventFeatureProperty(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeaturePropertyFormat, testFeatureID, testPropertyID),
	}

	if got := testEmptyEvent.FeatureProperty(testFeatureID, testPropertyID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.FeatureProperty() = %v want: %v", got, want)
	}
}

func TestEventFeatureDesiredProperties(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, testFeatureID),
	}

	if got := testEmptyEvent.FeatureDesiredProperties(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.FeatureDesiredProperties() = %v want: %v", got, want)
	}
}

func TestEventFeatureDesiredProperty(t *testing.T) {
	testEmptyEvent := &Event{}

	want := &Event{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, testFeatureID, testPropertyPath),
	}

	if got := testEmptyEvent.FeatureDesiredProperty(testFeatureID, testPropertyPath); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.FeatureDesiredProperty() = %v want: %v", got, want)
	}
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

	if got := testEvent.Live(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Live() = %v want: %v", got, want)
	}
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

	if got := testEvent.Twin(); !reflect.DeepEqual(got, want) {
		t.Errorf("Event.Twin() = %v want: %v", got, want)
	}
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
			if got := event.Envelope(testCase.arg...); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Event.Envelope() = %v want: %v", got, testCase.want)
			}
		})
	}
}
