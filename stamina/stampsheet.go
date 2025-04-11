package stamina

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

func VerifyStaminaValueByUserId(
	namespaceName string,
	staminaName string,
	verifyType string,
	value int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["verifyType"] = verifyType
	properties["value"] = value
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Stamina:VerifyStaminaValueByUserId",
		Request: properties,
	}
}

func VerifyStaminaMaxValueByUserId(
	namespaceName string,
	staminaName string,
	verifyType string,
	value int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["verifyType"] = verifyType
	properties["value"] = value
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Stamina:VerifyStaminaMaxValueByUserId",
		Request: properties,
	}
}

func VerifyStaminaRecoverIntervalMinutesByUserId(
	namespaceName string,
	staminaName string,
	verifyType string,
	value int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["verifyType"] = verifyType
	properties["value"] = value
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Stamina:VerifyStaminaRecoverIntervalMinutesByUserId",
		Request: properties,
	}
}

func VerifyStaminaRecoverValueByUserId(
	namespaceName string,
	staminaName string,
	verifyType string,
	value int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["verifyType"] = verifyType
	properties["value"] = value
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Stamina:VerifyStaminaRecoverValueByUserId",
		Request: properties,
	}
}

func VerifyStaminaOverflowValueByUserId(
	namespaceName string,
	staminaName string,
	verifyType string,
	value int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["verifyType"] = verifyType
	properties["value"] = value
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Stamina:VerifyStaminaOverflowValueByUserId",
		Request: properties,
	}
}

func DecreaseMaxValueByUserId(
	namespaceName string,
	staminaName string,
	decreaseValue int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["decreaseValue"] = decreaseValue
	return ConsumeAction{
		Action:  "Gs2Stamina:DecreaseMaxValueByUserId",
		Request: properties,
	}
}

func ConsumeStaminaByUserId(
	namespaceName string,
	staminaName string,
	consumeValue int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["consumeValue"] = consumeValue
	return ConsumeAction{
		Action:  "Gs2Stamina:ConsumeStaminaByUserId",
		Request: properties,
	}
}

func RecoverStaminaByUserId(
	namespaceName string,
	staminaName string,
	recoverValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["recoverValue"] = recoverValue
	return AcquireAction{
		Action:  "Gs2Stamina:RecoverStaminaByUserId",
		Request: properties,
	}
}

func RaiseMaxValueByUserId(
	namespaceName string,
	staminaName string,
	raiseValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["raiseValue"] = raiseValue
	return AcquireAction{
		Action:  "Gs2Stamina:RaiseMaxValueByUserId",
		Request: properties,
	}
}

func SetMaxValueByUserId(
	namespaceName string,
	staminaName string,
	maxValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["maxValue"] = maxValue
	return AcquireAction{
		Action:  "Gs2Stamina:SetMaxValueByUserId",
		Request: properties,
	}
}

func SetRecoverIntervalByUserId(
	namespaceName string,
	staminaName string,
	recoverIntervalMinutes int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["recoverIntervalMinutes"] = recoverIntervalMinutes
	return AcquireAction{
		Action:  "Gs2Stamina:SetRecoverIntervalByUserId",
		Request: properties,
	}
}

func SetRecoverValueByUserId(
	namespaceName string,
	staminaName string,
	recoverValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["staminaName"] = staminaName
	properties["recoverValue"] = recoverValue
	return AcquireAction{
		Action:  "Gs2Stamina:SetRecoverValueByUserId",
		Request: properties,
	}
}
