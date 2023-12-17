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

type EmitEvent struct {
	Event      string
	Parameters string
	Timestamp  int64
}

type EmitEventOptions struct {
}

func NewEmitEvent(
	event string,
	parameters string,
	timestamp int64,
	options EmitEventOptions,
) EmitEvent {
	data := EmitEvent{
		Event:      event,
		Parameters: parameters,
		Timestamp:  timestamp,
	}
	return data
}

func (p *EmitEvent) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Event"] = p.Event
	properties["Parameters"] = p.Parameters
	properties["Timestamp"] = p.Timestamp
	return properties
}
