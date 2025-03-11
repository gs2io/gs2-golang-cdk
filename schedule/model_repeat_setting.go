package schedule

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

type RepeatSettingRepeatType string

const RepeatSettingRepeatTypeAlways = RepeatSettingRepeatType("always")
const RepeatSettingRepeatTypeDaily = RepeatSettingRepeatType("daily")
const RepeatSettingRepeatTypeWeekly = RepeatSettingRepeatType("weekly")
const RepeatSettingRepeatTypeMonthly = RepeatSettingRepeatType("monthly")
const RepeatSettingRepeatTypeCustom = RepeatSettingRepeatType("custom")

func (p RepeatSettingRepeatType) Pointer() *RepeatSettingRepeatType {
	return &p
}

type RepeatSettingBeginDayOfWeek string

const RepeatSettingBeginDayOfWeekSunday = RepeatSettingBeginDayOfWeek("sunday")
const RepeatSettingBeginDayOfWeekMonday = RepeatSettingBeginDayOfWeek("monday")
const RepeatSettingBeginDayOfWeekTuesday = RepeatSettingBeginDayOfWeek("tuesday")
const RepeatSettingBeginDayOfWeekWednesday = RepeatSettingBeginDayOfWeek("wednesday")
const RepeatSettingBeginDayOfWeekThursday = RepeatSettingBeginDayOfWeek("thursday")
const RepeatSettingBeginDayOfWeekFriday = RepeatSettingBeginDayOfWeek("friday")
const RepeatSettingBeginDayOfWeekSaturday = RepeatSettingBeginDayOfWeek("saturday")

func (p RepeatSettingBeginDayOfWeek) Pointer() *RepeatSettingBeginDayOfWeek {
	return &p
}

type RepeatSettingEndDayOfWeek string

const RepeatSettingEndDayOfWeekSunday = RepeatSettingEndDayOfWeek("sunday")
const RepeatSettingEndDayOfWeekMonday = RepeatSettingEndDayOfWeek("monday")
const RepeatSettingEndDayOfWeekTuesday = RepeatSettingEndDayOfWeek("tuesday")
const RepeatSettingEndDayOfWeekWednesday = RepeatSettingEndDayOfWeek("wednesday")
const RepeatSettingEndDayOfWeekThursday = RepeatSettingEndDayOfWeek("thursday")
const RepeatSettingEndDayOfWeekFriday = RepeatSettingEndDayOfWeek("friday")
const RepeatSettingEndDayOfWeekSaturday = RepeatSettingEndDayOfWeek("saturday")

func (p RepeatSettingEndDayOfWeek) Pointer() *RepeatSettingEndDayOfWeek {
	return &p
}

type RepeatSetting struct {
	RepeatType      RepeatSettingRepeatType
	BeginDayOfMonth *int32
	EndDayOfMonth   *int32
	BeginDayOfWeek  *RepeatSettingBeginDayOfWeek
	EndDayOfWeek    *RepeatSettingEndDayOfWeek
	BeginHour       *int32
	EndHour         *int32
	AnchorTimestamp *int64
	ActiveDays      *int32
	InactiveDays    *int32
}

type RepeatSettingOptions struct {
	BeginDayOfMonth *int32
	EndDayOfMonth   *int32
	BeginDayOfWeek  *RepeatSettingBeginDayOfWeek
	EndDayOfWeek    *RepeatSettingEndDayOfWeek
	BeginHour       *int32
	EndHour         *int32
	AnchorTimestamp *int64
	ActiveDays      *int32
	InactiveDays    *int32
}

func NewRepeatSetting(
	repeatType RepeatSettingRepeatType,
	options RepeatSettingOptions,
) RepeatSetting {
	_data := RepeatSetting{
		RepeatType:      repeatType,
		BeginDayOfMonth: options.BeginDayOfMonth,
		EndDayOfMonth:   options.EndDayOfMonth,
		BeginDayOfWeek:  options.BeginDayOfWeek,
		EndDayOfWeek:    options.EndDayOfWeek,
		BeginHour:       options.BeginHour,
		EndHour:         options.EndHour,
		AnchorTimestamp: options.AnchorTimestamp,
		ActiveDays:      options.ActiveDays,
		InactiveDays:    options.InactiveDays,
	}
	return _data
}

type RepeatSettingRepeatTypeIsAlwaysOptions struct {
}

func NewRepeatSettingRepeatTypeIsAlways(
	options RepeatSettingRepeatTypeIsAlwaysOptions,
) RepeatSetting {
	return NewRepeatSetting(
		RepeatSettingRepeatTypeAlways,
		RepeatSettingOptions{},
	)
}

type RepeatSettingRepeatTypeIsDailyOptions struct {
}

func NewRepeatSettingRepeatTypeIsDaily(
	beginHour int32,
	endHour int32,
	options RepeatSettingRepeatTypeIsDailyOptions,
) RepeatSetting {
	return NewRepeatSetting(
		RepeatSettingRepeatTypeDaily,
		RepeatSettingOptions{
			BeginHour: &beginHour,
			EndHour:   &endHour,
		},
	)
}

type RepeatSettingRepeatTypeIsWeeklyOptions struct {
}

func NewRepeatSettingRepeatTypeIsWeekly(
	beginDayOfWeek RepeatSettingBeginDayOfWeek,
	endDayOfWeek RepeatSettingEndDayOfWeek,
	beginHour int32,
	endHour int32,
	options RepeatSettingRepeatTypeIsWeeklyOptions,
) RepeatSetting {
	return NewRepeatSetting(
		RepeatSettingRepeatTypeWeekly,
		RepeatSettingOptions{
			BeginDayOfWeek: &beginDayOfWeek,
			EndDayOfWeek:   &endDayOfWeek,
			BeginHour:      &beginHour,
			EndHour:        &endHour,
		},
	)
}

type RepeatSettingRepeatTypeIsMonthlyOptions struct {
}

func NewRepeatSettingRepeatTypeIsMonthly(
	beginDayOfMonth int32,
	endDayOfMonth int32,
	beginHour int32,
	endHour int32,
	options RepeatSettingRepeatTypeIsMonthlyOptions,
) RepeatSetting {
	return NewRepeatSetting(
		RepeatSettingRepeatTypeMonthly,
		RepeatSettingOptions{
			BeginDayOfMonth: &beginDayOfMonth,
			EndDayOfMonth:   &endDayOfMonth,
			BeginHour:       &beginHour,
			EndHour:         &endHour,
		},
	)
}

type RepeatSettingRepeatTypeIsCustomOptions struct {
}

func NewRepeatSettingRepeatTypeIsCustom(
	anchorTimestamp int64,
	activeDays int32,
	inactiveDays int32,
	options RepeatSettingRepeatTypeIsCustomOptions,
) RepeatSetting {
	return NewRepeatSetting(
		RepeatSettingRepeatTypeCustom,
		RepeatSettingOptions{
			AnchorTimestamp: &anchorTimestamp,
			ActiveDays:      &activeDays,
			InactiveDays:    &inactiveDays,
		},
	)
}

func (p *RepeatSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["RepeatType"] = p.RepeatType
	if p.BeginDayOfMonth != nil {
		properties["BeginDayOfMonth"] = p.BeginDayOfMonth
	}
	if p.EndDayOfMonth != nil {
		properties["EndDayOfMonth"] = p.EndDayOfMonth
	}
	if p.BeginDayOfWeek != nil {
		properties["BeginDayOfWeek"] = p.BeginDayOfWeek
	}
	if p.EndDayOfWeek != nil {
		properties["EndDayOfWeek"] = p.EndDayOfWeek
	}
	if p.BeginHour != nil {
		properties["BeginHour"] = p.BeginHour
	}
	if p.EndHour != nil {
		properties["EndHour"] = p.EndHour
	}
	if p.AnchorTimestamp != nil {
		properties["AnchorTimestamp"] = p.AnchorTimestamp
	}
	if p.ActiveDays != nil {
		properties["ActiveDays"] = p.ActiveDays
	}
	if p.InactiveDays != nil {
		properties["InactiveDays"] = p.InactiveDays
	}
	return properties
}
