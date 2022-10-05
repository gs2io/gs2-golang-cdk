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
	resetType       CounterScopeModelResetType
	resetDayOfMonth *int32
	resetDayOfWeek  *CounterScopeModelResetDayOfWeek
	resetHour       *int32
}

func NewCounterScopeModel(
	resetType CounterScopeModelResetType,
	resetDayOfMonth *int32,
	resetDayOfWeek *CounterScopeModelResetDayOfWeek,
	resetHour *int32,
) CounterScopeModel {
	data := CounterScopeModel{
		resetType:       resetType,
		resetDayOfMonth: resetDayOfMonth,
		resetDayOfWeek:  resetDayOfWeek,
		resetHour:       resetHour,
	}
	return data
}

func NewNotResetCounterScopeModel() CounterScopeModel {
	return CounterScopeModel{
		resetType: CounterScopeModelResetTypeNotReset,
	}
}

func NewDailyCounterScopeModel(
	resetHour int32,
) CounterScopeModel {
	return CounterScopeModel{
		resetType: CounterScopeModelResetTypeDaily,
		resetHour: &resetHour,
	}
}

func NewWeeklyCounterScopeModel(
	resetDayOfWeek CounterScopeModelResetDayOfWeek,
	resetHour int32,
) CounterScopeModel {
	return CounterScopeModel{
		resetType:      CounterScopeModelResetTypeWeekly,
		resetDayOfWeek: &resetDayOfWeek,
		resetHour:      &resetHour,
	}
}

func NewMonthlyCounterScopeModel(
	resetDayOfMonth int32,
	resetHour int32,
) CounterScopeModel {
	return CounterScopeModel{
		resetType:       CounterScopeModelResetTypeMonthly,
		resetDayOfMonth: &resetDayOfMonth,
		resetHour:       &resetHour,
	}
}

func (p *CounterScopeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
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

type ScopedValueResetType string

const ScopedValueResetTypeNotReset = ScopedValueResetType("notReset")
const ScopedValueResetTypeDaily = ScopedValueResetType("daily")
const ScopedValueResetTypeWeekly = ScopedValueResetType("weekly")
const ScopedValueResetTypeMonthly = ScopedValueResetType("monthly")

func (p ScopedValueResetType) Pointer() *ScopedValueResetType {
	return &p
}

type ScopedValue struct {
	resetType   ScopedValueResetType
	value       int64
	nextResetAt *int64
	updatedAt   int64
}

func NewScopedValue(
	resetType ScopedValueResetType,
	value int64,
	nextResetAt *int64,
	updatedAt int64,
) ScopedValue {
	data := ScopedValue{
		resetType:   resetType,
		value:       value,
		nextResetAt: nextResetAt,
		updatedAt:   updatedAt,
	}
	return data
}

func (p *ScopedValue) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ResetType"] = p.resetType
	properties["Value"] = p.value
	if p.nextResetAt != nil {
		properties["NextResetAt"] = p.nextResetAt
	}
	properties["UpdatedAt"] = p.updatedAt
	return properties
}
