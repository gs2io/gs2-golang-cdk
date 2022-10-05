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

type CounterModel struct {
	name                   string
	metadata               *string
	scopes                 []CounterScopeModel
	challengePeriodEventId *string
}

func NewCounterModel(
	name string,
	metadata *string,
	scopes []CounterScopeModel,
	challengePeriodEventId *string,
) *CounterModel {
	return &CounterModel{
		name:                   name,
		metadata:               metadata,
		scopes:                 scopes,
		challengePeriodEventId: challengePeriodEventId,
	}
}

func (p *CounterModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.scopes))
		for i, element := range p.scopes {
			values[i] = element.Properties()
		}
		properties["Scopes"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	return properties
}

func (p *CounterModel) Ref(
	namespaceName string,
) CounterModelRef {
	return CounterModelRef{
		namespaceName: namespaceName,
		counterName:   p.name,
	}
}

type MissionGroupModelResetType string

const MissionGroupModelResetTypeNotReset = MissionGroupModelResetType("notReset")
const MissionGroupModelResetTypeDaily = MissionGroupModelResetType("daily")
const MissionGroupModelResetTypeWeekly = MissionGroupModelResetType("weekly")
const MissionGroupModelResetTypeMonthly = MissionGroupModelResetType("monthly")

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
	name                            string
	metadata                        *string
	tasks                           []MissionTaskModel
	resetType                       MissionGroupModelResetType
	resetDayOfMonth                 *int32
	resetDayOfWeek                  *MissionGroupModelResetDayOfWeek
	resetHour                       *int32
	completeNotificationNamespaceId *string
}

func NewMissionGroupModel(
	name string,
	metadata *string,
	tasks []MissionTaskModel,
	resetType MissionGroupModelResetType,
	resetDayOfMonth *int32,
	resetDayOfWeek *MissionGroupModelResetDayOfWeek,
	resetHour *int32,
	completeNotificationNamespaceId *string,
) *MissionGroupModel {
	return &MissionGroupModel{
		name:                            name,
		metadata:                        metadata,
		tasks:                           tasks,
		resetType:                       resetType,
		resetDayOfMonth:                 resetDayOfMonth,
		resetDayOfWeek:                  resetDayOfWeek,
		resetHour:                       resetHour,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
}

func (p *MissionGroupModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.tasks))
		for i, element := range p.tasks {
			values[i] = element.Properties()
		}
		properties["Tasks"] = values
	}
	properties["ResetType"] = p.resetType
	if p.resetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.resetDayOfMonth
	}
	if p.resetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.resetDayOfWeek
	}
	if p.resetHour != nil {
		properties["ResetHour"] = p.resetHour
	}
	if p.completeNotificationNamespaceId != nil {
		properties["CompleteNotificationNamespaceId"] = p.completeNotificationNamespaceId
	}
	return properties
}

func (p *MissionGroupModel) Ref(
	namespaceName string,
) MissionGroupModelRef {
	return MissionGroupModelRef{
		namespaceName:    namespaceName,
		missionGroupName: p.name,
	}
}

func NewNotResetMissionGroupModel(
	name string,
	tasks []MissionTaskModel,
	metadata *string,
	completeNotificationNamespaceId *string,
) MissionGroupModel {
	return MissionGroupModel{
		resetType:                       MissionGroupModelResetTypeNotReset,
		name:                            name,
		metadata:                        metadata,
		tasks:                           tasks,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
}

func NewDailyMissionGroupModel(
	name string,
	tasks []MissionTaskModel,
	resetHour int32,
	metadata *string,
	completeNotificationNamespaceId *string,
) MissionGroupModel {
	return MissionGroupModel{
		resetType:                       MissionGroupModelResetTypeDaily,
		name:                            name,
		metadata:                        metadata,
		tasks:                           tasks,
		resetHour:                       &resetHour,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
}

func NewWeeklyMissionGroupModel(
	name string,
	tasks []MissionTaskModel,
	resetDayOfWeek MissionGroupModelResetDayOfWeek,
	resetHour int32,
	metadata *string,
	completeNotificationNamespaceId *string,
) MissionGroupModel {
	return MissionGroupModel{
		resetType:                       MissionGroupModelResetTypeWeekly,
		name:                            name,
		metadata:                        metadata,
		tasks:                           tasks,
		resetDayOfWeek:                  &resetDayOfWeek,
		resetHour:                       &resetHour,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
}

func NewMonthlyMissionGroupModel(
	name string,
	tasks []MissionTaskModel,
	resetDayOfMonth int32,
	resetHour int32,
	metadata *string,
	completeNotificationNamespaceId *string,
) MissionGroupModel {
	return MissionGroupModel{
		resetType:                       MissionGroupModelResetTypeMonthly,
		name:                            name,
		metadata:                        metadata,
		tasks:                           tasks,
		resetDayOfMonth:                 &resetDayOfMonth,
		resetHour:                       &resetHour,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
}

type MissionTaskModel struct {
	name                   string
	metadata               *string
	counterName            string
	targetValue            int64
	completeAcquireActions []AcquireAction
	challengePeriodEventId *string
	premiseMissionTaskName *string
}

func NewMissionTaskModel(
	name string,
	metadata *string,
	counterName string,
	targetValue int64,
	completeAcquireActions []AcquireAction,
	challengePeriodEventId *string,
	premiseMissionTaskName *string,
) *MissionTaskModel {
	return &MissionTaskModel{
		name:                   name,
		metadata:               metadata,
		counterName:            counterName,
		targetValue:            targetValue,
		completeAcquireActions: completeAcquireActions,
		challengePeriodEventId: challengePeriodEventId,
		premiseMissionTaskName: premiseMissionTaskName,
	}
}

func (p *MissionTaskModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["CounterName"] = p.counterName
	properties["TargetValue"] = p.targetValue
	{
		values := make([]map[string]interface{}, len(p.completeAcquireActions))
		for i, element := range p.completeAcquireActions {
			values[i] = element.Properties()
		}
		properties["CompleteAcquireActions"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	if p.premiseMissionTaskName != nil {
		properties["PremiseMissionTaskName"] = p.premiseMissionTaskName
	}
	return properties
}

func (p *MissionTaskModel) Ref(
	namespaceName string,
	missionGroupName string,
) MissionTaskModelRef {
	return MissionTaskModelRef{
		namespaceName:    namespaceName,
		missionGroupName: missionGroupName,
		missionTaskName:  p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack              *Stack
	version            string
	namespaceName      string
	missionGroupModels []MissionGroupModel
	counterModels      []CounterModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	missionGroupModels []MissionGroupModel,
	counterModels []CounterModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:              stack,
		version:            "2019-05-28",
		namespaceName:      namespaceName,
		missionGroupModels: missionGroupModels,
		counterModels:      counterModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Mission_CurrentMissionMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Mission::CurrentMissionMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	missionGroupModels := make([]map[string]interface{}, len(p.missionGroupModels))
	for i, item := range p.missionGroupModels {
		missionGroupModels[i] = item.Properties()
	}
	counterModels := make([]map[string]interface{}, len(p.counterModels))
	for i, item := range p.counterModels {
		counterModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":            p.version,
			"MissionGroupModels": missionGroupModels,
			"CounterModels":      counterModels,
		},
	}
}

type CounterModelMaster struct {
	CdkResource
	stack                  *Stack
	namespaceName          string
	name                   string
	metadata               *string
	description            *string
	scopes                 []CounterScopeModel
	challengePeriodEventId *string
}

func NewCounterModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	scopes []CounterScopeModel,
	metadata *string,
	description *string,
	challengePeriodEventId *string,
) *CounterModelMaster {

	self := CounterModelMaster{
		stack:                  stack,
		namespaceName:          namespaceName,
		name:                   name,
		metadata:               metadata,
		description:            description,
		scopes:                 scopes,
		challengePeriodEventId: challengePeriodEventId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *CounterModelMaster) ResourceName() string {
	return "Mission_CounterModelMaster_" + p.name
}

func (p *CounterModelMaster) ResourceType() string {
	return "GS2::Mission::CounterModelMaster"
}

func (p *CounterModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.description != nil {
		properties["Description"] = p.description
	}
	{
		values := make([]map[string]interface{}, len(p.scopes))
		for i, element := range p.scopes {
			values[i] = element.Properties()
		}
		properties["Scopes"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	return properties
}

func (p *CounterModelMaster) Ref(
	namespaceName string,
) CounterModelMasterRef {
	return CounterModelMasterRef{
		namespaceName: namespaceName,
		counterName:   p.name,
	}
}

func (p *CounterModelMaster) GetAttrCounterId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.CounterId",
	)
}

type MissionGroupModelMaster struct {
	CdkResource
	stack                           *Stack
	namespaceName                   string
	name                            string
	metadata                        *string
	description                     *string
	resetType                       string
	resetDayOfMonth                 *int32
	resetDayOfWeek                  *string
	resetHour                       *int32
	completeNotificationNamespaceId *string
}

func NewMissionGroupModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	resetType string,
	metadata *string,
	description *string,
	resetDayOfMonth *int32,
	resetDayOfWeek *string,
	resetHour *int32,
	completeNotificationNamespaceId *string,
) *MissionGroupModelMaster {

	self := MissionGroupModelMaster{
		stack:                           stack,
		namespaceName:                   namespaceName,
		name:                            name,
		metadata:                        metadata,
		description:                     description,
		resetType:                       resetType,
		resetDayOfMonth:                 resetDayOfMonth,
		resetDayOfWeek:                  resetDayOfWeek,
		resetHour:                       resetHour,
		completeNotificationNamespaceId: completeNotificationNamespaceId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *MissionGroupModelMaster) ResourceName() string {
	return "Mission_MissionGroupModelMaster_" + p.name
}

func (p *MissionGroupModelMaster) ResourceType() string {
	return "GS2::Mission::MissionGroupModelMaster"
}

func (p *MissionGroupModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["ResetType"] = p.resetType
	if p.resetDayOfMonth != nil {
		properties["ResetDayOfMonth"] = p.resetDayOfMonth
	}
	if p.resetDayOfWeek != nil {
		properties["ResetDayOfWeek"] = p.resetDayOfWeek
	}
	if p.resetHour != nil {
		properties["ResetHour"] = p.resetHour
	}
	if p.completeNotificationNamespaceId != nil {
		properties["CompleteNotificationNamespaceId"] = p.completeNotificationNamespaceId
	}
	return properties
}

func (p *MissionGroupModelMaster) Ref(
	namespaceName string,
) MissionGroupModelMasterRef {
	return MissionGroupModelMasterRef{
		namespaceName:    namespaceName,
		missionGroupName: p.name,
	}
}

func (p *MissionGroupModelMaster) GetAttrMissionGroupId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.MissionGroupId",
	)
}

type Namespace struct {
	CdkResource
	stack                  *Stack
	name                   string
	description            *string
	transactionSetting     TransactionSetting
	missionCompleteScript  *ScriptSetting
	counterIncrementScript *ScriptSetting
	receiveRewardsScript   *ScriptSetting
	completeNotification   *NotificationSetting
	logSetting             *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	transactionSetting TransactionSetting,
	description *string,
	missionCompleteScript *ScriptSetting,
	counterIncrementScript *ScriptSetting,
	receiveRewardsScript *ScriptSetting,
	completeNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                  stack,
		name:                   name,
		description:            description,
		transactionSetting:     transactionSetting,
		missionCompleteScript:  missionCompleteScript,
		counterIncrementScript: counterIncrementScript,
		receiveRewardsScript:   receiveRewardsScript,
		completeNotification:   completeNotification,
		logSetting:             logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Mission_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Mission::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.missionCompleteScript != nil {
		properties["MissionCompleteScript"] = p.missionCompleteScript.Properties()
	}
	if p.counterIncrementScript != nil {
		properties["CounterIncrementScript"] = p.counterIncrementScript.Properties()
	}
	if p.receiveRewardsScript != nil {
		properties["ReceiveRewardsScript"] = p.receiveRewardsScript.Properties()
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
	missionGroupModels []MissionGroupModel,
	counterModels []CounterModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		missionGroupModels,
		counterModels,
	).AddDependsOn(
		p,
	)
	return p
}

type MissionTaskModelMaster struct {
	CdkResource
	stack                  *Stack
	namespaceName          string
	missionGroupName       string
	name                   string
	metadata               *string
	description            *string
	counterName            string
	targetValue            int64
	completeAcquireActions []AcquireAction
	challengePeriodEventId *string
	premiseMissionTaskName *string
}

func NewMissionTaskModelMaster(
	stack *Stack,
	namespaceName string,
	missionGroupName string,
	name string,
	counterName string,
	targetValue int64,
	completeAcquireActions []AcquireAction,
	metadata *string,
	description *string,
	challengePeriodEventId *string,
	premiseMissionTaskName *string,
) *MissionTaskModelMaster {

	self := MissionTaskModelMaster{
		stack:                  stack,
		namespaceName:          namespaceName,
		missionGroupName:       missionGroupName,
		name:                   name,
		metadata:               metadata,
		description:            description,
		counterName:            counterName,
		targetValue:            targetValue,
		completeAcquireActions: completeAcquireActions,
		challengePeriodEventId: challengePeriodEventId,
		premiseMissionTaskName: premiseMissionTaskName,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *MissionTaskModelMaster) ResourceName() string {
	return "Mission_MissionTaskModelMaster_" + p.name
}

func (p *MissionTaskModelMaster) ResourceType() string {
	return "GS2::Mission::MissionTaskModelMaster"
}

func (p *MissionTaskModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["MissionGroupName"] = p.missionGroupName
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["CounterName"] = p.counterName
	properties["TargetValue"] = p.targetValue
	{
		values := make([]map[string]interface{}, len(p.completeAcquireActions))
		for i, element := range p.completeAcquireActions {
			values[i] = element.Properties()
		}
		properties["CompleteAcquireActions"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	if p.premiseMissionTaskName != nil {
		properties["PremiseMissionTaskName"] = p.premiseMissionTaskName
	}
	return properties
}

func (p *MissionTaskModelMaster) Ref(
	namespaceName string,
	missionGroupName string,
) MissionTaskModelMasterRef {
	return MissionTaskModelMasterRef{
		namespaceName:    namespaceName,
		missionGroupName: missionGroupName,
		missionTaskName:  p.name,
	}
}

func (p *MissionTaskModelMaster) GetAttrMissionTaskId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.MissionTaskId",
	)
}
