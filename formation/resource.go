package formation

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

type FormModel struct {
	name     string
	metadata *string
	slots    []SlotModel
}

func NewFormModel(
	name string,
	metadata *string,
	slots []SlotModel,
) *FormModel {
	return &FormModel{
		name:     name,
		metadata: metadata,
		slots:    slots,
	}
}

func (p *FormModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.slots))
		for i, element := range p.slots {
			values[i] = element.Properties()
		}
		properties["Slots"] = values
	}
	return properties
}

func (p *FormModel) Ref(
	namespaceName string,
) FormModelRef {
	return FormModelRef{
		namespaceName: namespaceName,
		formModelName: p.name,
	}
}

type MoldModel struct {
	name               string
	metadata           *string
	initialMaxCapacity int32
	maxCapacity        int32
	formModel          FormModel
}

func NewMoldModel(
	name string,
	metadata *string,
	initialMaxCapacity int32,
	maxCapacity int32,
	formModel FormModel,
) *MoldModel {
	return &MoldModel{
		name:               name,
		metadata:           metadata,
		initialMaxCapacity: initialMaxCapacity,
		maxCapacity:        maxCapacity,
		formModel:          formModel,
	}
}

func (p *MoldModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["InitialMaxCapacity"] = p.initialMaxCapacity
	properties["MaxCapacity"] = p.maxCapacity
	properties["FormModel"] = p.formModel.Properties()
	return properties
}

func (p *MoldModel) Ref(
	namespaceName string,
) MoldModelRef {
	return MoldModelRef{
		namespaceName: namespaceName,
		moldName:      p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	moldModels    []MoldModel
	formModels    []FormModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	moldModels []MoldModel,
	formModels []FormModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-09-09",
		namespaceName: namespaceName,
		moldModels:    moldModels,
		formModels:    formModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Formation_CurrentFormMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Formation::CurrentFormMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	moldModels := make([]map[string]interface{}, len(p.moldModels))
	for i, item := range p.moldModels {
		moldModels[i] = item.Properties()
	}
	formModels := make([]map[string]interface{}, len(p.formModels))
	for i, item := range p.formModels {
		formModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":    p.version,
			"MoldModels": moldModels,
			"FormModels": formModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack              *Stack
	name               string
	description        *string
	transactionSetting *TransactionSetting
	updateMoldScript   *ScriptSetting
	updateFormScript   *ScriptSetting
	logSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	transactionSetting *TransactionSetting,
	updateMoldScript *ScriptSetting,
	updateFormScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:              stack,
		name:               name,
		description:        description,
		transactionSetting: transactionSetting,
		updateMoldScript:   updateMoldScript,
		updateFormScript:   updateFormScript,
		logSetting:         logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Formation_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Formation::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.transactionSetting != nil {
		properties["TransactionSetting"] = p.transactionSetting.Properties()
	}
	if p.updateMoldScript != nil {
		properties["UpdateMoldScript"] = p.updateMoldScript.Properties()
	}
	if p.updateFormScript != nil {
		properties["UpdateFormScript"] = p.updateFormScript.Properties()
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
	moldModels []MoldModel,
	formModels []FormModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		moldModels,
		formModels,
	).AddDependsOn(
		p,
	)
	return p
}

type FormModelMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
	slots         []SlotModel
}

func NewFormModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	slots []SlotModel,
	description *string,
	metadata *string,
) *FormModelMaster {

	self := FormModelMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
		slots:         slots,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *FormModelMaster) ResourceName() string {
	return "Formation_FormModelMaster_" + p.name
}

func (p *FormModelMaster) ResourceType() string {
	return "GS2::Formation::FormModelMaster"
}

func (p *FormModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.slots))
		for i, element := range p.slots {
			values[i] = element.Properties()
		}
		properties["Slots"] = values
	}
	return properties
}

func (p *FormModelMaster) Ref(
	namespaceName string,
) FormModelMasterRef {
	return FormModelMasterRef{
		namespaceName: namespaceName,
		formModelName: p.name,
	}
}

func (p *FormModelMaster) GetAttrFormModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.FormModelId",
	)
}

type MoldModelMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	description        *string
	metadata           *string
	formModelName      string
	initialMaxCapacity int32
	maxCapacity        int32
}

func NewMoldModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	formModelName string,
	initialMaxCapacity int32,
	maxCapacity int32,
	description *string,
	metadata *string,
) *MoldModelMaster {

	self := MoldModelMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		description:        description,
		metadata:           metadata,
		formModelName:      formModelName,
		initialMaxCapacity: initialMaxCapacity,
		maxCapacity:        maxCapacity,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *MoldModelMaster) ResourceName() string {
	return "Formation_MoldModelMaster_" + p.name
}

func (p *MoldModelMaster) ResourceType() string {
	return "GS2::Formation::MoldModelMaster"
}

func (p *MoldModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["FormModelName"] = p.formModelName
	properties["InitialMaxCapacity"] = p.initialMaxCapacity
	properties["MaxCapacity"] = p.maxCapacity
	return properties
}

func (p *MoldModelMaster) Ref(
	namespaceName string,
) MoldModelMasterRef {
	return MoldModelMasterRef{
		namespaceName: namespaceName,
		moldName:      p.name,
	}
}

func (p *MoldModelMaster) GetAttrMoldModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.MoldModelId",
	)
}
