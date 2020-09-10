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

func WithCorrelationID(correlationId string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderCorrelationID] = correlationId
		return nil
	}
}
func WithReplyTo(replyTo string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderReplyTo] = replyTo
		return nil
	}
}
func WithReplyTarget(replyTarget string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderReplyTarget] = replyTarget
		return nil
	}
}

func WithChannel(channel string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderChannel] = channel
		return nil
	}
}

func WithResponseRequired(isResponseRequired bool) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderResponseRequired] = isResponseRequired
		return nil
	}
}

func WithOriginator(dittoOriginator string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderOriginator] = dittoOriginator
		return nil
	}
}

func WithOrigin(origin string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderOrigin] = origin
		return nil
	}
}
func WithDryRun(isDryRun bool) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderDryRun] = isDryRun
		return nil
	}
}

func WithETag(eTag string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderETag] = eTag
		return nil
	}
}

func WithIfMatch(ifMatch string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderIfMatch] = ifMatch
		return nil
	}
}

func WithIfNoneMatch(ifNoneMatch string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderIfNoneMatch] = ifNoneMatch
		return nil
	}
}

func WithTimeout(timeout string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderTimeout] = timeout
		return nil
	}
}
func WithSchemaVersion(schemaVersion string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderSchemaVersion] = schemaVersion
		return nil
	}
}

func WithContentType(contentType string) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[HeaderContentType] = contentType
		return nil
	}
}

func WithGeneric(headerID string, value interface{}) HeaderOpt {
	return func(headers *Headers) error {
		headers.Values[headerID] = value
		return nil
	}
}
