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

type MissionTaskModelTargetResetType string

const MissionTaskModelTargetResetTypeNotReset = MissionTaskModelTargetResetType("notReset")
const MissionTaskModelTargetResetTypeDaily = MissionTaskModelTargetResetType("daily")
const MissionTaskModelTargetResetTypeWeekly = MissionTaskModelTargetResetType("weekly")
const MissionTaskModelTargetResetTypeMonthly = MissionTaskModelTargetResetType("monthly")

func (p MissionTaskModelTargetResetType) Pointer() *MissionTaskModelTargetResetType {
	return &p
}

type MissionTaskModel struct {
	Name                   string
	Metadata               *string
	CounterName            string
	TargetResetType        *MissionTaskModelTargetResetType
	TargetValue            int64
	CompleteAcquireActions []AcquireAction
	ChallengePeriodEventId *string
	PremiseMissionTaskName *string
}

type MissionTaskModelOptions struct {
	Metadata               *string
	TargetResetType        *MissionTaskModelTargetResetType
	CompleteAcquireActions []AcquireAction
	ChallengePeriodEventId *string
	PremiseMissionTaskName *string
}

func NewMissionTaskModel(
	name string,
	counterName string,
	targetValue int64,
	options MissionTaskModelOptions,
) MissionTaskModel {
	data := MissionTaskModel{
		Name:                   name,
		CounterName:            counterName,
		TargetValue:            targetValue,
		Metadata:               options.Metadata,
		TargetResetType:        options.TargetResetType,
		CompleteAcquireActions: options.CompleteAcquireActions,
		ChallengePeriodEventId: options.ChallengePeriodEventId,
		PremiseMissionTaskName: options.PremiseMissionTaskName,
	}
	return data
}

func (p *MissionTaskModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["CounterName"] = p.CounterName
	if p.TargetResetType != nil {
		properties["TargetResetType"] = p.TargetResetType
	}
	properties["TargetValue"] = p.TargetValue
	{
		values := make([]map[string]interface{}, len(p.CompleteAcquireActions))
		for i, element := range p.CompleteAcquireActions {
			values[i] = element.Properties()
		}
		properties["CompleteAcquireActions"] = values
	}
	if p.ChallengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.ChallengePeriodEventId
	}
	if p.PremiseMissionTaskName != nil {
		properties["PremiseMissionTaskName"] = p.PremiseMissionTaskName
	}
	return properties
}
