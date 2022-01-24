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
	"encoding/json"
	"testing"
	"time"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestHeadersCorrelationID(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_correlation_id": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "correlation-id"},
			},
			want: "correlation-id",
		},
		"test_correlation_id_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{
					HeaderCorrelationID: nil,
				},
			},
			want: "",
		},
		"test_without_correlation_id": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.CorrelationID()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersTimeout(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       time.Duration
	}{
		"test_with_timeout": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderTimeout: "10s"},
			},
			want: 10 * time.Second,
		},
		"test_timeout_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderTimeout: ""},
			},
			want: 60 * time.Second,
		},
		"test_without_timeout": {
			testHeader: &Headers{},
			want:       60 * time.Second,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Timeout()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestTimeout(t *testing.T) {
	tests := map[string]struct {
		data string
		want time.Duration
	}{
		"test_seconds": {
			data: `{ "timeout": "10s" }`,
			want: 10 * time.Second,
		},
		"test_milliseconds": {
			data: `{ "timeout": "500ms" }`,
			want: 500 * time.Millisecond,
		},
		"test_minute": {
			data: `{ "timeout": "1m" }`,
			want: time.Minute,
		},
		"test_without_unit_symbol": {
			data: `{ "timeout": "10" }`,
			want: 10 * time.Second,
		},
		"test_with_60_m_timeout": {
			data: `{ "timeout": "60m" }`,
			want: 60 * time.Second,
		},
		"test_with_3600_s_timeout": {
			data: `{ "timeout": "3600" }`,
			want: 60 * time.Second,
		},
		"test_with_negative_timeout": {
			data: `{ "timeout": "-5" }`,
			want: 60 * time.Second,
		},
		"test_with_empty_timeout": {
			data: `{ "timeout": "" }`,
			want: 60 * time.Second,
		},
		"test_with_invalid_timeout": {
			data: `{ "timeout": "invalid" }`,
			want: 60 * time.Second,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var headers Headers
			json.Unmarshal([]byte(testCase.data), &headers)
			internal.AssertEqual(t, testCase.want, headers.Timeout())
		})
	}
}

func TestHeadersIsResponseRequired(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       bool
	}{
		"test_with_response_required": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderResponseRequired: true},
			},
			want: true,
		},
		"test_response_required_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderResponseRequired: nil},
			},
			want: false,
		},
		"test_without_response_required": {
			testHeader: &Headers{},
			want:       false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IsResponseRequired()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersChannel(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_channel": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderChannel: "1"},
			},
			want: "1",
		},
		"test_channel_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderChannel: nil},
			},
			want: "",
		},
		"test_without_channel": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Channel()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersIsDryRun(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       bool
	}{
		"test_with_dry_run": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderDryRun: true},
			},
			want: true,
		},
		"test_dry_run_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderDryRun: nil},
			},
			want: false,
		},
		"test_without_dry_run": {
			testHeader: &Headers{},
			want:       false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IsDryRun()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersOrigin(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_origin": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderOrigin: "origin"},
			},
			want: "origin",
		},
		"test_origin_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderOrigin: nil},
			},
			want: "",
		},
		"test_without_origin": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Origin()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersOriginator(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_ditto_originator": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderOriginator: "ditto-originator"},
			},
			want: "ditto-originator",
		},
		"test_ditto_originator_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderOriginator: nil},
			},
			want: "",
		},
		"test_without_ditto_originator": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Originator()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersETag(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_etag": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderETag: "test-etag"},
			},
			want: "test-etag",
		},
		"test_etag_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderETag: nil},
			},
			want: "",
		},
		"test_without_etag": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ETag()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersIfMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_if_match": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderIfMatch: "HeaderIfMatch"},
			},
			want: "HeaderIfMatch",
		},
		"test_if_match_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderIfMatch: nil},
			},
			want: "",
		},
		"test_without_if_match": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IfMatch()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersIfNoneMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_if_none_match": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderIfNoneMatch: "HeaderIfNoneMatch"},
			},
			want: "HeaderIfNoneMatch",
		},
		"test_if_none_match_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderIfNoneMatch: nil},
			},
			want: "",
		},
		"test_without_if_none_match": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IfNoneMatch()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersReplyTarget(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       int64
	}{
		"test_with_reply_target": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderReplyTarget: int64(123)},
			},
			want: int64(123),
		},
		"test_reply_target_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderReplyTarget: nil},
			},
			want: 0,
		},
		"test_without_reply_target": {
			testHeader: &Headers{},
			want:       0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ReplyTarget()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersReplyTo(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_reply_to": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderReplyTo: "someone"},
			},
			want: "someone",
		},
		"test_reply_to_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderReplyTo: nil},
			},
			want: "",
		},
		"test_without_reply_to": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ReplyTo()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersVersion(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       int64
	}{
		"test_with_version": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderSchemaVersion: int64(123)},
			},
			want: int64(123),
		},
		"test_version_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderSchemaVersion: nil},
			},
			want: 0,
		},
		"test_without_version": {
			testHeader: &Headers{},
			want:       0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Version()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersContentType(t *testing.T) {
	tests := map[string]struct {
		testHeader *Headers
		want       string
	}{
		"test_with_content_type": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderContentType: "HeaderContentType"},
			},
			want: "HeaderContentType",
		},
		"test_content_type_value_nil": {
			testHeader: &Headers{
				Values: map[string]interface{}{HeaderContentType: nil},
			},
			want: "",
		},
		"test_without_content_type": {
			testHeader: &Headers{},
			want:       "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ContentType()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestHeadersGeneric(t *testing.T) {
	t.Run("TestHeadersGeneric", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderContentType] = "HeaderContentType"
		h := &Headers{
			Values: arg,
		}

		got := h.Generic(HeaderContentType)
		internal.AssertEqual(t, arg[HeaderContentType], got)
	})
}

func TestHeadersMarshalJSON(t *testing.T) {
	argOk := make(map[string]interface{})
	argOk[HeaderContentType] = "application/json"
	argErr := make(map[string]interface{})
	someChannel := make(chan int)
	argErr["Channel"] = someChannel

	tests := map[string]struct {
		data    map[string]interface{}
		want    string
		wantErr bool
	}{
		"test_headers_marshal_JSON_ok": {
			data:    argOk,
			want:    "{\"content-type\":\"application/json\"}",
			wantErr: false,
		},
		"test_headers_marshal_JSON_error": {
			data:    argErr,
			wantErr: true,
		}}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := &Headers{testCase.data}
			got, err := h.MarshalJSON()
			if testCase.wantErr {
				if err == nil {
					t.Errorf("Headers.MarshalJSON() error must not be nil")
				}
			} else {
				if string(got) != testCase.want {
					t.Errorf("Headers.MarshalJSON() = %v, want %v", string(got), testCase.want)
				}
			}
		})
	}
}

func TestHeadersUnmarshalJSON(t *testing.T) {
	ct := "application/json"

	tests := map[string]struct {
		data    string
		wantErr bool
	}{
		"test_headers_unmarshal_JSON_ok": {
			data:    "{\"content-type\":\"application/json\"}",
			wantErr: false,
		},
		"test_headers_unmarshal_JSON_err": {
			data:    "",
			wantErr: true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeaders()
			err := got.UnmarshalJSON([]byte(testCase.data))
			if testCase.wantErr {
				if err == nil {
					t.Errorf("Headers.UnmarshalJSON() error must not be nil")
				}
			} else {
				if got.ContentType() != ct {
					t.Errorf("Headers.UnmarshalJSON() got = %v, want %v", got.ContentType(), ct)
				}
			}
		})
	}
}
