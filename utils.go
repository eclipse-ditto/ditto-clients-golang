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
	"fmt"
	"github.com/eclipse/ditto-clients-golang/protocol"
	"regexp"
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
