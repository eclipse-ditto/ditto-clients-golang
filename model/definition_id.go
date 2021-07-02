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

// DefinitionID represents an ID of a given definition entity.
// Compliant with the Ditto specification it consists of a namespace, name and a version
// in the form of 'namespace:name:version'.
// The DefinitionID is used to declare a Thing's model also it is used
// in declare the different models a Feature represents via its properties.
type DefinitionID struct {
	Namespace string
	Name      string
	Version   string
}

const definitionIdTemplate = "%s:%s:%s"

var regexDefinitionID = regexp.MustCompile("^([^:]+):([^:]+):([^:]+)$")

// NewDefinitionIDFrom creates a new DefinitionID instance from a provided string in the form of 'namespace:name:version'.
// Returns nil if the provided string doesn't match the form.
func NewDefinitionIDFrom(full string) *DefinitionID {
	if !regexDefinitionID.MatchString(full) {
		return nil
	}
	elements := strings.SplitN(full, ":", 3)
	return &DefinitionID{Namespace: elements[0], Name: elements[1], Version: elements[2]}
}

// NewDefinitionID creates a new DefinitionID instance with the namespace, name and version provided.
func NewDefinitionID(namespace string, name string, version string) *DefinitionID {
	return &DefinitionID{Namespace: namespace, Name: name, Version: version}
}

// String provides the string representation of a DefinitionID in the Ditto's specified form of 'namespace:name:version'.
func (definitionId *DefinitionID) String() string {
	return fmt.Sprintf(definitionIdTemplate, definitionId.Namespace, definitionId.Name, definitionId.Version)
}

func (definitionId *DefinitionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(definitionId.String())
}

func (definitionId *DefinitionID) UnmarshalJSON(data []byte) error {
	var defIDString = ""

	if err := json.Unmarshal(data, &defIDString); err != nil {
		return err
	}

	if !regexDefinitionID.MatchString(defIDString) {
		return errors.New("Invalid DefinitionID: " + defIDString)
	}
	elements := strings.SplitN(defIDString, ":", 3)
	definitionId.Namespace = elements[0]
	definitionId.Name = elements[1]
	definitionId.Version = elements[2]
	return nil
}

// WithNamespace sets the provided namespace to the current DefinitionID instance.
func (definitionId *DefinitionID) WithNamespace(namespace string) *DefinitionID {
	definitionId.Namespace = namespace
	return definitionId
}

// WithName sets the provided name to the current DefinitionID instance.
func (definitionId *DefinitionID) WithName(name string) *DefinitionID {
	definitionId.Name = name
	return definitionId
}

// WithVersion sets the provided version to the current DefinitionID instance.
func (definitionId *DefinitionID) WithVersion(version string) *DefinitionID {
	definitionId.Version = version
	return definitionId
}
