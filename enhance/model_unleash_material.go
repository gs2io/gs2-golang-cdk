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

type UnleashMaterialMaterialType string

const UnleashMaterialMaterialTypeIndividual = UnleashMaterialMaterialType("individual")
const UnleashMaterialMaterialTypeQuantity = UnleashMaterialMaterialType("quantity")

func (p UnleashMaterialMaterialType) Pointer() *UnleashMaterialMaterialType {
	return &p
}

type UnleashMaterial struct {
	Name              string
	MaterialType      UnleashMaterialMaterialType
	IndividualSetting *UnleashIndividualMaterialSetting
	QuantitySetting   *UnleashQuantityMaterialSetting
}

type UnleashMaterialOptions struct {
	IndividualSetting *UnleashIndividualMaterialSetting
	QuantitySetting   *UnleashQuantityMaterialSetting
}

func NewUnleashMaterial(
	name string,
	materialType UnleashMaterialMaterialType,
	options UnleashMaterialOptions,
) UnleashMaterial {
	_data := UnleashMaterial{
		Name:              name,
		MaterialType:      materialType,
		IndividualSetting: options.IndividualSetting,
		QuantitySetting:   options.QuantitySetting,
	}
	return _data
}

type UnleashMaterialMaterialTypeIsIndividualOptions struct {
}

func NewUnleashMaterialMaterialTypeIsIndividual(
	name string,
	individualSetting UnleashIndividualMaterialSetting,
	options UnleashMaterialMaterialTypeIsIndividualOptions,
) UnleashMaterial {
	return NewUnleashMaterial(
		name,
		UnleashMaterialMaterialTypeIndividual,
		UnleashMaterialOptions{
			IndividualSetting: &individualSetting,
		},
	)
}

type UnleashMaterialMaterialTypeIsQuantityOptions struct {
}

func NewUnleashMaterialMaterialTypeIsQuantity(
	name string,
	quantitySetting UnleashQuantityMaterialSetting,
	options UnleashMaterialMaterialTypeIsQuantityOptions,
) UnleashMaterial {
	return NewUnleashMaterial(
		name,
		UnleashMaterialMaterialTypeQuantity,
		UnleashMaterialOptions{
			QuantitySetting: &quantitySetting,
		},
	)
}

func (p *UnleashMaterial) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["MaterialType"] = p.MaterialType
	if p.IndividualSetting != nil {
		properties["IndividualSetting"] = p.IndividualSetting.Properties()
	}
	if p.QuantitySetting != nil {
		properties["QuantitySetting"] = p.QuantitySetting.Properties()
	}
	return properties
}
