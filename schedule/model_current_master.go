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
			"events":  events,
		},
	}
}
