package ranking2

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
	stack                  *Stack
	version                string
	namespaceName          string
	globalRankingModels    []GlobalRankingModel
	clusterRankingModels   []ClusterRankingModel
	subscribeRankingModels []SubscribeRankingModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	globalRankingModels []GlobalRankingModel,
	clusterRankingModels []ClusterRankingModel,
	subscribeRankingModels []SubscribeRankingModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:                  stack,
		version:                "2024-05-30",
		namespaceName:          namespaceName,
		globalRankingModels:    globalRankingModels,
		clusterRankingModels:   clusterRankingModels,
		subscribeRankingModels: subscribeRankingModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Ranking2_CurrentRankingMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Ranking2::CurrentRankingMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	globalRankingModels := make([]map[string]interface{}, len(p.globalRankingModels))
	for i, item := range p.globalRankingModels {
		globalRankingModels[i] = item.Properties()
	}
	clusterRankingModels := make([]map[string]interface{}, len(p.clusterRankingModels))
	for i, item := range p.clusterRankingModels {
		clusterRankingModels[i] = item.Properties()
	}
	subscribeRankingModels := make([]map[string]interface{}, len(p.subscribeRankingModels))
	for i, item := range p.subscribeRankingModels {
		subscribeRankingModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":                p.version,
			"globalRankingModels":    globalRankingModels,
			"clusterRankingModels":   clusterRankingModels,
			"subscribeRankingModels": subscribeRankingModels,
		},
	}
}
