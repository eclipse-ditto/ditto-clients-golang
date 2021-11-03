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

var (
	testNamespaceID = &model.NamespacedID{
		Namespace: "testNamespace",
		Name:      "testName",
	}
	testFeatureID    = "testFeatureID"
	testPropertyPath = "testProeprtyPath"
	testPropertyID   = "testPropertyID"
	testAttributeID  = "testAttributeID"
)

func TestNewCommand(t *testing.T) {
	want := &Command{
		Topic: &protocol.Topic{
			Namespace: testNamespaceID.Namespace,
			EntityID:  testNamespaceID.Name,
			Group:     protocol.GroupThings,
			Channel:   protocol.ChannelTwin,
			Criterion: protocol.CriterionCommands,
		},
		Path: pathThing,
	}

	if got := NewCommand(testNamespaceID); !reflect.DeepEqual(got, want) {
		t.Errorf("NewCommand() = %v want: %v", got, want)
	}
}

func TestCreate(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Action: protocol.ActionCreate,
		},
		Payload: &model.Thing{},
	}

	if got := testCommand.Create(&model.Thing{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Create() = %v want: %v", got, want)
	}
}

func TestModify(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Action: protocol.ActionModify,
		},
		Payload: &model.Feature{},
	}

	if got := testCommand.Modify(&model.Feature{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Modify() = %v want: %v", got, want)
	}
}

func TestMerge(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Action: protocol.ActionMerge,
		},
		Payload: &model.Feature{},
	}

	if got := testCommand.Merge(&model.Feature{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Merge() = %v want: %v", got, want)
	}
}

func TestRetrieve(t *testing.T) {
	payload := struct {
		ThingIDs []string `json:"thingIds"`
	}{
		ThingIDs: []string{"testNamespace:testName"},
	}

	tests := map[string]struct {
		arg         []model.NamespacedID
		testCommand *Command
		want        *Command
	}{
		"test_command_empty_arguments": {
			arg: []model.NamespacedID{},
			testCommand: &Command{
				Topic: &protocol.Topic{},
			},
			want: &Command{
				Topic: &protocol.Topic{
					Action: protocol.ActionRetrieve,
				},
			},
		},
		"test_command_without_arguments": {
			arg: nil,
			testCommand: &Command{
				Topic: &protocol.Topic{},
			},
			want: &Command{
				Topic: &protocol.Topic{
					Action: protocol.ActionRetrieve,
				},
			},
		},
		"test_command_not_empty_payload_without_arguments": {
			arg: []model.NamespacedID{},
			testCommand: &Command{
				Topic:   &protocol.Topic{},
				Payload: payload,
			},
			want: &Command{
				Topic: &protocol.Topic{
					Action: protocol.ActionRetrieve,
				},
				Payload: payload,
			},
		},
		"test_command_empty_payload_any_argument": {
			arg: []model.NamespacedID{
				{
					Namespace: "testNamespace",
					Name:      "testName",
				},
			},
			testCommand: &Command{
				Topic: &protocol.Topic{},
			},
			want: &Command{
				Topic: &protocol.Topic{
					Action: protocol.ActionRetrieve,
				},
				Payload: payload,
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := testCase.testCommand.Retrieve(testCase.arg...); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Command.Retrieve() = %v want: %v", got, testCase.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Action: protocol.ActionDelete,
		},
	}

	if got := testCommand.Delete(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Delete() = %v want: %v", got, want)
	}
}

func TestPolicyID(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingPolicyID,
	}

	if got := testEmptyCommand.PolicyID(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.PolicyID() = %v want: %v", got, want)
	}
}

func TestDefinition(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingDefinition,
	}

	if got := testEmptyCommand.Definition(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Definition() = %v want: %v", got, want)
	}
}

func TestAttributes(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingAttributes,
	}

	if got := testEmptyCommand.Attributes(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Attributes() = %v want: %v", got, want)
	}
}

func TestAttribute(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingAttributeFormat, testAttributeID),
	}

	if got := testEmptyCommand.Attribute(testAttributeID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Attribute() = %v want: %v", got, want)
	}
}

func TestFeatures(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingFeatures,
	}

	if got := testEmptyCommand.Features(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Features() = %v want: %v", got, want)
	}
}

func TestFeature(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureFormat, testFeatureID),
	}

	if got := testEmptyCommand.Feature(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Feature() = %v want: %v", got, want)
	}
}

func TestFeatureDefinition(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDefinitionFormat, testFeatureID),
	}

	if got := testEmptyCommand.FeatureDefinition(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.FeatureDefinition() = %v want: %v", got, want)
	}
}

func TestFeatureProperties(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeaturePropertiesFormat, testFeatureID),
	}

	if got := testEmptyCommand.FeatureProperties(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.FeatureProperties() = %v want: %v", got, want)
	}
}

func TestFeatureProperty(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeaturePropertyFormat, testFeatureID, testPropertyID),
	}

	if got := testEmptyCommand.FeatureProperty(testFeatureID, testPropertyID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.FeatureProperty() = %v want: %v", got, want)
	}
}

func TestFeatureDesiredProperties(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, testFeatureID),
	}

	if got := testEmptyCommand.FeatureDesiredProperties(testFeatureID); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.FeatureDesiredProperties() = %v want: %v", got, want)
	}
}

func TestFeatureDesiredProperty(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, testFeatureID, testPropertyPath),
	}

	if got := testEmptyCommand.FeatureDesiredProperty(testFeatureID, testPropertyPath); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.FeatureDesiredProperty() = %v want: %v", got, want)
	}
}

func TestLive(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Channel: protocol.ChannelLive,
		},
	}

	if got := testCommand.Live(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Live() = %v want: %v", got, want)
	}
}

func TestTwin(t *testing.T) {
	testCommand := &Command{
		Topic: &protocol.Topic{},
	}

	want := &Command{
		Topic: &protocol.Topic{
			Channel: protocol.ChannelTwin,
		},
	}

	if got := testCommand.Twin(); !reflect.DeepEqual(got, want) {
		t.Errorf("Command.Twin() = %v want: %v", got, want)
	}
}

func TestEnvelope(t *testing.T) {
	cmd := NewCommand(testNamespaceID)

	tests := map[string]struct {
		arg  []protocol.HeaderOpt
		want *protocol.Envelope
	}{
		"test_without_header": {
			arg: nil,
			want: &protocol.Envelope{
				Topic: cmd.Topic,
				Path:  cmd.Path,
				Value: cmd.Payload,
			},
		},
		"test_with_any_headers": {
			arg: []protocol.HeaderOpt{
				protocol.WithChannel("testChannel"),
			},
			want: &protocol.Envelope{
				Topic: cmd.Topic,
				Path:  cmd.Path,
				Value: cmd.Payload,
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
			if got := cmd.Envelope(testCase.arg...); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Command.Envelope() = %v want: %v", got, testCase.want)
			}
		})
	}
}
