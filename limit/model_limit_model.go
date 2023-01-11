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
	Name            string
	Metadata        *string
	ResetType       LimitModelResetType
	ResetDayOfMonth *int32
	ResetDayOfWeek  *LimitModelResetDayOfWeek
	ResetHour       *int32
}

type LimitModelOptions struct {
	Metadata        *string
	ResetDayOfMonth *int32
	ResetDayOfWeek  *LimitModelResetDayOfWeek
	ResetHour       *int32
}

func NewLimitModel(
	name string,
	resetType LimitModelResetType,
	options LimitModelOptions,
) LimitModel {
	data := LimitModel{
		Name:            name,
		ResetType:       resetType,
		Metadata:        options.Metadata,
		ResetDayOfMonth: options.ResetDayOfMonth,
		ResetDayOfWeek:  options.ResetDayOfWeek,
		ResetHour:       options.ResetHour,
	}
	return data
}

type LimitModelResetTypeIsNotResetOptions struct {
	Metadata *string
}

func NewLimitModelResetTypeIsNotReset(
	name string,
	options LimitModelResetTypeIsNotResetOptions,
) LimitModel {
	return NewLimitModel(
		name,
		LimitModelResetTypeNotReset,
		LimitModelOptions{
			Metadata: options.Metadata,
		},
	)
}

type LimitModelResetTypeIsDailyOptions struct {
	Metadata *string
}

func NewLimitModelResetTypeIsDaily(
	name string,
	resetHour int32,
	options LimitModelResetTypeIsDailyOptions,
) LimitModel {
	return NewLimitModel(
		name,
		LimitModelResetTypeDaily,
		LimitModelOptions{
			Metadata:  options.Metadata,
			ResetHour: &resetHour,
		},
	)
}

type LimitModelResetTypeIsWeeklyOptions struct {
	Metadata *string
}

func NewLimitModelResetTypeIsWeekly(
	name string,
	resetDayOfWeek LimitModelResetDayOfWeek,
	resetHour int32,
	options LimitModelResetTypeIsWeeklyOptions,
) LimitModel {
	return NewLimitModel(
		name,
		LimitModelResetTypeWeekly,
		LimitModelOptions{
			Metadata:       options.Metadata,
			ResetDayOfWeek: &resetDayOfWeek,
			ResetHour:      &resetHour,
		},
	)
}

type LimitModelResetTypeIsMonthlyOptions struct {
	Metadata *string
}

func NewLimitModelResetTypeIsMonthly(
	name string,
	resetDayOfMonth int32,
	resetHour int32,
	options LimitModelResetTypeIsMonthlyOptions,
) LimitModel {
	return NewLimitModel(
		name,
		LimitModelResetTypeMonthly,
		LimitModelOptions{
			Metadata:        options.Metadata,
			ResetDayOfMonth: &resetDayOfMonth,
			ResetHour:       &resetHour,
		},
	)
}

func (p *LimitModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
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
