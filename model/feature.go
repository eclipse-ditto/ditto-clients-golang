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

// Feature represents the Feature entity defined by the Ditto's Things specification.
// It is used to manage all data and functionality of a Thing that can be clustered in an outlined technical context.
type Feature struct {
	Definition        []*DefinitionID        `json:"definition,omitempty"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
	DesiredProperties map[string]interface{} `json:"desiredProperties,omitempty"`
}

// WithDefinitionFrom is an auxiliary method to set the Feature's definition from an array of strings converted into the proper DefinitionID instances.
func (feature *Feature) WithDefinitionFrom(definition ...string) *Feature {
	if definition != nil {
		feature.Definition = make([]*DefinitionID, len(definition))
		for i, def := range definition {
			feature.Definition[i] = NewDefinitionIDFrom(def)
		}
	}
	return feature
}

// WithDesiredProperties sets all desired properties of the current Feature instance.
func (feature *Feature) WithDesiredProperties(properties map[string]interface{}) *Feature {
	feature.DesiredProperties = properties
	return feature
}

// WithDesiredProperty sets/adds a desired property to the current Feature instance.
func (feature *Feature) WithDesiredProperty(id string, value interface{}) *Feature {
	if feature.DesiredProperties == nil {
		feature.DesiredProperties = make(map[string]interface{})
	}
	feature.DesiredProperties[id] = value
	return feature
}

// WithDefinition sets the definition of the current Feature instance to the provided set of DefinitionIDs.
func (feature *Feature) WithDefinition(definition ...*DefinitionID) *Feature {
	feature.Definition = definition
	return feature
}

// WithProperties sets all properties of the current Feature instance.
func (feature *Feature) WithProperties(properties map[string]interface{}) *Feature {
	feature.Properties = properties
	return feature
}

// WithProperty sets/adds a property to the current Feature instance.
func (feature *Feature) WithProperty(id string, value interface{}) *Feature {
	if feature.Properties == nil {
		feature.Properties = make(map[string]interface{})
	}
	feature.Properties[id] = value
	return feature
}
