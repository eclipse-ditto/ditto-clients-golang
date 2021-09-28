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

func TestFeatureWithDefinitionFrom(t *testing.T) {
	tests := []struct {
		name        string
		arg1        string
		arg2        string
		testFeature Feature
		want        []*DefinitionID
	}{
		{
			name:        "TestFeatureWithDefinitionFromWithoutExistingDefinitionID",
			arg1:        "test.namespace:testId:1.0.0",
			arg2:        "test.namespace:testId:1.0.1",
			testFeature: Feature{},
			want: []*DefinitionID{
				NewDefinitionIDFrom("test.namespace:testId:1.0.0"),
				NewDefinitionIDFrom("test.namespace:testId:1.0.1"),
			},
		},
		{
			name: "TestFeatureWithDefinitionFromWithExistingDefinitionID",
			arg1: "test.namespace:testId:1.0.0",
			arg2: "test.namespace:testId:1.0.1",
			testFeature: Feature{
				Definition: []*DefinitionID{
					NewDefinitionIDFrom("test.namespace:testId:0.0.0"),
				},
			},
			want: []*DefinitionID{
				NewDefinitionIDFrom("test.namespace:testId:1.0.0"),
				NewDefinitionIDFrom("test.namespace:testId:1.0.1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testFeature.WithDefinitionFrom(tt.arg1, tt.arg2); !reflect.DeepEqual(got.Definition, tt.want) {
				t.Errorf("Feature.WithDefinitionFrom() = %v, want %v", got.Definition, tt.want)
			}
		})
	}
}

func TestFeatureWithDefinition(t *testing.T) {
	arg1 := NewDefinitionIDFrom("test.namespace:testId:1.0.0")
	arg2 := NewDefinitionIDFrom("test.namespace:testId:1.0.0")

	testDefinitions := []*DefinitionID{arg1, arg2}

	testFeature := &Feature{}

	if got := testFeature.WithDefinition(arg1, arg2); !reflect.DeepEqual(got.Definition, testDefinitions) {
		t.Errorf("Feature.WithDefinition() = %v, want %v", got.Definition, testDefinitions)
	}
}

func TestFeatureWithProperties(t *testing.T) {
	arg := map[string]interface{}{
		"test.key1": "test.value1",
		"test.key2": 123,
	}

	testFeature := &Feature{}

	if got := testFeature.WithProperties(arg); !reflect.DeepEqual(got.Properties, arg) {
		t.Errorf("Feature.WithProperties() = %v, want %v", got.Properties, arg)
	}
}

func TestFeatureWithProperty(t *testing.T) {
	tests := []struct {
		name        string
		arg1        string
		arg2        string
		testFeature Feature
		want        map[string]interface{}
	}{
		{
			name:        "TestFeatureWithPropertyWithoutExistingProperty",
			arg1:        "test.key",
			arg2:        "test.value",
			testFeature: Feature{},
			want: map[string]interface{}{
				"test.key": "test.value",
			},
		},
		{
			name: "TestFeatureWithPropertyWithExistingProperty",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testFeature.WithProperty(tt.arg1, tt.arg2); !reflect.DeepEqual(got.Properties, tt.want) {
				t.Errorf("Feature.WithProperty() = %v, want %v", got.Properties, tt.want)
			}
		})
	}
}

func TestFeatureWithDesiredProperties(t *testing.T) {
	arg := map[string]interface{}{
		"test.key1": "test.value1",
		"test.key2": 123,
	}

	testFeature := &Feature{}

	if got := testFeature.WithDesiredProperties(arg); !reflect.DeepEqual(got.DesiredProperties, arg) {
		t.Errorf("Feature.WithDesiredProperties() = %v, want %v", got.DesiredProperties, arg)
	}
}

func TestFeatureWithDesiredProperty(t *testing.T) {
	tests := []struct {
		name        string
		arg1        string
		arg2        string
		testFeature Feature
		want        map[string]interface{}
	}{
		{
			name:        "TestFeatureWithDesiredPropertyWithoutExistingDesiredProperty",
			arg1:        "test.key",
			arg2:        "test.value",
			testFeature: Feature{},
			want: map[string]interface{}{
				"test.key": "test.value",
			},
		},
		{
			name: "TestFeatureWithDesiredPropertyExistingPropertyDesiredProperty",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testFeature.WithDesiredProperty(tt.arg1, tt.arg2); !reflect.DeepEqual(got.DesiredProperties, tt.want) {
				t.Errorf("Feature.WithDesiredProperty() = %v, want %v", got.DesiredProperties, tt.want)
			}
		})
	}
}
