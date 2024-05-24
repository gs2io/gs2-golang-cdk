package matchmaking

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

type NamespaceEnableDisconnectDetection string

const NamespaceEnableDisconnectDetectionDisable = NamespaceEnableDisconnectDetection("disable")
const NamespaceEnableDisconnectDetectionEnable = NamespaceEnableDisconnectDetection("enable")

func (p NamespaceEnableDisconnectDetection) Pointer() *NamespaceEnableDisconnectDetection {
	return &p
}

type NamespaceCreateGatheringTriggerType string

const NamespaceCreateGatheringTriggerTypeNone = NamespaceCreateGatheringTriggerType("none")
const NamespaceCreateGatheringTriggerTypeGs2Realtime = NamespaceCreateGatheringTriggerType("gs2_realtime")
const NamespaceCreateGatheringTriggerTypeGs2Script = NamespaceCreateGatheringTriggerType("gs2_script")

func (p NamespaceCreateGatheringTriggerType) Pointer() *NamespaceCreateGatheringTriggerType {
	return &p
}

type NamespaceCompleteMatchmakingTriggerType string

const NamespaceCompleteMatchmakingTriggerTypeNone = NamespaceCompleteMatchmakingTriggerType("none")
const NamespaceCompleteMatchmakingTriggerTypeGs2Realtime = NamespaceCompleteMatchmakingTriggerType("gs2_realtime")
const NamespaceCompleteMatchmakingTriggerTypeGs2Script = NamespaceCompleteMatchmakingTriggerType("gs2_script")

func (p NamespaceCompleteMatchmakingTriggerType) Pointer() *NamespaceCompleteMatchmakingTriggerType {
	return &p
}

type NamespaceEnableCollaborateSeasonRating string

const NamespaceEnableCollaborateSeasonRatingEnable = NamespaceEnableCollaborateSeasonRating("enable")
const NamespaceEnableCollaborateSeasonRatingDisable = NamespaceEnableCollaborateSeasonRating("disable")

func (p NamespaceEnableCollaborateSeasonRating) Pointer() *NamespaceEnableCollaborateSeasonRating {
	return &p
}

type Namespace struct {
	CdkResource
	stack                                         *Stack
	Name                                          string
	Description                                   *string
	EnableRating                                  bool
	EnableDisconnectDetection                     NamespaceEnableDisconnectDetection
	DisconnectDetectionTimeoutSeconds             *int32
	CreateGatheringTriggerType                    NamespaceCreateGatheringTriggerType
	CreateGatheringTriggerRealtimeNamespaceId     *string
	CreateGatheringTriggerScriptId                *string
	CompleteMatchmakingTriggerType                NamespaceCompleteMatchmakingTriggerType
	CompleteMatchmakingTriggerRealtimeNamespaceId *string
	CompleteMatchmakingTriggerScriptId            *string
	EnableCollaborateSeasonRating                 NamespaceEnableCollaborateSeasonRating
	CollaborateSeasonRatingNamespaceId            *string
	CollaborateSeasonRatingTtl                    *int32
	ChangeRatingScript                            *ScriptSetting
	JoinNotification                              *NotificationSetting
	LeaveNotification                             *NotificationSetting
	CompleteNotification                          *NotificationSetting
	ChangeRatingNotification                      *NotificationSetting
	LogSetting                                    *LogSetting
}

type NamespaceOptions struct {
	Description                                   *string
	EnableRating                                  bool
	EnableDisconnectDetection                     NamespaceEnableDisconnectDetection
	DisconnectDetectionTimeoutSeconds             *int32
	CreateGatheringTriggerRealtimeNamespaceId     *string
	CreateGatheringTriggerScriptId                *string
	CompleteMatchmakingTriggerRealtimeNamespaceId *string
	CompleteMatchmakingTriggerScriptId            *string
	EnableCollaborateSeasonRating                 NamespaceEnableCollaborateSeasonRating
	CollaborateSeasonRatingNamespaceId            *string
	CollaborateSeasonRatingTtl                    *int32
	ChangeRatingScript                            *ScriptSetting
	JoinNotification                              *NotificationSetting
	LeaveNotification                             *NotificationSetting
	CompleteNotification                          *NotificationSetting
	ChangeRatingNotification                      *NotificationSetting
	LogSetting                                    *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	createGatheringTriggerType NamespaceCreateGatheringTriggerType,
	completeMatchmakingTriggerType NamespaceCompleteMatchmakingTriggerType,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                             stack,
		Name:                              name,
		CreateGatheringTriggerType:        createGatheringTriggerType,
		CompleteMatchmakingTriggerType:    completeMatchmakingTriggerType,
		Description:                       options.Description,
		EnableRating:                      options.EnableRating,
		EnableDisconnectDetection:         options.EnableDisconnectDetection,
		DisconnectDetectionTimeoutSeconds: options.DisconnectDetectionTimeoutSeconds,
		CreateGatheringTriggerRealtimeNamespaceId:     options.CreateGatheringTriggerRealtimeNamespaceId,
		CreateGatheringTriggerScriptId:                options.CreateGatheringTriggerScriptId,
		CompleteMatchmakingTriggerRealtimeNamespaceId: options.CompleteMatchmakingTriggerRealtimeNamespaceId,
		CompleteMatchmakingTriggerScriptId:            options.CompleteMatchmakingTriggerScriptId,
		EnableCollaborateSeasonRating:                 options.EnableCollaborateSeasonRating,
		CollaborateSeasonRatingNamespaceId:            options.CollaborateSeasonRatingNamespaceId,
		CollaborateSeasonRatingTtl:                    options.CollaborateSeasonRatingTtl,
		ChangeRatingScript:                            options.ChangeRatingScript,
		JoinNotification:                              options.JoinNotification,
		LeaveNotification:                             options.LeaveNotification,
		CompleteNotification:                          options.CompleteNotification,
		ChangeRatingNotification:                      options.ChangeRatingNotification,
		LogSetting:                                    options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Matchmaking_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Matchmaking::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["EnableRating"] = p.EnableRating
	properties["EnableDisconnectDetection"] = p.EnableDisconnectDetection
	if p.DisconnectDetectionTimeoutSeconds != nil {
		properties["DisconnectDetectionTimeoutSeconds"] = p.DisconnectDetectionTimeoutSeconds
	}
	properties["CreateGatheringTriggerType"] = p.CreateGatheringTriggerType
	if p.CreateGatheringTriggerRealtimeNamespaceId != nil {
		properties["CreateGatheringTriggerRealtimeNamespaceId"] = p.CreateGatheringTriggerRealtimeNamespaceId
	}
	if p.CreateGatheringTriggerScriptId != nil {
		properties["CreateGatheringTriggerScriptId"] = p.CreateGatheringTriggerScriptId
	}
	properties["CompleteMatchmakingTriggerType"] = p.CompleteMatchmakingTriggerType
	if p.CompleteMatchmakingTriggerRealtimeNamespaceId != nil {
		properties["CompleteMatchmakingTriggerRealtimeNamespaceId"] = p.CompleteMatchmakingTriggerRealtimeNamespaceId
	}
	if p.CompleteMatchmakingTriggerScriptId != nil {
		properties["CompleteMatchmakingTriggerScriptId"] = p.CompleteMatchmakingTriggerScriptId
	}
	properties["EnableCollaborateSeasonRating"] = p.EnableCollaborateSeasonRating
	if p.CollaborateSeasonRatingNamespaceId != nil {
		properties["CollaborateSeasonRatingNamespaceId"] = p.CollaborateSeasonRatingNamespaceId
	}
	if p.CollaborateSeasonRatingTtl != nil {
		properties["CollaborateSeasonRatingTtl"] = p.CollaborateSeasonRatingTtl
	}
	if p.ChangeRatingScript != nil {
		properties["ChangeRatingScript"] = p.ChangeRatingScript.Properties()
	}
	if p.JoinNotification != nil {
		properties["JoinNotification"] = p.JoinNotification.Properties()
	}
	if p.LeaveNotification != nil {
		properties["LeaveNotification"] = p.LeaveNotification.Properties()
	}
	if p.CompleteNotification != nil {
		properties["CompleteNotification"] = p.CompleteNotification.Properties()
	}
	if p.ChangeRatingNotification != nil {
		properties["ChangeRatingNotification"] = p.ChangeRatingNotification.Properties()
	}
	if p.LogSetting != nil {
		properties["LogSetting"] = p.LogSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		NamespaceName: p.Name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}

func (p *Namespace) MasterData(
	ratingModels []RatingModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		ratingModels,
	).AddDependsOn(
		p,
	)
	return p
}
