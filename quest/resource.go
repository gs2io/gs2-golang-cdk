package quest

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

type QuestGroupModel struct {
	name                   string
	metadata               *string
	quests                 []QuestModel
	challengePeriodEventId *string
}

func NewQuestGroupModel(
	name string,
	metadata *string,
	quests []QuestModel,
	challengePeriodEventId *string,
) *QuestGroupModel {
	return &QuestGroupModel{
		name:                   name,
		metadata:               metadata,
		quests:                 quests,
		challengePeriodEventId: challengePeriodEventId,
	}
}

func (p *QuestGroupModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.quests))
		for i, element := range p.quests {
			values[i] = element.Properties()
		}
		properties["Quests"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	return properties
}

func (p *QuestGroupModel) Ref(
	namespaceName string,
) QuestGroupModelRef {
	return QuestGroupModelRef{
		namespaceName:  namespaceName,
		questGroupName: p.name,
	}
}

type QuestModel struct {
	name                        string
	metadata                    *string
	contents                    []Contents
	challengePeriodEventId      *string
	firstCompleteAcquireActions []AcquireAction
	consumeActions              []ConsumeAction
	failedAcquireActions        []AcquireAction
	premiseQuestNames           []string
}

func NewQuestModel(
	name string,
	metadata *string,
	contents []Contents,
	challengePeriodEventId *string,
	firstCompleteAcquireActions []AcquireAction,
	consumeActions []ConsumeAction,
	failedAcquireActions []AcquireAction,
	premiseQuestNames []string,
) *QuestModel {
	return &QuestModel{
		name:                        name,
		metadata:                    metadata,
		contents:                    contents,
		challengePeriodEventId:      challengePeriodEventId,
		firstCompleteAcquireActions: firstCompleteAcquireActions,
		consumeActions:              consumeActions,
		failedAcquireActions:        failedAcquireActions,
		premiseQuestNames:           premiseQuestNames,
	}
}

func (p *QuestModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.contents))
		for i, element := range p.contents {
			values[i] = element.Properties()
		}
		properties["Contents"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.firstCompleteAcquireActions))
		for i, element := range p.firstCompleteAcquireActions {
			values[i] = element.Properties()
		}
		properties["FirstCompleteAcquireActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.consumeActions))
		for i, element := range p.consumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.failedAcquireActions))
		for i, element := range p.failedAcquireActions {
			values[i] = element.Properties()
		}
		properties["FailedAcquireActions"] = values
	}
	properties["PremiseQuestNames"] = p.premiseQuestNames
	return properties
}

func (p *QuestModel) Ref(
	namespaceName string,
	questGroupName string,
) QuestModelRef {
	return QuestModelRef{
		namespaceName:  namespaceName,
		questGroupName: questGroupName,
		questName:      p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack            *Stack
	version          string
	namespaceName    string
	questGroupModels []QuestGroupModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	questGroupModels []QuestGroupModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:            stack,
		version:          "2019-05-14",
		namespaceName:    namespaceName,
		questGroupModels: questGroupModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Quest_CurrentQuestMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Quest::CurrentQuestMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	questGroupModels := make([]map[string]interface{}, len(p.questGroupModels))
	for i, item := range p.questGroupModels {
		questGroupModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":          p.version,
			"QuestGroupModels": questGroupModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack               *Stack
	name                string
	description         *string
	transactionSetting  TransactionSetting
	startQuestScript    *ScriptSetting
	completeQuestScript *ScriptSetting
	failedQuestScript   *ScriptSetting
	logSetting          *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	transactionSetting TransactionSetting,
	description *string,
	startQuestScript *ScriptSetting,
	completeQuestScript *ScriptSetting,
	failedQuestScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:               stack,
		name:                name,
		description:         description,
		transactionSetting:  transactionSetting,
		startQuestScript:    startQuestScript,
		completeQuestScript: completeQuestScript,
		failedQuestScript:   failedQuestScript,
		logSetting:          logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Quest_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Quest::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.startQuestScript != nil {
		properties["StartQuestScript"] = p.startQuestScript.Properties()
	}
	if p.completeQuestScript != nil {
		properties["CompleteQuestScript"] = p.completeQuestScript.Properties()
	}
	if p.failedQuestScript != nil {
		properties["FailedQuestScript"] = p.failedQuestScript.Properties()
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
	questGroupModels []QuestGroupModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		questGroupModels,
	).AddDependsOn(
		p,
	)
	return p
}

type QuestGroupModelMaster struct {
	CdkResource
	stack                  *Stack
	namespaceName          string
	name                   string
	description            *string
	metadata               *string
	challengePeriodEventId *string
}

func NewQuestGroupModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	description *string,
	metadata *string,
	challengePeriodEventId *string,
) *QuestGroupModelMaster {

	self := QuestGroupModelMaster{
		stack:                  stack,
		namespaceName:          namespaceName,
		name:                   name,
		description:            description,
		metadata:               metadata,
		challengePeriodEventId: challengePeriodEventId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *QuestGroupModelMaster) ResourceName() string {
	return "Quest_QuestGroupModelMaster_" + p.name
}

func (p *QuestGroupModelMaster) ResourceType() string {
	return "GS2::Quest::QuestGroupModelMaster"
}

func (p *QuestGroupModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	return properties
}

func (p *QuestGroupModelMaster) Ref(
	namespaceName string,
) QuestGroupModelMasterRef {
	return QuestGroupModelMasterRef{
		namespaceName:  namespaceName,
		questGroupName: p.name,
	}
}

func (p *QuestGroupModelMaster) GetAttrQuestGroupModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.QuestGroupModelId",
	)
}

type QuestModelMaster struct {
	CdkResource
	stack                       *Stack
	namespaceName               string
	questGroupName              string
	name                        string
	description                 *string
	metadata                    *string
	contents                    []Contents
	challengePeriodEventId      *string
	firstCompleteAcquireActions []AcquireAction
	consumeActions              []ConsumeAction
	failedAcquireActions        []AcquireAction
	premiseQuestNames           []string
}

func NewQuestModelMaster(
	stack *Stack,
	namespaceName string,
	questGroupName string,
	name string,
	contents []Contents,
	firstCompleteAcquireActions []AcquireAction,
	consumeActions []ConsumeAction,
	failedAcquireActions []AcquireAction,
	premiseQuestNames []string,
	description *string,
	metadata *string,
	challengePeriodEventId *string,
) *QuestModelMaster {

	self := QuestModelMaster{
		stack:                       stack,
		namespaceName:               namespaceName,
		questGroupName:              questGroupName,
		name:                        name,
		description:                 description,
		metadata:                    metadata,
		contents:                    contents,
		challengePeriodEventId:      challengePeriodEventId,
		firstCompleteAcquireActions: firstCompleteAcquireActions,
		consumeActions:              consumeActions,
		failedAcquireActions:        failedAcquireActions,
		premiseQuestNames:           premiseQuestNames,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *QuestModelMaster) ResourceName() string {
	return "Quest_QuestModelMaster_" + p.name
}

func (p *QuestModelMaster) ResourceType() string {
	return "GS2::Quest::QuestModelMaster"
}

func (p *QuestModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["QuestGroupName"] = p.questGroupName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.contents))
		for i, element := range p.contents {
			values[i] = element.Properties()
		}
		properties["Contents"] = values
	}
	if p.challengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.challengePeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.firstCompleteAcquireActions))
		for i, element := range p.firstCompleteAcquireActions {
			values[i] = element.Properties()
		}
		properties["FirstCompleteAcquireActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.consumeActions))
		for i, element := range p.consumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.failedAcquireActions))
		for i, element := range p.failedAcquireActions {
			values[i] = element.Properties()
		}
		properties["FailedAcquireActions"] = values
	}
	properties["PremiseQuestNames"] = p.premiseQuestNames
	return properties
}

func (p *QuestModelMaster) Ref(
	namespaceName string,
	questGroupName string,
) QuestModelMasterRef {
	return QuestModelMasterRef{
		namespaceName:  namespaceName,
		questGroupName: questGroupName,
		questName:      p.name,
	}
}

func (p *QuestModelMaster) GetAttrQuestModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.QuestModelId",
	)
}
