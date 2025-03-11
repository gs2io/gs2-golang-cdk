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

type MissionGroupModelResetType string

const MissionGroupModelResetTypeNotReset = MissionGroupModelResetType("notReset")
const MissionGroupModelResetTypeDaily = MissionGroupModelResetType("daily")
const MissionGroupModelResetTypeWeekly = MissionGroupModelResetType("weekly")
const MissionGroupModelResetTypeMonthly = MissionGroupModelResetType("monthly")
const MissionGroupModelResetTypeDays = MissionGroupModelResetType("days")

func (p MissionGroupModelResetType) Pointer() *MissionGroupModelResetType {
	return &p
}

type MissionGroupModelResetDayOfWeek string

const MissionGroupModelResetDayOfWeekSunday = MissionGroupModelResetDayOfWeek("sunday")
const MissionGroupModelResetDayOfWeekMonday = MissionGroupModelResetDayOfWeek("monday")
const MissionGroupModelResetDayOfWeekTuesday = MissionGroupModelResetDayOfWeek("tuesday")
const MissionGroupModelResetDayOfWeekWednesday = MissionGroupModelResetDayOfWeek("wednesday")
const MissionGroupModelResetDayOfWeekThursday = MissionGroupModelResetDayOfWeek("thursday")
const MissionGroupModelResetDayOfWeekFriday = MissionGroupModelResetDayOfWeek("friday")
const MissionGroupModelResetDayOfWeekSaturday = MissionGroupModelResetDayOfWeek("saturday")

func (p MissionGroupModelResetDayOfWeek) Pointer() *MissionGroupModelResetDayOfWeek {
	return &p
}

type MissionGroupModel struct {
	Name                            string
	Metadata                        *string
	Tasks                           []MissionTaskModel
	ResetType                       MissionGroupModelResetType
	ResetDayOfMonth                 *int32
	ResetDayOfWeek                  *MissionGroupModelResetDayOfWeek
	ResetHour                       *int32
	CompleteNotificationNamespaceId *string
	AnchorTimestamp                 *int64
	Days                            *int32
}

type MissionGroupModelOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	ResetDayOfMonth                 *int32
	ResetDayOfWeek                  *MissionGroupModelResetDayOfWeek
	ResetHour                       *int32
	CompleteNotificationNamespaceId *string
	AnchorTimestamp                 *int64
	Days                            *int32
}

func NewMissionGroupModel(
	name string,
	resetType MissionGroupModelResetType,
	options MissionGroupModelOptions,
) MissionGroupModel {
	_data := MissionGroupModel{
		Name:                            name,
		ResetType:                       resetType,
		Metadata:                        options.Metadata,
		Tasks:                           options.Tasks,
		ResetDayOfMonth:                 options.ResetDayOfMonth,
		ResetDayOfWeek:                  options.ResetDayOfWeek,
		ResetHour:                       options.ResetHour,
		CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
		AnchorTimestamp:                 options.AnchorTimestamp,
		Days:                            options.Days,
	}
	return _data
}

type MissionGroupModelResetTypeIsNotResetOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	CompleteNotificationNamespaceId *string
}

func NewMissionGroupModelResetTypeIsNotReset(
	name string,
	options MissionGroupModelResetTypeIsNotResetOptions,
) MissionGroupModel {
	return NewMissionGroupModel(
		name,
		MissionGroupModelResetTypeNotReset,
		MissionGroupModelOptions{
			Metadata:                        options.Metadata,
			Tasks:                           options.Tasks,
			CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
		},
	)
}

type MissionGroupModelResetTypeIsDailyOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	CompleteNotificationNamespaceId *string
}

func NewMissionGroupModelResetTypeIsDaily(
	name string,
	resetHour int32,
	options MissionGroupModelResetTypeIsDailyOptions,
) MissionGroupModel {
	return NewMissionGroupModel(
		name,
		MissionGroupModelResetTypeDaily,
		MissionGroupModelOptions{
			Metadata:                        options.Metadata,
			Tasks:                           options.Tasks,
			ResetHour:                       &resetHour,
			CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
		},
	)
}

type MissionGroupModelResetTypeIsWeeklyOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	CompleteNotificationNamespaceId *string
}

func NewMissionGroupModelResetTypeIsWeekly(
	name string,
	resetDayOfWeek MissionGroupModelResetDayOfWeek,
	resetHour int32,
	options MissionGroupModelResetTypeIsWeeklyOptions,
) MissionGroupModel {
	return NewMissionGroupModel(
		name,
		MissionGroupModelResetTypeWeekly,
		MissionGroupModelOptions{
			Metadata:                        options.Metadata,
			Tasks:                           options.Tasks,
			ResetDayOfWeek:                  &resetDayOfWeek,
			ResetHour:                       &resetHour,
			CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
		},
	)
}

type MissionGroupModelResetTypeIsMonthlyOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	CompleteNotificationNamespaceId *string
}

func NewMissionGroupModelResetTypeIsMonthly(
	name string,
	resetDayOfMonth int32,
	resetHour int32,
	options MissionGroupModelResetTypeIsMonthlyOptions,
) MissionGroupModel {
	return NewMissionGroupModel(
		name,
		MissionGroupModelResetTypeMonthly,
		MissionGroupModelOptions{
			Metadata:                        options.Metadata,
			Tasks:                           options.Tasks,
			ResetDayOfMonth:                 &resetDayOfMonth,
			ResetHour:                       &resetHour,
			CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
		},
	)
}

type MissionGroupModelResetTypeIsDaysOptions struct {
	Metadata                        *string
	Tasks                           []MissionTaskModel
	CompleteNotificationNamespaceId *string
}

func NewMissionGroupModelResetTypeIsDays(
	name string,
	anchorTimestamp int64,
	days int32,
	options MissionGroupModelResetTypeIsDaysOptions,
) MissionGroupModel {
	return NewMissionGroupModel(
		name,
		MissionGroupModelResetTypeDays,
		MissionGroupModelOptions{
			Metadata:                        options.Metadata,
			Tasks:                           options.Tasks,
			CompleteNotificationNamespaceId: options.CompleteNotificationNamespaceId,
			AnchorTimestamp:                 &anchorTimestamp,
			Days:                            &days,
		},
	)
}

func (p *MissionGroupModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.Tasks))
		for i, element := range p.Tasks {
			values[i] = element.Properties()
		}
		properties["Tasks"] = values
	}
	properties["ResetType"] = p.ResetType
	if p.ResetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.ResetDayOfMonth
	}
	if p.ResetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.ResetDayOfWeek
	}
	if p.ResetHour != nil {
		properties["ResetHour"] = p.ResetHour
	}
	if p.CompleteNotificationNamespaceId != nil {
		properties["CompleteNotificationNamespaceId"] = p.CompleteNotificationNamespaceId
	}
	if p.AnchorTimestamp != nil {
		properties["AnchorTimestamp"] = p.AnchorTimestamp
	}
	if p.Days != nil {
		properties["Days"] = p.Days
	}
	return properties
}
