package dictionary

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

type EntryModel struct {
	name     string
	metadata *string
}

func NewEntryModel(
	name string,
	metadata *string,
) *EntryModel {
	return &EntryModel{
		name:     name,
		metadata: metadata,
	}
}

func (p *EntryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

func (p *EntryModel) Ref(
	namespaceName string,
) EntryModelRef {
	return EntryModelRef{
		namespaceName: namespaceName,
		entryName:     p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	entryModels   []EntryModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	entryModels []EntryModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2020-04-30",
		namespaceName: namespaceName,
		entryModels:   entryModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Dictionary_CurrentEntryMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Dictionary::CurrentEntryMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	entryModels := make([]map[string]interface{}, len(p.entryModels))
	for i, item := range p.entryModels {
		entryModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":     p.version,
			"EntryModels": entryModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                *Stack
	name                 string
	description          *string
	entryScript          *ScriptSetting
	duplicateEntryScript *ScriptSetting
	logSetting           *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	entryScript *ScriptSetting,
	duplicateEntryScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                stack,
		name:                 name,
		description:          description,
		entryScript:          entryScript,
		duplicateEntryScript: duplicateEntryScript,
		logSetting:           logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Dictionary_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Dictionary::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.entryScript != nil {
		properties["EntryScript"] = p.entryScript.Properties()
	}
	if p.duplicateEntryScript != nil {
		properties["DuplicateEntryScript"] = p.duplicateEntryScript.Properties()
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
	entryModels []EntryModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		entryModels,
	).AddDependsOn(
		p,
	)
	return p
}

type EntryModelMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
}

func NewEntryModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	description *string,
	metadata *string,
) *EntryModelMaster {

	self := EntryModelMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *EntryModelMaster) ResourceName() string {
	return "Dictionary_EntryModelMaster_" + p.name
}

func (p *EntryModelMaster) ResourceType() string {
	return "GS2::Dictionary::EntryModelMaster"
}

func (p *EntryModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

func (p *EntryModelMaster) Ref(
	namespaceName string,
) EntryModelMasterRef {
	return EntryModelMasterRef{
		namespaceName: namespaceName,
		entryName:     p.name,
	}
}

func (p *EntryModelMaster) GetAttrEntryModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.EntryModelId",
	)
}
