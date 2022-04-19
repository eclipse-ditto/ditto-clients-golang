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
	"errors"
	"testing"
	"time"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func WithError() HeaderOpt {
	return func(headers Headers) error {
		return errors.New("this is an error example")
	}
}

func TestApplyOptsHeader(t *testing.T) {
	tests := map[string]struct {
		opts    []HeaderOpt
		wantErr bool
	}{
		"test_apply_opts_header": {
			opts:    []HeaderOpt{WithChannel("somethingNow")},
			wantErr: false,
		},
		"test_apply_opts_header_error": {
			opts:    []HeaderOpt{WithError()},
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			headers := Headers{HeaderChannel: "somethingBefore"}

			err := applyOptsHeader(headers, testCase.opts...)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("applyOptsHeader() must rise an error")
				}
			} else if headers[HeaderChannel] != "somethingNow" {
				t.Errorf("applyOptsHeader() Header want = \"somethingNow\" got %v", headers[HeaderChannel])
			}
		})
	}
}

func TestNewHeaders(t *testing.T) {
	tests := map[string]struct {
		opts []HeaderOpt
		want Headers
	}{
		"test_new_headers": {
			opts: []HeaderOpt{WithChannel("someChannel")},
			want: Headers{HeaderChannel: "someChannel"},
		},
		"test_new_headers_error": {
			opts: []HeaderOpt{WithError()},
			want: nil,
		},
		"test_new_headers_without_opts": {
			opts: nil,
			want: Headers{},
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeaders(testCase.opts...)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestNewHeadersFrom(t *testing.T) {
	tests := map[string]struct {
		arg1 Headers
		arg2 []HeaderOpt
		want Headers
	}{
		"test_copy_existing_empty_header_with_new_value": {
			arg1: Headers{},
			arg2: []HeaderOpt{WithCorrelationID("test-correlation-id")},
			want: Headers{HeaderCorrelationID: "test-correlation-id"},
		},
		"test_copy_existing_not_empty_haeder_with_new_value": {
			arg1: Headers{HeaderCorrelationID: "test-correlation-id"},
			arg2: []HeaderOpt{WithContentType("application/json")},
			want: Headers{
				HeaderCorrelationID: "test-correlation-id",
				HeaderContentType:   "application/json",
			},
		},
		"test_copy_existing_not_empty_header_nil_value": {
			arg1: Headers{HeaderCorrelationID: "test-correlation-id"},
			arg2: nil,
			want: Headers{HeaderCorrelationID: "test-correlation-id"},
		},
		"test_copy_existing_not_empty_header_empty_value": {
			arg1: Headers{HeaderCorrelationID: "test-correlation-id"},
			arg2: []HeaderOpt{},
			want: Headers{HeaderCorrelationID: "test-correlation-id"},
		},
		"test_copy_existing_empty_header_nil_value": {
			arg1: Headers{},
			arg2: nil,
			want: Headers{},
		},
		"test_copy_nil_header_with_values": {
			arg1: nil,
			arg2: []HeaderOpt{WithCorrelationID("correlation-id")},
			want: Headers{HeaderCorrelationID: "correlation-id"},
		},
		"test_copy_nil_header_nil_value": {
			arg1: nil,
			arg2: nil,
			want: Headers{},
		},
		"test_copy_nil_header_empty_value": {
			arg1: nil,
			arg2: []HeaderOpt{},
			want: Headers{},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.arg1, testCase.arg2...)
			internal.AssertEqual(t, testCase.want, got)
		})
	}
}

func TestWithCorrelationID(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_correlation_id": {
			testHeader: Headers{HeaderCorrelationID: "correlation-id"},
			arg:        "new-correlation-id",
		},
		"test_change_first_met_correlation_id": {
			testHeader: Headers{
				"Correlation-ID": "correlation-id-1",
				"CORRELATION-ID": "correlation-id-2",
			},
			arg: "correlation-id-3",
		},
		"test_set_new_correation_id": {
			testHeader: NewHeaders(),
			arg:        "new-correlation-id",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithCorrelationID(testCase.arg))
			internal.AssertEqual(t, testCase.arg, got.CorrelationID())
		})
	}
}

func TestWithReplyTo(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_reply_to": {
			testHeader: Headers{HeaderReplyTo: "reply-to"},
			arg:        "new-reply-to",
		},
		"test_change_first_met_reply-to": {
			testHeader: Headers{
				"Reply-To": "reply-to-1",
				"REPLY-TO": "reply-to-2",
			},
			arg: "reply-to-3",
		},
		"test_set_new_reply-to": {
			testHeader: NewHeaders(),
			arg:        "new-reply-to",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithReplyTo(testCase.arg))
			want := got.ReplyTo()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithReplyTarget(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        int64
	}{
		"test_change_existing_ditto_reply_target": {
			testHeader: Headers{HeaderReplyTarget: 1},
			arg:        2,
		},
		"test_change_first_met_ditto_reply_target": {
			testHeader: Headers{
				"Ditto-Reply-Target": 1,
				"DITTO-REPLY-TARGET": 2,
			},
			arg: 3,
		},
		"test_set_new_reply-target": {
			testHeader: NewHeaders(),
			arg:        1,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithReplyTarget(testCase.arg))
			want := got.ReplyTarget()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithChannel(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_ditto_channel": {
			testHeader: Headers{HeaderChannel: "test-ditto-channel"},
			arg:        "new-ditto-channel",
		},
		"test_change_first_met_ditto_channel": {
			testHeader: Headers{
				"Ditto-Channel": "test-ditto-channel-1",
				"DITTO-CHANNEL": "test-ditto-channel-2",
			},
			arg: "test-ditto-channel-3",
		},
		"test_set_new_ditto-channel": {
			testHeader: NewHeaders(),
			arg:        "test-ditto-channel",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithChannel(testCase.arg))
			want := got.Channel()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithResponseRequired(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        bool
	}{
		"test_change_existing_response_required": {
			testHeader: Headers{HeaderResponseRequired: true},
			arg:        false,
		},
		"test_change_first_met_response_required": {
			testHeader: Headers{
				"Response-Required": true,
				"RESPONSE-REQUIRED": true,
			},
			arg: false,
		},
		"test_set_new_response_required": {
			testHeader: NewHeaders(),
			arg:        false,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithResponseRequired(testCase.arg))
			want := got.IsResponseRequired()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithOriginator(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_ditto_originator": {
			testHeader: Headers{HeaderOriginator: "test-ditto-originator"},
			arg:        "new-ditto-originator",
		},
		"test_change_first_met_ditto_originator": {
			testHeader: Headers{
				"Ditto-Originator": "ditto-originator-1",
				"DITTO-ORIGINATOR": "ditto-originator-2",
			},
			arg: "ditto-originator-3",
		},
		"test_set_new_ditto_originator": {
			testHeader: NewHeaders(),
			arg:        "test-ditto-originator",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithOriginator(testCase.arg))
			want := got.Originator()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithOrigin(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_origin": {
			testHeader: Headers{HeaderOrigin: "test-origin"},
			arg:        "new-origin",
		},
		"test_change_first_met_origin": {
			testHeader: Headers{
				"Origin": "origin-1",
				"ORIGIN": "origin-2",
			},
			arg: "origin-3",
		},
		"test_set_new_origin": {
			testHeader: NewHeaders(),
			arg:        "test-origin",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithOrigin(testCase.arg))
			want := got.Origin()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithDryRun(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        bool
	}{
		"test_change_existing_dry_run": {
			testHeader: Headers{HeaderDryRun: true},
			arg:        false,
		},
		"test_change_first_met_dry_run": {
			testHeader: Headers{
				"Dry-Run": true,
				"DRY-RUN": true,
			},
			arg: false,
		},
		"test_set_new_dry_run": {
			testHeader: NewHeaders(),
			arg:        true,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithDryRun(testCase.arg))
			want := got.IsDryRun()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithETag(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_etag": {
			testHeader: Headers{HeaderETag: "test-etag"},
			arg:        "new-etag",
		},
		"test_change_first_met_etag": {
			testHeader: Headers{
				"ETag": "etag-1",
				"ETAG": "etag-2",
			},
			arg: "etag-3",
		},
		"test_set_new_etag": {
			testHeader: NewHeaders(),
			arg:        "test-etag",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithETag(testCase.arg))
			want := got.ETag()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithIfMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_if_match": {
			testHeader: Headers{HeaderIfMatch: "test-if-match"},
			arg:        "new-if-match",
		},
		"test_change_first_met_if_match": {
			testHeader: Headers{
				"If-Match": "if-match-1",
				"IF-MATCH": "if-match-2",
			},
			arg: "if-match-3",
		},
		"test_set_new_if_match": {
			testHeader: NewHeaders(),
			arg:        "test-if-match",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithIfMatch(testCase.arg))
			want := got.IfMatch()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithIfNoneMatch(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_if_none_match": {
			testHeader: Headers{HeaderIfNoneMatch: "test-if-none-match"},
			arg:        "new-if-none-match",
		},
		"test_change_first_met_if_none_match": {
			testHeader: Headers{
				"If-None-Match": "if-none-match-1",
				"IF-NONE-MATCH": "if-none-match-2",
			},
			arg: "if-none-match-3",
		},
		"test_set_new_if_none_match": {
			testHeader: NewHeaders(),
			arg:        "test-if_none_match",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithIfNoneMatch(testCase.arg))
			want := got.IfNoneMatch()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithTimeout(t *testing.T) {
	tests := map[string]struct {
		arg  time.Duration
		want time.Duration
	}{
		"test_with_seconds": {
			arg:  10 * time.Second,
			want: 10 * time.Second,
		},
		"test_with_milliseconds": {
			arg:  500 * time.Millisecond,
			want: 500 * time.Millisecond,
		},
		"test_with_minute": {
			arg:  1 * time.Minute,
			want: time.Minute,
		},
		"test_with_zero": {
			arg:  0,
			want: 0 * time.Second,
		},
		"test_with_1_hour_timeout": {
			arg:  time.Hour,
			want: 60 * time.Second,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeaders(WithTimeout(testCase.arg))
			internal.AssertEqual(t, testCase.want, got.Timeout())
		})
	}
}

func TestWithSchemaVersion(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        int64
	}{
		"test_change_existing_version": {
			testHeader: Headers{HeaderVersion: 1},
			arg:        int64(2),
		},
		"test_change_first_met_version": {
			testHeader: Headers{
				"Version": 0,
				"VERSION": int64(1),
			},
			arg: 2,
		},
		"test_set_new_etag": {
			testHeader: NewHeaders(),
			arg:        int64(2),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithVersion(testCase.arg))
			want := got.Version()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithContentType(t *testing.T) {
	tests := map[string]struct {
		testHeader Headers
		arg        string
	}{
		"test_change_existing_content_type": {
			testHeader: Headers{HeaderContentType: "test-content-type"},
			arg:        "new-content-type",
		},
		"test_change_first_met_content_type": {
			testHeader: Headers{
				"Content-Type": "content-type-1",
				"CONTENT-TYPE": "content-type-2",
			},
			arg: "content-type-3",
		},
		"test_set_new_content_type": {
			testHeader: NewHeaders(),
			arg:        "test-content-type",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := NewHeadersFrom(testCase.testHeader, WithContentType(testCase.arg))
			want := got.ContentType()
			internal.AssertEqual(t, testCase.arg, want)
		})
	}
}

func TestWithGeneric(t *testing.T) {
	t.Run("TestWithGeneric", func(t *testing.T) {
		hct := "contentType"

		got := NewHeaders(WithGeneric(HeaderContentType, hct))
		internal.AssertEqual(t, hct, got.ContentType())
	})
}
