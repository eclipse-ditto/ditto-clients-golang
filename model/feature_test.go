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

func TestFeatureWithDefinitionFrom(t *testing.T) {
	tests := map[string]struct {
		arg1        string
		arg2        string
		testFeature Feature
		want        []*DefinitionID
	}{
		"test_feature_with_definition_from_without_existing_definition_ID": {
			arg1:        "test.namespace:test-name:1.0.0",
			arg2:        "test.namespace:test-name:1.0.1",
			testFeature: Feature{},
			want: []*DefinitionID{
				NewDefinitionIDFrom("test.namespace:test-name:1.0.0"),
				NewDefinitionIDFrom("test.namespace:test-name:1.0.1"),
			},
		},
		"test_feature_with_definition_from_with_existing_definition_ID": {
			arg1: "test.namespace:test-name:1.0.0",
			arg2: "test.namespace:test-name:1.0.1",
			testFeature: Feature{
				Definition: []*DefinitionID{
					NewDefinitionIDFrom("test.namespace:test-name:0.0.0"),
				},
			},
			want: []*DefinitionID{
				NewDefinitionIDFrom("test.namespace:test-name:1.0.0"),
				NewDefinitionIDFrom("test.namespace:test-name:1.0.1"),
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testFeature.WithDefinitionFrom(testCase.arg1, testCase.arg2)
			internal.AssertEqual(t, testCase.want, got.Definition)
		})
	}
}

func TestFeatureWithDefinition(t *testing.T) {
	arg1 := NewDefinitionIDFrom("test.namespace:test-name:1.0.0")
	arg2 := NewDefinitionIDFrom("test.namespace:test-name:1.0.0")

	testDefinitions := []*DefinitionID{arg1, arg2}

	testFeature := &Feature{}

	got := testFeature.WithDefinition(arg1, arg2)
	internal.AssertEqual(t, testDefinitions, got.Definition)
}

func TestFeatureWithProperties(t *testing.T) {
	arg := map[string]interface{}{
		"test.key1": "test.value1",
		"test.key2": 123,
	}

	testFeature := &Feature{}
	got := testFeature.WithProperties(arg)
	internal.AssertEqual(t, arg, got.Properties)
}

func TestFeatureWithProperty(t *testing.T) {
	tests := map[string]struct {
		arg1        string
		arg2        string
		testFeature Feature
		want        map[string]interface{}
	}{
		"test_feature_with_property_without_existing_property": {
			arg1:        "test.key",
			arg2:        "test.value",
			testFeature: Feature{},
			want: map[string]interface{}{
				"test.key": "test.value",
			},
		},
		"test_feature_with_property_wWith_existing_property": {
			arg1: "test.key1",
			arg2: "test.value1",
			testFeature: Feature{
				Properties: map[string]interface{}{
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
			got := testCase.testFeature.WithProperty(testCase.arg1, testCase.arg2)
			internal.AssertEqual(t, testCase.want, got.Properties)
		})
	}
}

func TestFeatureWithDesiredProperties(t *testing.T) {
	arg := map[string]interface{}{
		"test.key1": "test.value1",
		"test.key2": 123,
	}

	testFeature := &Feature{}

	got := testFeature.WithDesiredProperties(arg)
	internal.AssertEqual(t, arg, got.DesiredProperties)
}

func TestFeatureWithDesiredProperty(t *testing.T) {
	tests := map[string]struct {
		arg1        string
		arg2        string
		testFeature Feature
		want        map[string]interface{}
	}{
		"test_feature_with_desired_property_without_existing_desired_property": {
			arg1:        "test.key",
			arg2:        "test.value",
			testFeature: Feature{},
			want: map[string]interface{}{
				"test.key": "test.value",
			},
		},
		"test_feature_with_desired_property_existing_property_desired_property": {
			arg1: "test.key1",
			arg2: "test.value1",
			testFeature: Feature{
				DesiredProperties: map[string]interface{}{
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
			got := testCase.testFeature.WithDesiredProperty(testCase.arg1, testCase.arg2)
			internal.AssertEqual(t, testCase.want, got.DesiredProperties)
		})
	}
}
