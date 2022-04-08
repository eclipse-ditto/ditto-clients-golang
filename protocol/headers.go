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

	// HeaderSchemaVersion represents 'version' header.
	HeaderSchemaVersion = "version"

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

// CorrelationID returns the 'correlation-id' header value if it is presented.
// If the header value is not presented, the 'correlation-id' header value will be generated in UUID format.
//
// If there are more than one headers differing only in capitalization, the CorrelationID returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) CorrelationID() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderCorrelationID) {
			return v.(string)
		}
	}
	h[HeaderCorrelationID] = uuid.New().String()
	return h[HeaderCorrelationID].(string)
}

// Timeout returns the 'timeout' header value if it is presented.
// The default and maximum value is duration of 60 seconds.
// If the header value is not presented, the Timout returns the default value.
//
// If there are more than one headers differing only in capitalization, the Timeout returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Timeout() time.Duration {
	for k := range h {
		if strings.EqualFold(k, HeaderTimeout) {
			if duration, err := parseTimeout(h[k].(string)); err == nil {
				return duration
			}
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
		if t >= 0 && t <= 60*time.Second {
			return t, nil
		}
	}
	return 60 * time.Second, fmt.Errorf("invalid timeout '%s'", timeout)
}

// IsResponseRequired returns the 'response-required' header value if it is presented.
// The default value is true.
// If the header value is not presented, the IsResponseRequired returns the default value.
//
// If there are more than one headers differing only in capitalization, the IsResponseRequired returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IsResponseRequired() bool {
	for k, v := range h {
		if strings.EqualFold(k, HeaderResponseRequired) {
			return v.(bool)
		}
	}
	return true
}

// Channel returns the 'ditto-channel' header value.
// If the header value is not presented, the Channel returns empty string.
//
// If there are more than one headers differing only in capitalization, the Channel returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Channel() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderChannel) {
			return v.(string)
		}
	}
	return ""
}

// IsDryRun returns the 'ditto-dry-run' header value if it is presented.
// The default value is false.
// If the header value is not presented, the IsDryRun returns the default value.
//
// If there are more than one headers differing only in capitalization, the IsDryRun returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IsDryRun() bool {
	for k, v := range h {
		if strings.EqualFold(k, HeaderDryRun) {
			return v.(bool)
		}
	}
	return false
}

// Origin returns the 'origin' header value if it is presented.
// If the header value is not presented, the Origin returns the empty string.
//
// If there are more than one headers differing only in capitalization, the Origin returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Origin() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderOrigin) {
			return v.(string)
		}
	}
	return ""
}

// Originator returns the 'ditto-originator' header value if it is presented.
// If the header value is not presented, the Originator returns the empty string.
//
// If there are more than one headers differing only in capitalization, the Originator returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Originator() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderOriginator) {
			return v.(string)
		}
	}
	return ""
}

// ETag returns the 'etag' header value if it is presented.
// If the header value is not presented, the ETag returns the empty string.
//
// If there are more than one headers differing only in capitalization, the ETag returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ETag() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderETag) {
			return v.(string)
		}
	}
	return ""
}

// IfMatch returns the 'if-match' header value if it is presented.
// If the header value is not presented, the IfMatch returns the empty string.
//
// If there are more than one headers differing only in capitalization, the IfMatch returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IfMatch() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderIfMatch) {
			return v.(string)
		}
	}
	return ""
}

// IfNoneMatch returns the 'if-none-match' header value if it is presented.
// If the header value is not presented, the IfNoneMatch returns the empty string.
//
// If there are more than one headers differing only in capitalization, the IfNoneMatch returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) IfNoneMatch() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderIfNoneMatch) {
			return v.(string)
		}
	}
	return ""
}

// ReplyTarget returns the 'ditto-reply-target' header value if it is presented.
// If the header value is not presented, the ReplyTarget returns 0.
//
// If there are more than one headers differing only in capitalization, the ReplyTarget returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ReplyTarget() int64 {
	for k, v := range h {
		if strings.EqualFold(k, HeaderReplyTarget) {
			return v.(int64)
		}
	}
	return 0
}

// ReplyTo returns the 'reply-to' header value if it is presented.
// If the header value is not presented, the ReplyTo returns the empty string.
//
// If there are more than one headers differing only in capitalization, the ReplyTo returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ReplyTo() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderReplyTo) {
			return v.(string)
		}
	}
	return ""
}

// Version returns the 'version' header value if it is presented.
// If the header value is not presented, the Version returns 2.
//
// If there are more than one headers differing only in capitalization, the Version returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Version() int64 {
	for k, v := range h {
		if strings.EqualFold(k, HeaderSchemaVersion) {
			return v.(int64)
		}
	}
	return 2
}

// ContentType returns the 'content-type' header value if it is presented.
// If the header value is not presented, the ContentType returns the empty string.
//
// If there are more than one headers differing only in capitalization, the ContentType returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) ContentType() string {
	for k, v := range h {
		if strings.EqualFold(k, HeaderContentType) {
			return v.(string)
		}
	}
	return ""
}

// Generic returns the first value of the provided key header. Capitalization of header names does not affect the header map.
// If there are no provided value, the Generic returns nil.
//
// If there are more than one headers differing only in capitalization Generic returns the first met value.
// To use the provided key to get the value, access the map directly.
func (h Headers) Generic(id string) interface{} {
	for k, v := range h {
		if strings.EqualFold(k, id) {
			return v
		}
	}
	return nil
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
