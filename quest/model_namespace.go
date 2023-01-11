package quest

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
	stack               *Stack
	Name                string
	Description         *string
	TransactionSetting  TransactionSetting
	StartQuestScript    *ScriptSetting
	CompleteQuestScript *ScriptSetting
	FailedQuestScript   *ScriptSetting
	LogSetting          *LogSetting
}

type NamespaceOptions struct {
	Description         *string
	StartQuestScript    *ScriptSetting
	CompleteQuestScript *ScriptSetting
	FailedQuestScript   *ScriptSetting
	LogSetting          *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	transactionSetting TransactionSetting,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:               stack,
		Name:                name,
		TransactionSetting:  transactionSetting,
		Description:         options.Description,
		StartQuestScript:    options.StartQuestScript,
		CompleteQuestScript: options.CompleteQuestScript,
		FailedQuestScript:   options.FailedQuestScript,
		LogSetting:          options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Quest_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Quest::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["TransactionSetting"] = p.TransactionSetting.Properties()
	if p.StartQuestScript != nil {
		properties["StartQuestScript"] = p.StartQuestScript.Properties()
	}
	if p.CompleteQuestScript != nil {
		properties["CompleteQuestScript"] = p.CompleteQuestScript.Properties()
	}
	if p.FailedQuestScript != nil {
		properties["FailedQuestScript"] = p.FailedQuestScript.Properties()
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
	questGroupModels []QuestGroupModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		questGroupModels,
	).AddDependsOn(
		p,
	)
	return p
}
