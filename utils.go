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

package ditto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/ditto-clients-golang/protocol"
	"reflect"
	"regexp"
	"runtime"
)

var regexHonoMQTTTopicRequest, _ = regexp.Compile("^command///req/([^/]+)/([^/]+)$")

const (
	honoMQTTTopicCommandResponseFormat = "command///res/%s/%d"
)

func extractHonoRequestId(honoTopic string) string {
	if regexHonoMQTTTopicRequest.MatchString(honoTopic) {
		reqIdInfo := regexHonoMQTTTopicRequest.FindStringSubmatch(honoTopic)
		return reqIdInfo[1]
	}
	return ""
}

func generateHonoResponseTopic(requestID string, status int) string {
	return fmt.Sprintf(honoMQTTTopicCommandResponseFormat, requestID, status)
}

func getEnvelope(mqttPayload []byte) (*protocol.Envelope, error) {
	env := &protocol.Envelope{}
	if err := json.Unmarshal(mqttPayload, env); err != nil {
		return nil, err
	}
	return env, nil
}

// Get the function name of a handler
func getHandlerName(handler Handler) string {
	return runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
}

func validateConfiguration(cfg *Configuration) error {
	if cfg == nil {
		return nil
	}
	if cfg.broker != "" {
		return errors.New("broker is not expected when using external MQTT client")
	} else if cfg.credentials != nil {
		return errors.New("credentials are not expected when using external MQTT client")
	} else if cfg.disconnectTimeout != defaultDisconnectTimeout && cfg.disconnectTimeout != 0 {
		return errors.New("disconnectTimeout is not expected when using external MQTT client")
	} else if cfg.keepAlive != defaultKeepAlive && cfg.keepAlive != 0 {
		return errors.New("keepAlive is not expected when using external MQTT client")
	} else if cfg.tlsConfig != nil {
		return errors.New("TLS configuration is not expected when using external MQTT client")
	}
	return nil
}
