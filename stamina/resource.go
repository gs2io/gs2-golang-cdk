package stamina

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

type StaminaModel struct {
	name                   string
	metadata               *string
	recoverIntervalMinutes int32
	recoverValue           int32
	initialCapacity        int32
	isOverflow             bool
	maxCapacity            *int32
	maxStaminaTable        *MaxStaminaTable
	recoverIntervalTable   *RecoverIntervalTable
	recoverValueTable      *RecoverValueTable
}

func NewStaminaModel(
	name string,
	metadata *string,
	recoverIntervalMinutes int32,
	recoverValue int32,
	initialCapacity int32,
	isOverflow bool,
	maxCapacity *int32,
	maxStaminaTable *MaxStaminaTable,
	recoverIntervalTable *RecoverIntervalTable,
	recoverValueTable *RecoverValueTable,
) *StaminaModel {
	return &StaminaModel{
		name:                   name,
		metadata:               metadata,
		recoverIntervalMinutes: recoverIntervalMinutes,
		recoverValue:           recoverValue,
		initialCapacity:        initialCapacity,
		isOverflow:             isOverflow,
		maxCapacity:            maxCapacity,
		maxStaminaTable:        maxStaminaTable,
		recoverIntervalTable:   recoverIntervalTable,
		recoverValueTable:      recoverValueTable,
	}
}

func (p *StaminaModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["RecoverIntervalMinutes"] = p.recoverIntervalMinutes
	properties["RecoverValue"] = p.recoverValue
	properties["InitialCapacity"] = p.initialCapacity
	properties["IsOverflow"] = p.isOverflow
	if p.maxCapacity != nil {
		properties["MaxCapacity"] = p.maxCapacity
	}
	if p.maxStaminaTable != nil {
		properties["MaxStaminaTable"] = p.maxStaminaTable.Properties()
	}
	if p.recoverIntervalTable != nil {
		properties["RecoverIntervalTable"] = p.recoverIntervalTable.Properties()
	}
	if p.recoverValueTable != nil {
		properties["RecoverValueTable"] = p.recoverValueTable.Properties()
	}
	return properties
}

func (p *StaminaModel) Ref(
	namespaceName string,
) StaminaModelRef {
	return StaminaModelRef{
		namespaceName: namespaceName,
		staminaName:   p.name,
	}
}

type MaxStaminaTable struct {
	name              string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewMaxStaminaTable(
	name string,
	metadata *string,
	experienceModelId string,
	values []int32,
) *MaxStaminaTable {
	return &MaxStaminaTable{
		name:              name,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
}

func (p *MaxStaminaTable) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *MaxStaminaTable) Ref(
	namespaceName string,
) MaxStaminaTableRef {
	return MaxStaminaTableRef{
		namespaceName:       namespaceName,
		maxStaminaTableName: p.name,
	}
}

type RecoverIntervalTable struct {
	name              string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewRecoverIntervalTable(
	name string,
	metadata *string,
	experienceModelId string,
	values []int32,
) *RecoverIntervalTable {
	return &RecoverIntervalTable{
		name:              name,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
}

func (p *RecoverIntervalTable) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *RecoverIntervalTable) Ref(
	namespaceName string,
) RecoverIntervalTableRef {
	return RecoverIntervalTableRef{
		namespaceName:            namespaceName,
		recoverIntervalTableName: p.name,
	}
}

type RecoverValueTable struct {
	name              string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewRecoverValueTable(
	name string,
	metadata *string,
	experienceModelId string,
	values []int32,
) *RecoverValueTable {
	return &RecoverValueTable{
		name:              name,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
}

func (p *RecoverValueTable) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *RecoverValueTable) Ref(
	namespaceName string,
) RecoverValueTableRef {
	return RecoverValueTableRef{
		namespaceName:         namespaceName,
		recoverValueTableName: p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	staminaModels []StaminaModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	staminaModels []StaminaModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-02-14",
		namespaceName: namespaceName,
		staminaModels: staminaModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Stamina_CurrentStaminaMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Stamina::CurrentStaminaMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	staminaModels := make([]map[string]interface{}, len(p.staminaModels))
	for i, item := range p.staminaModels {
		staminaModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":       p.version,
			"StaminaModels": staminaModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                 *Stack
	name                  string
	description           *string
	overflowTriggerScript *ScriptSetting
	logSetting            *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	overflowTriggerScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                 stack,
		name:                  name,
		description:           description,
		overflowTriggerScript: overflowTriggerScript,
		logSetting:            logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Stamina_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Stamina::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.overflowTriggerScript != nil {
		properties["OverflowTriggerScript"] = p.overflowTriggerScript.Properties()
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
	staminaModels []StaminaModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		staminaModels,
	).AddDependsOn(
		p,
	)
	return p
}

type StaminaModelMaster struct {
	CdkResource
	stack                    *Stack
	namespaceName            string
	name                     string
	description              *string
	metadata                 *string
	recoverIntervalMinutes   int32
	recoverValue             int32
	initialCapacity          int32
	isOverflow               bool
	maxCapacity              *int32
	maxStaminaTableName      *string
	recoverIntervalTableName *string
	recoverValueTableName    *string
}

func NewStaminaModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	recoverIntervalMinutes int32,
	recoverValue int32,
	initialCapacity int32,
	isOverflow bool,
	description *string,
	metadata *string,
	maxCapacity *int32,
	maxStaminaTableName *string,
	recoverIntervalTableName *string,
	recoverValueTableName *string,
) *StaminaModelMaster {

	self := StaminaModelMaster{
		stack:                    stack,
		namespaceName:            namespaceName,
		name:                     name,
		description:              description,
		metadata:                 metadata,
		recoverIntervalMinutes:   recoverIntervalMinutes,
		recoverValue:             recoverValue,
		initialCapacity:          initialCapacity,
		isOverflow:               isOverflow,
		maxCapacity:              maxCapacity,
		maxStaminaTableName:      maxStaminaTableName,
		recoverIntervalTableName: recoverIntervalTableName,
		recoverValueTableName:    recoverValueTableName,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *StaminaModelMaster) ResourceName() string {
	return "Stamina_StaminaModelMaster_" + p.name
}

func (p *StaminaModelMaster) ResourceType() string {
	return "GS2::Stamina::StaminaModelMaster"
}

func (p *StaminaModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["RecoverIntervalMinutes"] = p.recoverIntervalMinutes
	properties["RecoverValue"] = p.recoverValue
	properties["InitialCapacity"] = p.initialCapacity
	properties["IsOverflow"] = p.isOverflow
	if p.maxCapacity != nil {
		properties["MaxCapacity"] = p.maxCapacity
	}
	if p.maxStaminaTableName != nil {
		properties["MaxStaminaTableName"] = p.maxStaminaTableName
	}
	if p.recoverIntervalTableName != nil {
		properties["RecoverIntervalTableName"] = p.recoverIntervalTableName
	}
	if p.recoverValueTableName != nil {
		properties["RecoverValueTableName"] = p.recoverValueTableName
	}
	return properties
}

func (p *StaminaModelMaster) Ref(
	namespaceName string,
) StaminaModelMasterRef {
	return StaminaModelMasterRef{
		namespaceName: namespaceName,
		staminaName:   p.name,
	}
}

func (p *StaminaModelMaster) GetAttrStaminaModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.StaminaModelId",
	)
}

type MaxStaminaTableMaster struct {
	CdkResource
	stack             *Stack
	namespaceName     string
	name              string
	description       *string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewMaxStaminaTableMaster(
	stack *Stack,
	namespaceName string,
	name string,
	experienceModelId string,
	values []int32,
	description *string,
	metadata *string,
) *MaxStaminaTableMaster {

	self := MaxStaminaTableMaster{
		stack:             stack,
		namespaceName:     namespaceName,
		name:              name,
		description:       description,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *MaxStaminaTableMaster) ResourceName() string {
	return "Stamina_MaxStaminaTableMaster_" + p.name
}

func (p *MaxStaminaTableMaster) ResourceType() string {
	return "GS2::Stamina::MaxStaminaTableMaster"
}

func (p *MaxStaminaTableMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *MaxStaminaTableMaster) Ref(
	namespaceName string,
) MaxStaminaTableMasterRef {
	return MaxStaminaTableMasterRef{
		namespaceName:       namespaceName,
		maxStaminaTableName: p.name,
	}
}

func (p *MaxStaminaTableMaster) GetAttrMaxStaminaTableId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.MaxStaminaTableId",
	)
}

type RecoverIntervalTableMaster struct {
	CdkResource
	stack             *Stack
	namespaceName     string
	name              string
	description       *string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewRecoverIntervalTableMaster(
	stack *Stack,
	namespaceName string,
	name string,
	experienceModelId string,
	values []int32,
	description *string,
	metadata *string,
) *RecoverIntervalTableMaster {

	self := RecoverIntervalTableMaster{
		stack:             stack,
		namespaceName:     namespaceName,
		name:              name,
		description:       description,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *RecoverIntervalTableMaster) ResourceName() string {
	return "Stamina_RecoverIntervalTableMaster_" + p.name
}

func (p *RecoverIntervalTableMaster) ResourceType() string {
	return "GS2::Stamina::RecoverIntervalTableMaster"
}

func (p *RecoverIntervalTableMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *RecoverIntervalTableMaster) Ref(
	namespaceName string,
) RecoverIntervalTableMasterRef {
	return RecoverIntervalTableMasterRef{
		namespaceName:            namespaceName,
		recoverIntervalTableName: p.name,
	}
}

func (p *RecoverIntervalTableMaster) GetAttrRecoverIntervalTableId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RecoverIntervalTableId",
	)
}

type RecoverValueTableMaster struct {
	CdkResource
	stack             *Stack
	namespaceName     string
	name              string
	description       *string
	metadata          *string
	experienceModelId string
	values            []int32
}

func NewRecoverValueTableMaster(
	stack *Stack,
	namespaceName string,
	name string,
	experienceModelId string,
	values []int32,
	description *string,
	metadata *string,
) *RecoverValueTableMaster {

	self := RecoverValueTableMaster{
		stack:             stack,
		namespaceName:     namespaceName,
		name:              name,
		description:       description,
		metadata:          metadata,
		experienceModelId: experienceModelId,
		values:            values,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *RecoverValueTableMaster) ResourceName() string {
	return "Stamina_RecoverValueTableMaster_" + p.name
}

func (p *RecoverValueTableMaster) ResourceType() string {
	return "GS2::Stamina::RecoverValueTableMaster"
}

func (p *RecoverValueTableMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ExperienceModelId"] = p.experienceModelId
	properties["Values"] = p.values
	return properties
}

func (p *RecoverValueTableMaster) Ref(
	namespaceName string,
) RecoverValueTableMasterRef {
	return RecoverValueTableMasterRef{
		namespaceName:         namespaceName,
		recoverValueTableName: p.name,
	}
}

func (p *RecoverValueTableMaster) GetAttrRecoverValueTableId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RecoverValueTableId",
	)
}
