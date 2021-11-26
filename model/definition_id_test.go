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
	"errors"
	"reflect"
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestDefinitionIDNewDefinitionIDFrom(t *testing.T) {
	tests := map[string]struct {
		arg  string
		want *DefinitionID
	}{
		"test_new_definition_id_from_valid": {
			arg: "test.namespace_42:testId:1.0.0-qualifier",
			want: &DefinitionID{
				Namespace: "test.namespace_42",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
		"test_new_definition_id_from_without_namespace": {
			arg:  ":testId:1.0.0",
			want: nil,
		},
		"test_new_definition_id_from_without_name": {
			arg:  "testId:1.0.0",
			want: nil,
		},
		"test_new_definition_id_from_without_version": {
			arg:  "test.namespace:1.0.0",
			want: nil,
		},
		"test_new_definition_id_from_with_invalid_colon": {
			arg:  "test.namespace:testId:1.0.0:",
			want: nil,
		},
		"test_new_definition_id_from_with_invalid_character": {
			arg:  "test.name*space:testId:1.0.0",
			want: nil,
		},
		"test_new_definition_id_from_empty": {
			arg:  "",
			want: nil,
		},
		"test_new_definition_id_from_namespace_with_dash": {
			arg: "test-namespace:testId:1.0.0-qualifier",
			want: &DefinitionID{
				Namespace: "test-namespace",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
		"test_new_definition_id_from_namespace_with_dash_dot": {
			arg: "test.namespace-id:testId:1.0.0-qualifier",
			want: &DefinitionID{
				Namespace: "test.namespace-id",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewDefinitionIDFrom(testCase.arg)
			internal.AssertEqual(t, got, testCase.want)
		})
	}
}

func TestDefinitionIDNewDefinitionID(t *testing.T) {
	tests := map[string]struct {
		args []string
		want *DefinitionID
	}{
		"test_new_definition_id_valid": {
			args: []string{"test.namespace_42", "testId", "1.0.0-qualifier"},
			want: &DefinitionID{
				Namespace: "test.namespace_42",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
		"test_new_definition_id_invalid": {
			args: []string{"test/namespace", "testId", "1.0.0"},
			want: nil,
		},
		"test_new_definition_id_namespace_dash": {
			args: []string{"test-namespace", "testId", "1.0.0-qualifier"},
			want: &DefinitionID{
				Namespace: "test-namespace",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
		"test_new_definition_id_namespace_dash_dot": {
			args: []string{"test.namespace-id", "testId", "1.0.0-qualifier"},
			want: &DefinitionID{
				Namespace: "test.namespace-id",
				Name:      "testId",
				Version:   "1.0.0-qualifier",
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewDefinitionID(testCase.args[0], testCase.args[1], testCase.args[2])
			internal.AssertEqual(t, got, testCase.want)
		})
	}
}

func TestDefinitionIDString(t *testing.T) {
	testDefinitionID := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "testId",
		Version:   "1.0.0",
	}

	want := "test.namespace:testId:1.0.0"

	got := testDefinitionID.String()
	internal.AssertEqual(t, got, want)

	if got := testDefinitionID.String(); reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("DefinitionID.String() = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(want))
	}
}

func TestDefinitionIDMarshalJSON(t *testing.T) {
	testDefinitionID := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "testId",
		Version:   "1.0.0",
	}

	want := []byte("\"test.namespace:testId:1.0.0\"")

	got, _ := testDefinitionID.MarshalJSON()
	internal.AssertEqual(t, got, want)
}

func TestDefinitionIDUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		arg     []byte
		want    *DefinitionID
		wantErr error
	}{
		"test_definition_id_unmarshal_json_valid": {
			arg: []byte("\"test.namespace:testId:1.0.0\""),
			want: &DefinitionID{
				Namespace: "test.namespace",
				Name:      "testId",
				Version:   "1.0.0",
			},
			wantErr: nil,
		},
		"test_definition_id_unmarshal_json_invalid_namespace": {
			arg:     []byte("\"test:namespace:testId:1.0.0\""),
			wantErr: errors.New("invalid DefinitionID: test:namespace:testId:1.0.0"),
		},
		"test_definition_id_unmarshal_json_invalid_name": {
			arg:     []byte("\"test.namespace:1.0.0\""),
			wantErr: errors.New("invalid DefinitionID: test.namespace:1.0.0"),
		},
		"test_definition_id_unmarshal_json_invalid_version": {
			arg:     []byte("\"test.namespace:testId\""),
			wantErr: errors.New("invalid DefinitionID: test.namespace:testId"),
		},
		"test_definition_id_unmarshal_json_empty": {
			arg:     []byte(""),
			wantErr: errors.New("unexpected end of JSON input"),
		},
		"test_definition_id_unmarshal_json_invalid_argument": {
			arg:     []byte("test.namespace:testId:1.0.0"),
			wantErr: errors.New("invalid character 'e' in literal true (expecting 'r')"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := &DefinitionID{}
			err := got.UnmarshalJSON(testCase.arg)
			if testCase.wantErr != nil {
				internal.AssertError(t, err, testCase.wantErr)
			} else {
				internal.AssertEqual(t, got, testCase.want)
			}
		})
	}
}

func TestDefinitionIDWithNamespace(t *testing.T) {
	testDefinitionID := &DefinitionID{
		Name:    "testId",
		Version: "1.0.0",
	}

	arg := "test.namespace"

	want := &DefinitionID{
		Namespace: arg,
		Name:      "testId",
		Version:   "1.0.0",
	}

	got := testDefinitionID.WithNamespace(arg)
	internal.AssertEqual(t, got, want)
}

func TestDefinitionIDWithName(t *testing.T) {
	testDefinitionID := &DefinitionID{
		Namespace: "test.namespace",
		Version:   "1.0.0",
	}

	arg := "testId"

	want := &DefinitionID{
		Namespace: "test.namespace",
		Name:      arg,
		Version:   "1.0.0",
	}

	got := testDefinitionID.WithName(arg)
	internal.AssertEqual(t, got, want)
}

func TestDefinitionIDWithVersion(t *testing.T) {
	testDefinitionID := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	arg := "1.0.0"

	want := &DefinitionID{
		Namespace: "test.namespace",
		Name:      "testId",
		Version:   arg,
	}

	got := testDefinitionID.WithVersion(arg)
	internal.AssertEqual(t, got, want)
}
