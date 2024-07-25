package loginReward

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

type BonusModelMode string

const BonusModelModeSchedule = BonusModelMode("schedule")
const BonusModelModeStreaming = BonusModelMode("streaming")

func (p BonusModelMode) Pointer() *BonusModelMode {
	return &p
}

type BonusModelRepeat string

const BonusModelRepeatEnabled = BonusModelRepeat("enabled")
const BonusModelRepeatDisabled = BonusModelRepeat("disabled")

func (p BonusModelRepeat) Pointer() *BonusModelRepeat {
	return &p
}

type BonusModelMissedReceiveRelief string

const BonusModelMissedReceiveReliefEnabled = BonusModelMissedReceiveRelief("enabled")
const BonusModelMissedReceiveReliefDisabled = BonusModelMissedReceiveRelief("disabled")

func (p BonusModelMissedReceiveRelief) Pointer() *BonusModelMissedReceiveRelief {
	return &p
}

type BonusModel struct {
	Name                              string
	Metadata                          *string
	Mode                              BonusModelMode
	PeriodEventId                     *string
	ResetHour                         *int32
	Repeat                            *BonusModelRepeat
	Rewards                           []Reward
	MissedReceiveRelief               BonusModelMissedReceiveRelief
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

type BonusModelOptions struct {
	Metadata                          *string
	PeriodEventId                     *string
	ResetHour                         *int32
	Repeat                            *BonusModelRepeat
	Rewards                           []Reward
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

func NewBonusModel(
	name string,
	mode BonusModelMode,
	missedReceiveRelief BonusModelMissedReceiveRelief,
	options BonusModelOptions,
) BonusModel {
	data := BonusModel{
		Name:                              name,
		Mode:                              mode,
		MissedReceiveRelief:               missedReceiveRelief,
		Metadata:                          options.Metadata,
		PeriodEventId:                     options.PeriodEventId,
		ResetHour:                         options.ResetHour,
		Repeat:                            options.Repeat,
		Rewards:                           options.Rewards,
		MissedReceiveReliefVerifyActions:  options.MissedReceiveReliefVerifyActions,
		MissedReceiveReliefConsumeActions: options.MissedReceiveReliefConsumeActions,
	}
	return data
}

type BonusModelModeIsScheduleOptions struct {
	Metadata                          *string
	PeriodEventId                     *string
	Rewards                           []Reward
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

func NewBonusModelModeIsSchedule(
	name string,
	missedReceiveRelief BonusModelMissedReceiveRelief,
	options BonusModelModeIsScheduleOptions,
) BonusModel {
	return NewBonusModel(
		name,
		BonusModelModeSchedule,
		missedReceiveRelief,
		BonusModelOptions{
			Metadata:                          options.Metadata,
			PeriodEventId:                     options.PeriodEventId,
			Rewards:                           options.Rewards,
			MissedReceiveReliefVerifyActions:  options.MissedReceiveReliefVerifyActions,
			MissedReceiveReliefConsumeActions: options.MissedReceiveReliefConsumeActions,
		},
	)
}

type BonusModelModeIsStreamingOptions struct {
	Metadata                          *string
	PeriodEventId                     *string
	Rewards                           []Reward
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

func NewBonusModelModeIsStreaming(
	name string,
	missedReceiveRelief BonusModelMissedReceiveRelief,
	repeat BonusModelRepeat,
	options BonusModelModeIsStreamingOptions,
) BonusModel {
	return NewBonusModel(
		name,
		BonusModelModeStreaming,
		missedReceiveRelief,
		BonusModelOptions{
			Metadata:                          options.Metadata,
			PeriodEventId:                     options.PeriodEventId,
			Repeat:                            &repeat,
			Rewards:                           options.Rewards,
			MissedReceiveReliefVerifyActions:  options.MissedReceiveReliefVerifyActions,
			MissedReceiveReliefConsumeActions: options.MissedReceiveReliefConsumeActions,
		},
	)
}

type BonusModelMissedReceiveReliefIsEnabledOptions struct {
	Metadata                          *string
	PeriodEventId                     *string
	Rewards                           []Reward
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

func NewBonusModelMissedReceiveReliefIsEnabled(
	name string,
	mode BonusModelMode,
	options BonusModelMissedReceiveReliefIsEnabledOptions,
) BonusModel {
	return NewBonusModel(
		name,
		mode,
		BonusModelMissedReceiveReliefEnabled,
		BonusModelOptions{
			Metadata:                          options.Metadata,
			PeriodEventId:                     options.PeriodEventId,
			Rewards:                           options.Rewards,
			MissedReceiveReliefVerifyActions:  options.MissedReceiveReliefVerifyActions,
			MissedReceiveReliefConsumeActions: options.MissedReceiveReliefConsumeActions,
		},
	)
}

type BonusModelMissedReceiveReliefIsDisabledOptions struct {
	Metadata                          *string
	PeriodEventId                     *string
	Rewards                           []Reward
	MissedReceiveReliefVerifyActions  []VerifyAction
	MissedReceiveReliefConsumeActions []ConsumeAction
}

func NewBonusModelMissedReceiveReliefIsDisabled(
	name string,
	mode BonusModelMode,
	options BonusModelMissedReceiveReliefIsDisabledOptions,
) BonusModel {
	return NewBonusModel(
		name,
		mode,
		BonusModelMissedReceiveReliefDisabled,
		BonusModelOptions{
			Metadata:                          options.Metadata,
			PeriodEventId:                     options.PeriodEventId,
			Rewards:                           options.Rewards,
			MissedReceiveReliefVerifyActions:  options.MissedReceiveReliefVerifyActions,
			MissedReceiveReliefConsumeActions: options.MissedReceiveReliefConsumeActions,
		},
	)
}

func (p *BonusModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["Mode"] = p.Mode
	if p.PeriodEventId != nil {
		properties["PeriodEventId"] = p.PeriodEventId
	}
	if p.ResetHour != nil {
		properties["ResetHour"] = p.ResetHour
	}
	if p.Repeat != nil {
		properties["Repeat"] = p.Repeat
	}
	{
		values := make([]map[string]interface{}, len(p.Rewards))
		for i, element := range p.Rewards {
			values[i] = element.Properties()
		}
		properties["Rewards"] = values
	}
	properties["MissedReceiveRelief"] = p.MissedReceiveRelief
	{
		values := make([]map[string]interface{}, len(p.MissedReceiveReliefVerifyActions))
		for i, element := range p.MissedReceiveReliefVerifyActions {
			values[i] = element.Properties()
		}
		properties["MissedReceiveReliefVerifyActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.MissedReceiveReliefConsumeActions))
		for i, element := range p.MissedReceiveReliefConsumeActions {
			values[i] = element.Properties()
		}
		properties["MissedReceiveReliefConsumeActions"] = values
	}
	return properties
}
