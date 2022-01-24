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
	return func(headers *Headers) error {
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
			res := make(map[string]interface{})
			res[HeaderChannel] = "somethingBefore"
			headers := &Headers{res}
			err := applyOptsHeader(headers, testCase.opts...)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("applyOptsHeader() must rise an error")
				}
			} else if headers.Values[HeaderChannel] != "somethingNow" {
				t.Errorf("applyOptsHeader() Header want = \"somethingNow\" got %v", headers.Values[HeaderChannel])
			}
		})
	}
}

func TestNewHeaders(t *testing.T) {
	tests := map[string]struct {
		opts []HeaderOpt
		want *Headers
	}{
		"test_new_headers": {
			opts: []HeaderOpt{WithChannel("someChannel")},
			want: &Headers{
				Values: map[string]interface{}{
					HeaderChannel: "someChannel",
				},
			},
		},
		"test_new_headers_error": {
			opts: []HeaderOpt{WithError()},
			want: nil,
		},
		"test_new_headers_without_opts": {
			opts: nil,
			want: &Headers{
				Values: make(map[string]interface{}),
			},
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
		arg1 *Headers
		arg2 []HeaderOpt
		want *Headers
	}{
		"test_copy_existing_empty_header_with_new_value": {
			arg1: &Headers{},
			arg2: []HeaderOpt{WithCorrelationID("test-correlation-id")},
			want: &Headers{
				Values: map[string]interface{}{
					HeaderCorrelationID: "test-correlation-id",
				},
			},
		},
		"test_copy_existing_not_empty_haeder_with_new_value": {
			arg1: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "test-correlation-id"},
			},
			arg2: []HeaderOpt{WithContentType("application/json")},
			want: &Headers{
				Values: map[string]interface{}{
					HeaderCorrelationID: "test-correlation-id",
					HeaderContentType:   "application/json",
				},
			},
		},
		"test_copy_existing_not_empty_header_nil_value": {
			arg1: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "test-correlation-id"},
			},
			arg2: nil,
			want: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "test-correlation-id"},
			},
		},
		"test_copy_existing_not_empty_header_empty_value": {
			arg1: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "test-correlation-id"},
			},
			arg2: []HeaderOpt{},
			want: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "test-correlation-id"},
			},
		},
		"test_copy_existing_empty_header_nil_value": {
			arg1: &Headers{},
			arg2: nil,
			want: &Headers{
				Values: make(map[string]interface{}),
			},
		},
		"test_copy_nil_header_with_values": {
			arg1: nil,
			arg2: []HeaderOpt{WithCorrelationID("correlation-id")},
			want: &Headers{
				Values: map[string]interface{}{HeaderCorrelationID: "correlation-id"},
			},
		},
		"test_copy_nil_header_nil_value": {
			arg1: nil,
			arg2: nil,
			want: &Headers{
				make(map[string]interface{}),
			},
		},
		"test_copy_nil_header_empty_value": {
			arg1: nil,
			arg2: []HeaderOpt{},
			want: &Headers{
				Values: make(map[string]interface{}),
			},
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
	t.Run("TestWithCorrelationID", func(t *testing.T) {
		cid := "correlationId"

		got := NewHeaders(WithCorrelationID(cid))
		internal.AssertEqual(t, cid, got.CorrelationID())
	})
}

func TestWithReplyTo(t *testing.T) {
	t.Run("TestWithReplyTo", func(t *testing.T) {
		rto := "replyto"

		got := NewHeaders(WithReplyTo(rto))
		internal.AssertEqual(t, rto, got.ReplyTo())
	})
}

func TestWithReplyTarget(t *testing.T) {
	t.Run("TestWithReplyTarget", func(t *testing.T) {
		rt := "11111"

		got := NewHeaders(WithReplyTarget(rt))
		internal.AssertEqual(t, rt, got.Values[HeaderReplyTarget])
	})
}

func TestWithChannel(t *testing.T) {
	t.Run("TestWithChannel", func(t *testing.T) {
		cha := "channel"

		got := NewHeaders(WithChannel(cha))
		internal.AssertEqual(t, cha, got.Channel())
	})
}

func TestWithResponseRequired(t *testing.T) {
	t.Run("TestWithResponseRequired", func(t *testing.T) {
		rrq := true

		got := NewHeaders(WithResponseRequired(rrq))
		internal.AssertEqual(t, rrq, got.IsResponseRequired())
	})
}

func TestWithOriginator(t *testing.T) {
	t.Run("TestWithOriginator", func(t *testing.T) {
		org := "originator"

		got := NewHeaders(WithOriginator(org))
		internal.AssertEqual(t, org, got.Originator())
	})
}

func TestWithOrigin(t *testing.T) {
	t.Run("TestWithOrigin", func(t *testing.T) {
		org := "origin"

		got := NewHeaders(WithOrigin(org))
		internal.AssertEqual(t, org, got.Origin())
	})
}

func TestWithDryRun(t *testing.T) {
	t.Run("TestWithDryRun", func(t *testing.T) {
		dry := true

		got := NewHeaders(WithDryRun(dry))
		internal.AssertEqual(t, dry, got.IsDryRun())
	})
}

func TestWithETag(t *testing.T) {
	t.Run("TestWithETag", func(t *testing.T) {
		et := "etag"

		got := NewHeaders(WithETag(et))
		internal.AssertEqual(t, et, got.ETag())
	})
}

func TestWithIfMatch(t *testing.T) {
	t.Run("TestWithIfMatch", func(t *testing.T) {
		im := "ifMatch"

		got := NewHeaders(WithIfMatch(im))
		internal.AssertEqual(t, im, got.IfMatch())
	})
}

func TestWithIfNoneMatch(t *testing.T) {
	t.Run("TestWithIfNoneMatch", func(t *testing.T) {
		inm := "ifNoneMatch"

		got := NewHeaders(WithIfNoneMatch(inm))
		internal.AssertEqual(t, inm, got.IfNoneMatch())
	})
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
		"test_without_unit": {
			arg:  5,
			want: 1 * time.Millisecond,
		},
		"test_with_invalid_timeout": {
			arg:  -1,
			want: 60 * time.Second,
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
	t.Run("TestWithSchemaVersion", func(t *testing.T) {
		sv := "123456789"

		got := NewHeaders(WithSchemaVersion("123456789"))
		internal.AssertEqual(t, sv, got.Values[HeaderSchemaVersion])
	})
}

func TestWithContentType(t *testing.T) {
	t.Run("TestWithContentType", func(t *testing.T) {
		hct := "contentType"

		got := NewHeaders(WithContentType(hct))
		internal.AssertEqual(t, hct, got.ContentType())
	})
}

func TestWithGeneric(t *testing.T) {
	t.Run("TestWithGeneric", func(t *testing.T) {
		hct := "contentType"

		got := NewHeaders(WithGeneric(HeaderContentType, hct))
		internal.AssertEqual(t, hct, got.ContentType())
	})
}
