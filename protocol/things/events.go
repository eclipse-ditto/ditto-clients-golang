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

package things

import (
	"fmt"

	"github.com/eclipse/ditto-clients-golang/model"
	"github.com/eclipse/ditto-clients-golang/protocol"
)

// Event represents a message entity defined by the Ditto protocol for the Things group that defines a notification for a change that happened.
// This is a special Message that is always bound to a specific Thing instance along with providing the capabilities to configure:
// - the type of the change that happened - Created, Modified, Deleted
// - the channel used for the notification - Twin, Live
// - the entity that was affected - the whole Thing (the default), all features of the Thing (Features),
//                               a single Feature of the Thing (Feature), all attributes of the Thing (Attributes) or
//                               a single attribute of the Thing (Attribute), the Thing's policy (PolicyID)
//                               or the Thing's definition (Definition).
// Note: Only one change type can be configured to the event - if using the methods for configuring it - only the last one applies.
// Note: Only one channel can be configured to the event - if using the methods for configuring it - only the last one applies.
// Note: Only one entity that will b affected by the event can be configured - if using the methods for configuring it - only the last one applies.
type Event struct {
	Topic   *protocol.Topic
	Path    string
	Payload interface{}
}

// NewEvent creates a new Event instance for the defined by the provided NamespacedID Thing.
func NewEvent(thingID *model.NamespacedID) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithNamespace(thingID.Namespace).
			WithEntityName(thingID.Name).
			WithGroup(protocol.GroupThings).
			WithChannel(protocol.ChannelTwin).
			WithCriterion(protocol.CriterionEvents),
		Path: pathThing,
	}
}

// Created configures the Event to notify for a Thing that has been created using the provided payload instance.
func (event *Event) Created(thing *model.Thing) *Event {
	event.Topic.WithAction(protocol.ActionCreated)
	event.Payload = thing
	return event
}

// Modified configures the Event to notify for a modification with a new value applied defined by the provided payload.
func (event *Event) Modified(payload interface{}) *Event {
	event.Topic.WithAction(protocol.ActionModified)
	event.Payload = payload
	return event
}

// Merged configures the Event to notify for a modification with a merge patch defined by the provided payload.
func (event *Event) Merged(payload interface{}) *Event {
	event.Topic.WithAction(protocol.ActionMerged)
	event.Payload = payload
	return event
}

// Deleted configures the Event to notify for a deletion of a Thing or parts of the content it holds.
func (event *Event) Deleted() *Event {
	event.Topic.WithAction(protocol.ActionDeleted)
	return event
}

// PolicyID configures the Event to notify for a change in the Thing's policy.
func (event *Event) PolicyID() *Event {
	event.Path = pathThingPolicyID
	return event
}

// Definition configures the Event to notify for a change in the Thing's definition.
func (event *Event) Definition() *Event {
	event.Path = pathThingDefinition
	return event
}

// Attributes configures the Event to notify for a change in the Thing's attributes.
func (event *Event) Attributes() *Event {
	event.Path = pathThingAttributes
	return event
}

// Attribute configures the Event to notify for a change in the Thing's attribute
// defined by the provided attributePath as JSON pointer path (https://tools.ietf.org/html/rfc6901).
func (event *Event) Attribute(attributePath string) *Event {
	event.Path = fmt.Sprintf(pathThingAttributeFormat, attributePath)
	return event
}

// Features configures the Event to notify for a change in the Thing's features.
func (event *Event) Features() *Event {
	event.Path = pathThingFeatures
	return event
}

// Feature configures the Event to notify for a change in the Thing's feature defined by the provided featureID.
func (event *Event) Feature(featureID string) *Event {
	event.Path = fmt.Sprintf(pathThingFeatureFormat, featureID)
	return event
}

// FeatureDefinition configures the Event to notify for a change in the Thing's feature's definition for the feature
// defined by the provided featureID.
func (event *Event) FeatureDefinition(featureID string) *Event {
	event.Path = fmt.Sprintf(pathThingFeatureDefinitionFormat, featureID)
	return event
}

// FeatureProperties configures the Event to notify for a change in the Thing's feature's properties of the feature
// defined by the provided featureID.
func (event *Event) FeatureProperties(featureID string) *Event {
	event.Path = fmt.Sprintf(pathThingFeaturePropertiesFormat, featureID)
	return event
}

// FeatureProperty configures the Event to notify for a change in the Thing's feature's property
// defined by the provided featureID and propertyPath as JSON pointer path (https://tools.ietf.org/html/rfc6901).
func (event *Event) FeatureProperty(featureID, propertyPath string) *Event {
	event.Path = fmt.Sprintf(pathThingFeaturePropertyFormat, featureID, propertyPath)
	return event
}

// FeatureDesiredProperties configures the Event to notify for a change in the Thing's feature's desired properties
// of the feature defined by the provided featureID.
func (event *Event) FeatureDesiredProperties(featureID string) *Event {
	event.Path = fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, featureID)
	return event
}

// FeatureDesiredProperty configures the Event to notify for a change in the Thing's feature's desired property
// defined by the provided featureID and propertyPath as JSON pointer path (https://tools.ietf.org/html/rfc6901).
func (event *Event) FeatureDesiredProperty(featureID, propertyPath string) *Event {
	event.Path = fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, featureID, propertyPath)
	return event
}

// Live configures the channel of the Event accordingly.
func (event *Event) Live() *Event {
	event.Topic.WithChannel(protocol.ChannelLive)
	return event
}

// Twin configures the channel of the Event accordingly.
func (event *Event) Twin() *Event {
	event.Topic.WithChannel(protocol.ChannelTwin)
	return event
}

// Envelope generates the Ditto envelope with event's data applying all configurations and optionally all Headers provided.
func (event *Event) Envelope(headerOpts ...protocol.HeaderOpt) *protocol.Envelope {
	msg := &protocol.Envelope{
		Topic: event.Topic,
		Path:  event.Path,
		Value: event.Payload,
	}
	if headerOpts != nil {
		msg.Headers = protocol.NewHeaders(headerOpts...)
	}
	return msg
}
