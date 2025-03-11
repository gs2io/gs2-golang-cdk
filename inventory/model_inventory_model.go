package inventory

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

type InventoryModel struct {
	Name                  string
	Metadata              *string
	InitialCapacity       int32
	MaxCapacity           int32
	ProtectReferencedItem *bool
	ItemModels            []ItemModel
}

type InventoryModelOptions struct {
	Metadata              *string
	ProtectReferencedItem *bool
}

func NewInventoryModel(
	name string,
	initialCapacity int32,
	maxCapacity int32,
	itemModels []ItemModel,
	options InventoryModelOptions,
) InventoryModel {
	_data := InventoryModel{
		Name:                  name,
		InitialCapacity:       initialCapacity,
		MaxCapacity:           maxCapacity,
		ItemModels:            itemModels,
		Metadata:              options.Metadata,
		ProtectReferencedItem: options.ProtectReferencedItem,
	}
	return _data
}

func (p *InventoryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["InitialCapacity"] = p.InitialCapacity
	properties["MaxCapacity"] = p.MaxCapacity
	if p.ProtectReferencedItem != nil {
		properties["ProtectReferencedItem"] = p.ProtectReferencedItem
	}
	{
		values := make([]map[string]interface{}, len(p.ItemModels))
		for i, element := range p.ItemModels {
			values[i] = element.Properties()
		}
		properties["ItemModels"] = values
	}
	return properties
}
