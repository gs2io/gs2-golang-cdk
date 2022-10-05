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

type Namespace struct {
	CdkResource
	stack              *Stack
	name               string
	description        *string
	priority           string
	shareFree          bool
	currency           string
	appleKey           *string
	googleKey          *string
	enableFakeReceipt  bool
	createWalletScript *ScriptSetting
	depositScript      *ScriptSetting
	withdrawScript     *ScriptSetting
	logSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	priority string,
	shareFree bool,
	currency string,
	enableFakeReceipt bool,
	description *string,
	appleKey *string,
	googleKey *string,
	createWalletScript *ScriptSetting,
	depositScript *ScriptSetting,
	withdrawScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:              stack,
		name:               name,
		description:        description,
		priority:           priority,
		shareFree:          shareFree,
		currency:           currency,
		appleKey:           appleKey,
		googleKey:          googleKey,
		enableFakeReceipt:  enableFakeReceipt,
		createWalletScript: createWalletScript,
		depositScript:      depositScript,
		withdrawScript:     withdrawScript,
		logSetting:         logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Money_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Money::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["Priority"] = p.priority
	properties["ShareFree"] = p.shareFree
	properties["Currency"] = p.currency
	if p.appleKey != nil {
		properties["AppleKey"] = p.appleKey
	}
	if p.googleKey != nil {
		properties["GoogleKey"] = p.googleKey
	}
	properties["EnableFakeReceipt"] = p.enableFakeReceipt
	if p.createWalletScript != nil {
		properties["CreateWalletScript"] = p.createWalletScript.Properties()
	}
	if p.depositScript != nil {
		properties["DepositScript"] = p.depositScript.Properties()
	}
	if p.withdrawScript != nil {
		properties["WithdrawScript"] = p.withdrawScript.Properties()
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
