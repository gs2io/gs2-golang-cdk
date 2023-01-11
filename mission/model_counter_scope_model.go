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

type CounterScopeModelResetType string

const CounterScopeModelResetTypeNotReset = CounterScopeModelResetType("notReset")
const CounterScopeModelResetTypeDaily = CounterScopeModelResetType("daily")
const CounterScopeModelResetTypeWeekly = CounterScopeModelResetType("weekly")
const CounterScopeModelResetTypeMonthly = CounterScopeModelResetType("monthly")

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
	ResetType       CounterScopeModelResetType
	ResetDayOfMonth *int32
	ResetDayOfWeek  *CounterScopeModelResetDayOfWeek
	ResetHour       *int32
}

type CounterScopeModelOptions struct {
	ResetDayOfMonth *int32
	ResetDayOfWeek  *CounterScopeModelResetDayOfWeek
	ResetHour       *int32
}

func NewCounterScopeModel(
	resetType CounterScopeModelResetType,
	options CounterScopeModelOptions,
) CounterScopeModel {
	data := CounterScopeModel{
		ResetType:       resetType,
		ResetDayOfMonth: options.ResetDayOfMonth,
		ResetDayOfWeek:  options.ResetDayOfWeek,
		ResetHour:       options.ResetHour,
	}
	return data
}

type CounterScopeModelResetTypeIsNotResetOptions struct {
}

func NewCounterScopeModelResetTypeIsNotReset(
	options CounterScopeModelResetTypeIsNotResetOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		CounterScopeModelResetTypeNotReset,
		CounterScopeModelOptions{},
	)
}

type CounterScopeModelResetTypeIsDailyOptions struct {
}

func NewCounterScopeModelResetTypeIsDaily(
	resetHour int32,
	options CounterScopeModelResetTypeIsDailyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		CounterScopeModelResetTypeDaily,
		CounterScopeModelOptions{
			ResetHour: &resetHour,
		},
	)
}

type CounterScopeModelResetTypeIsWeeklyOptions struct {
}

func NewCounterScopeModelResetTypeIsWeekly(
	resetDayOfWeek CounterScopeModelResetDayOfWeek,
	resetHour int32,
	options CounterScopeModelResetTypeIsWeeklyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
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
	resetDayOfMonth int32,
	resetHour int32,
	options CounterScopeModelResetTypeIsMonthlyOptions,
) CounterScopeModel {
	return NewCounterScopeModel(
		CounterScopeModelResetTypeMonthly,
		CounterScopeModelOptions{
			ResetDayOfMonth: &resetDayOfMonth,
			ResetHour:       &resetHour,
		},
	)
}

func (p *CounterScopeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
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
	return properties
}
