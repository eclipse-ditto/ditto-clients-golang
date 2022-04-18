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
	"strings"
	"time"
)

// HeaderOpt represents a specific Headers option that can be applied to the Headers instance
// resulting in changing the value of a specific header of a set of headers.
type HeaderOpt func(headers Headers) error

func applyOptsHeader(headers Headers, opts ...HeaderOpt) error {
	for _, o := range opts {
		if err := o(headers); err != nil {
			return err
		}
	}
	return nil
}

// NewHeaders returns a new Headers instance.
func NewHeaders(opts ...HeaderOpt) Headers {
	res := Headers{}
	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

// NewHeadersFrom returns a new Headers instance using the provided header.
func NewHeadersFrom(orig Headers, opts ...HeaderOpt) Headers {
	if orig == nil {
		return NewHeaders(opts...)
	}
	res := Headers{}

	for key, value := range orig {
		res[key] = value
	}

	if err := applyOptsHeader(res, opts...); err != nil {
		return nil
	}
	return res
}

// WithCorrelationID sets a new value for header key 'correlation-id' if it is provided.
//
// If header key 'correlation-id' is not provided and there are more than one headers for 'correlation-id'
// differing only in capitalization, WithCorrelationID sets a new value for the first met header.
//
// If there aren't any headers for 'correlation-id', WithCorrelationID sets a new header with
// key 'correlation-id' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithCorrelationID(correlationID string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderCorrelationID, correlationID)
		return nil
	}
}

// WithReplyTo sets a new value for header key 'reply-to' if it is provided.
//
// If header key 'reply-to' is not provided and there are more than one headers for 'reply-to'
// differing only in capitalization, WithReplyTo sets a new value for the first met header.
//
// If there aren't any headers for 'reply-to', WithReplyTo sets a new header with
// key 'reply-to' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithReplyTo(replyTo string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderReplyTo, replyTo)
		return nil
	}
}

// WithReplyTarget sets a new value for header key 'ditto-reply-target' if it is provided.
//
// If header key 'ditto-reply-target' is not provided and there are more than one headers for 'ditto-reply-target'
// differing only in capitalization, WithReplyTarget sets a new value for the first met header.
//
// If there aren't any headers for 'ditto-reply-target', WithReplyTarget sets a new header with
// key 'ditto-reply-target' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithReplyTarget(replyTarget int64) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderReplyTarget, replyTarget)
		return nil
	}
}

// WithChannel sets a new value for header key 'ditto-channel' if it is provided.
//
// If header key 'ditto-channel' is not provided and there are more than one headers for 'ditto-channel'
// differing only in capitalization, WithChannel sets a new value for the first met header.
//
// If there aren't any headers for 'ditto-channel', WithChannel sets a new header with
// key 'ditto-channel' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithChannel(channel string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderChannel, channel)
		return nil
	}
}

// WithResponseRequired sets a new value for header key 'response-required' if it is provided.
//
// If header key 'response-required' is not provided and there are more than one headers for 'response-required'
// differing only in capitalization, WithResponseRequired sets a new value for the first met header.
//
// If there aren't any headers for 'response-required', WithResponseRequired sets a new header with
// key 'response-required' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithResponseRequired(isResponseRequired bool) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderResponseRequired, isResponseRequired)
		return nil
	}
}

// WithOriginator sets a new value for header key 'ditto-originator' if it is provided.
//
// If header key 'ditto-originator' is not provided and there are more than one headers for 'ditto-originator'
// differing only in capitalization, WithOriginator sets a new value for the first met header.
//
// If there aren't any headers for 'ditto-originator', WithOriginator sets a new header with
// key 'ditto-originator' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithOriginator(dittoOriginator string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderOriginator, dittoOriginator)
		return nil
	}
}

// WithOrigin sets a new value for header key 'origin' if it is provided.
//
// If header key 'origin' is not provided and there are more than one headers for 'origin'
// differing only in capitalization, WithOrigin sets a new value for the first met header.
//
// If there aren't any headers for 'origin', WithOrigin sets a new header with
// key 'origin' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithOrigin(origin string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderOrigin, origin)
		return nil
	}
}

// WithDryRun sets a new value for header key 'ditto-dry-run' if it is provided.
//
// If header key 'ditto-dry-run' is not provided and there are more than one headers for 'ditto-dry-run'
// differing only in capitalization, WithDryRun sets a new value for the first met header.
//
// If there aren't any headers for 'ditto-dry-run', WithDryRun sets a new header with
// key 'ditto-dry-run' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithDryRun(isDryRun bool) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderDryRun, isDryRun)
		return nil
	}
}

// WithETag sets a new value for header key 'etag' if it is provided.
//
// If header key 'etag' is not provided and there are more than one headers for 'etag'
// differing only in capitalization, WithETag sets a new value for the first met header.
//
// If there aren't any headers for 'etag', WithETag sets a new header with
// key 'etag' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithETag(eTag string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderETag, eTag)
		return nil
	}
}

// WithIfMatch sets a new value for header key 'if-match' if it is provided.
//
// If header key 'if-match' is not provided and there are more than one headers for 'if-match'
// differing only in capitalization, WithIfMatch sets a new value for the first met header.
//
// If there aren't any headers for 'if-match', WithIfMatch sets a new header with
// key 'if-match' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithIfMatch(ifMatch string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderIfMatch, ifMatch)
		return nil
	}
}

// WithIfNoneMatch sets a new value for header key 'if-none-match' if it is provided.
//
// If header key 'if-none-match' is not provided and there are more than one headers for 'if-none-match'
// differing only in capitalization, WithIfNoneMatch sets a new value for the first met header.
//
// If there aren't any headers for 'if-none-match', WithIfNoneMatch sets a new header with
// key 'if-none-match' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithIfNoneMatch(ifNoneMatch string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderIfNoneMatch, ifNoneMatch)
		return nil
	}
}

// WithTimeout sets a new value for header key 'timeout' if it is provided.
//
// If header key 'timeout' is not provided and there are more than one headers for 'timeout'
// differing only in capitalization, WithTimeout sets a new value for the first met header.
//
// If there aren't any headers for 'timeout', WithTimeout sets a new header with
// key 'timeout' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithTimeout(timeout time.Duration) HeaderOpt {
	return func(headers Headers) error {
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
		setNewValue(headers, HeaderTimeout, value)
		return nil
	}
}

// WithVersion sets a new value for header key 'version' if it is provided.
//
// If header key 'version' is not provided and there are more than one headers for 'version'
// differing only in capitalization, WithVersion sets a new value for the first met header.
//
// If there aren't any headers for 'version', WithVersion sets a new header with
// key 'version' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithVersion(version int64) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderVersion, version)
		return nil
	}
}

// WithContentType sets a new value for header key 'content-type' if it is provided.
//
// If header key 'content-type' is not provided and there are more than one headers for 'content-type'
// differing only in capitalization, WithContentType sets a new value for the first met header.
//
// If there aren't any headers for 'content-type', WithContentType sets a new header with
// key 'content-type' and the provided value.
//
// To use the provided key to set a new value, access the map directly.
func WithContentType(contentType string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderContentType, contentType)
		return nil
	}
}

// WithGeneric sets the value of the provided key header.
func WithGeneric(headerID string, value interface{}) HeaderOpt {
	return func(headers Headers) error {
		headers[headerID] = value
		return nil
	}
}

func setNewValue(headers Headers, headerKey string, headerValue interface{}) {
	if _, ok := headers[headerKey]; ok {
		headers[headerKey] = headerValue
		return
	}
	keys := sortHeadersKey(headers)
	for _, k := range keys {
		if strings.EqualFold(k, headerKey) {
			headers[k] = headerValue
			return
		}
	}
	headers[headerKey] = headerValue
}
