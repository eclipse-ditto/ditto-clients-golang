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
	"fmt"
	"strconv"
	"time"
)

const (
	// ContentTypeDitto defines the Ditto JSON 'content-type' header value for Ditto Protocol messages.
	ContentTypeDitto = "application/vnd.eclipse.ditto+json"

	// HeaderCorrelationID represents 'correlation-id' header.
	HeaderCorrelationID = "correlation-id"
	// HeaderResponseRequired represents 'response-required' header.
	HeaderResponseRequired = "response-required"
	// HeaderChannel represents 'ditto-channel' header.
	HeaderChannel = "ditto-channel"
	// HeaderDryRun represents 'ditto-dry-run' header.
	HeaderDryRun = "ditto-dry-run"
	// HeaderOrigin represents 'origin' header.
	HeaderOrigin = "origin"
	// HeaderOriginator represents 'ditto-originator' header.
	HeaderOriginator = "ditto-originator"
	// HeaderETag represents 'ETag' header.
	HeaderETag = "ETag"
	// HeaderIfMatch represents 'If-Match' header.
	HeaderIfMatch = "If-Match"
	// HeaderIfNoneMatch represents 'If-None-March' header.
	HeaderIfNoneMatch = "If-None-Match"
	// HeaderReplyTarget represents 'ditto-reply-target' header.
	HeaderReplyTarget = "ditto-reply-target"
	// HeaderReplyTo represents 'reply-to' header.
	HeaderReplyTo = "reply-to"
	// HeaderTimeout represents 'timeout' header.
	HeaderTimeout = "timeout"
	// HeaderSchemaVersion represents 'version' header.
	HeaderSchemaVersion = "version"
	// HeaderContentType represents 'content-type' header.
	HeaderContentType = "content-type"
)

// Headers represents all Ditto-specific headers along with additional HTTP/etc. headers
// that can be applied depending on the transport used.
// See https://www.eclipse.org/ditto/protocol-specification.html
type Headers struct {
	Values map[string]interface{}
}

// CorrelationID returns the 'correlation-id' header value or empty string if not set.
func (h *Headers) CorrelationID() string {
	if value, ok := h.Values[HeaderCorrelationID]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// Timeout returns the 'timeout' header value or duration of 60 seconds if not set.
func (h *Headers) Timeout() time.Duration {
	if value, ok := h.Values[HeaderTimeout]; ok {
		if duration, err := parseTimeout(value.(string)); err == nil {
			return duration
		}
	}
	return 60 * time.Second
}

func parseTimeout(timeout string) (time.Duration, error) {
	l := len(timeout)
	if l > 0 {
		t := time.Duration(-1)
		switch timeout[l-1] {
		case 'm':
			if i, err := strconv.Atoi(timeout[:l-1]); err == nil {
				t = time.Duration(i) * time.Minute
			}
		case 's':
			if timeout[l-2] == 'm' {
				if i, err := strconv.Atoi(timeout[:l-2]); err == nil {
					t = time.Duration(i) * time.Millisecond
				}
			} else {
				if i, err := strconv.Atoi(timeout[:l-1]); err == nil {
					t = time.Duration(i) * time.Second
				}
			}
		default:
			if i, err := strconv.Atoi(timeout); err == nil {
				t = time.Duration(i) * time.Second
			}
		}
		if inTimeoutRange(t) {
			return t, nil
		}
	}
	return -1, fmt.Errorf("invalid timeout '%s'", timeout)
}

func inTimeoutRange(timeout time.Duration) bool {
	return timeout >= 0 && timeout < time.Hour
}

// IsResponseRequired returns the 'response-required' header value or true if not set.
func (h *Headers) IsResponseRequired() bool {
	if value, ok := h.Values[HeaderResponseRequired]; ok && value != nil {
		return value.(bool)
	}
	return true
}

// Channel returns the 'ditto-channel' header value or empty string if not set.
func (h *Headers) Channel() string {
	if value, ok := h.Values[HeaderChannel]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// IsDryRun returns the 'ditto-dry-run' header value or empty string if not set.
func (h *Headers) IsDryRun() bool {
	if value, ok := h.Values[HeaderDryRun]; ok && value != nil {
		return value.(bool)
	}
	return false
}

// Origin returns the 'origin' header value or empty string if not set.
func (h *Headers) Origin() string {
	if value, ok := h.Values[HeaderOrigin]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// Originator returns the 'ditto-originator' header value or empty string if not set.
func (h *Headers) Originator() string {
	if value, ok := h.Values[HeaderOriginator]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// ETag returns the 'ETag' header value or empty string if not set.
func (h *Headers) ETag() string {
	if value, ok := h.Values[HeaderETag]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// IfMatch returns the 'If-Match' header value or empty string if not set.
func (h *Headers) IfMatch() string {
	if value, ok := h.Values[HeaderIfMatch]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// IfNoneMatch returns the 'If-None-Match' header value or empty string if not set.
func (h *Headers) IfNoneMatch() string {
	if value, ok := h.Values[HeaderIfNoneMatch]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// ReplyTarget returns the 'ditto-reply-target' header value or empty string if not set.
func (h *Headers) ReplyTarget() int64 {
	if value, ok := h.Values[HeaderReplyTarget]; ok && value != nil {
		return value.(int64)
	}
	return 0
}

// ReplyTo returns the 'reply-to' header value or empty string if not set.
func (h *Headers) ReplyTo() string {
	if value, ok := h.Values[HeaderReplyTo]; ok && value != nil {
		return value.(string)
	}
	return ""
}

// Version returns the 'version' header value or empty string if not set.
func (h *Headers) Version() int64 {
	if value, ok := h.Values[HeaderSchemaVersion]; ok && value != nil {
		return value.(int64)
	}
	return 0
}

// ContentType returns the 'content-type' header value or empty string if not set.
func (h *Headers) ContentType() string {
	if value, ok := h.Values[HeaderContentType]; ok && value != nil {
		return value.(string)
	}
	return ""
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
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if value, ok := m[HeaderTimeout]; ok {
		if _, err := parseTimeout(value.(string)); err != nil {
			return err
		}
	}

	h.Values = m
	return nil
}
