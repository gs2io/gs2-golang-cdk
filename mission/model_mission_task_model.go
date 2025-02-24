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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type MissionTaskModelVerifyCompleteType string

const MissionTaskModelVerifyCompleteTypeCounter = MissionTaskModelVerifyCompleteType("counter")
const MissionTaskModelVerifyCompleteTypeVerifyActions = MissionTaskModelVerifyCompleteType("verifyActions")

func (p MissionTaskModelVerifyCompleteType) Pointer() *MissionTaskModelVerifyCompleteType {
	return &p
}

type MissionTaskModelTargetResetType string

const MissionTaskModelTargetResetTypeNotReset = MissionTaskModelTargetResetType("notReset")
const MissionTaskModelTargetResetTypeDaily = MissionTaskModelTargetResetType("daily")
const MissionTaskModelTargetResetTypeWeekly = MissionTaskModelTargetResetType("weekly")
const MissionTaskModelTargetResetTypeMonthly = MissionTaskModelTargetResetType("monthly")
const MissionTaskModelTargetResetTypeDays = MissionTaskModelTargetResetType("days")

func (p MissionTaskModelTargetResetType) Pointer() *MissionTaskModelTargetResetType {
	return &p
}

type MissionTaskModel struct {
	Name                         string
	Metadata                     *string
	VerifyCompleteType           MissionTaskModelVerifyCompleteType
	TargetCounter                *TargetCounterModel
	VerifyCompleteConsumeActions []VerifyAction
	CompleteAcquireActions       []AcquireAction
	ChallengePeriodEventId       *string
	PremiseMissionTaskName       *string
	TargetResetType              *MissionTaskModelTargetResetType
}

type MissionTaskModelOptions struct {
	Metadata                     *string
	TargetCounter                *TargetCounterModel
	VerifyCompleteConsumeActions []VerifyAction
	CompleteAcquireActions       []AcquireAction
	ChallengePeriodEventId       *string
	PremiseMissionTaskName       *string
	TargetResetType              *MissionTaskModelTargetResetType
}

func NewMissionTaskModel(
	name string,
	verifyCompleteType MissionTaskModelVerifyCompleteType,
	options MissionTaskModelOptions,
) MissionTaskModel {
	data := MissionTaskModel{
		Name:                         name,
		VerifyCompleteType:           verifyCompleteType,
		Metadata:                     options.Metadata,
		TargetCounter:                options.TargetCounter,
		VerifyCompleteConsumeActions: options.VerifyCompleteConsumeActions,
		CompleteAcquireActions:       options.CompleteAcquireActions,
		ChallengePeriodEventId:       options.ChallengePeriodEventId,
		PremiseMissionTaskName:       options.PremiseMissionTaskName,
		TargetResetType:              options.TargetResetType,
	}
	return data
}

type MissionTaskModelVerifyCompleteTypeIsCounterOptions struct {
	Metadata                     *string
	VerifyCompleteConsumeActions []VerifyAction
	CompleteAcquireActions       []AcquireAction
	ChallengePeriodEventId       *string
	PremiseMissionTaskName       *string
	TargetResetType              *MissionTaskModelTargetResetType
}

func NewMissionTaskModelVerifyCompleteTypeIsCounter(
	name string,
	targetCounter TargetCounterModel,
	options MissionTaskModelVerifyCompleteTypeIsCounterOptions,
) MissionTaskModel {
	return NewMissionTaskModel(
		name,
		MissionTaskModelVerifyCompleteTypeCounter,
		MissionTaskModelOptions{
			Metadata:                     options.Metadata,
			TargetCounter:                &targetCounter,
			VerifyCompleteConsumeActions: options.VerifyCompleteConsumeActions,
			CompleteAcquireActions:       options.CompleteAcquireActions,
			ChallengePeriodEventId:       options.ChallengePeriodEventId,
			PremiseMissionTaskName:       options.PremiseMissionTaskName,
			TargetResetType:              options.TargetResetType,
		},
	)
}

type MissionTaskModelVerifyCompleteTypeIsVerifyActionsOptions struct {
	Metadata                     *string
	VerifyCompleteConsumeActions []VerifyAction
	CompleteAcquireActions       []AcquireAction
	ChallengePeriodEventId       *string
	PremiseMissionTaskName       *string
	TargetResetType              *MissionTaskModelTargetResetType
}

func NewMissionTaskModelVerifyCompleteTypeIsVerifyActions(
	name string,
	options MissionTaskModelVerifyCompleteTypeIsVerifyActionsOptions,
) MissionTaskModel {
	return NewMissionTaskModel(
		name,
		MissionTaskModelVerifyCompleteTypeVerifyActions,
		MissionTaskModelOptions{
			Metadata:                     options.Metadata,
			VerifyCompleteConsumeActions: options.VerifyCompleteConsumeActions,
			CompleteAcquireActions:       options.CompleteAcquireActions,
			ChallengePeriodEventId:       options.ChallengePeriodEventId,
			PremiseMissionTaskName:       options.PremiseMissionTaskName,
			TargetResetType:              options.TargetResetType,
		},
	)
}

func (p *MissionTaskModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["VerifyCompleteType"] = p.VerifyCompleteType
	if p.TargetCounter != nil {
		properties["TargetCounter"] = p.TargetCounter.Properties()
	}
	{
		values := make([]map[string]interface{}, len(p.VerifyCompleteConsumeActions))
		for i, element := range p.VerifyCompleteConsumeActions {
			values[i] = element.Properties()
		}
		properties["VerifyCompleteConsumeActions"] = values
	}
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
