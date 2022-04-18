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
		testHeader  Headers
		want        string
		isValidType bool
	}{
		"test_with_correlation_id": {
			testHeader:  Headers{HeaderCorrelationID: "correlation-id"},
			want:        "correlation-id",
			isValidType: true,
		},
		"test_empty_correlation_id": {
			testHeader:  Headers{HeaderCorrelationID: ""},
			want:        "",
			isValidType: true,
		},
		"test_corrlation_id_number": {
			testHeader:  Headers{HeaderCorrelationID: 1},
			want:        "",
			isValidType: false,
		},
		"test_same_corrlation_ids_invalid_value": {
			testHeader: Headers{
				HeaderCorrelationID: 1,
				"CORRELATION-ID":    "test",
			},
			want:        "",
			isValidType: false,
		},
		"test_same_corrlation_ids_valid_value": {
			testHeader: Headers{
				HeaderCorrelationID: "1",
				"CORRELATION-ID":    "test",
			},
			want:        "1",
			isValidType: true,
		},
		"test_same_corrlation_ids": {
			testHeader: Headers{
				"correlation-ID": "1",
				"CORRELATION-ID": "test",
			},
			want:        "test",
			isValidType: true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got, ok := testCase.testHeader.CorrelationID()
			internal.AssertEqual(t, testCase.want, got)
			if testCase.isValidType {
				internal.AssertTrue(t, ok)
			} else {
				internal.AssertFalse(t, ok)
				internal.AssertEqual(t, 1, testCase.testHeader[HeaderCorrelationID])
			}
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
		"test_empty_timeout": {
			testHeader: Headers{HeaderTimeout: ""},
			want:       60 * time.Second,
		},
		"test_timeout_number": {
			testHeader: Headers{HeaderTimeout: 1},
			want:       60 * time.Second,
		},
		"test_same_timeout_invalid_value": {
			testHeader: Headers{
				HeaderTimeout: 1,
				"TIMEOUT":     "10s",
			},
			want: 60 * time.Second,
		},
		"test_same_timeout_valid_value": {
			testHeader: Headers{
				HeaderTimeout: "1s",
				"TIMEOUT":     "1s",
			},
			want: time.Second,
		},
		"test_same_timeout": {
			testHeader: Headers{
				"Timeout": "1",
				"TIMEOUT": "5s",
			},
			want: 5 * time.Second,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Timeout()
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestTimeoutValue(t *testing.T) {
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
		valueInMap interface{}
	}{
		"test_with_response_required_false": {
			testHeader: Headers{HeaderResponseRequired: false},
			want:       false,
			valueInMap: false,
		},
		"test_with_response_required_true": {
			testHeader: Headers{HeaderResponseRequired: true},
			want:       true,
			valueInMap: true,
		},
		"test_without_response_required": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       true,
		},
		"test_response_required_number": {
			testHeader: Headers{HeaderResponseRequired: 1},
			valueInMap: 1,
			want:       true,
		},
		"test_same_response_required_invalid_value": {
			testHeader: Headers{
				HeaderResponseRequired: 1,
				"RESPONSE-REQUIRED":    false,
			},
			valueInMap: 1,
			want:       true,
		},
		"test_same_response_required_valid_value": {
			testHeader: Headers{
				HeaderResponseRequired: false,
				"RESPONSE-REQUIRED":    true,
			},
			valueInMap: false,
			want:       false,
		},
		"test_same_response_required": {
			testHeader: Headers{
				"Response-required": false,
				"RESPONSE-REQUIRED": true,
			},
			valueInMap: nil,
			want:       true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IsResponseRequired()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderResponseRequired])
		})
	}
}

func TestHeadersChannel(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		want       string
		valueInMap interface{}
	}{
		"test_with_channel": {
			testHeader: Headers{HeaderChannel: "1"},
			want:       "1",
			valueInMap: "1",
		},
		"test_without_channel": {
			testHeader: Headers{},
			want:       "",
			valueInMap: nil,
		},
		"test_channel_number": {
			testHeader: Headers{HeaderChannel: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_channel_invalid_value": {
			testHeader: Headers{
				HeaderChannel:   1,
				"DITTO-CHANNEL": "test-channel",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_channel_valid_value": {
			testHeader: Headers{
				HeaderChannel:   "test-channel",
				"DITTO-CHANNEL": "new-test-channel",
			},
			valueInMap: "test-channel",
			want:       "test-channel",
		},
		"test_same_channel": {
			testHeader: Headers{
				"Ditto-Channel": "test-channel",
				"DITTO-CHANNEL": "new-test-channel",
			},
			valueInMap: nil,
			want:       "new-test-channel",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Channel()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderChannel])
		})
	}
}

func TestHeadersIsDryRun(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		want       bool
		valueInMap interface{}
	}{
		"test_with_dry_run": {
			testHeader: Headers{HeaderDryRun: true},
			valueInMap: true,
			want:       true,
		},
		"test_without_dry_run": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       false,
		},
		"test_dry_run_number": {
			testHeader: Headers{HeaderDryRun: 1},
			valueInMap: 1,
			want:       false,
		},
		"test_same_dry_run_invalid_value": {
			testHeader: Headers{
				HeaderDryRun:    1,
				"DITTO-DRY-RUN": "false",
			},
			valueInMap: 1,
			want:       false,
		},
		"test_same_dry_run_valid_value": {
			testHeader: Headers{
				HeaderDryRun:    true,
				"DITTO-DRY-RUN": true,
			},
			valueInMap: true,
			want:       true,
		},
		"test_same_dry_run": {
			testHeader: Headers{
				"Ditto-Dry-Run": true,
				"DITTO-DRY-RUN": false,
			},
			valueInMap: nil,
			want:       false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IsDryRun()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderDryRun])
		})
	}
}

func TestHeadersOrigin(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_origin": {
			testHeader: Headers{HeaderOrigin: "origin"},
			valueInMap: "origin",
			want:       "origin",
		},
		"test_without_origin": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_origin_number": {
			testHeader: Headers{HeaderOrigin: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_origin_invalid_value": {
			testHeader: Headers{
				HeaderOrigin: 1,
				"ORIGIN":     "test-origin",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_origin_valid_value": {
			testHeader: Headers{
				HeaderOrigin: "test-origin",
				"ORIGIN":     "test-new-origin",
			},
			valueInMap: "test-origin",
			want:       "test-origin",
		},
		"test_same_origin": {
			testHeader: Headers{
				"Origin": "test-origin",
				"ORIGIN": "test-new-origin",
			},
			valueInMap: nil,
			want:       "test-new-origin",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Origin()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderOrigin])
		})
	}
}

func TestHeadersOriginator(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_ditto_originator": {
			testHeader: Headers{HeaderOriginator: "ditto-originator"},
			valueInMap: "ditto-originator",
			want:       "ditto-originator",
		},
		"test_without_ditto_originator": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_ditto_originator_number": {
			testHeader: Headers{HeaderOriginator: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_ditto_originator_invalid_value": {
			testHeader: Headers{
				HeaderOriginator:   1,
				"DITTO-ORIGINATOR": "test-ditto-originator",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_ditto_originator_valid_value": {
			testHeader: Headers{
				HeaderOriginator:   "test-ditto-originator",
				"DITTO-ORIGINATOR": "test-new-ditto-originator",
			},
			valueInMap: "test-ditto-originator",
			want:       "test-ditto-originator",
		},
		"test_same_ditto_originator": {
			testHeader: Headers{
				"Ditto-Originator": "test-ditto-originator",
				"DITTO-ORIGINATOR": "test-new-ditto-originator",
			},
			valueInMap: nil,
			want:       "test-new-ditto-originator",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Originator()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderOriginator])
		})
	}
}

func TestHeadersETag(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_etag": {
			testHeader: Headers{HeaderETag: "test-etag"},
			valueInMap: "test-etag",
			want:       "test-etag",
		},
		"test_without_etag": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_etag_number": {
			testHeader: Headers{HeaderETag: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_etag_invalid_value": {
			testHeader: Headers{
				HeaderETag: 1,
				"ETAG":     "test-etag",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_etag_valid_value": {
			testHeader: Headers{
				HeaderETag: "test-etag",
				"ETAG":     "test-new-etag",
			},
			valueInMap: "test-etag",
			want:       "test-etag",
		},
		"test_same_etag": {
			testHeader: Headers{
				"ETag": "test-etag",
				"ETAG": "test-new-etag",
			},
			valueInMap: nil,
			want:       "test-new-etag",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ETag()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderETag])
		})
	}
}

func TestHeadersIfMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_if_match": {
			testHeader: Headers{HeaderIfMatch: "test-if-match"},
			valueInMap: "test-if-match",
			want:       "test-if-match",
		},
		"test_without_if_match": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_if_match_number": {
			testHeader: Headers{HeaderIfMatch: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_if_match_invalid_value": {
			testHeader: Headers{
				HeaderIfMatch: 1,
				"IF-MATCH":    "test-if-match",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_if_match_valid_value": {
			testHeader: Headers{
				HeaderIfMatch: "test-if-match",
				"IF-MATCH":    "test-new-if-match",
			},
			valueInMap: "test-if-match",
			want:       "test-if-match",
		},
		"test_same_if_match": {
			testHeader: Headers{
				"If-Match": "test-if-match",
				"IF-MATCH": "test-new-if-match",
			},
			valueInMap: nil,
			want:       "test-new-if-match",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IfMatch()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderIfMatch])
		})
	}
}

func TestHeadersIfNoneMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_if_none_match": {
			testHeader: Headers{HeaderIfNoneMatch: "test-if-none-match"},
			valueInMap: "test-if-none-match",
			want:       "test-if-none-match",
		},
		"test_without_if_none_match": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_if_none_match_number": {
			testHeader: Headers{HeaderIfNoneMatch: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_if_none_match_invalid_value": {
			testHeader: Headers{
				HeaderIfNoneMatch: 1,
				"IF-NONE-MATCH":   "test-if-none-match",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_if_none_match_valid_value": {
			testHeader: Headers{
				HeaderIfNoneMatch: "test-if-none-match",
				"IF-NONE-mATCH":   "test-new-if-none-match",
			},
			valueInMap: "test-if-none-match",
			want:       "test-if-none-match",
		},
		"test_same_if_none_match": {
			testHeader: Headers{
				"If-None-Match": "test-if-none-match",
				"IF-NONE-MATCH": "test-new-if-none-match",
			},
			valueInMap: nil,
			want:       "test-new-if-none-match",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.IfNoneMatch()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderIfNoneMatch])
		})
	}
}

func TestHeadersReplyTarget(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       int64
	}{
		"test_with_reply_target": {
			testHeader: Headers{HeaderReplyTarget: int64(123)},
			valueInMap: int64(123),
			want:       int64(123),
		},
		"test_without_reply_target": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       0,
		},
		"test_reply_target_string": {
			testHeader: Headers{HeaderReplyTarget: "1"},
			valueInMap: "1",
			want:       0,
		},
		"test_same_reply_target_invalid_value": {
			testHeader: Headers{
				HeaderReplyTarget:    "1",
				"DITTO-REPLY-TARGET": 1,
			},
			valueInMap: "1",
			want:       0,
		},
		"test_same_reply_target_valid_value": {
			testHeader: Headers{
				HeaderReplyTarget:    int64(1),
				"DITTO-REPLY-TARGET": "1",
			},
			valueInMap: int64(1),
			want:       int64(1),
		},
		"test_same_reply_target": {
			testHeader: Headers{
				"Ditto-Reply-Target": int64(1),
				"DITTO-REPLY-TARGET": int64(2),
			},
			valueInMap: nil,
			want:       int64(2),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ReplyTarget()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderReplyTarget])
		})
	}
}

func TestHeadersReplyTo(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_reply_to": {
			testHeader: Headers{HeaderReplyTo: "someone"},
			valueInMap: "someone",
			want:       "someone",
		},
		"test_without_reply_to": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_reply_to_number": {
			testHeader: Headers{HeaderReplyTo: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_reply_to_invalid_value": {
			testHeader: Headers{
				HeaderReplyTo: 1,
				"REPLY-TO":    "test-reply-to",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_reply_to_valid_value": {
			testHeader: Headers{
				HeaderReplyTo: "test-reply-to",
				"REPLY-TO":    "test-new-reply-to",
			},
			valueInMap: "test-reply-to",
			want:       "test-reply-to",
		},
		"test_same_reply-to": {
			testHeader: Headers{
				"Reply-To": "test-reply-to",
				"REPLY-TO": "test-new-reply-to",
			},
			valueInMap: nil,
			want:       "test-new-reply-to",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.ReplyTo()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderReplyTo])
		})
	}
}

func TestHeadersVersion(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       int64
	}{
		"test_with_version": {
			testHeader: Headers{HeaderVersion: int64(123)},
			valueInMap: int64(123),
			want:       int64(123),
		},
		"test_without_version": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       int64(2),
		},
		"test_version_string": {
			testHeader: Headers{HeaderVersion: "1"},
			valueInMap: "1",
			want:       int64(2),
		},
		"test_same_version_invalid_value": {
			testHeader: Headers{
				HeaderVersion: "1",
				"VERSION":     int64(1),
			},
			valueInMap: "1",
			want:       int64(2),
		},
		"test_same_version_valid_value": {
			testHeader: Headers{
				HeaderVersion: int64(1),
				"VERSION":     int64(2),
			},
			valueInMap: int64(1),
			want:       int64(1),
		},
		"test_same_version": {
			testHeader: Headers{
				"Version": int64(12),
				"VERSION": int64(3),
			},
			valueInMap: nil,
			want:       int64(3),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := testCase.testHeader.Version()
			internal.AssertEqual(t, testCase.want, got)
			internal.AssertEqual(t, testCase.valueInMap, testCase.testHeader[HeaderVersion])
		})
	}
}

func TestHeadersContentType(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		valueInMap interface{}
		want       string
	}{
		"test_with_content_type": {
			testHeader: Headers{HeaderContentType: "test-content-type"},
			valueInMap: "test-content-type",
			want:       "test-content-type",
		},
		"test_without_content_type": {
			testHeader: Headers{},
			valueInMap: nil,
			want:       "",
		},
		"test_content_type_number": {
			testHeader: Headers{HeaderContentType: 1},
			valueInMap: 1,
			want:       "",
		},
		"test_same_content_type_invalid_value": {
			testHeader: Headers{
				HeaderContentType: 1,
				"CONTENT-TYPE":    "test-content-type",
			},
			valueInMap: 1,
			want:       "",
		},
		"test_same_content_type_valid_value": {
			testHeader: Headers{
				HeaderContentType: "test-content-type",
				"CONTENT-TYPE":    "test-new-content-type",
			},
			valueInMap: "test-content-type",
			want:       "test-content-type",
		},
		"test_same_content_type": {
			testHeader: Headers{
				"Content-Type": "test-content-type",
				"CONTENT-TYPE": "test-new-content-type",
			},
			valueInMap: nil,
			want:       "test-new-content-type",
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

	envelope.WithHeaders(NewHeaders(WithGeneric("coRRelation-ID", "correlation-id-2")))

	// return the first correlation-id (side effect from unmarshal JSON)
	res, _ := envelope.Headers.CorrelationID()
	internal.AssertEqual(t, "correlation-id-2", res)

	json.Marshal(envelope.Headers)

	data := `{
	    "correlation-iD":"correlation-id-3"
	}`
	json.Unmarshal([]byte(data), &envelope.Headers)
	internal.AssertEqual(t, "correlation-id-3", envelope.Headers["correlation-iD"])
}
