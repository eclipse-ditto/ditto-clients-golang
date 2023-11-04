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

// WithCorrelationID sets the HeaderCorrelationID value.
//
// If there is no HeaderCorrelationID value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithCorrelationID(correlationID string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderCorrelationID, correlationID)
		return nil
	}
}

// WithReplyTo sets the HeaderReplyTo value.
//
// If there is no HeaderReplyTo value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithReplyTo(replyTo string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderReplyTo, replyTo)
		return nil
	}
}

// WithReplyTarget sets the HeaderReplyTarget value.
//
// If there is no HeaderReplyTarget value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithReplyTarget(replyTarget int64) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderReplyTarget, replyTarget)
		return nil
	}
}

// WithChannel sets the HeaderChannel value.
//
// If there is no HeaderChannel value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithChannel(channel string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderChannel, channel)
		return nil
	}
}

// WithResponseRequired sets the HeaderResponseRequired value.
//
// If there is no HeaderResponseRequired value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithResponseRequired(isResponseRequired bool) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderResponseRequired, isResponseRequired)
		return nil
	}
}

// WithOriginator sets the HeaderOriginator value.
//
// If there is no HeaderOriginator value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithOriginator(dittoOriginator string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderOriginator, dittoOriginator)
		return nil
	}
}

// WithOrigin sets the HeaderOrigin value.
//
// If there is no HeaderOrigin value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithOrigin(origin string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderOrigin, origin)
		return nil
	}
}

// WithDryRun sets the HeaderDryRun value.
//
// If there is no HeaderDryRun value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithDryRun(isDryRun bool) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderDryRun, isDryRun)
		return nil
	}
}

// WithETag sets the HeaderETag value.
//
// If there is no HeaderETag value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithETag(eTag string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderETag, eTag)
		return nil
	}
}

// WithIfMatch sets the HeaderIfMatch value.
//
// If there is no HeaderIfMatch value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithIfMatch(ifMatch string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderIfMatch, ifMatch)
		return nil
	}
}

// WithIfNoneMatch sets the HeaderIfNoneMatch value.
//
// If there is no HeaderIfNoneMatch value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithIfNoneMatch(ifNoneMatch string) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderIfNoneMatch, ifNoneMatch)
		return nil
	}
}

// WithTimeout sets the HeaderTimeout value.
//
// If there is no HeaderTimeout value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
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

// WithVersion sets the HeaderVersion value.
//
// If there is no HeaderVersion value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
func WithVersion(version int64) HeaderOpt {
	return func(headers Headers) error {
		setNewValue(headers, HeaderVersion, version)
		return nil
	}
}

// WithContentType sets the HeaderContentType value.
//
// If there is no HeaderContentType value, but there is at least one which key differs only in capitalization,
// than the value would be set to the first such key(sorted in increasing order).
//
// Use WithGeneric to set a value to a specific key in regard to capitalization.
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
