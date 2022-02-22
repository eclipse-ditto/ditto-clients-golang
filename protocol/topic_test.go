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
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestTopicString(t *testing.T) {
	tests := map[string]struct {
		topic Topic
		want  string
	}{
		"test_topic_string_group_things": {
			topic: Topic{
				Namespace:  "namespace",
				EntityName: "entity_name",
				Group:      GroupThings,
				Channel:    ChannelTwin,
				Criterion:  CriterionMessages,
				Action:     ActionSubscribe,
			},
			want: "namespace/entity_name/" + string(GroupThings) + "/" + string(ChannelTwin) + "/" +
				string(CriterionMessages) + "/" + string(ActionSubscribe),
		},
		"test_topic_string_group_policies": {
			topic: Topic{
				Namespace:  "namespace",
				EntityName: "entity_name",
				Group:      GroupPolicies,
				Channel:    "",
				Criterion:  CriterionCommands,
				Action:     ActionCreate,
			},
			want: "namespace/entity_name/" + string(GroupPolicies) + "/" +
				string(CriterionCommands) + "/" + string(ActionCreate),
		},
		"test_topic_string_empty": {
			topic: Topic{
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
	t.Run("TestTopicMarshalJSON", func(t *testing.T) {
		topic := Topic{
			Namespace:  "namespace",
			EntityName: "entity_name",
			Group:      GroupThings,
			Channel:    ChannelTwin,
			Criterion:  CriterionMessages,
			Action:     ActionSubscribe,
		}
		got, err := topic.MarshalJSON()
		grp := "\"namespace/entity_name/" + string(GroupThings) +
			"/" + string(ChannelTwin) + "/" + string(CriterionMessages) + "/" +
			string(ActionSubscribe) + "\""

		internal.AssertNil(t, err)
		internal.AssertEqual(t, grp, string(got))
	})
}

func TestTopicUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		data                  string
		wantErr               bool
		onlyForUnmarshalError bool
	}{
		"test_topic_unmarshal_JSON_group_things": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelTwin) + "/" +
				string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_group_policies": {
			data: "namespace/entity_name/" +
				string(GroupPolicies) + "/" +
				string(CriterionCommands) + "/" +
				string(ActionCreate),
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_insufficient_elements": {
			data:                  "///",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_group_things_missing_channel_for_things": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" /*+ string(ChannelTwin) + "/"*/ +
				string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_group_things_insufficient_elements_for_things": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelTwin) + "/" +
				string(CriterionMessages),
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_topic_must_empty": {
			data:                  "/////",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_no_data": {
			data:                  "",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_only_for_internal_error": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelTwin) + "/" +
				string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               true,
			onlyForUnmarshalError: true,
		},
		"test_topic_unmarshal_JSON_set_configuration": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelLive) + "/" +
				string(CriterionMessages) + "/" +
				"$set.configuration/name",
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_refresh_property": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelLive) + "/" +
				string(CriterionMessages) + "/" +
				"$refresh.name",
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		"test_topic_unmarshal_JSON_refresh_action": {
			data: "namespace/entity_name/" +
				string(GroupThings) + "/" +
				string(ChannelLive) + "/" +
				string(CriterionMessages) + "/" +
				"$refresh",
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			topic := Topic{}
			defer func() {
				if r := recover(); r != nil {
					if testCase.wantErr {
						t.Logf("Topic.UnmarshalJSON() %v", r)
					} else {
						t.Errorf("Topic.UnmarshalJSON() unexpected error %v", r)
					}
				}
			}()
			var err error
			if testCase.onlyForUnmarshalError {
				err = topic.UnmarshalJSON([]byte(nil))
			} else {
				err = topic.UnmarshalJSON([]byte("\"" + testCase.data + "\""))
			}
			if err != nil {
				if testCase.wantErr {
					t.Logf("Topic.UnmarshalJSON() error = %v", err)
					return
				}
				t.Errorf("Topic.UnmarshalJSON() unexpected error = %v", err)
				return
			}
			if topic.String() == "" {
				if testCase.wantErr {
					t.Logf("Topic.UnmarshalJSON() topic is empty")
					return
				}
				t.Errorf("Topic.UnmarshalJSON() unexpected empty topic")
				return
			}
			if topic.String() != testCase.data {
				t.Errorf("Topic.UnmarshalJSON() want = %v, got %v", topic.String(), testCase.data)
				return
			}
		})
	}
}

func TestTopicNamespace(t *testing.T) {
	var (
		testValidNamespace    = "test.namespace"
		testInvalidNamespace  = "test:namespace"
		testValidEntityName   = "test.name"
		testInvalidEntityName = "test§name"
	)

	tests := map[string]struct {
		data       string
		namespace  string
		entityName string
		wantErr    error
	}{
		"test_topic_unmarshal_JSON_valid_namespace_entity_name": {
			data:       testValidNamespace + "/" + testValidEntityName + "/things/twin/retrieve",
			namespace:  testValidNamespace,
			entityName: testValidEntityName,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_empty_namespace_valid_entity_name": {
			data:       TopicPlaceholder + "/" + testValidEntityName + "/things/twin/retrieve",
			namespace:  TopicPlaceholder,
			entityName: testValidEntityName,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_valid_namespace_empty_entity_name": {
			data:       testValidNamespace + "/" + TopicPlaceholder + "/things/twin/retrieve",
			namespace:  testValidNamespace,
			entityName: TopicPlaceholder,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_empty_namespace_empty_entity_name": {
			data:       TopicPlaceholder + "/" + TopicPlaceholder + "/things/twin/retrieve",
			namespace:  TopicPlaceholder,
			entityName: TopicPlaceholder,
			wantErr:    nil,
		},
		"test_topic_unmarshal_JSON_invalid_namespace": {
			data:       testInvalidNamespace + "/" + testValidEntityName + "/things/twin/retrieve",
			namespace:  "",
			entityName: "",
			wantErr:    errors.New("invalid topic namespaced ID, namespace: " + testInvalidNamespace + ", entity name: " + testValidEntityName),
		},
		"test_topic_unmarshal_JSON_invalid_entity_name": {
			data:       testValidNamespace + "/" + testInvalidEntityName + "/things/twin/retrieve",
			namespace:  "",
			entityName: "",
			wantErr:    errors.New("invalid topic namespaced ID, namespace: " + testValidNamespace + ", entity name: " + testInvalidEntityName),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			topic := Topic{}
			err := topic.UnmarshalJSON([]byte("\"" + testCase.data + "\""))
			internal.AssertError(t, testCase.wantErr, err)
			internal.AssertEqual(t, testCase.namespace, topic.Namespace)
			internal.AssertEqual(t, testCase.entityName, topic.EntityName)
		})
	}
}

func TestTopicWithNamespace(t *testing.T) {
	t.Run("TestTopicWithNamespace", func(t *testing.T) {
		arg := "namespace"
		topic := Topic{}
		got := topic.WithNamespace(arg)
		internal.AssertEqual(t, arg, got.Namespace)
	})
}

func TestTopicWithEntityName(t *testing.T) {
	t.Run("TestTopicWithEntityName", func(t *testing.T) {
		arg := "EntityName"
		topic := Topic{}
		got := topic.WithEntityName(arg)
		internal.AssertEqual(t, arg, got.EntityName)
	})
}

func TestTopicWithGroup(t *testing.T) {
	t.Run("TestTopicWithGroup", func(t *testing.T) {
		arg := GroupThings
		topic := Topic{}
		got := topic.WithGroup(arg)
		internal.AssertEqual(t, arg, got.Group)
	})
}

func TestTopicWithChannel(t *testing.T) {
	t.Run("TestTopicWithChannel", func(t *testing.T) {
		arg := ChannelTwin
		topic := Topic{}
		got := topic.WithChannel(arg)
		internal.AssertEqual(t, arg, got.Channel)
	})
}

func TestTopicWithCriterion(t *testing.T) {
	t.Run("TestTopicWithCriterion", func(t *testing.T) {
		arg := CriterionMessages
		topic := Topic{}
		got := topic.WithCriterion(arg)
		internal.AssertEqual(t, arg, got.Criterion)
	})
}

func TestTopicWithAction(t *testing.T) {
	t.Run("TestTopicWithAction", func(t *testing.T) {
		arg := ActionSubscribe
		topic := Topic{}
		got := topic.WithAction(arg)
		internal.AssertEqual(t, arg, got.Action)
	})
}
