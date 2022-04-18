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

// CorrelationID returns the 'correlation-id' header value if it is presented and true if the type is string.
// CorrelationID returns an empty string and false if the header is presented, but the type is not a string.
//
// If the header value is not presented, the 'correlation-id' header value will be generated in UUID format.
//
// If there are more than one headers differing only in capitalization, the CorrelationID returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) CorrelationID() (string, bool) {
	if value, ok := h[HeaderCorrelationID]; ok {
		return h.stringValue(value, "")
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderCorrelationID) {
			return h.stringValue(h[k], "")
		}
	}
	h[HeaderCorrelationID] = uuid.New().String()
	return h[HeaderCorrelationID].(string), true
}

// Timeout returns the 'timeout' header value if it is presented.
// The default and maximum value is duration of 60 seconds.
//
// If the header value is not presented, the Timout returns the default value.
// If the header value is presented, but the type is not a string or the value is not valid, the Timeout returns the default value.
//
// If there are more than one headers differing only in capitalization, the Timeout returns the first met value.
// To use the provided key to get the value, access the map directly.
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

// IsResponseRequired returns the 'response-required' header value if it is presented.
// The default value is true.
//
// If the header value is not presented, the IsResponseRequired returns the default value.
// If the header value is presented, but the type is not a bool, the IsResponseRequired returns the default value.
//
// If there are more than one headers differing only in capitalization, the IsResponseRequired returns the first met value.
// To use the provided key to get the value, access the map directly.
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

// Channel returns the 'ditto-channel' header value.
//
// If the header value is not presented, the Channel returns the empty string.
// If the header value is presented, but the type is not a string, the Cannel returns the empty string.
//
// If there are more than one headers differing only in capitalization, the Channel returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Channel() string {
	if value, ok := h[HeaderChannel]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderChannel) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// IsDryRun returns the 'ditto-dry-run' header value if it is presented.
// The default value is false.
//
// If the header value is not presented, the IsDryRun returns the default value.
// If the header value is presented, but the type is not a bool, the IsDryRun returns the default value.
//
// If there are more than one headers differing only in capitalization, the IsDryRun returns the first met value.
// To use the provided key to get the value, access the map directly.
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

// Origin returns the 'origin' header value if it is presented.
//
// If the header value is not presented, the Origin returns the empty string.
// If the header value is presented, but the value is not a string, the Origin returns the empty string.
//
// If there are more than one headers differing only in capitalization, the Origin returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Origin() string {
	if value, ok := h[HeaderOrigin]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderOrigin) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// Originator returns the 'ditto-originator' header value if it is presented.
//
// If the header value is not presented, the Originator returns the empty string.
// If the header value is presented, but the type is not a string, the Originator returns the empty string.
//
// If there are more than one headers differing only in capitalization, the Originator returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Originator() string {
	if value, ok := h[HeaderOriginator]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderOriginator) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// ETag returns the 'etag' header value if it is presented.
//
// If the header value is not presented, the ETag returns the empty string.
// If the header value is presented, but the type is not a string, the ETag returns the empty string.
//
// If there are more than one headers for 'etag' differing only in capitalization
// the ETag returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ETag() string {
	if value, ok := h[HeaderETag]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderETag) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// IfMatch returns the 'if-match' header value if it is presented.
//
// If the header value is not presented, the IfMatch returns the empty string.
// If the header value is presented, but the type is not a string, the IfMatch returns the empty string.
//
// If there are more than one headers differing only in capitalization, the IfMatch returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IfMatch() string {
	if value, ok := h[HeaderIfMatch]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderIfMatch) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// IfNoneMatch returns the 'if-none-match' header value if it is presented.
//
// If the header value is not presented, the IfNoneMatch returns the empty string.
// If the header value is presented, but the type is not a string, the IfNonMatch returns the empty string.
//
// If there are more than one headers differing only in capitalization, the IfNoneMatch returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IfNoneMatch() string {
	if value, ok := h[HeaderIfNoneMatch]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderIfNoneMatch) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// ReplyTarget returns the 'ditto-reply-target' header value if it is presented.
//
// If the header value is not presented, the ReplyTarget returns 0.
// If the header value is presented, but the type is not an int64, the ReplyTarget returns 0.
//
// If there are more than one headers differing only in capitalization, the ReplyTarget returns the first met value.
// To use the provided key to get the value, access the map directly.
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

// ReplyTo returns the 'reply-to' header value if it is presented.
//
// If the header value is not presented, the ReplyTo returns the empty string.
// If the header value is presented, but the type is not a sting, the ReplyTo returns the empty string.
//
// If there are more than one headers differing only in capitalization, the ReplyTo returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ReplyTo() string {
	if value, ok := h[HeaderReplyTo]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderReplyTo) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// Version returns the 'version' header value if it is presented.
// The default value is 2.
//
// If the header value is not presented, the Version returns the default value.
// If the header value is presented, but the type is not an int64, he Version returns the default value.
//
// If there are more than one headers differing only in capitalization, the Version returns the first met value.
// To use the provided key to get the value, access the map directly.
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

// ContentType returns the 'content-type' header value if it is presented.
//
// If the header value is not presented, the ContentType returns the empty string.
// If the header value is not presented, the ContentType returns the empty string.
//
// If there are more than one headers differing only in capitalization, the ContentType returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ContentType() string {
	if value, ok := h[HeaderContentType]; ok {
		str, _ := h.stringValue(value, "")
		return str
	}
	keys := sortHeadersKey(h)
	for _, k := range keys {
		if strings.EqualFold(k, HeaderContentType) {
			str, _ := h.stringValue(h[k], "")
			return str
		}
	}
	return ""
}

// Generic returns the value of the provided key header.
func (h Headers) Generic(id string) interface{} {
	return h[id]
}

// With sets new Headers to the existing.
func (h Headers) With(opts ...HeaderOpt) Headers {
	res := make(map[string]interface{})

	for key, value := range h {
		res[key] = value
	}

	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

func (h Headers) stringValue(headerValue interface{}, defValue string) (string, bool) {
	if value, ok := headerValue.(string); ok {
		return value, true
	}
	return defValue, false
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
	var keys []string
	for k := range h {
		keys = append(keys, k)
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
