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
			Namespace:  testNamespaceID.Namespace,
			EntityName: testNamespaceID.Name,
			Group:      protocol.GroupThings,
			Channel:    protocol.ChannelTwin,
			Criterion:  protocol.CriterionCommands,
		},
		Path: pathThing,
	}

	got := NewCommand(testNamespaceID)
	internal.AssertEqual(t, want, got)
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

	got := testCommand.Create(&model.Thing{})
	internal.AssertEqual(t, want, got)
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

	got := testCommand.Modify(&model.Feature{})
	internal.AssertEqual(t, want, got)
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

	got := testCommand.Merge(&model.Feature{})
	internal.AssertEqual(t, want, got)
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
			got := testCase.testCommand.Retrieve(testCase.arg...)
			internal.AssertEqual(t, testCase.want, got)
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

	got := testCommand.Delete()
	internal.AssertEqual(t, want, got)
}

func TestPolicyID(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingPolicyID,
	}

	got := testEmptyCommand.PolicyID()
	internal.AssertEqual(t, want, got)
}

func TestDefinition(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingDefinition,
	}

	got := testEmptyCommand.Definition()
	internal.AssertEqual(t, want, got)
}

func TestAttributes(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingAttributes,
	}

	got := testEmptyCommand.Attributes()
	internal.AssertEqual(t, want, got)
}

func TestAttribute(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingAttributeFormat, testAttributeID),
	}

	got := testEmptyCommand.Attribute(testAttributeID)
	internal.AssertEqual(t, want, got)
}

func TestFeatures(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: pathThingFeatures,
	}

	got := testEmptyCommand.Features()
	internal.AssertEqual(t, want, got)
}

func TestFeature(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureFormat, testFeatureID),
	}

	got := testEmptyCommand.Feature(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestFeatureDefinition(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDefinitionFormat, testFeatureID),
	}

	got := testEmptyCommand.FeatureDefinition(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestFeatureProperties(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeaturePropertiesFormat, testFeatureID),
	}

	got := testEmptyCommand.FeatureProperties(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestFeatureProperty(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeaturePropertyFormat, testFeatureID, testPropertyID),
	}

	got := testEmptyCommand.FeatureProperty(testFeatureID, testPropertyID)
	internal.AssertEqual(t, want, got)
}

func TestFeatureDesiredProperties(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, testFeatureID),
	}

	got := testEmptyCommand.FeatureDesiredProperties(testFeatureID)
	internal.AssertEqual(t, want, got)
}

func TestFeatureDesiredProperty(t *testing.T) {
	testEmptyCommand := &Command{}

	want := &Command{
		Path: fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, testFeatureID, testPropertyPath),
	}

	got := testEmptyCommand.FeatureDesiredProperty(testFeatureID, testPropertyPath)
	internal.AssertEqual(t, want, got)
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

	got := testCommand.Live()
	internal.AssertEqual(t, want, got)
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

	got := testCommand.Twin()
	internal.AssertEqual(t, want, got)
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
				Topic:   cmd.Topic,
				Path:    cmd.Path,
				Value:   cmd.Payload,
				Headers: protocol.Headers{protocol.HeaderChannel: "testChannel"},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := cmd.Envelope(testCase.arg...)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}
