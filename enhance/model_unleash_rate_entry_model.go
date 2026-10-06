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

type UnleashRateEntryModelType string

const UnleashRateEntryModelTypeSimple = UnleashRateEntryModelType("simple")
const UnleashRateEntryModelTypeRecipe = UnleashRateEntryModelType("recipe")

func (p UnleashRateEntryModelType) Pointer() *UnleashRateEntryModelType {
	return &p
}

type UnleashRateEntryModel struct {
	GradeValue int64
	Type       UnleashRateEntryModelType
	NeedCount  *int32
	Recipes    []UnleashRecipe
}

type UnleashRateEntryModelOptions struct {
	NeedCount *int32
	Recipes   []UnleashRecipe
}

func NewUnleashRateEntryModel(
	gradeValue int64,
	type_ UnleashRateEntryModelType,
	options UnleashRateEntryModelOptions,
) UnleashRateEntryModel {
	_data := UnleashRateEntryModel{
		GradeValue: gradeValue,
		Type:       type_,
		NeedCount:  options.NeedCount,
		Recipes:    options.Recipes,
	}
	return _data
}

type UnleashRateEntryModelTypeIsSimpleOptions struct {
}

func NewUnleashRateEntryModelTypeIsSimple(
	gradeValue int64,
	needCount int32,
	options UnleashRateEntryModelTypeIsSimpleOptions,
) UnleashRateEntryModel {
	return NewUnleashRateEntryModel(
		gradeValue,
		UnleashRateEntryModelTypeSimple,
		UnleashRateEntryModelOptions{
			NeedCount: &needCount,
		},
	)
}

type UnleashRateEntryModelTypeIsRecipeOptions struct {
}

func NewUnleashRateEntryModelTypeIsRecipe(
	gradeValue int64,
	recipes []UnleashRecipe,
	options UnleashRateEntryModelTypeIsRecipeOptions,
) UnleashRateEntryModel {
	return NewUnleashRateEntryModel(
		gradeValue,
		UnleashRateEntryModelTypeRecipe,
		UnleashRateEntryModelOptions{
			Recipes: recipes,
		},
	)
}

func (p *UnleashRateEntryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GradeValue"] = p.GradeValue
	properties["Type"] = p.Type
	if p.NeedCount != nil {
		properties["NeedCount"] = p.NeedCount
	}
	{
		values := make([]map[string]interface{}, len(p.Recipes))
		for i, element := range p.Recipes {
			values[i] = element.Properties()
		}
		properties["Recipes"] = values
	}
	return properties
}
