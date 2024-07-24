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

func OpenMessageByUserId(
	namespaceName string,
	messageName *string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if messageName != nil {
		properties["messageName"] = messageName
	}
	return ConsumeAction{
		Action:  "Gs2Inbox:OpenMessageByUserId",
		Request: properties,
	}
}

func DeleteMessageByUserId(
	namespaceName string,
	messageName *string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if messageName != nil {
		properties["messageName"] = messageName
	}
	return ConsumeAction{
		Action:  "Gs2Inbox:DeleteMessageByUserId",
		Request: properties,
	}
}

func SendMessageByUserId(
	namespaceName string,
	metadata string,
	readAcquireActions *[]AcquireAction,
	expiresAt *int64,
	expiresTimeSpan *TimeSpan,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["metadata"] = metadata
	if readAcquireActions != nil {
		properties["readAcquireActions"] = readAcquireActions
	}
	if expiresAt != nil {
		properties["expiresAt"] = expiresAt
	}
	if expiresTimeSpan != nil {
		properties["expiresTimeSpan"] = expiresTimeSpan
	}
	return AcquireAction{
		Action:  "Gs2Inbox:SendMessageByUserId",
		Request: properties,
	}
}
