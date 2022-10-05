package enhance

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

type RateModel struct {
	name                       string
	description                *string
	metadata                   *string
	targetInventoryModelId     string
	acquireExperienceSuffix    string
	materialInventoryModelId   string
	acquireExperienceHierarchy []string
	experienceModelId          string
	bonusRates                 []BonusRate
}

func NewRateModel(
	name string,
	description *string,
	metadata *string,
	targetInventoryModelId string,
	acquireExperienceSuffix string,
	materialInventoryModelId string,
	acquireExperienceHierarchy []string,
	experienceModelId string,
	bonusRates []BonusRate,
) *RateModel {
	return &RateModel{
		name:                       name,
		description:                description,
		metadata:                   metadata,
		targetInventoryModelId:     targetInventoryModelId,
		acquireExperienceSuffix:    acquireExperienceSuffix,
		materialInventoryModelId:   materialInventoryModelId,
		acquireExperienceHierarchy: acquireExperienceHierarchy,
		experienceModelId:          experienceModelId,
		bonusRates:                 bonusRates,
	}
}

func (p *RateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["TargetInventoryModelId"] = p.targetInventoryModelId
	properties["AcquireExperienceSuffix"] = p.acquireExperienceSuffix
	properties["MaterialInventoryModelId"] = p.materialInventoryModelId
	properties["AcquireExperienceHierarchy"] = p.acquireExperienceHierarchy
	properties["ExperienceModelId"] = p.experienceModelId
	{
		values := make([]map[string]interface{}, len(p.bonusRates))
		for i, element := range p.bonusRates {
			values[i] = element.Properties()
		}
		properties["BonusRates"] = values
	}
	return properties
}

func (p *RateModel) Ref(
	namespaceName string,
) RateModelRef {
	return RateModelRef{
		namespaceName: namespaceName,
		rateName:      p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	rateModels    []RateModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	rateModels []RateModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2020-08-22",
		namespaceName: namespaceName,
		rateModels:    rateModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Enhance_CurrentRateMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Enhance::CurrentRateMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	rateModels := make([]map[string]interface{}, len(p.rateModels))
	for i, item := range p.rateModels {
		rateModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":    p.version,
			"RateModels": rateModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack               *Stack
	name                string
	description         *string
	enableDirectEnhance bool
	transactionSetting  TransactionSetting
	enhanceScript       *ScriptSetting
	logSetting          *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	enableDirectEnhance bool,
	transactionSetting TransactionSetting,
	description *string,
	enhanceScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:               stack,
		name:                name,
		description:         description,
		enableDirectEnhance: enableDirectEnhance,
		transactionSetting:  transactionSetting,
		enhanceScript:       enhanceScript,
		logSetting:          logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Enhance_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Enhance::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["EnableDirectEnhance"] = p.enableDirectEnhance
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.enhanceScript != nil {
		properties["EnhanceScript"] = p.enhanceScript.Properties()
	}
	if p.logSetting != nil {
		properties["LogSetting"] = p.logSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		namespaceName: p.name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}

func (p *Namespace) MasterData(
	rateModels []RateModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		rateModels,
	).AddDependsOn(
		p,
	)
	return p
}

type RateModelMaster struct {
	CdkResource
	stack                      *Stack
	namespaceName              string
	name                       string
	description                *string
	metadata                   *string
	targetInventoryModelId     string
	acquireExperienceSuffix    string
	materialInventoryModelId   string
	acquireExperienceHierarchy []string
	experienceModelId          string
	bonusRates                 []BonusRate
}

func NewRateModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	targetInventoryModelId string,
	acquireExperienceSuffix string,
	materialInventoryModelId string,
	acquireExperienceHierarchy []string,
	experienceModelId string,
	bonusRates []BonusRate,
	description *string,
	metadata *string,
) *RateModelMaster {

	self := RateModelMaster{
		stack:                      stack,
		namespaceName:              namespaceName,
		name:                       name,
		description:                description,
		metadata:                   metadata,
		targetInventoryModelId:     targetInventoryModelId,
		acquireExperienceSuffix:    acquireExperienceSuffix,
		materialInventoryModelId:   materialInventoryModelId,
		acquireExperienceHierarchy: acquireExperienceHierarchy,
		experienceModelId:          experienceModelId,
		bonusRates:                 bonusRates,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *RateModelMaster) ResourceName() string {
	return "Enhance_RateModelMaster_" + p.name
}

func (p *RateModelMaster) ResourceType() string {
	return "GS2::Enhance::RateModelMaster"
}

func (p *RateModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["TargetInventoryModelId"] = p.targetInventoryModelId
	properties["AcquireExperienceSuffix"] = p.acquireExperienceSuffix
	properties["MaterialInventoryModelId"] = p.materialInventoryModelId
	properties["AcquireExperienceHierarchy"] = p.acquireExperienceHierarchy
	properties["ExperienceModelId"] = p.experienceModelId
	{
		values := make([]map[string]interface{}, len(p.bonusRates))
		for i, element := range p.bonusRates {
			values[i] = element.Properties()
		}
		properties["BonusRates"] = values
	}
	return properties
}

func (p *RateModelMaster) Ref(
	namespaceName string,
) RateModelMasterRef {
	return RateModelMasterRef{
		namespaceName: namespaceName,
		rateName:      p.name,
	}
}

func (p *RateModelMaster) GetAttrRateModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RateModelId",
	)
}
