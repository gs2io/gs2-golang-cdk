package inbox

/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type GlobalMessage struct {
	Name                          string
	Metadata                      string
	ReadAcquireActions            []AcquireAction
	ExpiresTimeSpan               *TimeSpan
	ExpiresAt                     *int64
	MessageReceptionPeriodEventId *string
}

type GlobalMessageOptions struct {
	ReadAcquireActions            []AcquireAction
	ExpiresTimeSpan               *TimeSpan
	ExpiresAt                     *int64
	MessageReceptionPeriodEventId *string
}

func NewGlobalMessage(
	name string,
	metadata string,
	options GlobalMessageOptions,
) GlobalMessage {
	data := GlobalMessage{
		Name:                          name,
		Metadata:                      metadata,
		ReadAcquireActions:            options.ReadAcquireActions,
		ExpiresTimeSpan:               options.ExpiresTimeSpan,
		ExpiresAt:                     options.ExpiresAt,
		MessageReceptionPeriodEventId: options.MessageReceptionPeriodEventId,
	}
	return data
}

func (p *GlobalMessage) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Metadata"] = p.Metadata
	{
		values := make([]map[string]interface{}, len(p.ReadAcquireActions))
		for i, element := range p.ReadAcquireActions {
			values[i] = element.Properties()
		}
		properties["ReadAcquireActions"] = values
	}
	if p.ExpiresTimeSpan != nil {
		properties["ExpiresTimeSpan"] = p.ExpiresTimeSpan.Properties()
	}
	if p.ExpiresAt != nil {
		properties["ExpiresAt"] = p.ExpiresAt
	}
	if p.MessageReceptionPeriodEventId != nil {
		properties["MessageReceptionPeriodEventId"] = p.MessageReceptionPeriodEventId
	}
	return properties
}
