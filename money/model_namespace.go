package money

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

type NamespacePriority string

const NamespacePriorityFree = NamespacePriority("free")
const NamespacePriorityPaid = NamespacePriority("paid")

func (p NamespacePriority) Pointer() *NamespacePriority {
	return &p
}

type NamespaceCurrency string

const NamespaceCurrencyJpy = NamespaceCurrency("JPY")
const NamespaceCurrencyUsd = NamespaceCurrency("USD")
const NamespaceCurrencyTwd = NamespaceCurrency("TWD")

func (p NamespaceCurrency) Pointer() *NamespaceCurrency {
	return &p
}

type Namespace struct {
	CdkResource
	stack              *Stack
	Name               string
	Description        *string
	Priority           NamespacePriority
	ShareFree          bool
	Currency           NamespaceCurrency
	AppleKey           *string
	GoogleKey          *string
	EnableFakeReceipt  bool
	CreateWalletScript *ScriptSetting
	DepositScript      *ScriptSetting
	WithdrawScript     *ScriptSetting
	LogSetting         *LogSetting
}

type NamespaceOptions struct {
	Description        *string
	AppleKey           *string
	GoogleKey          *string
	EnableFakeReceipt  bool
	CreateWalletScript *ScriptSetting
	DepositScript      *ScriptSetting
	WithdrawScript     *ScriptSetting
	LogSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	priority NamespacePriority,
	shareFree bool,
	currency NamespaceCurrency,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:              stack,
		Name:               name,
		Priority:           priority,
		ShareFree:          shareFree,
		Currency:           currency,
		Description:        options.Description,
		AppleKey:           options.AppleKey,
		GoogleKey:          options.GoogleKey,
		EnableFakeReceipt:  options.EnableFakeReceipt,
		CreateWalletScript: options.CreateWalletScript,
		DepositScript:      options.DepositScript,
		WithdrawScript:     options.WithdrawScript,
		LogSetting:         options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Money_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Money::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["Priority"] = p.Priority
	properties["ShareFree"] = p.ShareFree
	properties["Currency"] = p.Currency
	if p.AppleKey != nil {
		properties["AppleKey"] = p.AppleKey
	}
	if p.GoogleKey != nil {
		properties["GoogleKey"] = p.GoogleKey
	}
	properties["EnableFakeReceipt"] = p.EnableFakeReceipt
	if p.CreateWalletScript != nil {
		properties["CreateWalletScript"] = p.CreateWalletScript.Properties()
	}
	if p.DepositScript != nil {
		properties["DepositScript"] = p.DepositScript.Properties()
	}
	if p.WithdrawScript != nil {
		properties["WithdrawScript"] = p.WithdrawScript.Properties()
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
