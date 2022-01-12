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
	"reflect"
	"testing"

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
		opts    []HeaderOpt
		wantErr bool
	}{
		"test_new_headers": {
			opts:    []HeaderOpt{WithChannel("someChannel")},
			wantErr: false,
		},
		"test_new_headers_error": {
			opts:    []HeaderOpt{WithError()},
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res := make(map[string]interface{})
			res[HeaderChannel] = "someChannel"
			want := &Headers{res}
			got := NewHeaders(testCase.opts...)
			if testCase.wantErr {
				if got != nil {
					t.Errorf("NewHeaders() must be nil")
				}
			} else if !reflect.DeepEqual(got, want) {
				t.Errorf("NewHeaders() = %v want %v", got, want)
			}
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
	t.Run("TestWithTimeout", func(t *testing.T) {
		tmo := "10"

		got := NewHeaders(WithTimeout(tmo))
		internal.AssertEqual(t, tmo, got.Timeout())
	})
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
