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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type EventScheduleType string

const EventScheduleTypeAbsolute = EventScheduleType("absolute")
const EventScheduleTypeRelative = EventScheduleType("relative")

func (p EventScheduleType) Pointer() *EventScheduleType {
	return &p
}

type EventRepeatType string

const EventRepeatTypeAlways = EventRepeatType("always")
const EventRepeatTypeDaily = EventRepeatType("daily")
const EventRepeatTypeWeekly = EventRepeatType("weekly")
const EventRepeatTypeMonthly = EventRepeatType("monthly")

func (p EventRepeatType) Pointer() *EventRepeatType {
	return &p
}

type EventRepeatBeginDayOfWeek string

const EventRepeatBeginDayOfWeekSunday = EventRepeatBeginDayOfWeek("sunday")
const EventRepeatBeginDayOfWeekMonday = EventRepeatBeginDayOfWeek("monday")
const EventRepeatBeginDayOfWeekTuesday = EventRepeatBeginDayOfWeek("tuesday")
const EventRepeatBeginDayOfWeekWednesday = EventRepeatBeginDayOfWeek("wednesday")
const EventRepeatBeginDayOfWeekThursday = EventRepeatBeginDayOfWeek("thursday")
const EventRepeatBeginDayOfWeekFriday = EventRepeatBeginDayOfWeek("friday")
const EventRepeatBeginDayOfWeekSaturday = EventRepeatBeginDayOfWeek("saturday")

func (p EventRepeatBeginDayOfWeek) Pointer() *EventRepeatBeginDayOfWeek {
	return &p
}

type EventRepeatEndDayOfWeek string

const EventRepeatEndDayOfWeekSunday = EventRepeatEndDayOfWeek("sunday")
const EventRepeatEndDayOfWeekMonday = EventRepeatEndDayOfWeek("monday")
const EventRepeatEndDayOfWeekTuesday = EventRepeatEndDayOfWeek("tuesday")
const EventRepeatEndDayOfWeekWednesday = EventRepeatEndDayOfWeek("wednesday")
const EventRepeatEndDayOfWeekThursday = EventRepeatEndDayOfWeek("thursday")
const EventRepeatEndDayOfWeekFriday = EventRepeatEndDayOfWeek("friday")
const EventRepeatEndDayOfWeekSaturday = EventRepeatEndDayOfWeek("saturday")

func (p EventRepeatEndDayOfWeek) Pointer() *EventRepeatEndDayOfWeek {
	return &p
}

type Event struct {
	Name                  string
	Metadata              *string
	ScheduleType          EventScheduleType
	AbsoluteBegin         *int64
	AbsoluteEnd           *int64
	RelativeTriggerName   *string
	RepeatSetting         RepeatSetting
	RepeatType            *EventRepeatType
	RepeatBeginDayOfMonth *int32
	RepeatEndDayOfMonth   *int32
	RepeatBeginDayOfWeek  *EventRepeatBeginDayOfWeek
	RepeatEndDayOfWeek    *EventRepeatEndDayOfWeek
	RepeatBeginHour       *int32
	RepeatEndHour         *int32
}

type EventOptions struct {
	RepeatType            *EventRepeatType
	Metadata              *string
	AbsoluteBegin         *int64
	AbsoluteEnd           *int64
	RelativeTriggerName   *string
	RepeatBeginDayOfMonth *int32
	RepeatEndDayOfMonth   *int32
	RepeatBeginDayOfWeek  *EventRepeatBeginDayOfWeek
	RepeatEndDayOfWeek    *EventRepeatEndDayOfWeek
	RepeatBeginHour       *int32
	RepeatEndHour         *int32
}

func NewEvent(
	name string,
	scheduleType EventScheduleType,
	repeatSetting RepeatSetting,
	options EventOptions,
) Event {
	data := Event{
		Name:                  name,
		ScheduleType:          scheduleType,
		RepeatSetting:         repeatSetting,
		RepeatType:            options.RepeatType,
		Metadata:              options.Metadata,
		AbsoluteBegin:         options.AbsoluteBegin,
		AbsoluteEnd:           options.AbsoluteEnd,
		RelativeTriggerName:   options.RelativeTriggerName,
		RepeatBeginDayOfMonth: options.RepeatBeginDayOfMonth,
		RepeatEndDayOfMonth:   options.RepeatEndDayOfMonth,
		RepeatBeginDayOfWeek:  options.RepeatBeginDayOfWeek,
		RepeatEndDayOfWeek:    options.RepeatEndDayOfWeek,
		RepeatBeginHour:       options.RepeatBeginHour,
		RepeatEndHour:         options.RepeatEndHour,
	}
	return data
}

type EventScheduleTypeIsAbsoluteOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventScheduleTypeIsAbsolute(
	name string,
	repeatSetting RepeatSetting,
	options EventScheduleTypeIsAbsoluteOptions,
) Event {
	return NewEvent(
		name,
		EventScheduleTypeAbsolute,
		repeatSetting,
		EventOptions{
			Metadata:      options.Metadata,
			AbsoluteBegin: options.AbsoluteBegin,
			AbsoluteEnd:   options.AbsoluteEnd,
		},
	)
}

type EventScheduleTypeIsRelativeOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventScheduleTypeIsRelative(
	name string,
	repeatSetting RepeatSetting,
	relativeTriggerName string,
	options EventScheduleTypeIsRelativeOptions,
) Event {
	return NewEvent(
		name,
		EventScheduleTypeRelative,
		repeatSetting,
		EventOptions{
			Metadata:            options.Metadata,
			AbsoluteBegin:       options.AbsoluteBegin,
			AbsoluteEnd:         options.AbsoluteEnd,
			RelativeTriggerName: &relativeTriggerName,
		},
	)
}

type EventRepeatTypeIsAlwaysOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventRepeatTypeIsAlways(
	name string,
	scheduleType EventScheduleType,
	repeatSetting RepeatSetting,
	options EventRepeatTypeIsAlwaysOptions,
) Event {
	eventRepeatTypeAlways := EventRepeatTypeAlways
	return NewEvent(
		name,
		scheduleType,
		repeatSetting,
		EventOptions{
			RepeatType:    &eventRepeatTypeAlways,
			Metadata:      options.Metadata,
			AbsoluteBegin: options.AbsoluteBegin,
			AbsoluteEnd:   options.AbsoluteEnd,
		},
	)
}

type EventRepeatTypeIsDailyOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventRepeatTypeIsDaily(
	name string,
	scheduleType EventScheduleType,
	repeatSetting RepeatSetting,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsDailyOptions,
) Event {
	eventRepeatTypeDaily := EventRepeatTypeDaily
	return NewEvent(
		name,
		scheduleType,
		repeatSetting,
		EventOptions{
			RepeatType:      &eventRepeatTypeDaily,
			Metadata:        options.Metadata,
			AbsoluteBegin:   options.AbsoluteBegin,
			AbsoluteEnd:     options.AbsoluteEnd,
			RepeatBeginHour: &repeatBeginHour,
			RepeatEndHour:   &repeatEndHour,
		},
	)
}

type EventRepeatTypeIsWeeklyOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventRepeatTypeIsWeekly(
	name string,
	scheduleType EventScheduleType,
	repeatSetting RepeatSetting,
	repeatBeginDayOfWeek EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek EventRepeatEndDayOfWeek,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsWeeklyOptions,
) Event {
	eventRepeatTypeWeekly := EventRepeatTypeWeekly
	return NewEvent(
		name,
		scheduleType,
		repeatSetting,
		EventOptions{
			RepeatType:           &eventRepeatTypeWeekly,
			Metadata:             options.Metadata,
			AbsoluteBegin:        options.AbsoluteBegin,
			AbsoluteEnd:          options.AbsoluteEnd,
			RepeatBeginDayOfWeek: &repeatBeginDayOfWeek,
			RepeatEndDayOfWeek:   &repeatEndDayOfWeek,
			RepeatBeginHour:      &repeatBeginHour,
			RepeatEndHour:        &repeatEndHour,
		},
	)
}

type EventRepeatTypeIsMonthlyOptions struct {
	Metadata      *string
	AbsoluteBegin *int64
	AbsoluteEnd   *int64
}

func NewEventRepeatTypeIsMonthly(
	name string,
	scheduleType EventScheduleType,
	repeatSetting RepeatSetting,
	repeatBeginDayOfMonth int32,
	repeatEndDayOfMonth int32,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsMonthlyOptions,
) Event {
	eventRepeatTypeMonthly := EventRepeatTypeMonthly
	return NewEvent(
		name,
		scheduleType,
		repeatSetting,
		EventOptions{
			RepeatType:            &eventRepeatTypeMonthly,
			Metadata:              options.Metadata,
			AbsoluteBegin:         options.AbsoluteBegin,
			AbsoluteEnd:           options.AbsoluteEnd,
			RepeatBeginDayOfMonth: &repeatBeginDayOfMonth,
			RepeatEndDayOfMonth:   &repeatEndDayOfMonth,
			RepeatBeginHour:       &repeatBeginHour,
			RepeatEndHour:         &repeatEndHour,
		},
	)
}

func (p *Event) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["ScheduleType"] = p.ScheduleType
	if p.AbsoluteBegin != nil {
		properties["AbsoluteBegin"] = p.AbsoluteBegin
	}
	if p.AbsoluteEnd != nil {
		properties["AbsoluteEnd"] = p.AbsoluteEnd
	}
	if p.RelativeTriggerName != nil {
		properties["RelativeTriggerName"] = p.RelativeTriggerName
	}
	properties["RepeatSetting"] = p.RepeatSetting.Properties()
	if p.RepeatType != nil {
		properties["RepeatType"] = p.RepeatType
	}
	if p.RepeatBeginDayOfMonth != nil {
		properties["RepeatBeginDayOfMonth"] = p.RepeatBeginDayOfMonth
	}
	if p.RepeatEndDayOfMonth != nil {
		properties["RepeatEndDayOfMonth"] = p.RepeatEndDayOfMonth
	}
	if p.RepeatBeginDayOfWeek != nil {
		properties["RepeatBeginDayOfWeek"] = p.RepeatBeginDayOfWeek
	}
	if p.RepeatEndDayOfWeek != nil {
		properties["RepeatEndDayOfWeek"] = p.RepeatEndDayOfWeek
	}
	if p.RepeatBeginHour != nil {
		properties["RepeatBeginHour"] = p.RepeatBeginHour
	}
	if p.RepeatEndHour != nil {
		properties["RepeatEndHour"] = p.RepeatEndHour
	}
	return properties
}
