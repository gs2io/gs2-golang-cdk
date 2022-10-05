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

type RatingModel struct {
	name       string
	metadata   *string
	volatility int32
}

func NewRatingModel(
	name string,
	metadata *string,
	volatility int32,
) *RatingModel {
	return &RatingModel{
		name:       name,
		metadata:   metadata,
		volatility: volatility,
	}
}

func (p *RatingModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Volatility"] = p.volatility
	return properties
}

func (p *RatingModel) Ref(
	namespaceName string,
) RatingModelRef {
	return RatingModelRef{
		namespaceName: namespaceName,
		ratingName:    p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	ratingModels  []RatingModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	ratingModels []RatingModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2020-06-24",
		namespaceName: namespaceName,
		ratingModels:  ratingModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Matchmaking_CurrentRatingModelMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Matchmaking::CurrentRatingModelMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	ratingModels := make([]map[string]interface{}, len(p.ratingModels))
	for i, item := range p.ratingModels {
		ratingModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":      p.version,
			"RatingModels": ratingModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                                         *Stack
	name                                          string
	description                                   *string
	enableRating                                  bool
	createGatheringTriggerType                    string
	createGatheringTriggerRealtimeNamespaceId     *string
	createGatheringTriggerScriptId                *string
	completeMatchmakingTriggerType                string
	completeMatchmakingTriggerRealtimeNamespaceId *string
	completeMatchmakingTriggerScriptId            *string
	joinNotification                              *NotificationSetting
	leaveNotification                             *NotificationSetting
	completeNotification                          *NotificationSetting
	logSetting                                    *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	enableRating bool,
	createGatheringTriggerType string,
	completeMatchmakingTriggerType string,
	description *string,
	createGatheringTriggerRealtimeNamespaceId *string,
	createGatheringTriggerScriptId *string,
	completeMatchmakingTriggerRealtimeNamespaceId *string,
	completeMatchmakingTriggerScriptId *string,
	joinNotification *NotificationSetting,
	leaveNotification *NotificationSetting,
	completeNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                      stack,
		name:                       name,
		description:                description,
		enableRating:               enableRating,
		createGatheringTriggerType: createGatheringTriggerType,
		createGatheringTriggerRealtimeNamespaceId:     createGatheringTriggerRealtimeNamespaceId,
		createGatheringTriggerScriptId:                createGatheringTriggerScriptId,
		completeMatchmakingTriggerType:                completeMatchmakingTriggerType,
		completeMatchmakingTriggerRealtimeNamespaceId: completeMatchmakingTriggerRealtimeNamespaceId,
		completeMatchmakingTriggerScriptId:            completeMatchmakingTriggerScriptId,
		joinNotification:                              joinNotification,
		leaveNotification:                             leaveNotification,
		completeNotification:                          completeNotification,
		logSetting:                                    logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Matchmaking_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Matchmaking::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["EnableRating"] = p.enableRating
	properties["CreateGatheringTriggerType"] = p.createGatheringTriggerType
	if p.createGatheringTriggerRealtimeNamespaceId != nil {
		properties["CreateGatheringTriggerRealtimeNamespaceId"] = p.createGatheringTriggerRealtimeNamespaceId
	}
	if p.createGatheringTriggerScriptId != nil {
		properties["CreateGatheringTriggerScriptId"] = p.createGatheringTriggerScriptId
	}
	properties["CompleteMatchmakingTriggerType"] = p.completeMatchmakingTriggerType
	if p.completeMatchmakingTriggerRealtimeNamespaceId != nil {
		properties["CompleteMatchmakingTriggerRealtimeNamespaceId"] = p.completeMatchmakingTriggerRealtimeNamespaceId
	}
	if p.completeMatchmakingTriggerScriptId != nil {
		properties["CompleteMatchmakingTriggerScriptId"] = p.completeMatchmakingTriggerScriptId
	}
	if p.joinNotification != nil {
		properties["JoinNotification"] = p.joinNotification.Properties()
	}
	if p.leaveNotification != nil {
		properties["LeaveNotification"] = p.leaveNotification.Properties()
	}
	if p.completeNotification != nil {
		properties["CompleteNotification"] = p.completeNotification.Properties()
	}
	if p.logSetting != nil {
		properties["LogSetting"] = p.logSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		namespaceName: p.name,
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
		p.name,
		ratingModels,
	).AddDependsOn(
		p,
	)
	return p
}

type Gathering struct {
	CdkResource
	stack             *Stack
	namespaceName     string
	userId            string
	player            Player
	attributeRanges   []AttributeRange
	capacityOfRoles   []CapacityOfRole
	allowUserIds      []string
	expiresAt         *int64
	expiresAtTimeSpan *TimeSpan
}

func NewGathering(
	stack *Stack,
	namespaceName string,
	userId string,
	player Player,
	attributeRanges []AttributeRange,
	capacityOfRoles []CapacityOfRole,
	allowUserIds []string,
	expiresAt *int64,
	expiresAtTimeSpan *TimeSpan,
) *Gathering {

	self := Gathering{
		stack:             stack,
		namespaceName:     namespaceName,
		userId:            userId,
		player:            player,
		attributeRanges:   attributeRanges,
		capacityOfRoles:   capacityOfRoles,
		allowUserIds:      allowUserIds,
		expiresAt:         expiresAt,
		expiresAtTimeSpan: expiresAtTimeSpan,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Gathering) ResourceName() string {
	return "Matchmaking_Gathering_" + p.name
}

func (p *Gathering) ResourceType() string {
	return "GS2::Matchmaking::Gathering"
}

func (p *Gathering) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["UserId"] = p.userId
	properties["Player"] = p.player.Properties()
	{
		values := make([]map[string]interface{}, len(p.attributeRanges))
		for i, element := range p.attributeRanges {
			values[i] = element.Properties()
		}
		properties["AttributeRanges"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.capacityOfRoles))
		for i, element := range p.capacityOfRoles {
			values[i] = element.Properties()
		}
		properties["CapacityOfRoles"] = values
	}
	properties["AllowUserIds"] = p.allowUserIds
	if p.expiresAt != nil {
		properties["ExpiresAt"] = p.expiresAt
	}
	if p.expiresAtTimeSpan != nil {
		properties["ExpiresAtTimeSpan"] = p.expiresAtTimeSpan.Properties()
	}
	return properties
}

func (p *Gathering) Ref(
	namespaceName string,
	userId string,
) GatheringRef {
	return GatheringRef{
		namespaceName: namespaceName,
		userId:        userId,
		gatheringName: p.name,
	}
}

func (p *Gathering) GetAttrGatheringId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.GatheringId",
	)
}

type RatingModelMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
	volatility    int32
}

func NewRatingModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	volatility int32,
	description *string,
	metadata *string,
) *RatingModelMaster {

	self := RatingModelMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
		volatility:    volatility,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *RatingModelMaster) ResourceName() string {
	return "Matchmaking_RatingModelMaster_" + p.name
}

func (p *RatingModelMaster) ResourceType() string {
	return "GS2::Matchmaking::RatingModelMaster"
}

func (p *RatingModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Volatility"] = p.volatility
	return properties
}

func (p *RatingModelMaster) Ref(
	namespaceName string,
) RatingModelMasterRef {
	return RatingModelMasterRef{
		namespaceName: namespaceName,
		ratingName:    p.name,
	}
}

func (p *RatingModelMaster) GetAttrRatingModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RatingModelId",
	)
}
