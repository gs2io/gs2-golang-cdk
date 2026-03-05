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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type BuffTargetModelTargetModelName string

const BuffTargetModelTargetModelNameGs2ExchangeIncrementalRateModel = BuffTargetModelTargetModelName("Gs2Exchange:IncrementalRateModel")
const BuffTargetModelTargetModelNameGs2ExchangeRateModel = BuffTargetModelTargetModelName("Gs2Exchange:RateModel")
const BuffTargetModelTargetModelNameGs2ExperienceExperienceModel = BuffTargetModelTargetModelName("Gs2Experience:ExperienceModel")
const BuffTargetModelTargetModelNameGs2ExperienceStatus = BuffTargetModelTargetModelName("Gs2Experience:Status")
const BuffTargetModelTargetModelNameGs2FormationMold = BuffTargetModelTargetModelName("Gs2Formation:Mold")
const BuffTargetModelTargetModelNameGs2IdleCategoryModel = BuffTargetModelTargetModelName("Gs2Idle:CategoryModel")
const BuffTargetModelTargetModelNameGs2IdleStatus = BuffTargetModelTargetModelName("Gs2Idle:Status")
const BuffTargetModelTargetModelNameGs2InventoryBigInventoryModel = BuffTargetModelTargetModelName("Gs2Inventory:BigInventoryModel")
const BuffTargetModelTargetModelNameGs2InventoryBigItemModel = BuffTargetModelTargetModelName("Gs2Inventory:BigItemModel")
const BuffTargetModelTargetModelNameGs2InventoryInventory = BuffTargetModelTargetModelName("Gs2Inventory:Inventory")
const BuffTargetModelTargetModelNameGs2InventoryInventoryModel = BuffTargetModelTargetModelName("Gs2Inventory:InventoryModel")
const BuffTargetModelTargetModelNameGs2InventoryItemModel = BuffTargetModelTargetModelName("Gs2Inventory:ItemModel")
const BuffTargetModelTargetModelNameGs2InventorySimpleInventoryModel = BuffTargetModelTargetModelName("Gs2Inventory:SimpleInventoryModel")
const BuffTargetModelTargetModelNameGs2InventorySimpleItemModel = BuffTargetModelTargetModelName("Gs2Inventory:SimpleItemModel")
const BuffTargetModelTargetModelNameGs2LimitCounter = BuffTargetModelTargetModelName("Gs2Limit:Counter")
const BuffTargetModelTargetModelNameGs2LimitLimitModel = BuffTargetModelTargetModelName("Gs2Limit:LimitModel")
const BuffTargetModelTargetModelNameGs2LoginRewardBonusModel = BuffTargetModelTargetModelName("Gs2LoginReward:BonusModel")
const BuffTargetModelTargetModelNameGs2MissionMissionTaskModel = BuffTargetModelTargetModelName("Gs2Mission:MissionTaskModel")
const BuffTargetModelTargetModelNameGs2Money2Wallet = BuffTargetModelTargetModelName("Gs2Money2:Wallet")
const BuffTargetModelTargetModelNameGs2MoneyWallet = BuffTargetModelTargetModelName("Gs2Money:Wallet")
const BuffTargetModelTargetModelNameGs2QuestQuestModel = BuffTargetModelTargetModelName("Gs2Quest:QuestModel")
const BuffTargetModelTargetModelNameGs2ShowcaseDisplayItem = BuffTargetModelTargetModelName("Gs2Showcase:DisplayItem")
const BuffTargetModelTargetModelNameGs2ShowcaseRandomDisplayItemModel = BuffTargetModelTargetModelName("Gs2Showcase:RandomDisplayItemModel")
const BuffTargetModelTargetModelNameGs2SkillTreeNodeModel = BuffTargetModelTargetModelName("Gs2SkillTree:NodeModel")
const BuffTargetModelTargetModelNameGs2StaminaStamina = BuffTargetModelTargetModelName("Gs2Stamina:Stamina")
const BuffTargetModelTargetModelNameGs2StaminaStaminaModel = BuffTargetModelTargetModelName("Gs2Stamina:StaminaModel")

func (p BuffTargetModelTargetModelName) Pointer() *BuffTargetModelTargetModelName {
	return &p
}

type BuffTargetModel struct {
	TargetModelName BuffTargetModelTargetModelName
	TargetFieldName string
	ConditionGrns   []BuffTargetGrn
	Rate            float32
}

type BuffTargetModelOptions struct {
}

func NewBuffTargetModel(
	targetModelName BuffTargetModelTargetModelName,
	targetFieldName string,
	conditionGrns []BuffTargetGrn,
	rate float32,
	options BuffTargetModelOptions,
) BuffTargetModel {
	_data := BuffTargetModel{
		TargetModelName: targetModelName,
		TargetFieldName: targetFieldName,
		ConditionGrns:   conditionGrns,
		Rate:            rate,
	}
	return _data
}

func (p *BuffTargetModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TargetModelName"] = p.TargetModelName
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
