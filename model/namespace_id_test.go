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
	"math/rand"
	"reflect"
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestNamespaceIDNewNamespacedID(t *testing.T) {
	tests := map[string]struct {
		args []string
		want *NamespacedID
	}{
		"test_new_namespaced_ID_valid": {
			args: []string{"test.namespace", "test-name"},
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_invalid_name_slash": {
			args: []string{"test.namespace", "test/name"},
			want: nil,
		},
		"test_new_namespaced_ID_invalid_name_control_character": {
			args: []string{"test.namespace", "test§name"},
			want: nil,
		},
		"test_new_namespaced_ID_invalid_empty_name": {
			args: []string{"test.namespace", ""},
			want: nil,
		},
		"test_new_namespaced_ID_namespace_with_colon": {
			args: []string{"test:namespace", "test-name"},
			want: nil,
		},
		"test_new_namespaced_ID_namespace_with_dash": {
			args: []string{"test-namespace", "test-name"},
			want: &NamespacedID{
				Namespace: "test-namespace",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_namespace_with_multiple_dash": {
			args: []string{"test-namespace-multiple-dash", "test-name"},
			want: &NamespacedID{
				Namespace: "test-namespace-multiple-dash",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_namespace_with_dash_dot": {
			args: []string{"test.namespace-dash-dot", "test-name"},
			want: &NamespacedID{
				Namespace: "test.namespace-dash-dot",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_invalid_namespace_control_character": {
			args: []string{"test.namespace§", "test-name"},
			want: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewNamespacedID(testCase.args[0], testCase.args[1])
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestNamespaceIDNewNamespacedIDFrom(t *testing.T) {
	tests := map[string]struct {
		arg  string
		want *NamespacedID
	}{
		"test_new_namespaced_ID_from_valid": {
			arg: "test.namespace_42:test-name",
			want: &NamespacedID{
				Namespace: "test.namespace_42",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_from_with_namespace_with_pascal_case": {
			arg: "Test.Namespace_42:test-name",
			want: &NamespacedID{
				Namespace: "Test.Namespace_42",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_from_without_namespace": {
			arg: ":test-name",
			want: &NamespacedID{
				Namespace: "",
				Name:      "test-name",
			},
		},
		"test_new_namespaced_ID_from_with_double_colon": {
			arg: "test.namespace:test-name:test-name",
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "test-name:test-name",
			},
		},
		"test_new_namespaced_ID_from_without_name": {
			arg:  "test.namsepsace",
			want: nil,
		},
		"test_new_namespaced_ID_from_with_name_with_slash": {
			arg:  "test.namespace:test-name/test-name",
			want: nil,
		},
		"test_new_namespaced_ID_from_with_name_with_control_character": {
			arg:  "test.namespace:test-name§test-name",
			want: nil,
		},
		"test_new_namespaced_ID_from_empty": {
			arg:  "",
			want: nil,
		},
		"test_new_namespaced_ID_from_invalid_length": {
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
		"test_new_namsepaced_ID_from_with_namespace_with_single_dash": {
			arg: "test-namespace:test-name",
			want: &NamespacedID{
				Namespace: "test-namespace",
				Name:      "test-name",
			},
		},
		"test_new_namsepaced_ID_from_with_namespace_with_multiple_dash": {
			arg: "test-namespace-multiple-dash:test-name",
			want: &NamespacedID{
				Namespace: "test-namespace-multiple-dash",
				Name:      "test-name",
			},
		},
		"test_new_namsepaced_ID_from_with_namespace_with_dash_dot": {
			arg: "test.namespace-dash-dot:test-name",
			want: &NamespacedID{
				Namespace: "test.namespace-dash-dot",
				Name:      "test-name",
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewNamespacedIDFrom(testCase.arg)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestNamespaceIDString(t *testing.T) {
	testNamespaceID := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "test-name",
	}

	want := "test.namespace:test-name"

	got := testNamespaceID.String()
	internal.AssertEqual(t, want, got)

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("NamespaceID.String() = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(want))
	}
}

func TestNamespaceIDMarshalJSON(t *testing.T) {
	testNamespace := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "test-name",
	}

	want := []byte("\"test.namespace:test-name\"")

	got, _ := testNamespace.MarshalJSON()
	internal.AssertEqual(t, want, got)
}

func TestNamespaceIDUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		arg     []byte
		want    *NamespacedID
		wantErr error
	}{
		"test_namespaced_ID_unmarshal_json_valid": {
			arg: []byte("\"test.namespace:test-name\""),
			want: &NamespacedID{
				Namespace: "test.namespace",
				Name:      "test-name",
			},
			wantErr: nil,
		},
		"test_namespace_ID_unmarshal_json_namespace_dash": {
			arg: []byte("\"test-namespace:test-name\""),
			want: &NamespacedID{
				Namespace: "test-namespace",
				Name:      "test-name",
			},
			wantErr: nil,
		},
		"test_namespace_ID_unmarshal_json_namespace_dash_dot": {
			arg: []byte("\"test.namespace-dash-dot:test-name\""),
			want: &NamespacedID{
				Namespace: "test.namespace-dash-dot",
				Name:      "test-name",
			},
			wantErr: nil,
		},
		"test_namespaced_ID_unmarshal_json_invalid": {
			arg: []byte("\"test:namespace\\test-name\""),
			wantErr: errors.New("invalid NamespacedID: test:namespace	est-name"),
		},
		"test_namespaced_ID_unmarshal_json_empty": {
			arg:     []byte(""),
			wantErr: errors.New("unexpected end of JSON input"),
		},
		"test_namespaced_ID_unmarshal_json_invalid_argument": {
			arg:     []byte("test.namespace:test-name"),
			wantErr: errors.New("invalid character 'e' in literal true (expecting 'r')"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := &NamespacedID{}
			err := got.UnmarshalJSON(testCase.arg)
			if testCase.wantErr != nil {
				internal.AssertError(t, testCase.wantErr, err)
			} else {
				internal.AssertEqual(t, testCase.want, got)
			}
		})
	}
}

func TestNamespaceIDWithNamespace(t *testing.T) {
	testNamespaceID := &NamespacedID{
		Name: "test-name",
	}

	arg := "test.namespace"

	want := &NamespacedID{
		Namespace: arg,
		Name:      "test-name",
	}

	got := testNamespaceID.WithNamespace(arg)
	internal.AssertEqual(t, want, got)
}

func TestNamespaceIDWithName(t *testing.T) {
	testNamespace := &NamespacedID{
		Namespace: "test.namespace",
	}

	arg := "test-name"

	want := &NamespacedID{
		Namespace: "test.namespace",
		Name:      "test-name",
	}

	got := testNamespace.WithName(arg)
	internal.AssertEqual(t, want, got)
}
