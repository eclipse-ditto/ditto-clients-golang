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

package model

import (
	"reflect"
	"testing"
)

func TestThingWithID(t *testing.T) {
	arg := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	testThing := &Thing{}

	if got := testThing.WithID(arg); got.ID != arg {
		t.Errorf("Thing.WithID() = %v, want %v", got.ID, arg)
	}
}

func TestThingWithIDFrom(t *testing.T) {
	arg := "test.namespace:testId"

	testThing := &Thing{}

	if got := testThing.WithIDFrom(arg); !reflect.DeepEqual(got.ID, NewNamespacedIDFrom(arg)) {
		t.Errorf("Thing.WithIDFrom() = %v, want %v", got.ID, arg)
	}
}

func TestThingWithDefinition(t *testing.T) {
	arg := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "testId",
		Version:   "1.0.0",
	}

	testThing := &Thing{}

	if got := testThing.WithDefinition(arg); got.DefinitionID != arg {
		t.Errorf("Thing.WithDefinition() = %v, want %v", got.DefinitionID, arg)
	}
}

func TestThingWithDefinitionFrom(t *testing.T) {
	arg := "test.namespace:testId:1.0.0"

	testThing := &Thing{}

	if got := testThing.WithDefinitionFrom(arg); !reflect.DeepEqual(got.DefinitionID, NewDefinitionIDFrom(arg)) {
		t.Errorf("Thing.WithDefinitionFrom() = %v, want %v", got.DefinitionID, arg)
	}
}

func TestThingWithPolicyID(t *testing.T) {
	arg := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	testThing := &Thing{}

	if got := testThing.WithPolicyID(arg); got.PolicyID != arg {
		t.Errorf("Thing.WithPolicyID() = %v, want %v", got.PolicyID, arg)
	}
}

func TestThingPolicyIDFrom(t *testing.T) {
	arg := "test.namespace:testId"

	testThing := &Thing{}

	if got := testThing.WithPolicyIDFrom(arg); !reflect.DeepEqual(got.PolicyID, NewNamespacedIDFrom(arg)) {
		t.Errorf("Thing.WithPolicyIDFrom() = %v, want %v", got.PolicyID, arg)
	}
}

func TestThingWithAttributes(t *testing.T) {
	arg := map[string]interface{}{
		"test.key": "test.value",
	}

	testThing := &Thing{}

	if got := testThing.WithAttributes(arg); !reflect.DeepEqual(got.Attributes, arg) {
		t.Errorf("Thing.WithAttributes() = %v, want %v", got.Attributes, arg)
	}
}

func TestThingWithAttribute(t *testing.T) {
	tests := []struct {
		name      string
		arg1      string
		arg2      interface{}
		testThing Thing
		want      map[string]interface{}
	}{
		{
			name:      "TestThingWithAttributeWithoutExistingAttributes",
			arg1:      "test.key1",
			arg2:      1.0,
			testThing: Thing{},
			want: map[string]interface{}{
				"test.key1": 1.0,
			},
		},
		{
			name: "TestThingWithAttributeWithExistingAttributes",
			arg1: "test.key1",
			arg2: "test.value1",
			testThing: Thing{
				Attributes: map[string]interface{}{
					"test.key2": "test.value2",
				},
			},
			want: map[string]interface{}{
				"test.key1": "test.value1",
				"test.key2": "test.value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testThing.WithAttribute(tt.arg1, tt.arg2); !reflect.DeepEqual(got.Attributes, tt.want) {
				t.Errorf("Thing.WithAttribute() = %v, want %v", got.Attributes, tt.want)
			}
		})
	}
}

func TestThingWitFeatures(t *testing.T) {
	arg := map[string]*Feature{
		"TestFeature": {
			Definition: []*DefinitionID{{
				Namespace: "test.namespace",
				Name:      "testId",
				Version:   "1.0.0",
			}},
			Properties: map[string]interface{}{
				"property": "propertyValue",
			},
		},
	}

	testThing := &Thing{}

	if got := testThing.WithFeatures(arg); !reflect.DeepEqual(got.Features, arg) {
		t.Errorf("Thing.WithFeatures() = %v, want %v", got.Features, arg)
	}
}

func TestThingWithFeature(t *testing.T) {
	tests := []struct {
		name      string
		arg1      string
		arg2      *Feature
		testThing Thing
		want      map[string]*Feature
	}{
		{
			name: "TestThingWithFeatureWithoutExistingFeature",
			arg1: "TestFeature",
			arg2: &Feature{
				Definition: []*DefinitionID{{
					Namespace: "test.namespace",
					Name:      "testId",
					Version:   "1.0.0",
				}},
				Properties: map[string]interface{}{
					"property": "propertyValue",
				},
			},
			testThing: Thing{},
			want: map[string]*Feature{
				"TestFeature": {
					Definition: []*DefinitionID{{
						Namespace: "test.namespace",
						Name:      "testId",
						Version:   "1.0.0",
					}},
					Properties: map[string]interface{}{
						"property": "propertyValue",
					},
				},
			},
		},
		{
			name: "TestThingWithFeatureWithExistingFeature",
			arg1: "TestFeature1",
			arg2: &Feature{
				Properties: map[string]interface{}{
					"property": "propertyValue",
				},
			},
			testThing: Thing{
				Features: map[string]*Feature{
					"TestFeature2": {},
				},
			},
			want: map[string]*Feature{
				"TestFeature1": {
					Properties: map[string]interface{}{
						"property": "propertyValue",
					},
				},
				"TestFeature2": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testThing.WithFeature(tt.arg1, tt.arg2); !reflect.DeepEqual(got.Features, tt.want) {
				t.Errorf("Thing.WithFeature() = %v, want %v", got.Features, tt.want)
			}
		})
	}
}
