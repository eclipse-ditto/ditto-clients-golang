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

import (
	"reflect"
	"testing"
)

func TestTopicString(t *testing.T) {
	tests := []struct {
		name  string
		topic Topic
		want  string
	}{
		{
			name: "TestTopicString GroupThings",
			topic: Topic{
				Namespace: "anamespace",
				EntityID:  "aentityid",
				Group:     GroupThings,
				Channel:   ChannelTwin,
				Criterion: CriterionMessages,
				Action:    ActionSubscribe,
			},
			want: "anamespace/aentityid/" + string(GroupThings) + "/" + string(ChannelTwin) + "/" +
				string(CriterionMessages) + "/" + string(ActionSubscribe),
		},
		{
			name: "TestTopicString GroupPolicies",
			topic: Topic{
				Namespace: "anamespace",
				EntityID:  "aentityid",
				Group:     GroupPolicies,
				Channel:   "",
				Criterion: CriterionCommands,
				Action:    ActionCreate,
			},
			want: "anamespace/aentityid/" + string(GroupPolicies) + "/" +
				string(CriterionCommands) + "/" + string(ActionCreate),
		},
		{
			name: "TestTopicString empty",
			topic: Topic{
				Namespace: "anamespace",
				EntityID:  "aentityid",
				Group:     "",
				Channel:   "",
				Criterion: CriterionCommands,
				Action:    ActionCreate,
			},
			want: "",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.topic.String()
			if got != tt.want {
				t.Errorf("Topic.String() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestTopicMarshalJSON(t *testing.T) {
	t.Run("TestTopicMarshalJSON", func(t *testing.T) {
		topic := Topic{
			Namespace: "namespace",
			EntityID:  "entityid",
			Group:     GroupThings,
			Channel:   ChannelTwin,
			Criterion: CriterionMessages,
			Action:    ActionSubscribe,
		}
		got, err := topic.MarshalJSON()
		grp := "\"namespace/entityid/" + string(GroupThings) +
			"/" + string(ChannelTwin) + "/" + string(CriterionMessages) + "/" +
			string(ActionSubscribe) + "\""
		if err != nil {
			t.Errorf("Topic.MarshalJSON() error = %v", err)
		}
		if !reflect.DeepEqual(grp, string(got)) {
			t.Errorf("Topic.MarshalJSON() want = %v , got %v", grp, string(got))
		}
	})
}

func TestTopicUnmarshalJSON(t *testing.T) {

	tests := []struct {
		name                  string
		data                  string
		wantErr               bool
		onlyForUnmarshalError bool
	}{
		{
			name: "TestTopicUnmarshalJSON GroupThings",
			data: "anamespace/aentityid/" +
				string(GroupThings) + "/" + string(ChannelTwin) + "/" + string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		{
			name: "TestTopicUnmarshalJSON GroupPolicies",
			data: "anamespace/aentityid/" +
				string(GroupPolicies) + "/" + string(CriterionCommands) + "/" + string(ActionCreate),
			wantErr:               false,
			onlyForUnmarshalError: false,
		},
		{
			name:                  "TestTopicUnmarshalJSON for insufficient elements",
			data:                  "///",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		{
			name: "TestTopicUnmarshalJSON GroupThings for missing channel for things",
			data: "anamespace/aentityid/" +
				string(GroupThings) + "/" /*+ string(ChannelTwin) + "/"*/ + string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		{
			name: "TestTopicUnmarshalJSON GroupThings insufficient elements for things",
			data: "anamespace/aentityid/" +
				string(GroupThings) + "/" + string(ChannelTwin) + "/" + string(CriterionMessages),
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		{
			name:                  "TestTopicUnmarshalJSON topic must be empty",
			data:                  "/////",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		{
			name:                  "TestTopicUnmarshalJSON no data",
			data:                  "",
			wantErr:               true,
			onlyForUnmarshalError: false,
		},
		{
			name: "TestTopicUnmarshalJSON only for internal error",
			data: "anamespace/aentityid/" +
				string(GroupThings) + "/" + string(ChannelTwin) + "/" + string(CriterionMessages) + "/" +
				string(ActionSubscribe),
			wantErr:               true,
			onlyForUnmarshalError: true,
		},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topic := Topic{}
			defer func() {
				if r := recover(); r != nil {
					if tt.wantErr {
						t.Logf("Topic.UnmarshalJSON() %v", r)
					} else {
						t.Errorf("Topic.UnmarshalJSON() unexpected error %v", r)
					}
				}
			}()
			var err error
			if tt.onlyForUnmarshalError {
				err = topic.UnmarshalJSON([]byte(nil))
			} else {
				err = topic.UnmarshalJSON([]byte("\"" + tt.data + "\""))
			}
			if err != nil {
				if tt.wantErr {
					t.Logf("Topic.UnmarshalJSON() error = %v", err)
					return
				} else {
					t.Errorf("Topic.UnmarshalJSON() unexpected error = %v", err)
					return
				}
			}
			if topic.String() == "" {
				if tt.wantErr {
					t.Logf("Topic.UnmarshalJSON() topic is empty")
					return
				} else {
					t.Errorf("Topic.UnmarshalJSON() unexpected empty topic")
					return
				}
			}
			if topic.String() != tt.data {
				t.Errorf("Topic.UnmarshalJSON() want = %v, got %v", topic.String(), tt.data)
				return
			}
		})
	}
}

func TestTopicWithNamespace(t *testing.T) {
	t.Run("TestTopicWithNamespace", func(t *testing.T) {
		arg := "anamespace"
		topic := Topic{}
		if got := topic.WithNamespace(arg); !reflect.DeepEqual(got.Namespace, arg) {
			t.Errorf("Topic.WithNamespace() = %v, want %v", got.Namespace, arg)
		}
	})
}

func TestTopicWithEntityID(t *testing.T) {
	t.Run("TestTopicWithEntityID", func(t *testing.T) {
		arg := "EntityID"
		topic := Topic{}
		if got := topic.WithEntityID(arg); !reflect.DeepEqual(got.EntityID, arg) {
			t.Errorf("Topic.WithEntityID() = %v, want %v", got.EntityID, arg)
		}
	})
}

func TestTopicWithGroup(t *testing.T) {
	t.Run("TestTopicWithGroup", func(t *testing.T) {
		arg := GroupThings
		topic := Topic{}
		if got := topic.WithGroup(arg); !reflect.DeepEqual(got.Group, arg) {
			t.Errorf("Topic.WithGroup() = %v, want %v", got.Group, arg)
		}
	})
}

func TestTopicWithChannel(t *testing.T) {
	t.Run("TestTopicWithChannel", func(t *testing.T) {
		arg := ChannelTwin
		topic := Topic{}
		if got := topic.WithChannel(arg); !reflect.DeepEqual(got.Channel, arg) {
			t.Errorf("Topic.WithChannel() = %v, want %v", got.Channel, arg)
		}
	})
}

func TestTopicWithCriterion(t *testing.T) {
	t.Run("TestTopicWithCriterion", func(t *testing.T) {
		arg := CriterionMessages
		topic := Topic{}
		if got := topic.WithCriterion(arg); !reflect.DeepEqual(got.Criterion, arg) {
			t.Errorf("Topic.WithCriterion() = %v, want %v", got.Criterion, arg)
		}
	})
}

func TestTopicWithAction(t *testing.T) {
	t.Run("TestTopicWithAction", func(t *testing.T) {
		arg := ActionSubscribe
		topic := Topic{}
		if got := topic.WithAction(arg); !reflect.DeepEqual(got.Action, arg) {
			t.Errorf("Topic.WithAction() = %v, want %v", got.Action, arg)
		}
	})
}
