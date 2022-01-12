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
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestThingWithID(t *testing.T) {
	arg := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "test-name",
	}

	testThing := &Thing{}

	got := testThing.WithID(arg)
	internal.AssertEqual(t, arg, got.ID)
}

func TestThingWithIDFrom(t *testing.T) {
	arg := "test.namespace:test-name"

	testThing := &Thing{}

	got := testThing.WithIDFrom(arg)
	internal.AssertEqual(t, NewNamespacedIDFrom(arg), got.ID)
}

func TestThingWithDefinition(t *testing.T) {
	arg := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "test-name",
		Version:   "1.0.0",
	}

	testThing := &Thing{}

	got := testThing.WithDefinition(arg)
	internal.AssertEqual(t, arg, got.DefinitionID)
}

func TestThingWithDefinitionFrom(t *testing.T) {
	arg := "test.namespace:test-name:1.0.0"

	testThing := &Thing{}

	got := testThing.WithDefinitionFrom(arg)
	internal.AssertEqual(t, NewDefinitionIDFrom(arg), got.DefinitionID)
}

func TestThingWithPolicyID(t *testing.T) {
	arg := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "test-name",
	}

	testThing := &Thing{}

	got := testThing.WithPolicyID(arg)
	internal.AssertEqual(t, arg, got.PolicyID)
}

func TestThingPolicyIDFrom(t *testing.T) {
	arg := "test.namespace:test-name"

	testThing := &Thing{}

	got := testThing.WithPolicyIDFrom(arg)
	internal.AssertEqual(t, NewNamespacedIDFrom(arg), got.PolicyID)
}

func TestThingWithAttributes(t *testing.T) {
	arg := map[string]interface{}{
		"test.key": "test.value",
	}

	testThing := &Thing{}

	got := testThing.WithAttributes(arg)
	internal.AssertEqual(t, arg, got.Attributes)
}

func TestThingWithAttribute(t *testing.T) {
	tests := map[string]struct {
		arg1      string
		arg2      interface{}
		testThing Thing
		want      map[string]interface{}
	}{
		"test_thing_with_attribute_without_existing_attributes": {
			arg1:      "test.key1",
			arg2:      1.0,
			testThing: Thing{},
			want: map[string]interface{}{
				"test.key1": 1.0,
			},
		},
		"test_thing_with_attribute_with_existing_attributes": {
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

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testThing.WithAttribute(testCase.arg1, testCase.arg2)
			internal.AssertEqual(t, testCase.want, got.Attributes)
		})
	}
}

func TestThingWitFeatures(t *testing.T) {
	arg := map[string]*Feature{
		"TestFeature": {
			Definition: []*DefinitionID{{
				Namespace: "test.namespace",
				Name:      "test-name",
				Version:   "1.0.0",
			}},
			Properties: map[string]interface{}{
				"property": "propertyValue",
			},
		},
	}

	testThing := &Thing{}

	got := testThing.WithFeatures(arg)
	internal.AssertEqual(t, arg, got.Features)
}

func TestThingWithFeature(t *testing.T) {
	tests := map[string]struct {
		arg1      string
		arg2      *Feature
		testThing Thing
		want      map[string]*Feature
	}{
		"test_thing_with_feature_without_existing_feature": {
			arg1: "TestFeature",
			arg2: &Feature{
				Definition: []*DefinitionID{{
					Namespace: "test.namespace",
					Name:      "test-name",
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
						Name:      "test-name",
						Version:   "1.0.0",
					}},
					Properties: map[string]interface{}{
						"property": "propertyValue",
					},
				},
			},
		},
		"test_thing_with_feature_with_existing_feature": {
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

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testThing.WithFeature(testCase.arg1, testCase.arg2)
			internal.AssertEqual(t, testCase.want, got.Features)
		})
	}
}
