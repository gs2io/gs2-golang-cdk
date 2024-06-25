package money2

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

type NamespaceCurrencyUsagePriority string

const NamespaceCurrencyUsagePriorityPrioritizeFree = NamespaceCurrencyUsagePriority("PrioritizeFree")
const NamespaceCurrencyUsagePriorityPrioritizePaid = NamespaceCurrencyUsagePriority("PrioritizePaid")

func (p NamespaceCurrencyUsagePriority) Pointer() *NamespaceCurrencyUsagePriority {
	return &p
}

type Namespace struct {
	CdkResource
	stack                 *Stack
	Name                  string
	CurrencyUsagePriority NamespaceCurrencyUsagePriority
	Description           *string
	SharedFreeCurrency    bool
	PlatformSetting       PlatformSetting
	ChangeBalanceScript   *ScriptSetting
	LogSetting            *LogSetting
}

type NamespaceOptions struct {
	Description         *string
	ChangeBalanceScript *ScriptSetting
	LogSetting          *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	currencyUsagePriority NamespaceCurrencyUsagePriority,
	sharedFreeCurrency bool,
	platformSetting PlatformSetting,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                 stack,
		Name:                  name,
		CurrencyUsagePriority: currencyUsagePriority,
		SharedFreeCurrency:    sharedFreeCurrency,
		PlatformSetting:       platformSetting,
		Description:           options.Description,
		ChangeBalanceScript:   options.ChangeBalanceScript,
		LogSetting:            options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Money2_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Money2::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["CurrencyUsagePriority"] = p.CurrencyUsagePriority
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["SharedFreeCurrency"] = p.SharedFreeCurrency
	properties["PlatformSetting"] = p.PlatformSetting.Properties()
	if p.ChangeBalanceScript != nil {
		properties["ChangeBalanceScript"] = p.ChangeBalanceScript.Properties()
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
	storeContentModels []StoreContentModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		storeContentModels,
	).AddDependsOn(
		p,
	)
	return p
}
