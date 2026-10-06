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

type UnleashQuantityMaterialSettingMatchType string

const UnleashQuantityMaterialSettingMatchTypeSameGroup = UnleashQuantityMaterialSettingMatchType("sameGroup")
const UnleashQuantityMaterialSettingMatchTypeSpecified = UnleashQuantityMaterialSettingMatchType("specified")

func (p UnleashQuantityMaterialSettingMatchType) Pointer() *UnleashQuantityMaterialSettingMatchType {
	return &p
}

type UnleashQuantityMaterialSetting struct {
	MatchType                UnleashQuantityMaterialSettingMatchType
	MaterialInventoryModelId *string
	ItemModelId              *string
	Count                    int32
}

type UnleashQuantityMaterialSettingOptions struct {
	MaterialInventoryModelId *string
	ItemModelId              *string
}

func NewUnleashQuantityMaterialSetting(
	matchType UnleashQuantityMaterialSettingMatchType,
	count int32,
	options UnleashQuantityMaterialSettingOptions,
) UnleashQuantityMaterialSetting {
	_data := UnleashQuantityMaterialSetting{
		MatchType:                matchType,
		Count:                    count,
		MaterialInventoryModelId: options.MaterialInventoryModelId,
		ItemModelId:              options.ItemModelId,
	}
	return _data
}

type UnleashQuantityMaterialSettingMatchTypeIsSameGroupOptions struct {
}

func NewUnleashQuantityMaterialSettingMatchTypeIsSameGroup(
	count int32,
	materialInventoryModelId string,
	options UnleashQuantityMaterialSettingMatchTypeIsSameGroupOptions,
) UnleashQuantityMaterialSetting {
	return NewUnleashQuantityMaterialSetting(
		UnleashQuantityMaterialSettingMatchTypeSameGroup,
		count,
		UnleashQuantityMaterialSettingOptions{
			MaterialInventoryModelId: &materialInventoryModelId,
		},
	)
}

type UnleashQuantityMaterialSettingMatchTypeIsSpecifiedOptions struct {
}

func NewUnleashQuantityMaterialSettingMatchTypeIsSpecified(
	count int32,
	itemModelId string,
	options UnleashQuantityMaterialSettingMatchTypeIsSpecifiedOptions,
) UnleashQuantityMaterialSetting {
	return NewUnleashQuantityMaterialSetting(
		UnleashQuantityMaterialSettingMatchTypeSpecified,
		count,
		UnleashQuantityMaterialSettingOptions{
			ItemModelId: &itemModelId,
		},
	)
}

func (p *UnleashQuantityMaterialSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["MatchType"] = p.MatchType
	if p.MaterialInventoryModelId != nil {
		properties["MaterialInventoryModelId"] = p.MaterialInventoryModelId
	}
	if p.ItemModelId != nil {
		properties["ItemModelId"] = p.ItemModelId
	}
	properties["Count"] = p.Count
	return properties
}
