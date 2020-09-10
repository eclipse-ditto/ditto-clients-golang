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

const (
	pathThing                        = "/"
	pathThingDefinition              = "/definition"
	pathThingPolicyID                = "/policyId"
	pathThingFeatures                = "/features"
	pathThingAttributes              = "/attributes"
	pathThingAttributeFormat         = pathThingAttributes + "/%s"
	pathThingFeatureFormat           = pathThingFeatures + "/%s"
	pathThingFeatureDefinitionFormat = pathThingFeatureFormat + "/definition"
	pathThingFeaturePropertiesFormat = pathThingFeatureFormat + "/properties"
	pathThingFeaturePropertyFormat   = pathThingFeaturePropertiesFormat + "/%s"
)

// Command represents a message entity defined by the Ditto protocol for the Things group that defines the execution of a certain action.
// This is a special Message that is always bound to a specific Thing instance along with providing the capabilities to configure:
// - the type of the action it will signal for execution - Create, Modify, Retrieve, Delete
// - the channel it will be sent - Twin, Live
// - the entity it will affect - the whole Thing (the default), all features of the Thing (Features),
//                               a single Feature of the Thing (Feature), all attributes of the Thing (Attributes) or
//                               a single attribute of the Thing (Attribute), the Thing's policy (PolicyID)
//                               or the Thing's definition (Definition).
// Note: Only one action can be configured to the command - if using the methods for configuring it - only the last one applies.
// Note: Only one channel can be configured to the command - if using the methods for configuring it - only the last one applies.
// Note: Only one entity that will b affected by the command can be configured - if using the methods for configuring it - only the last one applies.
type Command struct {
	Topic   *protocol.Topic
	Path    string
	Payload interface{}
}

// NewCommand creates a new Command instance for the defined by the provided NamespacedID Thing.
func NewCommand(thingID *model.NamespacedID) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithNamespace(thingID.Namespace).
			WithEntityID(thingID.Name).
			WithGroup(protocol.GroupThings).
			WithChannel(protocol.ChannelTwin).
			WithCriterion(protocol.CriterionCommands),
		Path: pathThing,
	}
}

// Create creates a new Thing entity based on the provided information.
func (cmd *Command) Create(thing *model.Thing) *Command {
	cmd.Topic.WithAction(protocol.ActionCreate)
	cmd.Payload = thing
	return cmd
}

// Modify sets the action of the command instance accordingly.
// The provided payload must be the new value to be used for modification
// compliant with the (part of) the Thing it is to be applied to.
func (cmd *Command) Modify(payload interface{}) *Command {
	cmd.Topic.WithAction(protocol.ActionModify)
	cmd.Payload = payload
	return cmd
}

// Retrieve sets the action of the command instance accordingly.
// If thingIDs are provided the response will contain the information for these Things only.
// Further Headers can be added via the Message method to adjust the response even more.
// The topic placeholder for the Thing ID's namespace and/or name can be used to perform the multiple Things request, e.g.:
// '_:_', '_:thing.id' are valid Thing NamespacedIDs to perform the multiple Things retrieve call.
func (cmd *Command) Retrieve(thingIDs ...model.NamespacedID) *Command {
	cmd.Topic.WithAction(protocol.ActionRetrieve)
	if thingIDs != nil && len(thingIDs) > 0 {
		var thingIDsStruct interface{}
		thingIDsArray := make([]string, len(thingIDs))
		for i, id := range thingIDs {
			thingIDsArray[i] = id.String()
		}
		thingIDsStruct = struct {
			ThingIDs []string `json:"thingIds"`
		}{
			ThingIDs: thingIDsArray,
		}
		cmd.Payload = thingIDsStruct
	}
	return cmd
}

// Delete sets the action of the command instance accordingly.
func (cmd *Command) Delete() *Command {
	cmd.Topic.WithAction(protocol.ActionDelete)
	return cmd
}

// PolicyID configures the command to affect the Thing's Policy.
func (cmd *Command) PolicyID() *Command {
	cmd.Path = pathThingPolicyID
	return cmd
}

// Definition configures the command to affect the Thing's definition.
func (cmd *Command) Definition() *Command {
	cmd.Path = pathThingDefinition
	return cmd
}

// Attributes configures the command to affect the Thing's attributes.
func (cmd *Command) Attributes() *Command {
	cmd.Path = pathThingAttributes
	return cmd
}

// Attribute configures the command to affect a specified by the provided attributeID attribute of the Thing.
func (cmd *Command) Attribute(attributeID string) *Command {
	cmd.Path = fmt.Sprintf(pathThingAttributeFormat, attributeID)
	return cmd
}

// Features configures the command to affect all the features of the Thing.
func (cmd *Command) Features() *Command {
	cmd.Path = pathThingFeatures
	return cmd
}

// Feature configures the command to affect a specified by the provided featureID feature of the Thing.
func (cmd *Command) Feature(featureID string) *Command {
	cmd.Path = fmt.Sprintf(pathThingFeatureFormat, featureID)
	return cmd
}

// FeatureDefinition configures the command to affect the definition of a specified by the provided featureID feature of the Thing.
func (cmd *Command) FeatureDefinition(featureID string) *Command {
	cmd.Path = fmt.Sprintf(pathThingFeatureDefinitionFormat, featureID)
	return cmd
}

// FeatureProperties configures the command to affect all properties of a specified by the provided featureID feature of the Thing.
func (cmd *Command) FeatureProperties(featureID string) *Command {
	cmd.Path = fmt.Sprintf(pathThingFeaturePropertiesFormat, featureID)
	return cmd
}

// FeatureProperty configures the command to affect a specified property via the provided propertyID of a specified by the provided featureID feature of the Thing.
func (cmd *Command) FeatureProperty(featureID, propertyID string) *Command {
	cmd.Path = fmt.Sprintf(pathThingFeaturePropertyFormat, featureID, propertyID)
	return cmd
}

// Live configures the channel of the command accordingly.
func (cmd *Command) Live() *Command {
	cmd.Topic.WithChannel(protocol.ChannelLive)
	return cmd
}

// Twin configures the channel of the command accordingly.
func (cmd *Command) Twin() *Command {
	cmd.Topic.WithChannel(protocol.ChannelTwin)
	return cmd
}

// Message generates the Ditto message applying all configurations and optionally all Headers provided.
func (cmd *Command) Message(headerOpts ...protocol.HeaderOpt) *protocol.Message {
	msg := &protocol.Message{
		Topic: cmd.Topic,
		Path:  cmd.Path,
		Value: cmd.Payload,
	}
	if headerOpts != nil {
		msg.Headers = protocol.NewHeaders(headerOpts...)
	}
	return msg
}
