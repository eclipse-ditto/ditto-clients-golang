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
	"fmt"
	"strings"
)

// NamespacedID represents the namespaced entity ID defined by the Ditto specification.
// It is a unique identifier representing a Thing compliant with the Ditto requirements:
// - namespace and name separated by a : (colon)
// - have a maximum length of 256 characters.
type NamespacedID struct {
	Namespace string
	Name      string
}

// NewNamespacedID creates a new NamespacedID instance using the provided namespace and name.
func NewNamespacedID(namespace string, name string) *NamespacedID {
	return &NamespacedID{Namespace: namespace, Name: name}
}

// NewNamespacedIDFrom creates a new NamespacedID instance using the provided string in the valid form of 'namespace:name'.
func NewNamespacedIDFrom(full string) *NamespacedID {
	elements := strings.Split(full, ":")
	return &NamespacedID{Namespace: elements[0], Name: strings.Join(elements[1:], ":")}
}

// String provides the string representation of the NamespacedID entity in the form of 'namespace:name'.
func (nsID *NamespacedID) String() string {
	return fmt.Sprintf("%s:%s", nsID.Namespace, nsID.Name)
}

func (nsID *NamespacedID) MarshalJSON() ([]byte, error) {
	return json.Marshal(nsID.String())
}

func (nsID *NamespacedID) UnmarshalJSON(data []byte) error {
	var nsIDString = ""
	if err := json.Unmarshal(data, &nsIDString); err != nil {
		return err
	}
	elements := strings.Split(nsIDString, ":")
	nsID.Namespace = elements[0]
	nsID.Name = strings.Join(elements[1:], ":")
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
