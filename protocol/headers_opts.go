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
	"strconv"
	"time"
)

// HeaderOpt represents a specific Headers option that can be applied to the Headers instance
// resulting in changing the value of a specific header of a set of headers.
type HeaderOpt func(headers *Headers) error

func applyOptsHeader(headers *Headers, opts ...HeaderOpt) error {
	for _, o := range opts {
		if err := o(headers); err != nil {
			return err
		}
	}
	return nil
}

// NewHeaders returns a new Headers instance.
func NewHeaders(opts ...HeaderOpt) *Headers {
	res := &Headers{}
	res.Values = make(map[string]interface{})
	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

// NewHeadersFrom returns a new Headers instance using the provided header.
func NewHeadersFrom(orig *Headers, opts ...HeaderOpt) *Headers {
	if orig == nil {
		return NewHeaders(opts...)
	}
	res := &Headers{
		Values: make(map[string]interface{}),
	}
	for key, value := range orig.Values {
		res.Values[key] = value
	}
	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

// WithCorrelationID sets the 'correlation-id' header value.
func WithCorrelationID(correlationID string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderCorrelationID] = correlationID
		return nil
	}
}

// WithReplyTo sets the 'reply-to' header value.
func WithReplyTo(replyTo string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderReplyTo] = replyTo
		return nil
	}
}

// WithReplyTarget sets the 'ditto-reply-target' header value.
func WithReplyTarget(replyTarget string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderReplyTarget] = replyTarget
		return nil
	}
}

// WithChannel sets the 'ditto-channel' header value.
func WithChannel(channel string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderChannel] = channel
		return nil
	}
}

// WithResponseRequired sets the 'response-required' header value.
func WithResponseRequired(isResponseRequired bool) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderResponseRequired] = isResponseRequired
		return nil
	}
}

// WithOriginator sets the 'ditto-originator' header value.
func WithOriginator(dittoOriginator string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderOriginator] = dittoOriginator
		return nil
	}
}

// WithOrigin sets the 'origin' header value.
func WithOrigin(origin string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderOrigin] = origin
		return nil
	}
}

// WithDryRun sets the 'ditto-dry-run' header value.
func WithDryRun(isDryRun bool) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderDryRun] = isDryRun
		return nil
	}
}

// WithETag sets the 'ETag' header value.
func WithETag(eTag string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderETag] = eTag
		return nil
	}
}

// WithIfMatch sets the 'If-Match' header value.
func WithIfMatch(ifMatch string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderIfMatch] = ifMatch
		return nil
	}
}

// WithIfNoneMatch sets the 'If-None-Match' header value.
func WithIfNoneMatch(ifNoneMatch string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderIfNoneMatch] = ifNoneMatch
		return nil
	}
}

// WithTimeout sets the 'timeout' header value.
//
// The provided value should be a non-negative duration in Millisecond, Second or Minute unit.
// The change results in timeout header string value containing the duration
// and the unit symbol (ms, s or m), e.g. '45s' or '250ms' or '1m'.
//
// The default value is '60s'.
// If a negative duration or duration of an hour or more is provided, the timeout header value
// is removed, i.e. the default one is used.
func WithTimeout(timeout time.Duration) HeaderOpt {
	return func(headers *Headers) error {
		if inTimeoutRange(timeout) {
			var value string

			if timeout > time.Second {
				div := int64(timeout / time.Second)
				rem := timeout % time.Second
				if rem == 0 {
					value = strconv.FormatInt(div, 10)
				} else {
					value = strconv.FormatInt(div+1, 10)
				}
			} else {
				div := int64(timeout / time.Millisecond)
				rem := timeout % time.Millisecond
				if rem == 0 {
					value = strconv.FormatInt(div, 10) + "ms"
				} else {
					value = strconv.FormatInt(div+1, 10) + "ms"
				}
			}

			headers.Values[HeaderTimeout] = value
		} else {
			delete(headers.Values, HeaderTimeout)
		}
		return nil
	}
}

// WithSchemaVersion sets the 'version' header value.
func WithSchemaVersion(schemaVersion string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderSchemaVersion] = schemaVersion
		return nil
	}
}

// WithContentType sets the 'content-type' header value.
func WithContentType(contentType string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderContentType] = contentType
		return nil
	}
}

// WithGeneric sets the value of the provided key header.
func WithGeneric(headerID string, value interface{}) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[headerID] = value
		return nil
	}
}
