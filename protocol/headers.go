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

// Ditto-specific headers constants.
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

// CorrelationID returns the 'correlation-id' header value or empty string if not set.
func (h *Headers) CorrelationID() string {
	if h.Values[HeaderCorrelationID] == nil {
		return ""
	}
	return h.Values[HeaderCorrelationID].(string)
}

// Timeout returns the 'timeout' header value or empty string if not set.
func (h *Headers) Timeout() string {
	if h.Values[HeaderTimeout] == nil {
		return ""
	}
	return h.Values[HeaderTimeout].(string)
}

// IsResponseRequired returns the 'response-required' header value or empty string if not set.
func (h *Headers) IsResponseRequired() bool {
	if h.Values[HeaderResponseRequired] == nil {
		return false
	}
	return h.Values[HeaderResponseRequired].(bool)
}

// Channel returns the 'ditto-channel' header value or empty string if not set.
func (h *Headers) Channel() string {
	if h.Values[HeaderChannel] == nil {
		return ""
	}
	return h.Values[HeaderChannel].(string)
}

// IsDryRun returns the 'ditto-dry-run' header value or empty string if not set.
func (h *Headers) IsDryRun() bool {
	if h.Values[HeaderDryRun] == nil {
		return false
	}
	return h.Values[HeaderDryRun].(bool)
}

// Origin returns the 'origin' header value or empty string if not set.
func (h *Headers) Origin() string {
	if h.Values[HeaderOrigin] == nil {
		return ""
	}
	return h.Values[HeaderOrigin].(string)
}

// Originator returns the 'ditto-originator' header value or empty string if not set.
func (h *Headers) Originator() string {
	if h.Values[HeaderOriginator] == nil {
		return ""
	}
	return h.Values[HeaderOriginator].(string)
}

// ETag returns the 'ETag' header value or empty string if not set.
func (h *Headers) ETag() string {
	if h.Values[HeaderETag] == nil {
		return ""
	}
	return h.Values[HeaderETag].(string)
}

// IfMatch returns the 'If-Match' header value or empty string if not set.
func (h *Headers) IfMatch() string {
	if h.Values[HeaderIfMatch] == nil {
		return ""
	}
	return h.Values[HeaderIfMatch].(string)
}

// IfNoneMatch returns the 'If-None-Match' header value or empty string if not set.
func (h *Headers) IfNoneMatch() string {
	if h.Values[HeaderIfNoneMatch] == nil {
		return ""
	}
	return h.Values[HeaderIfNoneMatch].(string)
}

// ReplyTarget returns the 'ditto-reply-target' header value or empty string if not set.
func (h *Headers) ReplyTarget() int64 {
	if h.Values[HeaderReplyTarget] == nil {
		return 0
	}
	return h.Values[HeaderReplyTarget].(int64)
}

// ReplyTo returns the 'reply-to' header value or empty string if not set.
func (h *Headers) ReplyTo() string {
	if h.Values[HeaderReplyTo] == nil {
		return ""
	}
	return h.Values[HeaderReplyTo].(string)
}

// Version returns the 'version' header value or empty string if not set.
func (h *Headers) Version() int64 {
	if h.Values[HeaderSchemaVersion] == nil {
		return 0
	}
	return h.Values[HeaderSchemaVersion].(int64)
}

// ContentType returns the 'content-type' header value or empty string if not set.
func (h *Headers) ContentType() string {
	if h.Values[HeaderContentType] == nil {
		return ""
	}
	return h.Values[HeaderContentType].(string)
}

// Generic returns the value of the provided key header and if a header with such key is present.
func (h *Headers) Generic(id string) interface{} {
	return h.Values[id]
}

// MarshalJSON marshels Headers.
func (h *Headers) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Values)
}

// UnmarshalJSON unmarshels Headers.
func (h *Headers) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	h.Values = v
	return nil
}
