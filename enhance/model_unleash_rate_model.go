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

type UnleashRateModel struct {
	Name                   string
	Description            *string
	Metadata               *string
	TargetInventoryModelId string
	GradeModelId           string
	GradeEntries           []UnleashRateEntryModel
}

type UnleashRateModelOptions struct {
	Description *string
	Metadata    *string
}

func NewUnleashRateModel(
	name string,
	targetInventoryModelId string,
	gradeModelId string,
	gradeEntries []UnleashRateEntryModel,
	options UnleashRateModelOptions,
) UnleashRateModel {
	_data := UnleashRateModel{
		Name:                   name,
		TargetInventoryModelId: targetInventoryModelId,
		GradeModelId:           gradeModelId,
		GradeEntries:           gradeEntries,
		Description:            options.Description,
		Metadata:               options.Metadata,
	}
	return _data
}

func (p *UnleashRateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["TargetInventoryModelId"] = p.TargetInventoryModelId
	properties["GradeModelId"] = p.GradeModelId
	{
		values := make([]map[string]interface{}, len(p.GradeEntries))
		for i, element := range p.GradeEntries {
			values[i] = element.Properties()
		}
		properties["GradeEntries"] = values
	}
	return properties
}
