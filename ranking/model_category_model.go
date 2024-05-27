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
	Name                       string
	Metadata                   *string
	MinimumValue               *int64
	MaximumValue               *int64
	Sum                        bool
	OrderDirection             CategoryModelOrderDirection
	Scope                      CategoryModelScope
	GlobalRankingSetting       *GlobalRankingSetting
	EntryPeriodEventId         *string
	AccessPeriodEventId        *string
	UniqueByUserId             *bool
	CalculateFixedTimingHour   *int32
	CalculateFixedTimingMinute *int32
	CalculateIntervalMinutes   *int32
	AdditionalScopes           []Scope
	IgnoreUserIds              []string
	Generation                 *string
}

type CategoryModelOptions struct {
	Metadata                   *string
	MinimumValue               *int64
	MaximumValue               *int64
	GlobalRankingSetting       *GlobalRankingSetting
	EntryPeriodEventId         *string
	AccessPeriodEventId        *string
	UniqueByUserId             *bool
	CalculateFixedTimingHour   *int32
	CalculateFixedTimingMinute *int32
	CalculateIntervalMinutes   *int32
	AdditionalScopes           []Scope
	IgnoreUserIds              []string
	Generation                 *string
}

func NewCategoryModel(
	name string,
	sum bool,
	orderDirection CategoryModelOrderDirection,
	scope CategoryModelScope,
	options CategoryModelOptions,
) CategoryModel {
	data := CategoryModel{
		Name:                       name,
		Sum:                        sum,
		OrderDirection:             orderDirection,
		Scope:                      scope,
		Metadata:                   options.Metadata,
		MinimumValue:               options.MinimumValue,
		MaximumValue:               options.MaximumValue,
		GlobalRankingSetting:       options.GlobalRankingSetting,
		EntryPeriodEventId:         options.EntryPeriodEventId,
		AccessPeriodEventId:        options.AccessPeriodEventId,
		UniqueByUserId:             options.UniqueByUserId,
		CalculateFixedTimingHour:   options.CalculateFixedTimingHour,
		CalculateFixedTimingMinute: options.CalculateFixedTimingMinute,
		CalculateIntervalMinutes:   options.CalculateIntervalMinutes,
		AdditionalScopes:           options.AdditionalScopes,
		IgnoreUserIds:              options.IgnoreUserIds,
		Generation:                 options.Generation,
	}
	return data
}

type CategoryModelScopeIsGlobalOptions struct {
	Metadata                   *string
	MinimumValue               *int64
	MaximumValue               *int64
	EntryPeriodEventId         *string
	AccessPeriodEventId        *string
	CalculateFixedTimingHour   *int32
	CalculateFixedTimingMinute *int32
	AdditionalScopes           []Scope
	IgnoreUserIds              []string
	Generation                 *string
}

func NewCategoryModelScopeIsGlobal(
	name string,
	sum bool,
	orderDirection CategoryModelOrderDirection,
	globalRankingSetting GlobalRankingSetting,
	uniqueByUserId bool,
	calculateIntervalMinutes int32,
	options CategoryModelScopeIsGlobalOptions,
) CategoryModel {
	return NewCategoryModel(
		name,
		sum,
		orderDirection,
		CategoryModelScopeGlobal,
		CategoryModelOptions{
			Metadata:                   options.Metadata,
			MinimumValue:               options.MinimumValue,
			MaximumValue:               options.MaximumValue,
			GlobalRankingSetting:       &globalRankingSetting,
			EntryPeriodEventId:         options.EntryPeriodEventId,
			AccessPeriodEventId:        options.AccessPeriodEventId,
			UniqueByUserId:             &uniqueByUserId,
			CalculateFixedTimingHour:   options.CalculateFixedTimingHour,
			CalculateFixedTimingMinute: options.CalculateFixedTimingMinute,
			CalculateIntervalMinutes:   &calculateIntervalMinutes,
			AdditionalScopes:           options.AdditionalScopes,
			IgnoreUserIds:              options.IgnoreUserIds,
			Generation:                 options.Generation,
		},
	)
}

type CategoryModelScopeIsScopedOptions struct {
	Metadata                   *string
	MinimumValue               *int64
	MaximumValue               *int64
	EntryPeriodEventId         *string
	AccessPeriodEventId        *string
	CalculateFixedTimingHour   *int32
	CalculateFixedTimingMinute *int32
	AdditionalScopes           []Scope
	IgnoreUserIds              []string
	Generation                 *string
}

func NewCategoryModelScopeIsScoped(
	name string,
	sum bool,
	orderDirection CategoryModelOrderDirection,
	options CategoryModelScopeIsScopedOptions,
) CategoryModel {
	return NewCategoryModel(
		name,
		sum,
		orderDirection,
		CategoryModelScopeScoped,
		CategoryModelOptions{
			Metadata:                   options.Metadata,
			MinimumValue:               options.MinimumValue,
			MaximumValue:               options.MaximumValue,
			EntryPeriodEventId:         options.EntryPeriodEventId,
			AccessPeriodEventId:        options.AccessPeriodEventId,
			CalculateFixedTimingHour:   options.CalculateFixedTimingHour,
			CalculateFixedTimingMinute: options.CalculateFixedTimingMinute,
			AdditionalScopes:           options.AdditionalScopes,
			IgnoreUserIds:              options.IgnoreUserIds,
			Generation:                 options.Generation,
		},
	)
}

func (p *CategoryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	if p.MinimumValue != nil {
		properties["MinimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		properties["MaximumValue"] = p.MaximumValue
	}
	properties["Sum"] = p.Sum
	properties["OrderDirection"] = p.OrderDirection
	properties["Scope"] = p.Scope
	if p.GlobalRankingSetting != nil {
		properties["GlobalRankingSetting"] = p.GlobalRankingSetting.Properties()
	}
	if p.EntryPeriodEventId != nil {
		properties["EntryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.AccessPeriodEventId != nil {
		properties["AccessPeriodEventId"] = p.AccessPeriodEventId
	}
	if p.UniqueByUserId != nil {
		properties["UniqueByUserId"] = p.UniqueByUserId
	}
	if p.CalculateFixedTimingHour != nil {
		properties["CalculateFixedTimingHour"] = p.CalculateFixedTimingHour
	}
	if p.CalculateFixedTimingMinute != nil {
		properties["CalculateFixedTimingMinute"] = p.CalculateFixedTimingMinute
	}
	if p.CalculateIntervalMinutes != nil {
		properties["CalculateIntervalMinutes"] = p.CalculateIntervalMinutes
	}
	{
		values := make([]map[string]interface{}, len(p.AdditionalScopes))
		for i, element := range p.AdditionalScopes {
			values[i] = element.Properties()
		}
		properties["AdditionalScopes"] = values
	}
	properties["IgnoreUserIds"] = p.IgnoreUserIds
	if p.Generation != nil {
		properties["Generation"] = p.Generation
	}
	return properties
}
