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

// Thing represents the Thing entity model form the Ditto's specification.
// Things are very generic entities and are mostly used as a “handle” for multiple features belonging to this Thing.
type Thing struct {
	ID           *NamespacedID          `json:"thingId"`
	PolicyID     *NamespacedID          `json:"policyId,omitempty"`
	DefinitionID *DefinitionID          `json:"definitionId,omitempty"`
	Attributes   map[string]interface{} `json:"attributes,omitempty"`
	Features     map[string]*Feature    `json:"features,omitempty"`
	Revision     int64                  `json:"revision,omitempty"`
	Timestamp    string                 `json:"timestamp,omitempty"`
}

// WithID sets the provided NamespacedID as the current Thing's instance ID value.
func (thing *Thing) WithID(id *NamespacedID) *Thing {
	thing.ID = id
	return thing
}

// WithIDFrom is an auxiliary method that sets the ID value of the current Thing instance based on the provided string n the form of 'namespace:name'.
func (thing *Thing) WithIDFrom(id string) *Thing {
	thing.ID = NewNamespacedIDFrom(id)
	return thing
}

// WithDefinition sets the current Thing instance's definition to the provided value.
func (thing *Thing) WithDefinition(definition *DefinitionID) *Thing {
	thing.DefinitionID = definition
	return thing
}

// WithDefinitionFrom is an auxiliary method to set the current Thing instance's definition to the provided one in the form of 'namespace:name:version'.
func (thing *Thing) WithDefinitionFrom(definitionId string) *Thing {
	thing.DefinitionID = NewDefinitionIDFrom(definitionId)
	return thing
}

// WithPolicyID sets the provided Policy ID to the current Thing instance.
func (thing *Thing) WithPolicyID(policyID *NamespacedID) *Thing {
	thing.PolicyID = policyID
	return thing
}

// WithPolicyIDFrom is an auxiliary method that sets the Policy ID of the current Thing instance
// from a string NamespacedID representation in the form of 'namespace:name'.
func (thing *Thing) WithPolicyIDFrom(policyId string) *Thing {
	thing.PolicyID = NewNamespacedIDFrom(policyId)
	return thing
}

// WithAttributes sets all attributes to the current Thing instance.
func (thing *Thing) WithAttributes(attrs map[string]interface{}) *Thing {
	thing.Attributes = attrs
	return thing
}

// WithAttribute sets/add an attribute to the current Thing instance.
func (thing *Thing) WithAttribute(id string, value interface{}) *Thing {
	if thing.Attributes == nil {
		thing.Attributes = make(map[string]interface{})
	}
	thing.Attributes[id] = value
	return thing
}

// WithFeatures sets all features to the current Thing instance.
func (thing *Thing) WithFeatures(features map[string]*Feature) *Thing {
	thing.Features = features
	return thing
}

// WithFeature sets/adds a Feature to the current features set of the Thing instance.
func (thing *Thing) WithFeature(id string, value *Feature) *Thing {
	if thing.Features == nil {
		thing.Features = make(map[string]*Feature)
	}
	thing.Features[id] = value
	return thing
}
