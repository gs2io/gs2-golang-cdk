package stateMachine

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

type EventEventType string

const EventEventTypeChangeState = EventEventType("change_state")
const EventEventTypeEmit = EventEventType("emit")

func (p EventEventType) Pointer() *EventEventType {
	return &p
}

type Event struct {
	EventType        EventEventType
	ChangeStateEvent *ChangeStateEvent
	EmitEvent        *EmitEvent
}

type EventOptions struct {
	ChangeStateEvent *ChangeStateEvent
	EmitEvent        *EmitEvent
}

func NewEvent(
	eventType EventEventType,
	options EventOptions,
) Event {
	data := Event{
		EventType:        eventType,
		ChangeStateEvent: options.ChangeStateEvent,
		EmitEvent:        options.EmitEvent,
	}
	return data
}

type EventEventTypeIsChangeStateOptions struct {
}

func NewEventEventTypeIsChangeState(
	changeStateEvent ChangeStateEvent,
	options EventEventTypeIsChangeStateOptions,
) Event {
	return NewEvent(
		EventEventTypeChangeState,
		EventOptions{
			ChangeStateEvent: &changeStateEvent,
		},
	)
}

type EventEventTypeIsEmitOptions struct {
}

func NewEventEventTypeIsEmit(
	emitEvent EmitEvent,
	options EventEventTypeIsEmitOptions,
) Event {
	return NewEvent(
		EventEventTypeEmit,
		EventOptions{
			EmitEvent: &emitEvent,
		},
	)
}

func (p *Event) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["EventType"] = p.EventType
	if p.ChangeStateEvent != nil {
		properties["ChangeStateEvent"] = p.ChangeStateEvent.Properties()
	}
	if p.EmitEvent != nil {
		properties["EmitEvent"] = p.EmitEvent.Properties()
	}
	return properties
}
