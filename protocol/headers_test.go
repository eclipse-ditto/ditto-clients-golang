// Copyright (c) 2020 Contributors to the Eclipse Foundation
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
	"reflect"
	"testing"
)

func TestHeadersCorrelationID(t *testing.T) {
	t.Run("TestHeadersCorrelationID", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderCorrelationID] = "correlation-id"
		h := &Headers{
			Values: arg,
		}
		if got := h.CorrelationID(); got != "correlation-id" {
			t.Errorf("Headers.CorrelationID() = %v, want %v", got, "correlation-id")
		}
		arg[HeaderCorrelationID] = nil
		if got := h.CorrelationID(); got != "" {
			t.Errorf("Headers.CorrelationID() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersTimeout(t *testing.T) {
	t.Run("TestHeadersTimeout", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderTimeout] = "10"
		h := &Headers{
			Values: arg,
		}
		if got := h.Timeout(); got != "10" {
			t.Errorf("Headers.Timeout() = %v, want %v", got, "10")
		}
		arg[HeaderTimeout] = nil
		if got := h.Timeout(); got != "" {
			t.Errorf("Headers.Timeout() nil = %v, want \"\"", got)
		}

	})
}

func TestHeadersIsResponseRequired(t *testing.T) {
	t.Run("TestHeadersIsResponseRequired", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderResponseRequired] = false
		h := &Headers{
			Values: arg,
		}
		if got := h.IsResponseRequired(); got != false {
			t.Errorf("Headers.IsResponseRequired() = %v, want %v", got, false)
		}
		arg[HeaderResponseRequired] = nil
		if got := h.IsResponseRequired(); got != false {
			t.Errorf("Headers.IsResponseRequired() nil = %v, want %v", got, false)
		}
	})
}

func TestHeadersChannel(t *testing.T) {
	t.Run("TestHeadersChannel", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderChannel] = "1"
		h := &Headers{
			Values: arg,
		}
		if got := h.Channel(); got != "1" {
			t.Errorf("Headers.Channel() = %v, want %v", got, "1")
		}
		arg[HeaderChannel] = nil
		if got := h.Channel(); got != "" {
			t.Errorf("Headers.Channel() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersIsDryRun(t *testing.T) {
	t.Run("TestHeadersIsDryRun", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderDryRun] = false
		h := &Headers{
			Values: arg,
		}
		if got := h.IsDryRun(); got != false {
			t.Errorf("Headers.IsDryRun() = %v, want %v", got, false)
		}
		arg[HeaderDryRun] = nil
		if got := h.IsDryRun(); got != false {
			t.Errorf("Headers.IsDryRun() nil = %v, want %v", got, false)
		}
	})
}

func TestHeadersOrigin(t *testing.T) {
	t.Run("TestHeadersOrigin", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderOrigin] = "origin"
		h := &Headers{
			Values: arg,
		}
		if got := h.Origin(); got != "origin" {
			t.Errorf("Headers.Origin() = %v, want %v", got, "origin")
		}
		arg[HeaderOrigin] = nil
		if got := h.Origin(); got != "" {
			t.Errorf("Headers.Origin() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersOriginator(t *testing.T) {
	t.Run("TestHeadersOriginator", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderOriginator] = "ditto-originator"
		h := &Headers{
			Values: arg,
		}
		if got := h.Originator(); got != "ditto-originator" {
			t.Errorf("Headers.Originator() = %v, want %v", got, "ditto-originator")
		}
		arg[HeaderOriginator] = nil
		if got := h.Originator(); got != "" {
			t.Errorf("Headers.Originator() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersETag(t *testing.T) {
	t.Run("TestHeadersETag", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderETag] = "1"
		h := &Headers{
			Values: arg,
		}
		if got := h.ETag(); got != "1" {
			t.Errorf("Headers.ETag() = %v, want %v", got, "1")
		}
		arg[HeaderETag] = nil
		if got := h.ETag(); got != "" {
			t.Errorf("Headers.ETag() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersIfMatch(t *testing.T) {
	t.Run("TestHeadersIfMatch", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderIfMatch] = "HeaderIfMatch"
		h := &Headers{
			Values: arg,
		}
		if got := h.IfMatch(); got != "HeaderIfMatch" {
			t.Errorf("Headers.IfMatch() = %v, want %v", got, "HeaderIfMatch")
		}
		arg[HeaderIfMatch] = nil
		if got := h.IfMatch(); got != "" {
			t.Errorf("Headers.IfMatch() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersIfNoneMatch(t *testing.T) {
	t.Run("TestHeadersIfNoneMatch", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderIfNoneMatch] = "123"
		h := &Headers{
			Values: arg,
		}
		if got := h.IfNoneMatch(); got != "123" {
			t.Errorf("Headers.IfNoneMatch() = %v, want %v", got, "123")
		}
		arg[HeaderIfNoneMatch] = nil
		if got := h.IfNoneMatch(); got != "" {
			t.Errorf("Headers.IfNoneMatch() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersReplyTarget(t *testing.T) {
	t.Run("TestHeadersReplyTarget", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderReplyTarget] = int64(123)
		h := &Headers{
			Values: arg,
		}
		if got := h.ReplyTarget(); got != int64(123) {
			t.Errorf("Headers.ReplyTarget() = %v, want %v", got, int64(123))
		}
		arg[HeaderReplyTarget] = nil
		if got := h.ReplyTarget(); got != int64(0) {
			t.Errorf("Headers.ReplyTarget() nil = %v, want %v", got, int64(0))
		}
	})
}

func TestHeadersReplyTo(t *testing.T) {
	t.Run("TestHeadersReplyTo", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderReplyTo] = "someone"
		h := &Headers{
			Values: arg,
		}
		if got := h.ReplyTo(); got != "someone" {
			t.Errorf("Headers.ReplyTo() = %v, want %v", got, "someone")
		}
		arg[HeaderReplyTo] = nil
		if got := h.ReplyTo(); got != "" {
			t.Errorf("Headers.ReplyTo() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersVersion(t *testing.T) {
	t.Run("TestHeadersVersion", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderSchemaVersion] = int64(1111)
		h := &Headers{
			Values: arg,
		}
		if got := h.Version(); got != int64(1111) {
			t.Errorf("Headers.Version() = %v, want %v", got, int64(1111))
		}
		arg[HeaderSchemaVersion] = nil
		if got := h.Version(); got != int64(0) {
			t.Errorf("Headers.Version() nil = %v, want %v", got, int64(0))
		}
	})
}

func TestHeadersContentType(t *testing.T) {
	t.Run("TestHeadersContentType", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderContentType] = "HeaderContentType"
		h := &Headers{
			Values: arg,
		}
		if got := h.ContentType(); got != "HeaderContentType" {
			t.Errorf("Headers.ContentType() = %v, want %v", got, "HeaderContentType")
		}
		arg[HeaderContentType] = nil
		if got := h.ContentType(); got != "" {
			t.Errorf("Headers.ContentType() nil = %v, want \"\"", got)
		}
	})
}

func TestHeadersGeneric(t *testing.T) {
	t.Run("TestHeadersGeneric", func(t *testing.T) {
		arg := make(map[string]interface{})
		arg[HeaderContentType] = "HeaderContentType"
		h := &Headers{
			Values: arg,
		}
		if got := h.Generic(HeaderContentType); !reflect.DeepEqual(got, arg[HeaderContentType]) {
			t.Errorf("Headers.Generic() = %v, want %v", got, arg[HeaderContentType])
		}
	})
}

func TestHeadersMarshalJSON(t *testing.T) {
	argOk := make(map[string]interface{})
	argOk[HeaderContentType] = "application/json"
	argErr := make(map[string]interface{})
	someChannel := make(chan int)
	argErr["Channel"] = someChannel
	tests := []struct {
		name    string
		data    map[string]interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "TestHeadersMarshalJSON ok",
			data:    argOk,
			want:    "{\"content-type\":\"application/json\"}",
			wantErr: false,
		},
		{
			name:    "TestHeadersMarshalJSON error",
			data:    argErr,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Headers{tt.data}
			got, err := h.MarshalJSON()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Headers.MarshalJSON() error must not be nil")
				}
			} else {
				if string(got) != tt.want {
					t.Errorf("Headers.MarshalJSON() = %v, want %v", string(got), tt.want)
				}
			}
		})
	}
}

func TestHeadersUnmarshalJSON(t *testing.T) {
	ct := "application/json"
	tests := []struct {
		name    string
		data    string
		wantErr bool
	}{
		{
			name:    "TestHeadersUnmarshalJSON ok",
			data:    "{\"content-type\":\"application/json\"}",
			wantErr: false,
		},
		{
			name:    "TestHeadersUnmarshalJSON err",
			data:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHeaders()
			err := got.UnmarshalJSON([]byte(tt.data))
			if tt.wantErr {
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
