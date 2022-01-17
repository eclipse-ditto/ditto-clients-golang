// Copyright (c) 2021 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// http://www.eclipse.org/legal/epl-2.0
//
// SPDX-License-Identifier: EPL-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eclipse/paho.mqtt.golang (interfaces: Token)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockToken is a mock of Token interface.
type MockToken struct {
	ctrl     *gomock.Controller
	recorder *MockTokenMockRecorder
}

// MockTokenMockRecorder is the mock recorder for MockToken.
type MockTokenMockRecorder struct {
	mock *MockToken
}

// NewMockToken creates a new mock instance.
func NewMockToken(ctrl *gomock.Controller) *MockToken {
	mock := &MockToken{ctrl: ctrl}
	mock.recorder = &MockTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToken) EXPECT() *MockTokenMockRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockToken) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockTokenMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockToken)(nil).Done))
}

// Error mocks base method.
func (m *MockToken) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockTokenMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockToken)(nil).Error))
}

// Wait mocks base method.
func (m *MockToken) Wait() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockTokenMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockToken)(nil).Wait))
}

// WaitTimeout mocks base method.
func (m *MockToken) WaitTimeout(arg0 time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitTimeout", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// WaitTimeout indicates an expected call of WaitTimeout.
func (mr *MockTokenMockRecorder) WaitTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitTimeout", reflect.TypeOf((*MockToken)(nil).WaitTimeout), arg0)
}
