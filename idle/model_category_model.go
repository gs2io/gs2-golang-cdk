package idle

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

type CategoryModelRewardResetMode string

const CategoryModelRewardResetModeReset = CategoryModelRewardResetMode("Reset")
const CategoryModelRewardResetModeCarryOver = CategoryModelRewardResetMode("CarryOver")

func (p CategoryModelRewardResetMode) Pointer() *CategoryModelRewardResetMode {
	return &p
}

type CategoryModel struct {
	Name                      string
	Metadata                  *string
	RewardIntervalMinutes     int32
	DefaultMaximumIdleMinutes int32
	RewardResetMode           CategoryModelRewardResetMode
	AcquireActions            []AcquireActionList
	IdlePeriodScheduleId      *string
	ReceivePeriodScheduleId   *string
}

type CategoryModelOptions struct {
	Metadata                *string
	IdlePeriodScheduleId    *string
	ReceivePeriodScheduleId *string
}

func NewCategoryModel(
	name string,
	rewardIntervalMinutes int32,
	defaultMaximumIdleMinutes int32,
	rewardResetMode CategoryModelRewardResetMode,
	acquireActions []AcquireActionList,
	options CategoryModelOptions,
) CategoryModel {
	_data := CategoryModel{
		Name:                      name,
		RewardIntervalMinutes:     rewardIntervalMinutes,
		DefaultMaximumIdleMinutes: defaultMaximumIdleMinutes,
		RewardResetMode:           rewardResetMode,
		AcquireActions:            acquireActions,
		Metadata:                  options.Metadata,
		IdlePeriodScheduleId:      options.IdlePeriodScheduleId,
		ReceivePeriodScheduleId:   options.ReceivePeriodScheduleId,
	}
	return _data
}

func (p *CategoryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["RewardIntervalMinutes"] = p.RewardIntervalMinutes
	properties["DefaultMaximumIdleMinutes"] = p.DefaultMaximumIdleMinutes
	properties["RewardResetMode"] = p.RewardResetMode
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	if p.IdlePeriodScheduleId != nil {
		properties["IdlePeriodScheduleId"] = p.IdlePeriodScheduleId
	}
	if p.ReceivePeriodScheduleId != nil {
		properties["ReceivePeriodScheduleId"] = p.ReceivePeriodScheduleId
	}
	return properties
}
