package limit

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

type LimitModelResetType string

const LimitModelResetTypeNotReset = LimitModelResetType("notReset")
const LimitModelResetTypeDaily = LimitModelResetType("daily")
const LimitModelResetTypeWeekly = LimitModelResetType("weekly")
const LimitModelResetTypeMonthly = LimitModelResetType("monthly")

func (p LimitModelResetType) Pointer() *LimitModelResetType {
	return &p
}

type LimitModelResetDayOfWeek string

const LimitModelResetDayOfWeekSunday = LimitModelResetDayOfWeek("sunday")
const LimitModelResetDayOfWeekMonday = LimitModelResetDayOfWeek("monday")
const LimitModelResetDayOfWeekTuesday = LimitModelResetDayOfWeek("tuesday")
const LimitModelResetDayOfWeekWednesday = LimitModelResetDayOfWeek("wednesday")
const LimitModelResetDayOfWeekThursday = LimitModelResetDayOfWeek("thursday")
const LimitModelResetDayOfWeekFriday = LimitModelResetDayOfWeek("friday")
const LimitModelResetDayOfWeekSaturday = LimitModelResetDayOfWeek("saturday")

func (p LimitModelResetDayOfWeek) Pointer() *LimitModelResetDayOfWeek {
	return &p
}

type LimitModel struct {
	name            string
	metadata        *string
	resetType       LimitModelResetType
	resetDayOfMonth *int32
	resetDayOfWeek  *LimitModelResetDayOfWeek
	resetHour       *int32
}

func NewLimitModel(
	name string,
	metadata *string,
	resetType LimitModelResetType,
	resetDayOfMonth *int32,
	resetDayOfWeek *LimitModelResetDayOfWeek,
	resetHour *int32,
) *LimitModel {
	return &LimitModel{
		name:            name,
		metadata:        metadata,
		resetType:       resetType,
		resetDayOfMonth: resetDayOfMonth,
		resetDayOfWeek:  resetDayOfWeek,
		resetHour:       resetHour,
	}
}

func (p *LimitModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ResetType"] = p.resetType
	if p.resetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.resetDayOfMonth
	}
	if p.resetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.resetDayOfWeek
	}
	if p.resetHour != nil {
		properties["ResetHour"] = p.resetHour
	}
	return properties
}

func (p *LimitModel) Ref(
	namespaceName string,
) LimitModelRef {
	return LimitModelRef{
		namespaceName: namespaceName,
		limitName:     p.name,
	}
}

func NewNotResetLimitModel(
	name string,
	metadata *string,
) LimitModel {
	return LimitModel{
		resetType: LimitModelResetTypeNotReset,
		name:      name,
		metadata:  metadata,
	}
}

func NewDailyLimitModel(
	name string,
	resetHour int32,
	metadata *string,
) LimitModel {
	return LimitModel{
		resetType: LimitModelResetTypeDaily,
		name:      name,
		metadata:  metadata,
		resetHour: &resetHour,
	}
}

func NewWeeklyLimitModel(
	name string,
	resetDayOfWeek LimitModelResetDayOfWeek,
	resetHour int32,
	metadata *string,
) LimitModel {
	return LimitModel{
		resetType:      LimitModelResetTypeWeekly,
		name:           name,
		metadata:       metadata,
		resetDayOfWeek: &resetDayOfWeek,
		resetHour:      &resetHour,
	}
}

func NewMonthlyLimitModel(
	name string,
	resetDayOfMonth int32,
	resetHour int32,
	metadata *string,
) LimitModel {
	return LimitModel{
		resetType:       LimitModelResetTypeMonthly,
		name:            name,
		metadata:        metadata,
		resetDayOfMonth: &resetDayOfMonth,
		resetHour:       &resetHour,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	limitModels   []LimitModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	limitModels []LimitModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-04-05",
		namespaceName: namespaceName,
		limitModels:   limitModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Limit_CurrentLimitMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Limit::CurrentLimitMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	limitModels := make([]map[string]interface{}, len(p.limitModels))
	for i, item := range p.limitModels {
		limitModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":     p.version,
			"LimitModels": limitModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack       *Stack
	name        string
	description *string
	logSetting  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:       stack,
		name:        name,
		description: description,
		logSetting:  logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Limit_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Limit::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
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
	limitModels []LimitModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		limitModels,
	).AddDependsOn(
		p,
	)
	return p
}

type LimitModelMaster struct {
	CdkResource
	stack           *Stack
	namespaceName   string
	name            string
	description     *string
	metadata        *string
	resetType       string
	resetDayOfMonth *int32
	resetDayOfWeek  *string
	resetHour       *int32
}

func NewLimitModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	resetType string,
	description *string,
	metadata *string,
	resetDayOfMonth *int32,
	resetDayOfWeek *string,
	resetHour *int32,
) *LimitModelMaster {

	self := LimitModelMaster{
		stack:           stack,
		namespaceName:   namespaceName,
		name:            name,
		description:     description,
		metadata:        metadata,
		resetType:       resetType,
		resetDayOfMonth: resetDayOfMonth,
		resetDayOfWeek:  resetDayOfWeek,
		resetHour:       resetHour,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *LimitModelMaster) ResourceName() string {
	return "Limit_LimitModelMaster_" + p.name
}

func (p *LimitModelMaster) ResourceType() string {
	return "GS2::Limit::LimitModelMaster"
}

func (p *LimitModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ResetType"] = p.resetType
	if p.resetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.resetDayOfMonth
	}
	if p.resetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.resetDayOfWeek
	}
	if p.resetHour != nil {
		properties["ResetHour"] = p.resetHour
	}
	return properties
}

func (p *LimitModelMaster) Ref(
	namespaceName string,
) LimitModelMasterRef {
	return LimitModelMasterRef{
		namespaceName: namespaceName,
		limitName:     p.name,
	}
}

func (p *LimitModelMaster) GetAttrLimitModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.LimitModelId",
	)
}
