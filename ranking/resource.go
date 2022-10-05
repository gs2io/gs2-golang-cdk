package ranking

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

type CategoryModelOrderDirection string

const CategoryModelOrderDirectionAsc = CategoryModelOrderDirection("asc")
const CategoryModelOrderDirectionDesc = CategoryModelOrderDirection("desc")

func (p CategoryModelOrderDirection) Pointer() *CategoryModelOrderDirection {
	return &p
}

type CategoryModelScope string

const CategoryModelScopeGlobal = CategoryModelScope("global")
const CategoryModelScopeScoped = CategoryModelScope("scoped")

func (p CategoryModelScope) Pointer() *CategoryModelScope {
	return &p
}

type CategoryModel struct {
	name                       string
	metadata                   *string
	minimumValue               *int64
	maximumValue               *int64
	orderDirection             CategoryModelOrderDirection
	scope                      CategoryModelScope
	uniqueByUserId             bool
	calculateFixedTimingHour   *int32
	calculateFixedTimingMinute *int32
	calculateIntervalMinutes   *int32
	entryPeriodEventId         *string
	accessPeriodEventId        *string
	generation                 *string
}

func NewCategoryModel(
	name string,
	metadata *string,
	minimumValue *int64,
	maximumValue *int64,
	orderDirection CategoryModelOrderDirection,
	scope CategoryModelScope,
	uniqueByUserId bool,
	calculateFixedTimingHour *int32,
	calculateFixedTimingMinute *int32,
	calculateIntervalMinutes *int32,
	entryPeriodEventId *string,
	accessPeriodEventId *string,
	generation *string,
) *CategoryModel {
	return &CategoryModel{
		name:                       name,
		metadata:                   metadata,
		minimumValue:               minimumValue,
		maximumValue:               maximumValue,
		orderDirection:             orderDirection,
		scope:                      scope,
		uniqueByUserId:             uniqueByUserId,
		calculateFixedTimingHour:   calculateFixedTimingHour,
		calculateFixedTimingMinute: calculateFixedTimingMinute,
		calculateIntervalMinutes:   calculateIntervalMinutes,
		entryPeriodEventId:         entryPeriodEventId,
		accessPeriodEventId:        accessPeriodEventId,
		generation:                 generation,
	}
}

func (p *CategoryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.minimumValue != nil {
		properties["MinimumValue"] = p.minimumValue
	}
	if p.maximumValue != nil {
		properties["MaximumValue"] = p.maximumValue
	}
	properties["OrderDirection"] = p.orderDirection
	properties["Scope"] = p.scope
	properties["UniqueByUserId"] = p.uniqueByUserId
	if p.calculateFixedTimingHour != nil {
		properties["CalculateFixedTimingHour"] = p.calculateFixedTimingHour
	}
	if p.calculateFixedTimingMinute != nil {
		properties["CalculateFixedTimingMinute"] = p.calculateFixedTimingMinute
	}
	if p.calculateIntervalMinutes != nil {
		properties["CalculateIntervalMinutes"] = p.calculateIntervalMinutes
	}
	if p.entryPeriodEventId != nil {
		properties["EntryPeriodEventId"] = p.entryPeriodEventId
	}
	if p.accessPeriodEventId != nil {
		properties["AccessPeriodEventId"] = p.accessPeriodEventId
	}
	if p.generation != nil {
		properties["Generation"] = p.generation
	}
	return properties
}

func (p *CategoryModel) Ref(
	namespaceName string,
) CategoryModelRef {
	return CategoryModelRef{
		namespaceName: namespaceName,
		categoryName:  p.name,
	}
}

func NewGlobalCategoryModel(
	name string,
	orderDirection CategoryModelOrderDirection,
	uniqueByUserId bool,
	calculateIntervalMinutes int32,
	metadata *string,
	minimumValue *int64,
	maximumValue *int64,
	calculateFixedTimingHour *int32,
	calculateFixedTimingMinute *int32,
	entryPeriodEventId *string,
	accessPeriodEventId *string,
	generation *string,
) CategoryModel {
	return CategoryModel{
		scope:                      CategoryModelScopeGlobal,
		name:                       name,
		metadata:                   metadata,
		minimumValue:               minimumValue,
		maximumValue:               maximumValue,
		orderDirection:             orderDirection,
		uniqueByUserId:             uniqueByUserId,
		calculateFixedTimingHour:   calculateFixedTimingHour,
		calculateFixedTimingMinute: calculateFixedTimingMinute,
		calculateIntervalMinutes:   &calculateIntervalMinutes,
		entryPeriodEventId:         entryPeriodEventId,
		accessPeriodEventId:        accessPeriodEventId,
		generation:                 generation,
	}
}

func NewScopedCategoryModel(
	name string,
	orderDirection CategoryModelOrderDirection,
	uniqueByUserId bool,
	metadata *string,
	minimumValue *int64,
	maximumValue *int64,
	calculateFixedTimingHour *int32,
	calculateFixedTimingMinute *int32,
	entryPeriodEventId *string,
	accessPeriodEventId *string,
	generation *string,
) CategoryModel {
	return CategoryModel{
		scope:                      CategoryModelScopeScoped,
		name:                       name,
		metadata:                   metadata,
		minimumValue:               minimumValue,
		maximumValue:               maximumValue,
		orderDirection:             orderDirection,
		uniqueByUserId:             uniqueByUserId,
		calculateFixedTimingHour:   calculateFixedTimingHour,
		calculateFixedTimingMinute: calculateFixedTimingMinute,
		entryPeriodEventId:         entryPeriodEventId,
		accessPeriodEventId:        accessPeriodEventId,
		generation:                 generation,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack          *Stack
	version        string
	namespaceName  string
	categoryModels []CategoryModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	categoryModels []CategoryModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:          stack,
		version:        "2019-09-17",
		namespaceName:  namespaceName,
		categoryModels: categoryModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Ranking_CurrentRankingMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Ranking::CurrentRankingMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	categoryModels := make([]map[string]interface{}, len(p.categoryModels))
	for i, item := range p.categoryModels {
		categoryModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":        p.version,
			"CategoryModels": categoryModels,
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
	return "Ranking_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Ranking::Namespace"
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
	categoryModels []CategoryModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		categoryModels,
	).AddDependsOn(
		p,
	)
	return p
}

type CategoryModelMaster struct {
	CdkResource
	stack                      *Stack
	namespaceName              string
	name                       string
	description                *string
	metadata                   *string
	minimumValue               *int64
	maximumValue               *int64
	orderDirection             string
	scope                      string
	uniqueByUserId             *bool
	calculateFixedTimingHour   *int32
	calculateFixedTimingMinute *int32
	calculateIntervalMinutes   *int32
	entryPeriodEventId         *string
	accessPeriodEventId        *string
	generation                 *string
}

func NewCategoryModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	orderDirection string,
	scope string,
	description *string,
	metadata *string,
	minimumValue *int64,
	maximumValue *int64,
	uniqueByUserId *bool,
	calculateFixedTimingHour *int32,
	calculateFixedTimingMinute *int32,
	calculateIntervalMinutes *int32,
	entryPeriodEventId *string,
	accessPeriodEventId *string,
	generation *string,
) *CategoryModelMaster {

	self := CategoryModelMaster{
		stack:                      stack,
		namespaceName:              namespaceName,
		name:                       name,
		description:                description,
		metadata:                   metadata,
		minimumValue:               minimumValue,
		maximumValue:               maximumValue,
		orderDirection:             orderDirection,
		scope:                      scope,
		uniqueByUserId:             uniqueByUserId,
		calculateFixedTimingHour:   calculateFixedTimingHour,
		calculateFixedTimingMinute: calculateFixedTimingMinute,
		calculateIntervalMinutes:   calculateIntervalMinutes,
		entryPeriodEventId:         entryPeriodEventId,
		accessPeriodEventId:        accessPeriodEventId,
		generation:                 generation,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *CategoryModelMaster) ResourceName() string {
	return "Ranking_CategoryModelMaster_" + p.name
}

func (p *CategoryModelMaster) ResourceType() string {
	return "GS2::Ranking::CategoryModelMaster"
}

func (p *CategoryModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.minimumValue != nil {
		properties["MinimumValue"] = p.minimumValue
	}
	if p.maximumValue != nil {
		properties["MaximumValue"] = p.maximumValue
	}
	properties["OrderDirection"] = p.orderDirection
	properties["Scope"] = p.scope
	if p.uniqueByUserId != nil {
		properties["UniqueByUserId"] = p.uniqueByUserId
	}
	if p.calculateFixedTimingHour != nil {
		properties["CalculateFixedTimingHour"] = p.calculateFixedTimingHour
	}
	if p.calculateFixedTimingMinute != nil {
		properties["CalculateFixedTimingMinute"] = p.calculateFixedTimingMinute
	}
	if p.calculateIntervalMinutes != nil {
		properties["CalculateIntervalMinutes"] = p.calculateIntervalMinutes
	}
	if p.entryPeriodEventId != nil {
		properties["EntryPeriodEventId"] = p.entryPeriodEventId
	}
	if p.accessPeriodEventId != nil {
		properties["AccessPeriodEventId"] = p.accessPeriodEventId
	}
	if p.generation != nil {
		properties["Generation"] = p.generation
	}
	return properties
}

func (p *CategoryModelMaster) Ref(
	namespaceName string,
) CategoryModelMasterRef {
	return CategoryModelMasterRef{
		namespaceName: namespaceName,
		categoryName:  p.name,
	}
}

func (p *CategoryModelMaster) GetAttrCategoryModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.CategoryModelId",
	)
}
