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

type CurrentMasterData struct {
	CdkResource
	stack              *Stack
	version            string
	namespaceName      string
	missionGroupModels []MissionGroupModel
	counterModels      []CounterModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	missionGroupModels []MissionGroupModel,
	counterModels []CounterModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:              stack,
		version:            "2019-05-28",
		namespaceName:      namespaceName,
		missionGroupModels: missionGroupModels,
		counterModels:      counterModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Mission_CurrentMissionMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Mission::CurrentMissionMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	missionGroupModels := make([]map[string]interface{}, len(p.missionGroupModels))
	for i, item := range p.missionGroupModels {
		missionGroupModels[i] = item.Properties()
	}
	counterModels := make([]map[string]interface{}, len(p.counterModels))
	for i, item := range p.counterModels {
		counterModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":            p.version,
			"missionGroupModels": missionGroupModels,
			"counterModels":      counterModels,
		},
	}
}
