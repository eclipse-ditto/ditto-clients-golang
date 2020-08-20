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
	"encoding/json"
)

const (
	HeaderCorrelationID    = "correlation-id"
	HeaderResponseRequired = "response-required"
	HeaderChannel          = "ditto-channel"
	HeaderDryRun           = "ditto-dry-run"
	HeaderOrigin           = "origin"
	HeaderOriginator       = "ditto-originator"
	HeaderETag             = "ETag"
	HeaderIfMatch          = "If-Match"
	HeaderIfNoneMatch      = "If-None-Match"
	HeaderReplyTarget      = "ditto-reply-target"
	HeaderReplyTo          = "reply-to"
	HeaderTimeout          = "timeout"
	HeaderSchemaVersion    = "version"
	HeaderContentType      = "content-type"
)

// Headers represents all Ditto-specific headers along with additional HTTP/etc. headers
// that can be applied depending on the transport used.
// See https://www.eclipse.org/ditto/protocol-specification.html
type Headers struct {
	Values map[string]interface{}
}

func (h *Headers) CorrelationID() string {
	if h.Values[HeaderCorrelationID] == nil {
		return ""
	}
	return h.Values[HeaderCorrelationID].(string)
}

func (h *Headers) Timeout() string {
	if h.Values[HeaderTimeout] == nil {
		return ""
	}
	return h.Values[HeaderTimeout].(string)
}

func (h *Headers) IsResponseRequired() bool {
	if h.Values[HeaderResponseRequired] == nil {
		return false
	}
	return h.Values[HeaderResponseRequired].(bool)
}
func (h *Headers) Channel() string {
	if h.Values[HeaderChannel] == nil {
		return ""
	}
	return h.Values[HeaderChannel].(string)
}
func (h *Headers) IsDryRun() bool {
	if h.Values[HeaderDryRun] == nil {
		return false
	}
	return h.Values[HeaderDryRun].(bool)
}
func (h *Headers) Origin() string {
	if h.Values[HeaderOrigin] == nil {
		return ""
	}
	return h.Values[HeaderOrigin].(string)
}
func (h *Headers) Originator() string {
	if h.Values[HeaderOriginator] == nil {
		return ""
	}
	return h.Values[HeaderOriginator].(string)
}
func (h *Headers) ETag() string {
	if h.Values[HeaderETag] == nil {
		return ""
	}
	return h.Values[HeaderETag].(string)
}
func (h *Headers) IfMatch() string {
	if h.Values[HeaderIfMatch] == nil {
		return ""
	}
	return h.Values[HeaderIfMatch].(string)
}
func (h *Headers) IfNoneMatch() string {
	if h.Values[HeaderIfNoneMatch] == nil {
		return ""
	}
	return h.Values[HeaderIfNoneMatch].(string)
}
func (h *Headers) ReplyTarget() int64 {
	if h.Values[HeaderReplyTarget] == nil {
		return 0
	}
	return h.Values[HeaderReplyTarget].(int64)
}
func (h *Headers) ReplyTo() string {
	if h.Values[HeaderReplyTo] == nil {
		return ""
	}
	return h.Values[HeaderReplyTo].(string)
}
func (h *Headers) Version() int64 {
	if h.Values[HeaderSchemaVersion] == nil {
		return 0
	}
	return h.Values[HeaderSchemaVersion].(int64)
}
func (h *Headers) ContentType() string {
	if h.Values[HeaderContentType] == nil {
		return ""
	}
	return h.Values[HeaderContentType].(string)
}
func (h *Headers) Generic(id string) interface{} {
	return h.Values[id]
}

func (h *Headers) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Values)
}

func (h *Headers) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	h.Values = v
	return nil
}
