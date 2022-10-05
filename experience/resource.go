package experience

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

type ExperienceModel struct {
	name              string
	metadata          *string
	defaultExperience int64
	defaultRankCap    int64
	maxRankCap        int64
	rankThreshold     Threshold
}

func NewExperienceModel(
	name string,
	metadata *string,
	defaultExperience int64,
	defaultRankCap int64,
	maxRankCap int64,
	rankThreshold Threshold,
) *ExperienceModel {
	return &ExperienceModel{
		name:              name,
		metadata:          metadata,
		defaultExperience: defaultExperience,
		defaultRankCap:    defaultRankCap,
		maxRankCap:        maxRankCap,
		rankThreshold:     rankThreshold,
	}
}

func (p *ExperienceModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["DefaultExperience"] = p.defaultExperience
	properties["DefaultRankCap"] = p.defaultRankCap
	properties["MaxRankCap"] = p.maxRankCap
	properties["RankThreshold"] = p.rankThreshold.Properties()
	return properties
}

func (p *ExperienceModel) Ref(
	namespaceName string,
) ExperienceModelRef {
	return ExperienceModelRef{
		namespaceName:  namespaceName,
		experienceName: p.name,
	}
}

type Threshold struct {
	metadata *string
	values   []int64
}

func NewThreshold(
	metadata *string,
	values []int64,
) *Threshold {
	return &Threshold{
		metadata: metadata,
		values:   values,
	}
}

func (p *Threshold) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Values"] = p.values
	return properties
}

type CurrentMasterData struct {
	CdkResource
	stack            *Stack
	version          string
	namespaceName    string
	experienceModels []ExperienceModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	experienceModels []ExperienceModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:            stack,
		version:          "2019-01-11",
		namespaceName:    namespaceName,
		experienceModels: experienceModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Experience_CurrentExperienceMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Experience::CurrentExperienceMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	experienceModels := make([]map[string]interface{}, len(p.experienceModels))
	for i, item := range p.experienceModels {
		experienceModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":          p.version,
			"ExperienceModels": experienceModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                    *Stack
	name                     string
	description              *string
	experienceCapScriptId    *string
	changeExperienceScript   *ScriptSetting
	changeRankScript         *ScriptSetting
	changeRankCapScript      *ScriptSetting
	overflowExperienceScript *ScriptSetting
	logSetting               *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	experienceCapScriptId *string,
	changeExperienceScript *ScriptSetting,
	changeRankScript *ScriptSetting,
	changeRankCapScript *ScriptSetting,
	overflowExperienceScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                    stack,
		name:                     name,
		description:              description,
		experienceCapScriptId:    experienceCapScriptId,
		changeExperienceScript:   changeExperienceScript,
		changeRankScript:         changeRankScript,
		changeRankCapScript:      changeRankCapScript,
		overflowExperienceScript: overflowExperienceScript,
		logSetting:               logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Experience_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Experience::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.experienceCapScriptId != nil {
		properties["ExperienceCapScriptId"] = p.experienceCapScriptId
	}
	if p.changeExperienceScript != nil {
		properties["ChangeExperienceScript"] = p.changeExperienceScript.Properties()
	}
	if p.changeRankScript != nil {
		properties["ChangeRankScript"] = p.changeRankScript.Properties()
	}
	if p.changeRankCapScript != nil {
		properties["ChangeRankCapScript"] = p.changeRankCapScript.Properties()
	}
	if p.overflowExperienceScript != nil {
		properties["OverflowExperienceScript"] = p.overflowExperienceScript.Properties()
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
	experienceModels []ExperienceModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		experienceModels,
	).AddDependsOn(
		p,
	)
	return p
}

type ExperienceModelMaster struct {
	CdkResource
	stack             *Stack
	namespaceName     string
	name              string
	description       *string
	metadata          *string
	defaultExperience int64
	defaultRankCap    int64
	maxRankCap        int64
	rankThresholdName string
}

func NewExperienceModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	defaultExperience int64,
	defaultRankCap int64,
	maxRankCap int64,
	rankThresholdName string,
	description *string,
	metadata *string,
) *ExperienceModelMaster {

	self := ExperienceModelMaster{
		stack:             stack,
		namespaceName:     namespaceName,
		name:              name,
		description:       description,
		metadata:          metadata,
		defaultExperience: defaultExperience,
		defaultRankCap:    defaultRankCap,
		maxRankCap:        maxRankCap,
		rankThresholdName: rankThresholdName,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *ExperienceModelMaster) ResourceName() string {
	return "Experience_ExperienceModelMaster_" + p.name
}

func (p *ExperienceModelMaster) ResourceType() string {
	return "GS2::Experience::ExperienceModelMaster"
}

func (p *ExperienceModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["DefaultExperience"] = p.defaultExperience
	properties["DefaultRankCap"] = p.defaultRankCap
	properties["MaxRankCap"] = p.maxRankCap
	properties["RankThresholdName"] = p.rankThresholdName
	return properties
}

func (p *ExperienceModelMaster) Ref(
	namespaceName string,
) ExperienceModelMasterRef {
	return ExperienceModelMasterRef{
		namespaceName:  namespaceName,
		experienceName: p.name,
	}
}

func (p *ExperienceModelMaster) GetAttrExperienceModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ExperienceModelId",
	)
}

type ThresholdMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
	values        []int64
}

func NewThresholdMaster(
	stack *Stack,
	namespaceName string,
	name string,
	values []int64,
	description *string,
	metadata *string,
) *ThresholdMaster {

	self := ThresholdMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
		values:        values,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *ThresholdMaster) ResourceName() string {
	return "Experience_ThresholdMaster_" + p.name
}

func (p *ThresholdMaster) ResourceType() string {
	return "GS2::Experience::ThresholdMaster"
}

func (p *ThresholdMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Values"] = p.values
	return properties
}

func (p *ThresholdMaster) Ref(
	namespaceName string,
) ThresholdMasterRef {
	return ThresholdMasterRef{
		namespaceName: namespaceName,
		thresholdName: p.name,
	}
}

func (p *ThresholdMaster) GetAttrThresholdId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ThresholdId",
	)
}
