package buff

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
	stack           *Stack
	version         string
	namespaceName   string
	buffEntryModels []BuffEntryModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	buffEntryModels []BuffEntryModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:           stack,
		version:         "2024-04-15",
		namespaceName:   namespaceName,
		buffEntryModels: buffEntryModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Buff_CurrentBuffMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Buff::CurrentBuffMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	buffEntryModels := make([]map[string]interface{}, len(p.buffEntryModels))
	for i, item := range p.buffEntryModels {
		buffEntryModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":         p.version,
			"buffEntryModels": buffEntryModels,
		},
	}
}
