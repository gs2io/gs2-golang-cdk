package schedule

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

func VerifyTriggerByUserId(
	namespaceName string,
	triggerName string,
	verifyType string,
	elapsedMinutes *int32,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["triggerName"] = triggerName
	properties["verifyType"] = verifyType
	if elapsedMinutes != nil {
		properties["elapsedMinutes"] = elapsedMinutes
	}
	return VerifyAction{
		Action:  "Gs2Schedule:VerifyTriggerByUserId",
		Request: properties,
	}
}

func DeleteTriggerByUserId(
	namespaceName string,
	triggerName string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["triggerName"] = triggerName
	return ConsumeAction{
		Action:  "Gs2Schedule:DeleteTriggerByUserId",
		Request: properties,
	}
}

func TriggerByUserId(
	namespaceName string,
	triggerName string,
	triggerStrategy string,
	ttl *int32,
	eventId *string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["triggerName"] = triggerName
	properties["triggerStrategy"] = triggerStrategy
	if ttl != nil {
		properties["ttl"] = ttl
	}
	if eventId != nil {
		properties["eventId"] = eventId
	}
	return AcquireAction{
		Action:  "Gs2Schedule:TriggerByUserId",
		Request: properties,
	}
}

func ExtendTriggerByUserId(
	namespaceName string,
	triggerName string,
	extendSeconds int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["triggerName"] = triggerName
	properties["extendSeconds"] = extendSeconds
	return AcquireAction{
		Action:  "Gs2Schedule:ExtendTriggerByUserId",
		Request: properties,
	}
}

func VerifyEventByUserId(
	namespaceName string,
	eventName string,
	verifyType string,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["eventName"] = eventName
	properties["verifyType"] = verifyType
	return VerifyAction{
		Action:  "Gs2Schedule:VerifyEventByUserId",
		Request: properties,
	}
}
