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
	name                  string
	metadata              *string
	initialCapacity       int32
	maxCapacity           int32
	protectReferencedItem *bool
	itemModels            []ItemModel
}

func NewInventoryModel(
	name string,
	metadata *string,
	initialCapacity int32,
	maxCapacity int32,
	protectReferencedItem *bool,
	itemModels []ItemModel,
) *InventoryModel {
	return &InventoryModel{
		name:                  name,
		metadata:              metadata,
		initialCapacity:       initialCapacity,
		maxCapacity:           maxCapacity,
		protectReferencedItem: protectReferencedItem,
		itemModels:            itemModels,
	}
}

func (p *InventoryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["InitialCapacity"] = p.initialCapacity
	properties["MaxCapacity"] = p.maxCapacity
	if p.protectReferencedItem != nil {
		properties["ProtectReferencedItem"] = p.protectReferencedItem
	}
	{
		values := make([]map[string]interface{}, len(p.itemModels))
		for i, element := range p.itemModels {
			values[i] = element.Properties()
		}
		properties["ItemModels"] = values
	}
	return properties
}

func (p *InventoryModel) Ref(
	namespaceName string,
) InventoryModelRef {
	return InventoryModelRef{
		namespaceName: namespaceName,
		inventoryName: p.name,
	}
}

type ItemModel struct {
	name                string
	metadata            *string
	stackingLimit       int64
	allowMultipleStacks bool
	sortValue           int32
}

func NewItemModel(
	name string,
	metadata *string,
	stackingLimit int64,
	allowMultipleStacks bool,
	sortValue int32,
) *ItemModel {
	return &ItemModel{
		name:                name,
		metadata:            metadata,
		stackingLimit:       stackingLimit,
		allowMultipleStacks: allowMultipleStacks,
		sortValue:           sortValue,
	}
}

func (p *ItemModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["StackingLimit"] = p.stackingLimit
	properties["AllowMultipleStacks"] = p.allowMultipleStacks
	properties["SortValue"] = p.sortValue
	return properties
}

func (p *ItemModel) Ref(
	namespaceName string,
	inventoryName string,
) ItemModelRef {
	return ItemModelRef{
		namespaceName: namespaceName,
		inventoryName: inventoryName,
		itemName:      p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack           *Stack
	version         string
	namespaceName   string
	inventoryModels []InventoryModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	inventoryModels []InventoryModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:           stack,
		version:         "2019-02-05",
		namespaceName:   namespaceName,
		inventoryModels: inventoryModels,
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
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":         p.version,
			"InventoryModels": inventoryModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack          *Stack
	name           string
	description    *string
	acquireScript  *ScriptSetting
	overflowScript *ScriptSetting
	consumeScript  *ScriptSetting
	logSetting     *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	acquireScript *ScriptSetting,
	overflowScript *ScriptSetting,
	consumeScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:          stack,
		name:           name,
		description:    description,
		acquireScript:  acquireScript,
		overflowScript: overflowScript,
		consumeScript:  consumeScript,
		logSetting:     logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Inventory_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Inventory::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.acquireScript != nil {
		properties["AcquireScript"] = p.acquireScript.Properties()
	}
	if p.overflowScript != nil {
		properties["OverflowScript"] = p.overflowScript.Properties()
	}
	if p.consumeScript != nil {
		properties["ConsumeScript"] = p.consumeScript.Properties()
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
	inventoryModels []InventoryModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		inventoryModels,
	).AddDependsOn(
		p,
	)
	return p
}

type InventoryModelMaster struct {
	CdkResource
	stack                 *Stack
	namespaceName         string
	name                  string
	description           *string
	metadata              *string
	initialCapacity       int32
	maxCapacity           int32
	protectReferencedItem bool
}

func NewInventoryModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	initialCapacity int32,
	maxCapacity int32,
	protectReferencedItem bool,
	description *string,
	metadata *string,
) *InventoryModelMaster {

	self := InventoryModelMaster{
		stack:                 stack,
		namespaceName:         namespaceName,
		name:                  name,
		description:           description,
		metadata:              metadata,
		initialCapacity:       initialCapacity,
		maxCapacity:           maxCapacity,
		protectReferencedItem: protectReferencedItem,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *InventoryModelMaster) ResourceName() string {
	return "Inventory_InventoryModelMaster_" + p.name
}

func (p *InventoryModelMaster) ResourceType() string {
	return "GS2::Inventory::InventoryModelMaster"
}

func (p *InventoryModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["InitialCapacity"] = p.initialCapacity
	properties["MaxCapacity"] = p.maxCapacity
	properties["ProtectReferencedItem"] = p.protectReferencedItem
	return properties
}

func (p *InventoryModelMaster) Ref(
	namespaceName string,
) InventoryModelMasterRef {
	return InventoryModelMasterRef{
		namespaceName: namespaceName,
		inventoryName: p.name,
	}
}

func (p *InventoryModelMaster) GetAttrInventoryModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.InventoryModelId",
	)
}

type ItemModelMaster struct {
	CdkResource
	stack               *Stack
	namespaceName       string
	inventoryName       string
	name                string
	description         *string
	metadata            *string
	stackingLimit       int64
	allowMultipleStacks bool
	sortValue           int32
}

func NewItemModelMaster(
	stack *Stack,
	namespaceName string,
	inventoryName string,
	name string,
	stackingLimit int64,
	allowMultipleStacks bool,
	sortValue int32,
	description *string,
	metadata *string,
) *ItemModelMaster {

	self := ItemModelMaster{
		stack:               stack,
		namespaceName:       namespaceName,
		inventoryName:       inventoryName,
		name:                name,
		description:         description,
		metadata:            metadata,
		stackingLimit:       stackingLimit,
		allowMultipleStacks: allowMultipleStacks,
		sortValue:           sortValue,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *ItemModelMaster) ResourceName() string {
	return "Inventory_ItemModelMaster_" + p.name
}

func (p *ItemModelMaster) ResourceType() string {
	return "GS2::Inventory::ItemModelMaster"
}

func (p *ItemModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["InventoryName"] = p.inventoryName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["StackingLimit"] = p.stackingLimit
	properties["AllowMultipleStacks"] = p.allowMultipleStacks
	properties["SortValue"] = p.sortValue
	return properties
}

func (p *ItemModelMaster) Ref(
	namespaceName string,
	inventoryName string,
) ItemModelMasterRef {
	return ItemModelMasterRef{
		namespaceName: namespaceName,
		inventoryName: inventoryName,
		itemName:      p.name,
	}
}

func (p *ItemModelMaster) GetAttrItemModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ItemModelId",
	)
}
