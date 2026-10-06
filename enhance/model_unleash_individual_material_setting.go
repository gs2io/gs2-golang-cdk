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

type UnleashIndividualMaterialSettingMatchType string

const UnleashIndividualMaterialSettingMatchTypeSameItem = UnleashIndividualMaterialSettingMatchType("sameItem")
const UnleashIndividualMaterialSettingMatchTypeSameGroup = UnleashIndividualMaterialSettingMatchType("sameGroup")

func (p UnleashIndividualMaterialSettingMatchType) Pointer() *UnleashIndividualMaterialSettingMatchType {
	return &p
}

type UnleashIndividualMaterialSettingGradeCondition string

const UnleashIndividualMaterialSettingGradeConditionAny = UnleashIndividualMaterialSettingGradeCondition("any")
const UnleashIndividualMaterialSettingGradeConditionSameAsTarget = UnleashIndividualMaterialSettingGradeCondition("sameAsTarget")
const UnleashIndividualMaterialSettingGradeConditionEqual = UnleashIndividualMaterialSettingGradeCondition("equal")

func (p UnleashIndividualMaterialSettingGradeCondition) Pointer() *UnleashIndividualMaterialSettingGradeCondition {
	return &p
}

type UnleashIndividualMaterialSetting struct {
	MatchType      UnleashIndividualMaterialSettingMatchType
	GradeCondition UnleashIndividualMaterialSettingGradeCondition
	GradeValue     *int64
	Count          int32
}

type UnleashIndividualMaterialSettingOptions struct {
	GradeValue *int64
}

func NewUnleashIndividualMaterialSetting(
	matchType UnleashIndividualMaterialSettingMatchType,
	gradeCondition UnleashIndividualMaterialSettingGradeCondition,
	count int32,
	options UnleashIndividualMaterialSettingOptions,
) UnleashIndividualMaterialSetting {
	_data := UnleashIndividualMaterialSetting{
		MatchType:      matchType,
		GradeCondition: gradeCondition,
		Count:          count,
		GradeValue:     options.GradeValue,
	}
	return _data
}

type UnleashIndividualMaterialSettingGradeConditionIsAnyOptions struct {
}

func NewUnleashIndividualMaterialSettingGradeConditionIsAny(
	matchType UnleashIndividualMaterialSettingMatchType,
	count int32,
	options UnleashIndividualMaterialSettingGradeConditionIsAnyOptions,
) UnleashIndividualMaterialSetting {
	return NewUnleashIndividualMaterialSetting(
		matchType,
		UnleashIndividualMaterialSettingGradeConditionAny,
		count,
		UnleashIndividualMaterialSettingOptions{},
	)
}

type UnleashIndividualMaterialSettingGradeConditionIsSameAsTargetOptions struct {
}

func NewUnleashIndividualMaterialSettingGradeConditionIsSameAsTarget(
	matchType UnleashIndividualMaterialSettingMatchType,
	count int32,
	options UnleashIndividualMaterialSettingGradeConditionIsSameAsTargetOptions,
) UnleashIndividualMaterialSetting {
	return NewUnleashIndividualMaterialSetting(
		matchType,
		UnleashIndividualMaterialSettingGradeConditionSameAsTarget,
		count,
		UnleashIndividualMaterialSettingOptions{},
	)
}

type UnleashIndividualMaterialSettingGradeConditionIsEqualOptions struct {
}

func NewUnleashIndividualMaterialSettingGradeConditionIsEqual(
	matchType UnleashIndividualMaterialSettingMatchType,
	count int32,
	gradeValue int64,
	options UnleashIndividualMaterialSettingGradeConditionIsEqualOptions,
) UnleashIndividualMaterialSetting {
	return NewUnleashIndividualMaterialSetting(
		matchType,
		UnleashIndividualMaterialSettingGradeConditionEqual,
		count,
		UnleashIndividualMaterialSettingOptions{
			GradeValue: &gradeValue,
		},
	)
}

func (p *UnleashIndividualMaterialSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["MatchType"] = p.MatchType
	properties["GradeCondition"] = p.GradeCondition
	if p.GradeValue != nil {
		properties["GradeValue"] = p.GradeValue
	}
	properties["Count"] = p.Count
	return properties
}
