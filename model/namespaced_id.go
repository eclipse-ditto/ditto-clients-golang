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

package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const namespacedIDTemplate = "%s:%s"

var regexNamespacedID = regexp.MustCompile("^(|(?:[a-zA-Z]\\w*)(?:[.\\-][a-zA-Z]\\w*)*):([^\\x00-\\x1F\\x7F-\\xFF/]+)$")

// NamespacedID represents the namespaced ID defined by the Ditto specification.
// It is a unique identifier representing a Thing compliant with the Ditto requirements:
// - namespace and name separated by a : (colon)
// - have a maximum length of 256 characters.
type NamespacedID struct {
	Namespace string
	Name      string
}

// NewNamespacedID creates a new NamespacedID instance using the provided namespace and name.
// Returns nil if the provided string doesn't match the form.
func NewNamespacedID(namespace string, name string) *NamespacedID {
	if strings.Contains(namespace, ":") {
		return nil
	}
	if _, err := isValidNamespacedID(fmt.Sprintf(namespacedIDTemplate, namespace, name)); err == nil {
		return &NamespacedID{Namespace: namespace, Name: name}
	}
	return nil
}

// NewNamespacedIDFrom creates a new NamespacedID instance using the provided string in the valid form of 'namespace:name'.
// Returns nil if the provided string doesn't match the form.
func NewNamespacedIDFrom(full string) *NamespacedID {
	if matches, err := isValidNamespacedID(full); err == nil {
		return &NamespacedID{Namespace: matches[1], Name: matches[2]}
	}
	return nil
}

// String provides the string representation of the NamespacedID entity in the form of 'namespace:name'.
func (nsID *NamespacedID) String() string {
	return fmt.Sprintf(namespacedIDTemplate, nsID.Namespace, nsID.Name)
}

// MarshalJSON marshals NamespacedID.
func (nsID *NamespacedID) MarshalJSON() ([]byte, error) {
	return json.Marshal(nsID.String())
}

// UnmarshalJSON unmarshals NamespacedID.
func (nsID *NamespacedID) UnmarshalJSON(data []byte) error {
	var nsIDString = ""
	if err := json.Unmarshal(data, &nsIDString); err != nil {
		return err
	}
	matches, err := isValidNamespacedID(nsIDString)
	if err != nil {
		return err
	}
	nsID.Namespace = matches[1]
	nsID.Name = matches[2]
	return nil
}

// WithNamespace sets the provided namespace to the current NamespacedID instance.
func (nsID *NamespacedID) WithNamespace(namespace string) *NamespacedID {
	nsID.Namespace = namespace
	return nsID
}

// WithName sets the provided name to the current NamespacedID instance.
func (nsID *NamespacedID) WithName(name string) *NamespacedID {
	nsID.Name = name
	return nsID
}

func isValidNamespacedID(nsIDString string) ([]string, error) {
	if len(nsIDString) > 256 {
		return nil, errors.New("length exceeds 256, invalid NamespacedID: " + nsIDString)
	}
	if matches := regexNamespacedID.FindStringSubmatch(nsIDString); len(matches) == 3 {
		return matches, nil
	}
	return nil, errors.New("invalid NamespacedID: " + nsIDString)
}
