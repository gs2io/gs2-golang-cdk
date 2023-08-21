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

type CurrentMasterData struct {
	CdkResource
	stack                 *Stack
	version               string
	namespaceName         string
	inventoryModels       []InventoryModel
	simpleInventoryModels []SimpleInventoryModel
	bigInventoryModels    []BigInventoryModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	inventoryModels []InventoryModel,
	simpleInventoryModels []SimpleInventoryModel,
	bigInventoryModels []BigInventoryModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:                 stack,
		version:               "2019-02-05",
		namespaceName:         namespaceName,
		inventoryModels:       inventoryModels,
		simpleInventoryModels: simpleInventoryModels,
		bigInventoryModels:    bigInventoryModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Inventory_CurrentItemModelMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Inventory::CurrentItemModelMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	inventoryModels := make([]map[string]interface{}, len(p.inventoryModels))
	for i, item := range p.inventoryModels {
		inventoryModels[i] = item.Properties()
	}
	simpleInventoryModels := make([]map[string]interface{}, len(p.simpleInventoryModels))
	for i, item := range p.simpleInventoryModels {
		simpleInventoryModels[i] = item.Properties()
	}
	bigInventoryModels := make([]map[string]interface{}, len(p.bigInventoryModels))
	for i, item := range p.bigInventoryModels {
		bigInventoryModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":               p.version,
			"inventoryModels":       inventoryModels,
			"simpleInventoryModels": simpleInventoryModels,
			"bigInventoryModels":    bigInventoryModels,
		},
	}
}
