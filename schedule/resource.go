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
	name                  string
	metadata              *string
	scheduleType          EventScheduleType
	repeatType            *EventRepeatType
	absoluteBegin         *int64
	absoluteEnd           *int64
	repeatBeginDayOfMonth *int32
	repeatEndDayOfMonth   *int32
	repeatBeginDayOfWeek  *EventRepeatBeginDayOfWeek
	repeatEndDayOfWeek    *EventRepeatEndDayOfWeek
	repeatBeginHour       *int32
	repeatEndHour         *int32
	relativeTriggerName   *string
	relativeDuration      *int32
}

func NewEvent(
	name string,
	metadata *string,
	scheduleType EventScheduleType,
	repeatType *EventRepeatType,
	absoluteBegin *int64,
	absoluteEnd *int64,
	repeatBeginDayOfMonth *int32,
	repeatEndDayOfMonth *int32,
	repeatBeginDayOfWeek *EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek *EventRepeatEndDayOfWeek,
	repeatBeginHour *int32,
	repeatEndHour *int32,
	relativeTriggerName *string,
	relativeDuration *int32,
) *Event {
	return &Event{
		name:                  name,
		metadata:              metadata,
		scheduleType:          scheduleType,
		repeatType:            repeatType,
		absoluteBegin:         absoluteBegin,
		absoluteEnd:           absoluteEnd,
		repeatBeginDayOfMonth: repeatBeginDayOfMonth,
		repeatEndDayOfMonth:   repeatEndDayOfMonth,
		repeatBeginDayOfWeek:  repeatBeginDayOfWeek,
		repeatEndDayOfWeek:    repeatEndDayOfWeek,
		repeatBeginHour:       repeatBeginHour,
		repeatEndHour:         repeatEndHour,
		relativeTriggerName:   relativeTriggerName,
		relativeDuration:      relativeDuration,
	}
}

func (p *Event) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ScheduleType"] = p.scheduleType
	if p.repeatType != nil {
		properties["RepeatType"] = p.repeatType
	}
	if p.absoluteBegin != nil {
		properties["AbsoluteBegin"] = p.absoluteBegin
	}
	if p.absoluteEnd != nil {
		properties["AbsoluteEnd"] = p.absoluteEnd
	}
	if p.repeatBeginDayOfMonth != nil {
		properties["RepeatBeginDayOfMonth"] = p.repeatBeginDayOfMonth
	}
	if p.repeatEndDayOfMonth != nil {
		properties["RepeatEndDayOfMonth"] = p.repeatEndDayOfMonth
	}
	if p.repeatBeginDayOfWeek != nil {
		properties["RepeatBeginDayOfWeek"] = p.repeatBeginDayOfWeek
	}
	if p.repeatEndDayOfWeek != nil {
		properties["RepeatEndDayOfWeek"] = p.repeatEndDayOfWeek
	}
	if p.repeatBeginHour != nil {
		properties["RepeatBeginHour"] = p.repeatBeginHour
	}
	if p.repeatEndHour != nil {
		properties["RepeatEndHour"] = p.repeatEndHour
	}
	if p.relativeTriggerName != nil {
		properties["RelativeTriggerName"] = p.relativeTriggerName
	}
	if p.relativeDuration != nil {
		properties["RelativeDuration"] = p.relativeDuration
	}
	return properties
}

func (p *Event) Ref(
	namespaceName string,
) EventRef {
	return EventRef{
		namespaceName: namespaceName,
		eventName:     p.name,
	}
}

func NewAbsoluteEvent(
	name string,
	repeatType EventRepeatType,
	absoluteBegin int64,
	absoluteEnd int64,
	metadata *string,
	repeatBeginDayOfMonth *int32,
	repeatEndDayOfMonth *int32,
	repeatBeginDayOfWeek *EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek *EventRepeatEndDayOfWeek,
	repeatBeginHour *int32,
	repeatEndHour *int32,
) Event {
	return Event{
		scheduleType:          EventScheduleTypeAbsolute,
		name:                  name,
		metadata:              metadata,
		repeatType:            &repeatType,
		absoluteBegin:         &absoluteBegin,
		absoluteEnd:           &absoluteEnd,
		repeatBeginDayOfMonth: repeatBeginDayOfMonth,
		repeatEndDayOfMonth:   repeatEndDayOfMonth,
		repeatBeginDayOfWeek:  repeatBeginDayOfWeek,
		repeatEndDayOfWeek:    repeatEndDayOfWeek,
		repeatBeginHour:       repeatBeginHour,
		repeatEndHour:         repeatEndHour,
	}
}

func NewRelativeEvent(
	name string,
	relativeTriggerName string,
	relativeDuration int32,
	metadata *string,
	repeatBeginDayOfMonth *int32,
	repeatEndDayOfMonth *int32,
	repeatBeginDayOfWeek *EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek *EventRepeatEndDayOfWeek,
	repeatBeginHour *int32,
	repeatEndHour *int32,
) Event {
	return Event{
		scheduleType:          EventScheduleTypeRelative,
		name:                  name,
		metadata:              metadata,
		repeatBeginDayOfMonth: repeatBeginDayOfMonth,
		repeatEndDayOfMonth:   repeatEndDayOfMonth,
		repeatBeginDayOfWeek:  repeatBeginDayOfWeek,
		repeatEndDayOfWeek:    repeatEndDayOfWeek,
		repeatBeginHour:       repeatBeginHour,
		repeatEndHour:         repeatEndHour,
		relativeTriggerName:   &relativeTriggerName,
		relativeDuration:      &relativeDuration,
	}
}

func NewAlwaysEvent(
	name string,
	scheduleType EventScheduleType,
	metadata *string,
	absoluteBegin *int64,
	absoluteEnd *int64,
	relativeTriggerName *string,
	relativeDuration *int32,
) Event {
	return Event{
		repeatType:          EventRepeatTypeAlways.Pointer(),
		name:                name,
		metadata:            metadata,
		scheduleType:        scheduleType,
		absoluteBegin:       absoluteBegin,
		absoluteEnd:         absoluteEnd,
		relativeTriggerName: relativeTriggerName,
		relativeDuration:    relativeDuration,
	}
}

func NewDailyEvent(
	name string,
	scheduleType EventScheduleType,
	repeatBeginHour int32,
	repeatEndHour int32,
	metadata *string,
	absoluteBegin *int64,
	absoluteEnd *int64,
	relativeTriggerName *string,
	relativeDuration *int32,
) Event {
	return Event{
		repeatType:          EventRepeatTypeDaily.Pointer(),
		name:                name,
		metadata:            metadata,
		scheduleType:        scheduleType,
		absoluteBegin:       absoluteBegin,
		absoluteEnd:         absoluteEnd,
		repeatBeginHour:     &repeatBeginHour,
		repeatEndHour:       &repeatEndHour,
		relativeTriggerName: relativeTriggerName,
		relativeDuration:    relativeDuration,
	}
}

func NewWeeklyEvent(
	name string,
	scheduleType EventScheduleType,
	repeatBeginDayOfWeek EventRepeatBeginDayOfWeek,
	repeatEndDayOfWeek EventRepeatEndDayOfWeek,
	repeatBeginHour int32,
	repeatEndHour int32,
	metadata *string,
	absoluteBegin *int64,
	absoluteEnd *int64,
	relativeTriggerName *string,
	relativeDuration *int32,
) Event {
	return Event{
		repeatType:           EventRepeatTypeWeekly.Pointer(),
		name:                 name,
		metadata:             metadata,
		scheduleType:         scheduleType,
		absoluteBegin:        absoluteBegin,
		absoluteEnd:          absoluteEnd,
		repeatBeginDayOfWeek: &repeatBeginDayOfWeek,
		repeatEndDayOfWeek:   &repeatEndDayOfWeek,
		repeatBeginHour:      &repeatBeginHour,
		repeatEndHour:        &repeatEndHour,
		relativeTriggerName:  relativeTriggerName,
		relativeDuration:     relativeDuration,
	}
}

func NewMonthlyEvent(
	name string,
	scheduleType EventScheduleType,
	repeatBeginDayOfMonth int32,
	repeatEndDayOfMonth int32,
	repeatBeginHour int32,
	repeatEndHour int32,
	metadata *string,
	absoluteBegin *int64,
	absoluteEnd *int64,
	relativeTriggerName *string,
	relativeDuration *int32,
) Event {
	return Event{
		repeatType:            EventRepeatTypeMonthly.Pointer(),
		name:                  name,
		metadata:              metadata,
		scheduleType:          scheduleType,
		absoluteBegin:         absoluteBegin,
		absoluteEnd:           absoluteEnd,
		repeatBeginDayOfMonth: &repeatBeginDayOfMonth,
		repeatEndDayOfMonth:   &repeatEndDayOfMonth,
		repeatBeginHour:       &repeatBeginHour,
		repeatEndHour:         &repeatEndHour,
		relativeTriggerName:   relativeTriggerName,
		relativeDuration:      relativeDuration,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	events        []Event
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	events []Event,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-03-31",
		namespaceName: namespaceName,
		events:        events,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Schedule_CurrentEventMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Schedule::CurrentEventMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	events := make([]map[string]interface{}, len(p.events))
	for i, item := range p.events {
		events[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version": p.version,
			"Events":  events,
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
	return "Schedule_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Schedule::Namespace"
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
	events []Event,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		events,
	).AddDependsOn(
		p,
	)
	return p
}

type EventMaster struct {
	CdkResource
	stack                 *Stack
	namespaceName         string
	name                  string
	description           *string
	metadata              *string
	scheduleType          string
	absoluteBegin         *int64
	absoluteEnd           *int64
	repeatType            *string
	repeatBeginDayOfMonth *int32
	repeatEndDayOfMonth   *int32
	repeatBeginDayOfWeek  *string
	repeatEndDayOfWeek    *string
	repeatBeginHour       *int32
	repeatEndHour         *int32
	relativeTriggerName   *string
	relativeDuration      *int32
}

func NewEventMaster(
	stack *Stack,
	namespaceName string,
	name string,
	scheduleType string,
	description *string,
	metadata *string,
	absoluteBegin *int64,
	absoluteEnd *int64,
	repeatType *string,
	repeatBeginDayOfMonth *int32,
	repeatEndDayOfMonth *int32,
	repeatBeginDayOfWeek *string,
	repeatEndDayOfWeek *string,
	repeatBeginHour *int32,
	repeatEndHour *int32,
	relativeTriggerName *string,
	relativeDuration *int32,
) *EventMaster {

	self := EventMaster{
		stack:                 stack,
		namespaceName:         namespaceName,
		name:                  name,
		description:           description,
		metadata:              metadata,
		scheduleType:          scheduleType,
		absoluteBegin:         absoluteBegin,
		absoluteEnd:           absoluteEnd,
		repeatType:            repeatType,
		repeatBeginDayOfMonth: repeatBeginDayOfMonth,
		repeatEndDayOfMonth:   repeatEndDayOfMonth,
		repeatBeginDayOfWeek:  repeatBeginDayOfWeek,
		repeatEndDayOfWeek:    repeatEndDayOfWeek,
		repeatBeginHour:       repeatBeginHour,
		repeatEndHour:         repeatEndHour,
		relativeTriggerName:   relativeTriggerName,
		relativeDuration:      relativeDuration,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *EventMaster) ResourceName() string {
	return "Schedule_EventMaster_" + p.name
}

func (p *EventMaster) ResourceType() string {
	return "GS2::Schedule::EventMaster"
}

func (p *EventMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["ScheduleType"] = p.scheduleType
	if p.absoluteBegin != nil {
		properties["AbsoluteBegin"] = p.absoluteBegin
	}
	if p.absoluteEnd != nil {
		properties["AbsoluteEnd"] = p.absoluteEnd
	}
	if p.repeatType != nil {
		properties["RepeatType"] = p.repeatType
	}
	if p.repeatBeginDayOfMonth != nil {
		properties["RepeatBeginDayOfMonth"] = p.repeatBeginDayOfMonth
	}
	if p.repeatEndDayOfMonth != nil {
		properties["RepeatEndDayOfMonth"] = p.repeatEndDayOfMonth
	}
	if p.repeatBeginDayOfWeek != nil {
		properties["RepeatBeginDayOfWeek"] = p.repeatBeginDayOfWeek
	}
	if p.repeatEndDayOfWeek != nil {
		properties["RepeatEndDayOfWeek"] = p.repeatEndDayOfWeek
	}
	if p.repeatBeginHour != nil {
		properties["RepeatBeginHour"] = p.repeatBeginHour
	}
	if p.repeatEndHour != nil {
		properties["RepeatEndHour"] = p.repeatEndHour
	}
	if p.relativeTriggerName != nil {
		properties["RelativeTriggerName"] = p.relativeTriggerName
	}
	if p.relativeDuration != nil {
		properties["RelativeDuration"] = p.relativeDuration
	}
	return properties
}

func (p *EventMaster) Ref(
	namespaceName string,
) EventMasterRef {
	return EventMasterRef{
		namespaceName: namespaceName,
		eventName:     p.name,
	}
}

func (p *EventMaster) GetAttrEventId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.EventId",
	)
}
