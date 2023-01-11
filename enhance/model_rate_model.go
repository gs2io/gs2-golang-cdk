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
	Name                       string
	Description                *string
	Metadata                   *string
	TargetInventoryModelId     string
	AcquireExperienceSuffix    string
	MaterialInventoryModelId   string
	AcquireExperienceHierarchy []string
	ExperienceModelId          string
	BonusRates                 []BonusRate
}

type RateModelOptions struct {
	Description                *string
	Metadata                   *string
	AcquireExperienceHierarchy []string
	BonusRates                 []BonusRate
}

func NewRateModel(
	name string,
	targetInventoryModelId string,
	acquireExperienceSuffix string,
	materialInventoryModelId string,
	experienceModelId string,
	options RateModelOptions,
) RateModel {
	data := RateModel{
		Name:                       name,
		TargetInventoryModelId:     targetInventoryModelId,
		AcquireExperienceSuffix:    acquireExperienceSuffix,
		MaterialInventoryModelId:   materialInventoryModelId,
		ExperienceModelId:          experienceModelId,
		Description:                options.Description,
		Metadata:                   options.Metadata,
		AcquireExperienceHierarchy: options.AcquireExperienceHierarchy,
		BonusRates:                 options.BonusRates,
	}
	return data
}

func (p *RateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["TargetInventoryModelId"] = p.TargetInventoryModelId
	properties["AcquireExperienceSuffix"] = p.AcquireExperienceSuffix
	properties["MaterialInventoryModelId"] = p.MaterialInventoryModelId
	properties["AcquireExperienceHierarchy"] = p.AcquireExperienceHierarchy
	properties["ExperienceModelId"] = p.ExperienceModelId
	{
		values := make([]map[string]interface{}, len(p.BonusRates))
		for i, element := range p.BonusRates {
			values[i] = element.Properties()
		}
		properties["BonusRates"] = values
	}
	return properties
}
