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
	"testing"

	"github.com/eclipse/ditto-clients-golang/internal"
)

func TestHeadersCorrelationID(t *testing.T) {
	t.Run("TestHeadersCorrelationID", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderCorrelationID] = "correlation-id"
		h := &Headers{
			Values: arg,
		}

		got := h.CorrelationID()
		internal.AssertEqual(t, "correlation-id", got)

		arg[HeaderCorrelationID] = nil

		got = h.CorrelationID()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersTimeout(t *testing.T) {
	t.Run("TestHeadersTimeout", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderTimeout] = "10"
		h := &Headers{
			Values: arg,
		}

		got := h.Timeout()
		internal.AssertEqual(t, "10", got)

		arg[HeaderTimeout] = nil
		got = h.Timeout()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersIsResponseRequired(t *testing.T) {
	t.Run("TestHeadersIsResponseRequired", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderResponseRequired] = false
		h := &Headers{
			Values: arg,
		}

		got := h.IsResponseRequired()
		internal.AssertFalse(t, got)

		arg[HeaderResponseRequired] = nil
		got = h.IsResponseRequired()
		internal.AssertFalse(t, got)
	})
}

func TestHeadersChannel(t *testing.T) {
	t.Run("TestHeadersChannel", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderChannel] = "1"
		h := &Headers{
			Values: arg,
		}

		got := h.Channel()
		internal.AssertEqual(t, "1", got)

		arg[HeaderChannel] = nil
		got = h.Channel()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersIsDryRun(t *testing.T) {
	t.Run("TestHeadersIsDryRun", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderDryRun] = false
		h := &Headers{
			Values: arg,
		}

		got := h.IsDryRun()
		internal.AssertFalse(t, got)

		arg[HeaderDryRun] = nil
		got = h.IsDryRun()
		internal.AssertFalse(t, got)
	})
}

func TestHeadersOrigin(t *testing.T) {
	t.Run("TestHeadersOrigin", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderOrigin] = "origin"
		h := &Headers{
			Values: arg,
		}

		got := h.Origin()
		internal.AssertEqual(t, "origin", got)

		arg[HeaderOrigin] = nil
		got = h.Origin()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersOriginator(t *testing.T) {
	t.Run("TestHeadersOriginator", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderOriginator] = "ditto-originator"
		h := &Headers{
			Values: arg,
		}

		got := h.Originator()
		internal.AssertEqual(t, "ditto-originator", got)

		arg[HeaderOriginator] = nil
		got = h.Originator()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersETag(t *testing.T) {
	t.Run("TestHeadersETag", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderETag] = "1"
		h := &Headers{
			Values: arg,
		}

		got := h.ETag()
		internal.AssertEqual(t, "1", got)

		arg[HeaderETag] = nil
		got = h.ETag()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersIfMatch(t *testing.T) {
	t.Run("TestHeadersIfMatch", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderIfMatch] = "HeaderIfMatch"
		h := &Headers{
			Values: arg,
		}

		got := h.IfMatch()
		internal.AssertEqual(t, "HeaderIfMatch", got)

		arg[HeaderIfMatch] = nil
		got = h.IfMatch()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersIfNoneMatch(t *testing.T) {
	t.Run("TestHeadersIfNoneMatch", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderIfNoneMatch] = "123"
		h := &Headers{
			Values: arg,
		}

		got := h.IfNoneMatch()
		internal.AssertEqual(t, "123", got)

		arg[HeaderIfNoneMatch] = nil
		got = h.IfNoneMatch()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersReplyTarget(t *testing.T) {
	t.Run("TestHeadersReplyTarget", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderReplyTarget] = int64(123)
		h := &Headers{
			Values: arg,
		}

		got := h.ReplyTarget()
		internal.AssertEqual(t, int64(123), got)

		arg[HeaderReplyTarget] = nil
		got = h.ReplyTarget()
		internal.AssertEqual(t, int64(0), got)
	})
}

func TestHeadersReplyTo(t *testing.T) {
	t.Run("TestHeadersReplyTo", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderReplyTo] = "someone"
		h := &Headers{
			Values: arg,
		}

		got := h.ReplyTo()
		internal.AssertEqual(t, "someone", got)

		arg[HeaderReplyTo] = nil
		got = h.ReplyTo()
		internal.AssertEqual(t, "", got)
	})
}

func TestHeadersVersion(t *testing.T) {
	t.Run("TestHeadersVersion", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderSchemaVersion] = int64(1111)
		h := &Headers{
			Values: arg,
		}

		got := h.Version()
		internal.AssertEqual(t, int64(1111), got)

		arg[HeaderSchemaVersion] = nil
		got = h.Version()
		internal.AssertEqual(t, int64(0), got)
	})
}

func TestHeadersContentType(t *testing.T) {
	t.Run("TestHeadersContentType", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderContentType] = "HeaderContentType"
		h := &Headers{
			Values: arg,
		}

		got := h.ContentType()
		internal.AssertEqual(t, "HeaderContentType", got)

		arg[HeaderContentType] = nil
		got = h.ContentType()
		internal.AssertEqual(t, "", got)
	})
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
