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
	"errors"
	"reflect"
	"testing"
)

func WithError() HeaderOpt {
	return func(headers *Headers) error {
		return errors.New("this is an error example")
	}
}

func TestApplyOptsHeader(t *testing.T) {
	tests := map[string]struct {
		opts    []HeaderOpt
		wantErr bool
	}{
		"test_apply_opts_header": {
			opts:    []HeaderOpt{WithChannel("somethingNow")},
			wantErr: false,
		},
		"test_apply_opts_header_error": {
			opts:    []HeaderOpt{WithError()},
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res := make(map[string]interface{})
			res[HeaderChannel] = "somethingBefore"
			headers := &Headers{res}
			err := applyOptsHeader(headers, testCase.opts...)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("applyOptsHeader() must rise an error")
				}
			} else if headers.Values[HeaderChannel] != "somethingNow" {
				t.Errorf("applyOptsHeader() Header want = \"somethingNow\" got %v", headers.Values[HeaderChannel])
			}
		})
	}
}

func TestNewHeaders(t *testing.T) {
	tests := map[string]struct {
		opts    []HeaderOpt
		wantErr bool
	}{
		"test_new_headers": {
			opts:    []HeaderOpt{WithChannel("someChannel")},
			wantErr: false,
		},
		"test_new_headers_error": {
			opts:    []HeaderOpt{WithError()},
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res := make(map[string]interface{})
			res[HeaderChannel] = "someChannel"
			want := &Headers{res}
			got := NewHeaders(testCase.opts...)
			if testCase.wantErr {
				if got != nil {
					t.Errorf("NewHeaders() must be nil")
				}
			} else if !reflect.DeepEqual(got, want) {
				t.Errorf("NewHeaders() = %v want %v", got, want)
			}
		})
	}
}

func TestWithCorrelationID(t *testing.T) {
	t.Run("TestWithCorrelationID", func(t *testing.T) {
		cid := "correlationId"
		if got := NewHeaders(WithCorrelationID(cid)); !reflect.DeepEqual(got.CorrelationID(), cid) {
			t.Errorf("WithCorrelationID() = %v, want %v", got.CorrelationID(), cid)
		}
	})
}

func TestWithReplyTo(t *testing.T) {
	t.Run("TestWithReplyTo", func(t *testing.T) {
		rto := "replyto"
		if got := NewHeaders(WithReplyTo(rto)); !reflect.DeepEqual(got.ReplyTo(), rto) {
			t.Errorf("WithReplyTo() = %v, want %v", got.ReplyTo(), rto)
		}
	})
}

func TestWithReplyTarget(t *testing.T) {
	t.Run("TestWithReplyTarget", func(t *testing.T) {
		rt := "11111"
		if got := NewHeaders(WithReplyTarget(rt)); !reflect.DeepEqual(got.Values[HeaderReplyTarget], rt) {
			t.Errorf("WithReplyTarget() = %v, want %v", got.Values[HeaderReplyTarget], rt)
		}
	})
}

func TestWithChannel(t *testing.T) {
	t.Run("TestWithChannel", func(t *testing.T) {
		cha := "channel"
		if got := NewHeaders(WithChannel(cha)); !reflect.DeepEqual(got.Channel(), cha) {
			t.Errorf("WithChannel() = %v, want %v", got.Channel(), cha)
		}
	})
}

func TestWithResponseRequired(t *testing.T) {
	t.Run("TestWithResponseRequired", func(t *testing.T) {
		rrq := true
		if got := NewHeaders(WithResponseRequired(rrq)); !reflect.DeepEqual(got.IsResponseRequired(), rrq) {
			t.Errorf("WithResponseRequired() = %v, want %v", got.IsResponseRequired(), rrq)
		}
	})
}

func TestWithOriginator(t *testing.T) {
	t.Run("TestWithOriginator", func(t *testing.T) {
		org := "originator"
		if got := NewHeaders(WithOriginator(org)); !reflect.DeepEqual(got.Originator(), org) {
			t.Errorf("WithOriginator() = %v, want %v", got.Originator(), org)
		}
	})
}

func TestWithOrigin(t *testing.T) {
	t.Run("TestWithOrigin", func(t *testing.T) {
		org := "origin"
		if got := NewHeaders(WithOrigin(org)); !reflect.DeepEqual(got.Origin(), org) {
			t.Errorf("WithOrigin() = %v, want %v", got.Origin(), org)
		}
	})
}

func TestWithDryRun(t *testing.T) {
	t.Run("TestWithDryRun", func(t *testing.T) {
		dry := true
		if got := NewHeaders(WithDryRun(dry)); !reflect.DeepEqual(got.IsDryRun(), dry) {
			t.Errorf("WithDryRun() = %v, want %v", got.IsDryRun(), dry)
		}
	})
}

func TestWithETag(t *testing.T) {
	t.Run("TestWithETag", func(t *testing.T) {
		et := "etag"
		if got := NewHeaders(WithETag(et)); !reflect.DeepEqual(got.ETag(), et) {
			t.Errorf("WithETag() = %v, want %v", got.ETag(), et)
		}
	})
}

func TestWithIfMatch(t *testing.T) {
	t.Run("TestWithIfMatch", func(t *testing.T) {
		im := "ifMatch"
		if got := NewHeaders(WithIfMatch(im)); !reflect.DeepEqual(got.IfMatch(), im) {
			t.Errorf("WithIfMatch() = %v, want %v", got.IfMatch(), im)
		}
	})
}

func TestWithIfNoneMatch(t *testing.T) {
	t.Run("TestWithIfNoneMatch", func(t *testing.T) {
		inm := "ifNoneMatch"
		if got := NewHeaders(WithIfNoneMatch(inm)); !reflect.DeepEqual(got.IfNoneMatch(), inm) {
			t.Errorf("WithIfNoneMatch() = %v, want %v", got.IfNoneMatch(), inm)
		}
	})
}

func TestWithTimeout(t *testing.T) {
	t.Run("TestWithTimeout", func(t *testing.T) {
		tmo := "10"
		if got := NewHeaders(WithTimeout(tmo)); !reflect.DeepEqual(got.Timeout(), tmo) {
			t.Errorf("WithTimeout() = %v, want %v", got.Timeout(), tmo)
		}
	})
}

func TestWithSchemaVersion(t *testing.T) {
	t.Run("TestWithSchemaVersion", func(t *testing.T) {
		sv := "123456789"
		if got := NewHeaders(WithSchemaVersion("123456789")); !reflect.DeepEqual(got.Values[HeaderSchemaVersion], sv) {
			t.Errorf("WithSchemaVersion() = %v, want %v", got.Values[HeaderSchemaVersion], sv)
		}
	})
}

func TestWithContentType(t *testing.T) {
	t.Run("TestWithContentType", func(t *testing.T) {
		hct := "contentType"
		if got := NewHeaders(WithContentType(hct)); !reflect.DeepEqual(got.ContentType(), hct) {
			t.Errorf("WithContentType() = %v, want %v", got.ContentType(), hct)
		}
	})
}

func TestWithGeneric(t *testing.T) {
	t.Run("TestWithGeneric", func(t *testing.T) {
		hct := "contentType"
		if got := NewHeaders(WithGeneric(HeaderContentType, hct)); !reflect.DeepEqual(got.ContentType(), hct) {
			t.Errorf("WithGeneric() = %v, want %v", got.ContentType(), hct)
		}
	})
}
