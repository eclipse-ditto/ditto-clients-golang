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
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	// ContentTypeDitto defines the Ditto JSON 'content-type' header value for Ditto Protocol messages.
	ContentTypeDitto = "application/vnd.eclipse.ditto+json"

	// ContentTypeJSON defines the JSON 'content-type' header value for Ditto Protocol messages.
	ContentTypeJSON = "application/json"

	// ContentTypeJSONMerge defines the JSON merge patch 'content-type' header value for Ditto Protocol messages,
	// as specified with RFC 7396 (https://datatracker.ietf.org/doc/html/rfc7396).
	ContentTypeJSONMerge = "application/merge-patch+json"

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

	// HeaderETag represents 'etag' header.
	HeaderETag = "etag"

	// HeaderIfMatch represents 'if-match' header.
	HeaderIfMatch = "if-match"

	// HeaderIfNoneMatch represents 'if-none-march' header.
	HeaderIfNoneMatch = "if-none-match"

	// HeaderReplyTarget represents 'ditto-reply-target' header.
	HeaderReplyTarget = "ditto-reply-target"

	// HeaderReplyTo represents 'reply-to' header.
	HeaderReplyTo = "reply-to"

	// HeaderTimeout represents 'timeout' header.
	HeaderTimeout = "timeout"

	// HeaderVersion represents 'version' header.
	HeaderVersion = "version"

	// HeaderContentType represents 'content-type' header.
	HeaderContentType = "content-type"
)

// Headers represents all Ditto-specific headers along with additional HTTP/etc. Headers
// that can be applied depending on the transport used.
//
// The header values in this map should be serialized.
// The provided getter methods returns the header values which is associated with this definition's key.
// See https://www.eclipse.org/ditto/protocol-specification.html
type Headers map[string]interface{}

// CorrelationID returns the HeaderCorrelationID header value if it is presented.
//
// If there is no HeaderCorrelationID value, but there is at least one value which key differs only in capitalization,
// the CorrelationID returns the value corresponding to the first such key(sorted in increasing order).
//
// If there is no match about for this header, the CorrelationID will generate HeaderCorrelationID value in UUID format.
//
// If the type of the HeaderCorrelationID header (or the first met header) is not a string, the CorrelationID returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) CorrelationID() string {
	if value, ok := h[HeaderCorrelationID]; ok {
		return getStr(value, "")
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderCorrelationID) {
			return getStr(h[k], "")
		}
	}
	h[HeaderCorrelationID] = uuid.New().String()
	return h[HeaderCorrelationID].(string)
}

// Timeout returns the HeaderTimeout header value if it is presented.
// The default and maximum value is duration of 60 seconds.
//
// If there is no HeaderTimeout value, but there is at least one value which key differs only in capitalization,
// the Timeout returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderTimeout header (or the first met header) is not a string, the Timeout returns the default value.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) Timeout() time.Duration {
	if value, ok := h[HeaderTimeout]; ok {
		return h.timeoutValue(value)
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderTimeout) {
			return h.timeoutValue(h[k])
		}
	}
	return 60 * time.Second
}

// IsResponseRequired returns the HeaderResponseRequired header value if it is presented.
// The default value is true.
//
// If there is no HeaderResponseRequired value, but there is at least one value which key differs only in capitalization,
// the IsResponseRequired returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderResponseRequired header (or the first met header) is not a bool, the IsResponseRequired
// returns the default value.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) IsResponseRequired() bool {
	if value, ok := h[HeaderResponseRequired]; ok {
		return h.boolValue(value, true)
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderResponseRequired) {
			return h.boolValue(h[k], true)
		}
	}
	return true
}

// Channel returns the HeaderChannel header value.
//
// If there is no HeaderChannel value, but there is at least one value which key differs only in capitalization,
// the Channel returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderChannel header (or the first met header) is not a string, the Channel returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) Channel() string {
	return h.stringValue(HeaderChannel, "")

}

// IsDryRun returns the HeaderDryRun header value if it is presented.
// The default value is false.
//
// If there is no HeaderDryRun value, but there is at least one value which key differs only in capitalization,
// the IsDryRun returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderDryRun header (or the first met header) is not a bool, the IsDryRun returns the default value.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) IsDryRun() bool {
	if value, ok := h[HeaderDryRun]; ok {
		return h.boolValue(value, false)
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderDryRun) {
			return h.boolValue(h[k], false)
		}
	}
	return false
}

// Origin returns the HeaderOrigin header value if it is presented.
//
// If there is no HeaderOrigin value, but there is at least one value which key differs only in capitalization,
// the Origin returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderOrigin header (or the first met header) is not a string, the Origin returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) Origin() string {
	return h.stringValue(HeaderOrigin, "")

}

// Originator returns the HeaderOriginator header value if it is presented.
//
// If there is no HeaderOriginator value, but there is at least one value which key differs only in capitalization,
// the Originator returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderOriginator header (or the first met header) is not a string, the Originator returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) Originator() string {
	return h.stringValue(HeaderOriginator, "")

}

// ETag returns the HeaderETag header value if it is presented.
//
// If there is no HeaderETag value, but there is at least one value which key differs only in capitalization,
// the ETag returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderETag header (or the first met header) is not a string, the ETag returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) ETag() string {
	return h.stringValue(HeaderETag, "")

}

// IfMatch returns the HeaderIfMatch header value if it is presented.
//
// If there is no HeaderIfMatch value, but there is at least one value which key differs only in capitalization,
// the IfMatch returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderIfMatch header (or the first met header) is not a string, the IfMatch returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) IfMatch() string {
	return h.stringValue(HeaderIfMatch, "")

}

// IfNoneMatch returns the HeaderIfNoneMatch header value if it is presented.
//
// If there is no HeaderIfNoneMatch value, but there is at least one value which key differs only in capitalization,
// the IfNoneMatch returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderIfNoneMatch header (or the first met header) is not a string, the IfNoneMatch returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) IfNoneMatch() string {
	return h.stringValue(HeaderIfNoneMatch, "")

}

// ReplyTarget returns the HeaderReplyTarget header value if it is presented.
//
// If there is no HeaderReplyTarget value, but there is at least one value which key differs only in capitalization,
// the ReplyTarget returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderReplyTarget header (or the first met header) is not an int64, the ReplyTarget returns 0.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) ReplyTarget() int64 {
	if value, ok := h[HeaderReplyTarget]; ok {
		return h.intValue(value, 0)
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderReplyTarget) {
			return h.intValue(h[k], 0)
		}
	}
	return 0
}

// ReplyTo returns the HeaderReplyTo header value if it is presented.
//
// If there is no HeaderReplyTo value, but there is at least one value which key differs only in capitalization,
// the ReplyTo returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderReplyTo header (or the first met header) is not a string, the ReplyTo returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) ReplyTo() string {
	return h.stringValue(HeaderReplyTo, "")
}

// Version returns the HeaderVersion header value if it is presented.
// The default value is 2.
//
// If there is no HeaderVersion value, but there is at least one value which key differs only in capitalization,
// the Version returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderVersion header (or the first met header) is not an int 64, the Version returns the default value.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) Version() int64 {
	if value, ok := h[HeaderVersion]; ok {
		return h.intValue(value, int64(2))
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderVersion) {
			return h.intValue(h[k], int64(2))
		}
	}
	return int64(2)
}

// ContentType returns the HeaderContentType header value if it is presented.
//
// If there is no HeaderContentType value, but there is at least one value which key differs only in capitalization,
// the ContentType returns the value corresponding to the first such key(sorted in increasing order).
//
// If the type of the HeaderContentType header (or the first met header) is not a string, the ContentType returns the empty string.
//
// Use Generic or access the map directly to get a value to a specific key in regard to capitalization.
func (h Headers) ContentType() string {
	return h.stringValue(HeaderContentType, "")
}

// Generic returns the value of the provided key header.
func (h Headers) Generic(id string) interface{} {
	return h[id]
}

func (h Headers) stringValue(headerKey, defValue string) string {
	if value, ok := h[headerKey]; ok {
		return getStr(value, defValue)
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, headerKey) {
			return getStr(h[k], defValue)
		}
	}
	return defValue
}

func getStr(value interface{}, defValue string) string {
	if str, ok := value.(string); ok {
		return str
	}
	return defValue
}

func (h Headers) timeoutValue(headerValue interface{}) time.Duration {
	if value, ok := headerValue.(string); ok {
		if duration, err := parseTimeout(value); err == nil {
			return duration
		}
	}
	return 60 * time.Second
}

func (h Headers) intValue(headerValue interface{}, defValue int64) int64 {
	if value, ok := headerValue.(int64); ok {
		return value
	}
	return defValue
}

func (h Headers) boolValue(headerValue interface{}, defValue bool) bool {
	if value, ok := headerValue.(bool); ok {
		return value
	}
	return defValue
}

func sortHeadersKey(h Headers) []string {
	keys := make([]string, len(h))
	i := 0
	for k := range h {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
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
		if t >= 0 && t <= 60*time.Second {
			return t, nil
		}
	}
	return 60 * time.Second, fmt.Errorf("invalid timeout '%s'", timeout)
}
