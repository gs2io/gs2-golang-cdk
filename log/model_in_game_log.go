package log

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

type InGameLog struct {
	Timestamp int64
	RequestId string
	UserId    *string
	Tags      []InGameLogTag
	Payload   string
}

type InGameLogOptions struct {
	UserId *string
	Tags   []InGameLogTag
}

func NewInGameLog(
	timestamp int64,
	requestId string,
	payload string,
	options InGameLogOptions,
) InGameLog {
	data := InGameLog{
		Timestamp: timestamp,
		RequestId: requestId,
		Payload:   payload,
		UserId:    options.UserId,
		Tags:      options.Tags,
	}
	return data
}

func (p *InGameLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.Timestamp
	properties["RequestId"] = p.RequestId
	if p.UserId != nil {
		properties["UserId"] = p.UserId
	}
	{
		values := make([]map[string]interface{}, len(p.Tags))
		for i, element := range p.Tags {
			values[i] = element.Properties()
		}
		properties["Tags"] = values
	}
	properties["Payload"] = p.Payload
	return properties
}
