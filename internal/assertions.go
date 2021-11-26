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

package internal

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func AssertError(t *testing.T, expected error, actual error) {
	if expected == nil {
		if actual != nil {
			t.Errorf("expected nil , got %v", actual)
			t.Fail()
		}
	} else {
		if actual == nil {
			t.Errorf("expected %v , got nil", expected)
			t.Fail()
		} else {
			if expected.Error() != actual.Error() {
				t.Errorf("expected %v , got %v", expected, actual)
				t.Fail()
			}
		}
	}
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v , got %v", expected, actual)
		t.Fail()
	}
}

func AssertWithTimeout(t *testing.T, waitGroup *sync.WaitGroup, testTimeout time.Duration) {
	testWaitChan := make(chan struct{})
	go func() {
		defer close(testWaitChan)
		waitGroup.Wait()
	}()
	select {
	case <-testWaitChan:
		return // completed normally
	case <-time.After(testTimeout * time.Second):
		t.Fatal("timed out waiting for ", testTimeout)
	}
}
