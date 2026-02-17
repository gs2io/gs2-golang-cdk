package buff

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

type BuffTargetActionTargetActionName string

const BuffTargetActionTargetActionNameGs2ExperienceAddExperienceByUserId = BuffTargetActionTargetActionName("Gs2Experience:AddExperienceByUserId")
const BuffTargetActionTargetActionNameGs2ExperienceSubExperience = BuffTargetActionTargetActionName("Gs2Experience:SubExperience")
const BuffTargetActionTargetActionNameGs2ExperienceSubExperienceByUserId = BuffTargetActionTargetActionName("Gs2Experience:SubExperienceByUserId")
const BuffTargetActionTargetActionNameGs2InventoryAcquireBigItemByUserId = BuffTargetActionTargetActionName("Gs2Inventory:AcquireBigItemByUserId")
const BuffTargetActionTargetActionNameGs2InventoryAcquireItemSetByUserId = BuffTargetActionTargetActionName("Gs2Inventory:AcquireItemSetByUserId")
const BuffTargetActionTargetActionNameGs2InventoryAcquireSimpleItemsByUserId = BuffTargetActionTargetActionName("Gs2Inventory:AcquireSimpleItemsByUserId")
const BuffTargetActionTargetActionNameGs2InventoryConsumeBigItem = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeBigItem")
const BuffTargetActionTargetActionNameGs2InventoryConsumeBigItemByUserId = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeBigItemByUserId")
const BuffTargetActionTargetActionNameGs2InventoryConsumeItemSet = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeItemSet")
const BuffTargetActionTargetActionNameGs2InventoryConsumeItemSetByUserId = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeItemSetByUserId")
const BuffTargetActionTargetActionNameGs2InventoryConsumeSimpleItems = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeSimpleItems")
const BuffTargetActionTargetActionNameGs2InventoryConsumeSimpleItemsByUserId = BuffTargetActionTargetActionName("Gs2Inventory:ConsumeSimpleItemsByUserId")
const BuffTargetActionTargetActionNameGs2LimitCountUp = BuffTargetActionTargetActionName("Gs2Limit:CountUp")
const BuffTargetActionTargetActionNameGs2LimitCountUpByUserId = BuffTargetActionTargetActionName("Gs2Limit:CountUpByUserId")
const BuffTargetActionTargetActionNameGs2Money2DepositByUserId = BuffTargetActionTargetActionName("Gs2Money2:DepositByUserId")
const BuffTargetActionTargetActionNameGs2Money2Withdraw = BuffTargetActionTargetActionName("Gs2Money2:Withdraw")
const BuffTargetActionTargetActionNameGs2Money2WithdrawByUserId = BuffTargetActionTargetActionName("Gs2Money2:WithdrawByUserId")
const BuffTargetActionTargetActionNameGs2MoneyDepositByUserId = BuffTargetActionTargetActionName("Gs2Money:DepositByUserId")
const BuffTargetActionTargetActionNameGs2MoneyWithdraw = BuffTargetActionTargetActionName("Gs2Money:Withdraw")
const BuffTargetActionTargetActionNameGs2MoneyWithdrawByUserId = BuffTargetActionTargetActionName("Gs2Money:WithdrawByUserId")
const BuffTargetActionTargetActionNameGs2StaminaConsumeStamina = BuffTargetActionTargetActionName("Gs2Stamina:ConsumeStamina")
const BuffTargetActionTargetActionNameGs2StaminaConsumeStaminaByUserId = BuffTargetActionTargetActionName("Gs2Stamina:ConsumeStaminaByUserId")
const BuffTargetActionTargetActionNameGs2StaminaRecoverStaminaByUserId = BuffTargetActionTargetActionName("Gs2Stamina:RecoverStaminaByUserId")

func (p BuffTargetActionTargetActionName) Pointer() *BuffTargetActionTargetActionName {
	return &p
}

type BuffTargetAction struct {
	TargetActionName string
	TargetFieldName  string
	ConditionGrns    []BuffTargetGrn
	Rate             float32
}

type BuffTargetActionOptions struct {
}

func NewBuffTargetAction(
	targetActionName string,
	targetFieldName string,
	conditionGrns []BuffTargetGrn,
	rate float32,
	options BuffTargetActionOptions,
) BuffTargetAction {
	_data := BuffTargetAction{
		TargetActionName: targetActionName,
		TargetFieldName:  targetFieldName,
		ConditionGrns:    conditionGrns,
		Rate:             rate,
	}
	return _data
}

func (p *BuffTargetAction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TargetActionName"] = p.TargetActionName
	properties["TargetFieldName"] = p.TargetFieldName
	{
		values := make([]map[string]interface{}, len(p.ConditionGrns))
		for i, element := range p.ConditionGrns {
			values[i] = element.Properties()
		}
		properties["ConditionGrns"] = values
	}
	properties["Rate"] = p.Rate
	return properties
}
