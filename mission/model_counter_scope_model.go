package mission

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

type CounterScopeModelScopeType string

const CounterScopeModelScopeTypeResetTiming = CounterScopeModelScopeType("resetTiming")
const CounterScopeModelScopeTypeVerifyAction = CounterScopeModelScopeType("verifyAction")

func (p CounterScopeModelScopeType) Pointer() *CounterScopeModelScopeType {
	return &p
}

type CounterScopeModelResetType string

const CounterScopeModelResetTypeNotReset = CounterScopeModelResetType("notReset")
const CounterScopeModelResetTypeDaily = CounterScopeModelResetType("daily")
const CounterScopeModelResetTypeWeekly = CounterScopeModelResetType("weekly")
const CounterScopeModelResetTypeMonthly = CounterScopeModelResetType("monthly")
const CounterScopeModelResetTypeDays = CounterScopeModelResetType("days")

func (p CounterScopeModelResetType) Pointer() *CounterScopeModelResetType {
	return &p
}

type CounterScopeModelResetDayOfWeek string

const CounterScopeModelResetDayOfWeekSunday = CounterScopeModelResetDayOfWeek("sunday")
const CounterScopeModelResetDayOfWeekMonday = CounterScopeModelResetDayOfWeek("monday")
const CounterScopeModelResetDayOfWeekTuesday = CounterScopeModelResetDayOfWeek("tuesday")
const CounterScopeModelResetDayOfWeekWednesday = CounterScopeModelResetDayOfWeek("wednesday")
const CounterScopeModelResetDayOfWeekThursday = CounterScopeModelResetDayOfWeek("thursday")
const CounterScopeModelResetDayOfWeekFriday = CounterScopeModelResetDayOfWeek("friday")
const CounterScopeModelResetDayOfWeekSaturday = CounterScopeModelResetDayOfWeek("saturday")

func (p CounterScopeModelResetDayOfWeek) Pointer() *CounterScopeModelResetDayOfWeek {
	return &p
}

type CounterScopeModel struct {
	ScopeType       CounterScopeModelScopeType
	ResetType       CounterScopeModelResetType
	ResetDayOfMonth *int32
	ResetDayOfWeek  *CounterScopeModelResetDayOfWeek
	ResetHour       *int32
	ConditionName   *string
	Condition       *VerifyAction
	AnchorTimestamp *int64
	Days            *int32
}

type CounterScopeModelOptions struct {
	ResetDayOfMonth *int32
	ResetDayOfWeek  *CounterScopeModelResetDayOfWeek
	ResetHour       *int32
	ConditionName   *string
	Condition       *VerifyAction
	AnchorTimestamp *int64
	Days            *int32
}

func NewCounterScopeModel(
	scopeType CounterScopeModelScopeType,
	resetType CounterScopeModelResetType,
	options CounterScopeModelOptions,
) CounterScopeModel {
	_data := CounterScopeModel{
		ScopeType:       scopeType,
		ResetType:       resetType,
		ResetDayOfMonth: options.ResetDayOfMonth,
		ResetDayOfWeek:  options.ResetDayOfWeek,
		ResetHour:       options.ResetHour,
		ConditionName:   options.ConditionName,
		Condition:       options.Condition,
		AnchorTimestamp: options.AnchorTimestamp,
		Days:            options.Days,
	}
	return _data
}

type CounterScopeModelResetTypeIsNotResetOptions struct {
}

func NewCounterScopeModelResetTypeIsNotReset(
	scopeType CounterScopeModelScopeType,
	options CounterScopeModelResetTypeIsNotResetOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		scopeType,
		CounterScopeModelResetTypeNotReset,
		CounterScopeModelOptions{},
	)
}

type CounterScopeModelResetTypeIsDailyOptions struct {
}

func NewCounterScopeModelResetTypeIsDaily(
	scopeType CounterScopeModelScopeType,
	resetHour int32,
	options CounterScopeModelResetTypeIsDailyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		scopeType,
		CounterScopeModelResetTypeDaily,
		CounterScopeModelOptions{
			ResetHour: &resetHour,
		},
	)
}

type CounterScopeModelResetTypeIsWeeklyOptions struct {
}

func NewCounterScopeModelResetTypeIsWeekly(
	scopeType CounterScopeModelScopeType,
	resetDayOfWeek CounterScopeModelResetDayOfWeek,
	resetHour int32,
	options CounterScopeModelResetTypeIsWeeklyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		scopeType,
		CounterScopeModelResetTypeWeekly,
		CounterScopeModelOptions{
			ResetDayOfWeek: &resetDayOfWeek,
			ResetHour:      &resetHour,
		},
	)
}

type CounterScopeModelResetTypeIsMonthlyOptions struct {
}

func NewCounterScopeModelResetTypeIsMonthly(
	scopeType CounterScopeModelScopeType,
	resetDayOfMonth int32,
	resetHour int32,
	options CounterScopeModelResetTypeIsMonthlyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		scopeType,
		CounterScopeModelResetTypeMonthly,
		CounterScopeModelOptions{
			ResetDayOfMonth: &resetDayOfMonth,
			ResetHour:       &resetHour,
		},
	)
}

type CounterScopeModelResetTypeIsDaysOptions struct {
}

func NewCounterScopeModelResetTypeIsDays(
	scopeType CounterScopeModelScopeType,
	anchorTimestamp int64,
	days int32,
	options CounterScopeModelResetTypeIsDaysOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		scopeType,
		CounterScopeModelResetTypeDays,
		CounterScopeModelOptions{
			AnchorTimestamp: &anchorTimestamp,
			Days:            &days,
		},
	)
}

type CounterScopeModelScopeTypeIsResetTimingOptions struct {
}

func NewCounterScopeModelScopeTypeIsResetTiming(
	resetType CounterScopeModelResetType,
	options CounterScopeModelScopeTypeIsResetTimingOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		CounterScopeModelScopeTypeResetTiming,
		resetType,
		CounterScopeModelOptions{},
	)
}

type CounterScopeModelScopeTypeIsVerifyActionOptions struct {
}

func NewCounterScopeModelScopeTypeIsVerifyAction(
	resetType CounterScopeModelResetType,
	conditionName string,
	condition VerifyAction,
	options CounterScopeModelScopeTypeIsVerifyActionOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		CounterScopeModelScopeTypeVerifyAction,
		resetType,
		CounterScopeModelOptions{
			ConditionName: &conditionName,
			Condition:     &condition,
		},
	)
}

func (p *CounterScopeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ScopeType"] = p.ScopeType
	properties["ResetType"] = p.ResetType
	if p.ResetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.ResetDayOfMonth
	}
	if p.ResetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.ResetDayOfWeek
	}
	if p.ResetHour != nil {
		properties["ResetHour"] = p.ResetHour
	}
	if p.ConditionName != nil {
		properties["ConditionName"] = p.ConditionName
	}
	if p.Condition != nil {
		properties["Condition"] = p.Condition.Properties()
	}
	if p.AnchorTimestamp != nil {
		properties["AnchorTimestamp"] = p.AnchorTimestamp
	}
	if p.Days != nil {
		properties["Days"] = p.Days
	}
	return properties
}
