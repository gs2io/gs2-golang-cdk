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
	RepeatType            *EventRepeatType
	AbsoluteBegin         *int64
	AbsoluteEnd           *int64
	RepeatBeginDayOfMonth *int32
	RepeatEndDayOfMonth   *int32
	RepeatBeginDayOfWeek  *EventRepeatBeginDayOfWeek
	RepeatEndDayOfWeek    *EventRepeatEndDayOfWeek
	RepeatBeginHour       *int32
	RepeatEndHour         *int32
	RelativeTriggerName   *string
	RelativeDuration      *int32
}

type EventOptions struct {
	Metadata              *string
	RepeatType            *EventRepeatType
	AbsoluteBegin         *int64
	AbsoluteEnd           *int64
	RepeatBeginDayOfMonth *int32
	RepeatEndDayOfMonth   *int32
	RepeatBeginDayOfWeek  *EventRepeatBeginDayOfWeek
	RepeatEndDayOfWeek    *EventRepeatEndDayOfWeek
	RepeatBeginHour       *int32
	RepeatEndHour         *int32
	RelativeTriggerName   *string
	RelativeDuration      *int32
}

func NewEvent(
	name string,
	scheduleType EventScheduleType,
	options EventOptions,
) Event {
	data := Event{
		Name:                  name,
		ScheduleType:          scheduleType,
		Metadata:              options.Metadata,
		RepeatType:            options.RepeatType,
		AbsoluteBegin:         options.AbsoluteBegin,
		AbsoluteEnd:           options.AbsoluteEnd,
		RepeatBeginDayOfMonth: options.RepeatBeginDayOfMonth,
		RepeatEndDayOfMonth:   options.RepeatEndDayOfMonth,
		RepeatBeginDayOfWeek:  options.RepeatBeginDayOfWeek,
		RepeatEndDayOfWeek:    options.RepeatEndDayOfWeek,
		RepeatBeginHour:       options.RepeatBeginHour,
		RepeatEndHour:         options.RepeatEndHour,
		RelativeTriggerName:   options.RelativeTriggerName,
		RelativeDuration:      options.RelativeDuration,
	}
	return data
}

type EventScheduleTypeIsAbsoluteOptions struct {
	Metadata *string
}

func NewEventScheduleTypeIsAbsolute(
	name string,
	repeatType EventRepeatType,
	absoluteBegin int64,
	absoluteEnd int64,
	options EventScheduleTypeIsAbsoluteOptions,
) Event {
	return NewEvent(
		name,
		EventScheduleTypeAbsolute,
		EventOptions{
			Metadata:      options.Metadata,
			RepeatType:    &repeatType,
			AbsoluteBegin: &absoluteBegin,
			AbsoluteEnd:   &absoluteEnd,
		},
	)
}

type EventScheduleTypeIsRelativeOptions struct {
	Metadata *string
}

func NewEventScheduleTypeIsRelative(
	name string,
	relativeTriggerName string,
	relativeDuration int32,
	options EventScheduleTypeIsRelativeOptions,
) Event {
	return NewEvent(
		name,
		EventScheduleTypeRelative,
		EventOptions{
			Metadata:            options.Metadata,
			RelativeTriggerName: &relativeTriggerName,
			RelativeDuration:    &relativeDuration,
		},
	)
}

type EventRepeatTypeIsAlwaysOptions struct {
	Metadata *string
}

func NewEventRepeatTypeIsAlways(
	name string,
	scheduleType EventScheduleType,
	options EventRepeatTypeIsAlwaysOptions,
) Event {
	return NewEvent(
		name,
		scheduleType,
		EventOptions{
			Metadata: options.Metadata,
			RepeatType: func() *EventRepeatType {
				v := EventRepeatTypeAlways
				return &v
			}(),
		},
	)
}

type EventRepeatTypeIsDailyOptions struct {
	Metadata *string
}

func NewEventRepeatTypeIsDaily(
	name string,
	scheduleType EventScheduleType,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsDailyOptions,
) Event {
	return NewEvent(
		name,
		scheduleType,
		EventOptions{
			Metadata: options.Metadata,
			RepeatType: func() *EventRepeatType {
				v := EventRepeatTypeDaily
				return &v
			}(),
			RepeatBeginHour: &repeatBeginHour,
			RepeatEndHour:   &repeatEndHour,
		},
	)
}

type EventRepeatTypeIsWeeklyOptions struct {
	Metadata *string
}

func NewEventRepeatTypeIsWeekly(
	name string,
	scheduleType EventScheduleType,
	repeatBeginDayOfWeek EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek EventRepeatEndDayOfWeek,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsWeeklyOptions,
) Event {
	return NewEvent(
		name,
		scheduleType,
		EventOptions{
			Metadata: options.Metadata,
			RepeatType: func() *EventRepeatType {
				v := EventRepeatTypeWeekly
				return &v
			}(),
			RepeatBeginDayOfWeek: &repeatBeginDayOfWeek,
			RepeatEndDayOfWeek:   &repeatEndDayOfWeek,
			RepeatBeginHour:      &repeatBeginHour,
			RepeatEndHour:        &repeatEndHour,
		},
	)
}

type EventRepeatTypeIsMonthlyOptions struct {
	Metadata *string
}

func NewEventRepeatTypeIsMonthly(
	name string,
	scheduleType EventScheduleType,
	repeatBeginDayOfMonth int32,
	repeatEndDayOfMonth int32,
	repeatBeginHour int32,
	repeatEndHour int32,
	options EventRepeatTypeIsMonthlyOptions,
) Event {
	return NewEvent(
		name,
		scheduleType,
		EventOptions{
			Metadata: options.Metadata,
			RepeatType: func() *EventRepeatType {
				v := EventRepeatTypeMonthly
				return &v
			}(),
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
	if p.RepeatType != nil {
		properties["RepeatType"] = p.RepeatType
	}
	if p.AbsoluteBegin != nil {
		properties["AbsoluteBegin"] = p.AbsoluteBegin
	}
	if p.AbsoluteEnd != nil {
		properties["AbsoluteEnd"] = p.AbsoluteEnd
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
	if p.RelativeTriggerName != nil {
		properties["RelativeTriggerName"] = p.RelativeTriggerName
	}
	if p.RelativeDuration != nil {
		properties["RelativeDuration"] = p.RelativeDuration
	}
	return properties
}
