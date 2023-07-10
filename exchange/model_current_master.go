package exchange

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
	stack                 *Stack
	version               string
	namespaceName         string
	rateModels            []RateModel
	incrementalRateModels []IncrementalRateModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	rateModels []RateModel,
	incrementalRateModels []IncrementalRateModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:                 stack,
		version:               "2019-08-19",
		namespaceName:         namespaceName,
		rateModels:            rateModels,
		incrementalRateModels: incrementalRateModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Exchange_CurrentRateMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Exchange::CurrentRateMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	rateModels := make([]map[string]interface{}, len(p.rateModels))
	for i, item := range p.rateModels {
		rateModels[i] = item.Properties()
	}
	incrementalRateModels := make([]map[string]interface{}, len(p.incrementalRateModels))
	for i, item := range p.incrementalRateModels {
		incrementalRateModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":               p.version,
			"rateModels":            rateModels,
			"incrementalRateModels": incrementalRateModels,
		},
	}
}
