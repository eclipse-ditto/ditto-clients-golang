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

package ditto

type (
	// Logger interface allows plugging of a logger implementation that
	// fits best the needs of the application that is to use the Ditto library.
	Logger interface {
		Println(v ...interface{})
		Printf(format string, v ...interface{})
	}

	// LoggerStub provides an empty default implementation.
	LoggerStub struct{}
)

// Println provides an empty default implementation for logging.
func (LoggerStub) Println(v ...interface{}) {}

// Printf provides an empty default implementation for formatted logging.
func (LoggerStub) Printf(format string, v ...interface{}) {}

// Levels of the library's output that can be configured during package initialization in init().
var (
	INFO  Logger = LoggerStub{}
	WARN  Logger = LoggerStub{}
	DEBUG Logger = LoggerStub{}
	ERROR Logger = LoggerStub{}
)
