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
		testHeader Headers
		want       string
	}{
		"test_with_correlation_id": {
			testHeader: Headers{HeaderCorrelationID: "correlation-id"},
			want:       "correlation-id",
		},
		"test_without_correlation_id": {
			testHeader: Headers{},
			want:       "",
		},

		"test_empty_correlation_id": {
			testHeader: Headers{HeaderCorrelationID: ""},
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
		testHeader Headers
		want       time.Duration
	}{
		"test_with_timeout": {
			testHeader: Headers{HeaderTimeout: "10s"},
			want:       10 * time.Second,
		},
		"test_without_timeout": {
			testHeader: Headers{},
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
			headers := NewHeaders()
			json.Unmarshal([]byte(testCase.data), &headers)
			internal.AssertEqual(t, testCase.want, headers.Timeout())
		})
	}
}

func TestHeadersIsResponseRequired(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		want       bool
	}{
		"test_with_response_required_false": {
			testHeader: Headers{HeaderResponseRequired: false},
			want:       false,
		},
		"test_with_response_required_true": {
			testHeader: Headers{HeaderResponseRequired: true},
			want:       true,
		},
		"test_without_response_required": {
			testHeader: Headers{},
			want:       true,
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
		testHeader Headers
		want       string
	}{
		"test_with_channel": {
			testHeader: Headers{HeaderChannel: "1"},
			want:       "1",
		},
		"test_without_channel": {
			testHeader: Headers{},
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
		testHeader Headers
		want       bool
	}{
		"test_with_dry_run": {
			testHeader: Headers{HeaderDryRun: true},
			want:       true,
		},
		"test_without_dry_run": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_origin": {
			testHeader: Headers{HeaderOrigin: "origin"},
			want:       "origin",
		},
		"test_without_origin": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_ditto_originator": {
			testHeader: Headers{HeaderOriginator: "ditto-originator"},
			want:       "ditto-originator",
		},
		"test_without_ditto_originator": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_etag": {
			testHeader: Headers{HeaderETag: "test-etag"},
			want:       "test-etag",
		},
		"test_without_etag": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_if_match": {
			testHeader: Headers{HeaderIfMatch: "HeaderIfMatch"},
			want:       "HeaderIfMatch",
		},
		"test_without_if_match": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_if_none_match": {
			testHeader: Headers{HeaderIfNoneMatch: "HeaderIfNoneMatch"},
			want:       "HeaderIfNoneMatch",
		},
		"test_without_if_none_match": {
			testHeader: Headers{},
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
		testHeader Headers
		want       int64
	}{
		"test_with_reply_target": {
			testHeader: Headers{HeaderReplyTarget: int64(123)},
			want:       int64(123),
		},
		"test_without_reply_target": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_reply_to": {
			testHeader: Headers{HeaderReplyTo: "someone"},
			want:       "someone",
		},
		"test_without_reply_to": {
			testHeader: Headers{},
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
		testHeader Headers
		want       int64
	}{
		"test_with_version": {
			testHeader: Headers{HeaderSchemaVersion: int64(123)},
			want:       int64(123),
		},
		"test_without_version": {
			testHeader: Headers{},
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
		testHeader Headers
		want       string
	}{
		"test_with_content_type": {
			testHeader: Headers{HeaderContentType: "HeaderContentType"},
			want:       "HeaderContentType",
		},
		"test_without_content_type": {
			testHeader: Headers{},
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
		h := Headers{HeaderContentType: "HeaderContentType"}

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
			got, err := json.Marshal(testCase.data)
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
	tests := map[string]struct {
		data string
		want Headers
	}{
		"test_headers_unmarshal_JSON_with_one_heder": {
			data: `{"content-type":"application/json"}`,
			want: Headers{HeaderContentType: "application/json"},
		},
		"test_headers_unmarshal_JSON_with_many_headers": {
			data: `{
				"content-type":"application/json",
				"timeout": "30ms",
				"response-required":false
			}`,
			want: Headers{
				HeaderContentType:      "application/json",
				HeaderTimeout:          "30ms",
				HeaderResponseRequired: false,
			},
		},
		"test_headers_unmarshal_JSON_err": {
			data: "",
			want: Headers{},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeaders()
			json.Unmarshal([]byte(testCase.data), &got)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestCaseInsensitiveKey(t *testing.T) {
	headers := Headers{HeaderCorrelationID: "correlation-id-1"}
	envelope := &Envelope{
		Headers: headers,
	}
	// override correlation-id instead of custom header - header key is the last set one
	internal.AssertEqual(t, "correlation-id-1", envelope.Headers.Generic("correlation-ID"))

	envelope.WithHeaders(NewHeaders(WithGeneric("coRRelation-ID", "correlation-id-2")))

	// return the first correlation-id (side effect from unmarshal JSON)
	res := envelope.Headers.CorrelationID()
	internal.AssertEqual(t, "correlation-id-2", res)

	json.Marshal(envelope.Headers)

	data := `{
	    "correlation-iD":"correlation-id-3"
	}`
	json.Unmarshal([]byte(data), &envelope.Headers)
	internal.AssertEqual(t, "correlation-id-3", envelope.Headers["correlation-iD"])
}
