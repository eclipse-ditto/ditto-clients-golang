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
	"errors"
	"fmt"
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestTopicString(t *testing.T) {
	tests := map[string]struct {
		topic *Topic
		want  string
	}{
		"test_topic_string_group_things": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "entity_name",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelTwin) + "/" +
				string(CriterionMessages) + "/" +
				string(ActionSubscribe),
		},
		"test_topic_string_group_policies": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "entity_name",
				Group:      GroupPolicies,
				Channel:    "",
				Criterion:  CriterionCommands,
				Action:     ActionCreate,
			},
			want: "namespace/entity_name/" +
				string(GroupPolicies) + "/" +
				string(CriterionCommands) + "/" +
				string(ActionCreate),
		},
		"test_topic_string_empty": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "entity_name",
				Group:      "",
				Channel:    "",
				Criterion:  CriterionCommands,
				Action:     ActionCreate,
			},
			want: "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.topic.String()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestTopicMarshalJSON(t *testing.T) {

	tests := map[string]struct {
		topic         *Topic
		want          string
		expectedError error
	}{
		"test_marshalJSON_with_all_entities_things": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want:          `"namespace/test/things/twin/messages/subscribe"`,
			expectedError: nil,
		},
		"test_marshalJSON_with_all_entities_policies": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Group:      GroupPolicies,
				Criterion:  CriterionCommands,
				Action:     ActionModify,
			},
			want:          `"namespace/test/policies/commands/modify"`,
			expectedError: nil,
		},
		"test_marshalJSON_without_namespace": {
			topic: &Topic{
				EntityName: "test",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want:          ``,
			expectedError: errors.New("invalid topic: /test/things/twin/messages/subscribe"),
		},
		"test_marshalJSON_without_name": {
			topic: &Topic{
				Namespace: "namespace",
				Group:     GroupThings,
				Channel:   ChannelTwin,
				Criterion: CriterionMessages,
				Action:    ActionSubscribe,
			},
			want:          ``,
			expectedError: errors.New("invalid topic: namespace//things/twin/messages/subscribe"),
		},
		"test_marshalJSON_without_group": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want:          ``,
			expectedError: errors.New("invalid topic: "), // for a missing group the string representation of the Topi is ""
		},
		"test_marshalJSON_without_channel": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Group:      GroupThings,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want:          ``,
			expectedError: errors.New("invalid topic: namespace/test/things//messages/subscribe"),
		},
		"test_marshalJSON_without_criterion": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Action:     ActionSubscribe,
			},
			want:          ``,
			expectedError: errors.New("invalid topic: namespace/test/things/twin//subscribe"),
		},
		"test_marshalJSON_without_action": {
			topic: &Topic{
				Namespace:  "namespace",
				EntityName: "test",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
			},
			want:          `"namespace/test/things/twin/messages"`,
			expectedError: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Log(testName)
			mTopic, mError := testCase.topic.MarshalJSON()
			internal.AssertEqual(t, testCase.expectedError, mError)
			internal.AssertEqual(t, testCase.want, string(mTopic))
		})
	}
}

func TestTopicValidUnmarshalJSON(t *testing.T) {
	tests := []string{
		`"namespace/test/things/twin/messages/subscribe"`,
		`"namespace/test/policies/commands/create"`,
		`"namespace/test/things/live/messages/$set.configuration/name"`,
		`"namespace/test/things/live/messages/$refresh.name"`,
		`"namespace/test/things/live/messages/$refresh"`,
		`"namespace/test/things/live/messages/a"`,
	}

	var topic *Topic
	for _, test := range tests {
		topic = &Topic{}
		topic.UnmarshalJSON([]byte(test))
		internal.AssertEqual(t, test, fmt.Sprintf("%q", topic.String()))
	}
}

func TestTopicInvalidUnmarshalJSON(t *testing.T) {
	tests := []string{
		`"//////"`,
		`"/////"`,
		`"namespace/name/"`,
		`"namespace/name"`,
		`"namespace/name/things"`,
		`"namespace/name/things/commands"`,
		`"namespace/name/things/live/events//create"`,
		`"namespace/name/random_group/commands/modify"`,
		`"namespace/name/random_group/live/events/create"`,
	}

	var topic *Topic
	for _, test := range tests {
		topic = &Topic{}
		err := topic.UnmarshalJSON([]byte(test))
		fmt.Println(err)
		internal.AssertNotNil(t, err)
	}
}

func TestTopicNamespace(t *testing.T) {
	var (
		testValidNamespace    = "namespace"
		testInvalidNamespace  = ":namespace"
		testValidEntityName   = "test"
		testInvalidEntityName = "test§name"
	)

	tests := map[string]struct {
		data       string
		namespace  string
		entityName string
		wantErr    error
	}{
		"test_topic_unmarshal_JSON_valid_namespace_entity_name": {
			data:       `"namespace/test/things/twin/retrieve"`,
			namespace:  testValidNamespace,
			entityName: testValidEntityName,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_empty_namespace_valid_entity_name": {
			data:       `"_/test/things/twin/retrieve"`,
			namespace:  TopicPlaceholder,
			entityName: testValidEntityName,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_valid_namespace_empty_entity_name": {
			data:       `"namespace/_/things/twin/retrieve"`,
			namespace:  testValidNamespace,
			entityName: TopicPlaceholder,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_empty_namespace_empty_entity_name": {
			data:       `"_/_/things/twin/retrieve"`,
			namespace:  TopicPlaceholder,
			entityName: TopicPlaceholder,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_invalid_namespace": {
			data:       `":namespace/test/things/twin/retrieve"`,
			namespace:  "",
			entityName: "",
			wantErr:    errors.New("invalid topic namespaced ID, namespace: " + testInvalidNamespace + ", entity name: " + testValidEntityName),
		},
		"test_topic_unmarshal_JSON_invalid_entity_name": {
			data:       `"namespace/test§name/things/twin/retrieve"`,
			namespace:  "",
			entityName: "",
			wantErr:    errors.New("invalid topic namespaced ID, namespace: " + testValidNamespace + ", entity name: " + testInvalidEntityName),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			topic := &Topic{}
			err := topic.UnmarshalJSON([]byte(testCase.data))
			internal.AssertError(t, testCase.wantErr, err)
			internal.AssertEqual(t, testCase.namespace, topic.Namespace)
			internal.AssertEqual(t, testCase.entityName, topic.EntityName)
		})
	}
}

func TestTopicWithNamespace(t *testing.T) {
	t.Run("TestTopicWithNamespace", func(t *testing.T) {
		arg := "namespace"
		topic := &Topic{
			EntityName: "test",
			Group:      GroupThings,
			Channel:    ChannelTwin,
			Criterion:  CriterionCommands,
			Action:     ActionCreate,
		}
		got := topic.WithNamespace(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.Namespace)
	})
}

func TestTopicWithEntityName(t *testing.T) {
	t.Run("TestTopicWithEntityName", func(t *testing.T) {
		arg := "test"
		topic := &Topic{
			Namespace: "namespace",
			Group:     GroupThings,
			Channel:   ChannelTwin,
			Criterion: CriterionCommands,
			Action:    ActionCreate,
		}
		got := topic.WithEntityName(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.EntityName)
	})
}

func TestTopicWithGroup(t *testing.T) {
	t.Run("TestTopicWithGroup", func(t *testing.T) {
		arg := GroupThings
		topic := &Topic{
			Namespace:  "namespace",
			EntityName: "test",
			Channel:    ChannelTwin,
			Criterion:  CriterionCommands,
			Action:     ActionCreate,
		}
		got := topic.WithGroup(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.Group)
	})
}

func TestTopicWithChannel(t *testing.T) {
	t.Run("TestTopicWithChannel", func(t *testing.T) {
		arg := ChannelTwin
		topic := &Topic{
			Namespace:  "namespace",
			EntityName: "test",
			Group:      GroupThings,
			Criterion:  CriterionCommands,
			Action:     ActionCreate,
		}
		got := topic.WithChannel(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.Channel)
	})
}

func TestTopicWithCriterion(t *testing.T) {
	t.Run("TestTopicWithCriterion", func(t *testing.T) {
		arg := CriterionCommands
		topic := &Topic{
			Namespace:  "namespace",
			EntityName: "test",
			Group:      GroupThings,
			Channel:    ChannelTwin,
			Action:     ActionCreate,
		}
		got := topic.WithCriterion(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.Criterion)
	})
}

func TestTopicWithAction(t *testing.T) {
	t.Run("TestTopicWithAction", func(t *testing.T) {
		arg := ActionCreate
		topic := &Topic{
			Namespace:  "namespace",
			EntityName: "test",
			Group:      GroupThings,
			Channel:    ChannelTwin,
			Criterion:  CriterionCommands,
		}
		got := topic.WithAction(arg)
		internal.AssertEqual(t, "namespace/test/things/twin/commands/create", got.String())
		internal.AssertEqual(t, arg, got.Action)
	})
}
