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

type Namespace struct {
	CdkResource
	stack                   *Stack
	Name                    string
	Description             *string
	AcquireScript           *ScriptSetting
	OverflowScript          *ScriptSetting
	ConsumeScript           *ScriptSetting
	SimpleItemAcquireScript *ScriptSetting
	SimpleItemConsumeScript *ScriptSetting
	BigItemAcquireScript    *ScriptSetting
	BigItemConsumeScript    *ScriptSetting
	LogSetting              *LogSetting
}

type NamespaceOptions struct {
	Description             *string
	AcquireScript           *ScriptSetting
	OverflowScript          *ScriptSetting
	ConsumeScript           *ScriptSetting
	SimpleItemAcquireScript *ScriptSetting
	SimpleItemConsumeScript *ScriptSetting
	BigItemAcquireScript    *ScriptSetting
	BigItemConsumeScript    *ScriptSetting
	LogSetting              *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                   stack,
		Name:                    name,
		Description:             options.Description,
		AcquireScript:           options.AcquireScript,
		OverflowScript:          options.OverflowScript,
		ConsumeScript:           options.ConsumeScript,
		SimpleItemAcquireScript: options.SimpleItemAcquireScript,
		SimpleItemConsumeScript: options.SimpleItemConsumeScript,
		BigItemAcquireScript:    options.BigItemAcquireScript,
		BigItemConsumeScript:    options.BigItemConsumeScript,
		LogSetting:              options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Inventory_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Inventory::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.AcquireScript != nil {
		properties["AcquireScript"] = p.AcquireScript.Properties()
	}
	if p.OverflowScript != nil {
		properties["OverflowScript"] = p.OverflowScript.Properties()
	}
	if p.ConsumeScript != nil {
		properties["ConsumeScript"] = p.ConsumeScript.Properties()
	}
	if p.SimpleItemAcquireScript != nil {
		properties["SimpleItemAcquireScript"] = p.SimpleItemAcquireScript.Properties()
	}
	if p.SimpleItemConsumeScript != nil {
		properties["SimpleItemConsumeScript"] = p.SimpleItemConsumeScript.Properties()
	}
	if p.BigItemAcquireScript != nil {
		properties["BigItemAcquireScript"] = p.BigItemAcquireScript.Properties()
	}
	if p.BigItemConsumeScript != nil {
		properties["BigItemConsumeScript"] = p.BigItemConsumeScript.Properties()
	}
	if p.LogSetting != nil {
		properties["LogSetting"] = p.LogSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		NamespaceName: p.Name,
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
	simpleInventoryModels []SimpleInventoryModel,
	bigInventoryModels []BigInventoryModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		inventoryModels,
		simpleInventoryModels,
		bigInventoryModels,
	).AddDependsOn(
		p,
	)
	return p
}
