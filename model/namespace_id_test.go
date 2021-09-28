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
	"math/rand"
	"reflect"
	"testing"
)

func TestNamespaceIDNewNamespacedID(t *testing.T) {
	tests := map[string]struct {
		args []string
		want *NamespacedID
	}{
		"test_new_namespaced_id_valid": {
			args: []string{"test.namespace", "testId"},
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "testId",
			},
		},
		"test_new_namespaced_id_invalid": {
			args: []string{"test.namespace", "test/Id"},
			want: nil,
		},
		"test_new_namespaced_id_namespace_with_colon": {
			args: []string{"test:namespace", "testId"},
			want: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := NewNamespacedID(testCase.args[0], testCase.args[1]); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("NewNamespacedID() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestNamespaceIDNewNamespacedIDFrom(t *testing.T) {
	tests := map[string]struct {
		arg  string
		want *NamespacedID
	}{
		"test_new_namespaced_id_from_valid": {
			arg:  "test.namespace_42:testId",
			want: &NamespacedID{
				Namespace: "test.namespace_42",
				Name:      "testId",
			},
		},
		"test_new_namespaced_id_from_with_namespace_with_pascal_case": {
			arg:  "Test.Namespace_42:testId",
			want: &NamespacedID{
				Namespace: "Test.Namespace_42",
				Name:      "testId",
			},
		},
		"test_new_namespaced_id_from_without_namespace": {
			arg:  ":testId",
			want: &NamespacedID{
				Namespace: "",
				Name:      "testId",
			},
		},
		"test_new_namespaced_id_from_with_double_colon": {
			arg:  "test.namespace:testId:testId",
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "testId:testId",
			},
		},
		"test_new_namespaced_id_from_without_name": {
			arg:  "test.namsepsaced",
			want: nil,
		},
		"test_new_namespaced_id_from_with_name_with_slash": {
			arg:  "test.namespace:testId/testId",
			want: nil,
		},
		"test_new_namespaced_id_from_with_name_with_control_character": {
			arg:  "test.namespace:testId\ntestId",
			want: nil,
		},
		"test_new_namespaced_id_from_empty": {
			arg:  "",
			want: nil,
		},
		"test_new_namespaced_id_from_invalid_length": {
			arg: func() string {
				letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				generated := make([]byte, 257)
				for i := range generated {
					generated[i] = letters[rand.Intn(len(letters))]
				}
				generated[rand.Intn(len(generated)-2)] = ':'
				return string(generated)
			}(),
			want: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := NewNamespacedIDFrom(testCase.arg); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("NewNamespacedIDFrom() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestNamespaceIDString(t *testing.T) {
	testNamespaceID := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	want := "test.namespace:testId"

	if got := testNamespaceID.String(); got != want {
		t.Errorf("NamespaceID.String() = %v, want %v", got, want)
	}

	if got := testNamespaceID.String(); reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("NamespaceID.String() = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(want))
	}
}

func TestNamespaceIDMarshalJSON(t *testing.T) {
	testNamespace := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	want := []byte("\"test.namespace:testId\"")

	if got, _ := testNamespace.MarshalJSON(); !reflect.DeepEqual(got, want) {
		t.Errorf("NamespacedID.MarshaJSON() = %v, want %v", got, want)
	}
}

func TestNamespaceIDUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		arg     []byte
		want    *NamespacedID
		wantErr bool
	}{
		"test_namespaced_id_unmarshal_json_valid": {
			arg:  []byte("\"test.namespace:testId\""),
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "testId",
			},
			wantErr: false,
		},
		"test_namespaced_id_unmarshal_json_invalid": {
			arg:     []byte("\"test:namespace\\testId\""),
			wantErr: true,
		},
		"test_namespaced_id_unmarshal_json_empty": {
			arg:     []byte(""),
			wantErr: true,
		},
		"test_namespaced_id_unmarshal_json_invalid_argument": {
			arg:     []byte("test.namespace:testId"),
			wantErr: true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := &NamespacedID{}
			err := got.UnmarshalJSON(testCase.arg)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("NamespacedID.UnmarshalJSON() error must not be nil")
				}
			} else {
				if !reflect.DeepEqual(got, testCase.want) {
					t.Errorf("NamespacedID.UnmarshalJSON() = %v, want %v", got, testCase.want)
				}
			}
		})
	}
}

func TestNamespaceIDWithNamespace(t *testing.T) {
	testNamespaceID := &NamespacedID{
		Name: "testId",
	}

	arg := "test.namespace"

	want := &NamespacedID{
		Namespace: arg,
		Name:      "testId",
	}

	if got := testNamespaceID.WithNamespace(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("NamespacedID.WithNamespace() = %v, want %v", got, want)
	}
}

func TestNamespaceIDWithName(t *testing.T) {
	testNamespace := &NamespacedID{
		Namespace: "test.namespace",
	}

	arg := "testId"

	want := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "testId",
	}

	if got := testNamespace.WithName(arg); !reflect.DeepEqual(got, want) {
		t.Errorf("NamespacedID.WithName() = %v, want %v", got, want)
	}
}
