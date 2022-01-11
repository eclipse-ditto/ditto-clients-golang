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

const (
	definitionIDTemplate     = "%s:%s:%s"
	definitionElementPattern = "([_a-zA-Z0-9\\-.]+)"
)

var regexDefinitionID = regexp.MustCompile("^" + fmt.Sprintf(definitionIDTemplate, definitionElementPattern, definitionElementPattern, definitionElementPattern) + "$")

// NewDefinitionIDFrom creates a new DefinitionID instance from a provided string in the form of 'namespace:name:version'.
// Returns nil if the provided string doesn't match the form.
func NewDefinitionIDFrom(full string) *DefinitionID {
	if matches, err := validateDefinitionID(full); err == nil {
		return &DefinitionID{Namespace: matches[1], Name: matches[2], Version: matches[3]}
	}
	return nil
}

// NewDefinitionID creates a new DefinitionID instance with the namespace, name and version provided.
// Returns nil if the provided string doesn't match the form.
func NewDefinitionID(namespace string, name string, version string) *DefinitionID {
	if _, err := validateDefinitionID(fmt.Sprintf(definitionIDTemplate, namespace, name, version)); err == nil {
		return &DefinitionID{Namespace: namespace, Name: name, Version: version}
	}
	return nil
}

// String provides the string representation of a DefinitionID in the Ditto's specified form of 'namespace:name:version'.
func (definitionID *DefinitionID) String() string {
	return fmt.Sprintf(definitionIDTemplate, definitionID.Namespace, definitionID.Name, definitionID.Version)
}

func (definitionID *DefinitionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(definitionID.String())
}

func (definitionID *DefinitionID) UnmarshalJSON(data []byte) error {
	var defIDString = ""

	if err := json.Unmarshal(data, &defIDString); err != nil {
		return err
	}

	matches, err := validateDefinitionID(defIDString)
	if err != nil {
		return err
	}

	definitionID.Namespace = matches[1]
	definitionID.Name = matches[2]
	definitionID.Version = matches[3]
	return nil
}

// WithNamespace sets the provided namespace to the current DefinitionID instance.
func (definitionID *DefinitionID) WithNamespace(namespace string) *DefinitionID {
	definitionID.Namespace = namespace
	return definitionID
}

// WithName sets the provided name to the current DefinitionID instance.
func (definitionID *DefinitionID) WithName(name string) *DefinitionID {
	definitionID.Name = name
	return definitionID
}

// WithVersion sets the provided version to the current DefinitionID instance.
func (definitionID *DefinitionID) WithVersion(version string) *DefinitionID {
	definitionID.Version = version
	return definitionID
}

func validateDefinitionID(defIDString string) ([]string, error) {
	if matches := regexDefinitionID.FindStringSubmatch(defIDString); len(matches) == 4 {
		return matches, nil
	}
	return nil, errors.New("invalid DefinitionID: " + defIDString)
}
