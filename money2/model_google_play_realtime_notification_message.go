package money2

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

type GooglePlayRealtimeNotificationMessage struct {
	Data        string
	MessageId   string
	PublishTime string
}

type GooglePlayRealtimeNotificationMessageOptions struct {
}

func NewGooglePlayRealtimeNotificationMessage(
	data string,
	messageId string,
	publishTime string,
	options GooglePlayRealtimeNotificationMessageOptions,
) GooglePlayRealtimeNotificationMessage {
	data := GooglePlayRealtimeNotificationMessage{
		Data:        data,
		MessageId:   messageId,
		PublishTime: publishTime,
	}
	return data
}

func (p *GooglePlayRealtimeNotificationMessage) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Data"] = p.Data
	properties["MessageId"] = p.MessageId
	properties["PublishTime"] = p.PublishTime
	return properties
}
