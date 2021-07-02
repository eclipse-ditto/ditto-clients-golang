# Eclipse Ditto Client SDK for Golang

[![Join the chat at https://gitter.im/eclipse/ditto](https://badges.gitter.im/eclipse/ditto.svg)](https://gitter.im/eclipse/ditto?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![License](https://img.shields.io/badge/License-EPL%202.0-green.svg)](https://opensource.org/licenses/EPL-2.0)

This repository contains the Golang client SDK for [Eclipse Ditto](https://eclipse.org/ditto/).

Currently, [Eclipse Hono MQTT](https://www.eclipse.org/hono/docs/user-guide/mqtt-adapter/) is the only one supported transport.

Table of Contents
-----------------
* [Installation](#Installation)
* [Creating and connecting a client](#Creating-and-connecting-a-client)
* [Working with features](#Working-with-features)
    * [Create a new feature instance](#Create-new-feature-instance)
    * [Modify a feature's property](#Modify-a-feature's-property)
* [Subscribing and handling messages](#Subscribing-and-handling-messages)

## Installation

```
go get github.com/eclipse/ditto-clients-golang
```

## Creating and connecting a client

Each client instance requires a ditto.Configuration object.

```go
config := ditto.NewConfiguration().
    WithKeepAlive(30 * time.Second). // default keep alive is 30 seconds
    // WithCredentials(&ditto.Credentials{Username: "John", Password: "qwerty"}). if such are available or required
    WithBroker("mqtt-host:1883").
    WithConnectHandler(connectHandler)

func connectHandler(client *ditto.Client) {
    // add logic to be executed when the client is connected
}
```

With this configuration a client instance could be created.

```go
client = ditto.NewClient(config)
```
**_NOTE:_** In some cases an external Paho instance could be provided for the communication. If this is the case, there is a ditto.NewClientMqtt() create function available.

After you have configured and created your client instance, it's ready to be connected.
```go
if err := client.Connect(); err != nil {
    fmt.Errorf("%v", err)
    panic("cannot connect to broker")
}
defer disconnect(client)
```


## Working with features

### Create a new feature instance

Define the feature to be created.

```go
myFeature := &model.Feature{}
myFeature.
    WithDefinitionFrom("my.model.namespace:FeatureModel:1.0.0"). // you can provide a semantic definition of your feature
    WithProperty("myProperty", "myValue")
```

Create your Ditto command. Modify acts as an upsert - it either updates or creates features.

```go
command := things.
    NewCommand(model.NewNamespacedIDFrom("my.namespace:thing.id")). // specify which thing you will send the command to
    Twin().
    Feature("MyFeature").
    Modify(myFeature) // the payload for the modification - i.e. the feature's JSON representation
```

Send the Ditto command.

```go
envelope := command.Envelope(protocol.WithResponseRequired(false))
if err := client.Send(envelope); err != nil {
    fmt.Errorf("could not send Ditto message: %v", err)
}
```

### Modify a feature's property

Modify overrides the current feature's property.

```go
command = things.
    NewCommand(model.NewNamespacedIDFrom("my.namespace:thing.id")). // specify which thing you will send the command to
    Twin().
    FeatureProperty("MyFeature", "myProperty").
    Modify("myNewValue") // the payload for the modification - i.e. the new property's value JSON representation
```

## Subscribing and handling messages

Subscribe for incoming Ditto messages.

```go
func connectHandler(client *ditto.Client) {
    // it's a good practise to subscribe after the client is connected
    client.Subscribe(messagesHandler)
}
```
**_NOTE:_** You can add multiple handlers for Ditto messages processing.

It's a good practice to clear all subscriptions on client disconnect.
```go
func disconnect(client *ditto.Client) {
    // add any resources clearing logic
    client.Unsubscribe()
    client.Disconnect()
}
```
**_NOTE:_** If no message handler is provided then all would be removed.

Handle and reply to Ditto messages.

```go
func messagesHandler(requestID string, msg *protocol.Envelope) {
    if msg.Topic.Namespace == "my.namespace" && msg.Topic.EntityID == "thing.id" &&
            msg.Path == "/features/MyFeature/inbox/messages/myCommand" {
        // respond to the message by using the outbox
        response := things.NewMessage(model.NewNamespacedID(msg.Topic.Namespace, msg.Topic.EntityID)).
            Feature("MyFeature").Outbox("myCommand").WithPayload("responsePayload")
        responseMsg := response.Envelope(protocol.WithCorrelationID(msg.Headers.CorrelationID()), protocol.WithResponseRequired(false))
        responseMsg.Status = 200
        if replyErr := client.Reply(requestID, responseMsg); replyErr != nil {
            fmt.Errorf("failed to send response to request Id %s: %v", requestID, replyErr)
        }
    }
}
```