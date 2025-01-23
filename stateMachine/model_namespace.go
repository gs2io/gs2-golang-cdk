package stateMachine

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

type NamespaceSupportSpeculativeExecution string

const NamespaceSupportSpeculativeExecutionEnable = NamespaceSupportSpeculativeExecution("enable")
const NamespaceSupportSpeculativeExecutionDisable = NamespaceSupportSpeculativeExecution("disable")

func (p NamespaceSupportSpeculativeExecution) Pointer() *NamespaceSupportSpeculativeExecution {
	return &p
}

type Namespace struct {
	CdkResource
	stack                       *Stack
	Name                        string
	Description                 *string
	SupportSpeculativeExecution NamespaceSupportSpeculativeExecution
	TransactionSetting          *TransactionSetting
	StartScript                 *ScriptSetting
	PassScript                  *ScriptSetting
	ErrorScript                 *ScriptSetting
	LowestStateMachineVersion   *int64
	LogSetting                  *LogSetting
}

type NamespaceOptions struct {
	Description                 *string
	SupportSpeculativeExecution NamespaceSupportSpeculativeExecution
	TransactionSetting          *TransactionSetting
	StartScript                 *ScriptSetting
	PassScript                  *ScriptSetting
	ErrorScript                 *ScriptSetting
	LowestStateMachineVersion   *int64
	LogSetting                  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                       stack,
		Name:                        name,
		Description:                 options.Description,
		SupportSpeculativeExecution: options.SupportSpeculativeExecution,
		TransactionSetting:          options.TransactionSetting,
		StartScript:                 options.StartScript,
		PassScript:                  options.PassScript,
		ErrorScript:                 options.ErrorScript,
		LowestStateMachineVersion:   options.LowestStateMachineVersion,
		LogSetting:                  options.LogSetting,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "StateMachine_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::StateMachine::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["SupportSpeculativeExecution"] = p.SupportSpeculativeExecution
	if p.TransactionSetting != nil {
		properties["TransactionSetting"] = p.TransactionSetting.Properties()
	}
	if p.StartScript != nil {
		properties["StartScript"] = p.StartScript.Properties()
	}
	if p.PassScript != nil {
		properties["PassScript"] = p.PassScript.Properties()
	}
	if p.ErrorScript != nil {
		properties["ErrorScript"] = p.ErrorScript.Properties()
	}
	if p.LowestStateMachineVersion != nil {
		properties["LowestStateMachineVersion"] = p.LowestStateMachineVersion
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
