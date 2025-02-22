package mission

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

func VerifyCompleteByUserId(
	namespaceName string,
	missionGroupName string,
	verifyType string,
	missionTaskName string,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["missionGroupName"] = missionGroupName
	properties["verifyType"] = verifyType
	properties["missionTaskName"] = missionTaskName
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Mission:VerifyCompleteByUserId",
		Request: properties,
	}
}

func ReceiveByUserId(
	namespaceName string,
	missionGroupName string,
	missionTaskName string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["missionGroupName"] = missionGroupName
	properties["missionTaskName"] = missionTaskName
	return ConsumeAction{
		Action:  "Gs2Mission:ReceiveByUserId",
		Request: properties,
	}
}

func BatchReceiveByUserId(
	namespaceName string,
	missionGroupName string,
	missionTaskNames []string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["missionGroupName"] = missionGroupName
	properties["missionTaskNames"] = missionTaskNames
	return ConsumeAction{
		Action:  "Gs2Mission:BatchReceiveByUserId",
		Request: properties,
	}
}

func RevertReceiveByUserId(
	namespaceName string,
	missionGroupName string,
	missionTaskName string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["missionGroupName"] = missionGroupName
	properties["missionTaskName"] = missionTaskName
	return AcquireAction{
		Action:  "Gs2Mission:RevertReceiveByUserId",
		Request: properties,
	}
}

func VerifyCounterValueByUserId(
	namespaceName string,
	counterName string,
	verifyType string,
	scopeType *string,
	resetType *string,
	conditionName *string,
	value *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["counterName"] = counterName
	properties["verifyType"] = verifyType
	if scopeType != nil {
		properties["scopeType"] = scopeType
	}
	if resetType != nil {
		properties["resetType"] = resetType
	}
	if conditionName != nil {
		properties["conditionName"] = conditionName
	}
	if value != nil {
		properties["value"] = value
	}
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Mission:VerifyCounterValueByUserId",
		Request: properties,
	}
}

func DecreaseCounterByUserId(
	namespaceName string,
	counterName string,
	value int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["counterName"] = counterName
	properties["value"] = value
	return ConsumeAction{
		Action:  "Gs2Mission:DecreaseCounterByUserId",
		Request: properties,
	}
}

func ResetCounterByUserId(
	namespaceName string,
	counterName string,
	scopes []ScopedValue,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["counterName"] = counterName
	properties["scopes"] = scopes
	return ConsumeAction{
		Action:  "Gs2Mission:ResetCounterByUserId",
		Request: properties,
	}
}

func IncreaseCounterByUserId(
	namespaceName string,
	counterName string,
	value int64,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["counterName"] = counterName
	properties["value"] = value
	return AcquireAction{
		Action:  "Gs2Mission:IncreaseCounterByUserId",
		Request: properties,
	}
}

func SetCounterByUserId(
	namespaceName string,
	counterName string,
	values *[]ScopedValue,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["counterName"] = counterName
	if values != nil {
		properties["values"] = values
	}
	return AcquireAction{
		Action:  "Gs2Mission:SetCounterByUserId",
		Request: properties,
	}
}
