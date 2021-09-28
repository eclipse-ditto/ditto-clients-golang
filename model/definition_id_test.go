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
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := NewDefinitionIDFrom(testCase.arg); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("NewDefinitionIDFrom() = %v, want %v", got, testCase.want)
			}
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
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := NewDefinitionID(testCase.args[0], testCase.args[1], testCase.args[2]); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("NewDefinitionID() = %v, want %v", got, testCase.want)
			}
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

	if got := testDefinitionID.String(); got != want {
		t.Errorf("DefinitionID.String() = %v, want %v", got, want)
	}

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

	if got, _ := testDefinitionID.MarshalJSON(); !reflect.DeepEqual(got, want) {
		t.Errorf("DefinitionID.MarshalJSON() = %v, want %v", got, want)
	}
}

func TestDefinitionIDUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		arg     []byte
		want    *DefinitionID
		wantErr bool
	}{
		"test_definition_id_unmarshal_json_valid": {
			arg: []byte("\"test.namespace:testId:1.0.0\""),
			want: &DefinitionID{
				Namespace: "test.namespace",
				Name:      "testId",
				Version:   "1.0.0",
			},
			wantErr: false,
		},
		"test_definition_id_unmarshal_json_invalid_namespace": {
			arg:     []byte("\"test:namespace:testId:1.0.0\""),
			wantErr: true,
		},
		"test_definition_id_unmarshal_json_invalid_name": {
			arg:     []byte("\"test.namespace:1.0.0\""),
			wantErr: true,
		},
		"test_definition_id_unmarshal_json_invalid_version": {
			arg:     []byte("\"test.namespace:testId\""),
			wantErr: true,
		},
		"test_definition_id_unmarshal_json_empty": {
			arg:     []byte(""),
			wantErr: true,
		},
		"test_definition_id_unmarshal_json_invalid_argument": {
			arg:     []byte("test.namespace:testId:1.0.0"),
			wantErr: true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := &DefinitionID{}
			err := got.UnmarshalJSON(testCase.arg)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("DefinitionID.UnmarshalJSON() error must not be nil")
				}
			} else {
				if !reflect.DeepEqual(got, testCase.want) {
					t.Errorf("DefinitionID.UnmarshalJSON() = %v, want %v", got, testCase.want)
				}
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

	if got := testDefinitionID.WithNamespace(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("DefinitionID.WithNamespace() = %v, want %v", got, want)
	}
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

	if got := testDefinitionID.WithName(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("DefinitionID.WithName() = %v, want %v", got, want)
	}
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

	if got := testDefinitionID.WithVersion(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("DefinitionID.WithVerson() = %v, want %v", got, want)
	}
}
