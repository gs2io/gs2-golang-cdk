package serialKey

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

type CampaignModel struct {
	name               string
	metadata           *string
	enableCampaignCode bool
}

func NewCampaignModel(
	name string,
	metadata *string,
	enableCampaignCode bool,
) *CampaignModel {
	return &CampaignModel{
		name:               name,
		metadata:           metadata,
		enableCampaignCode: enableCampaignCode,
	}
}

func (p *CampaignModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["EnableCampaignCode"] = p.enableCampaignCode
	return properties
}

func (p *CampaignModel) Ref(
	namespaceName string,
) CampaignModelRef {
	return CampaignModelRef{
		namespaceName:     namespaceName,
		campaignModelName: p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack          *Stack
	version        string
	namespaceName  string
	campaignModels []CampaignModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	campaignModels []CampaignModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:          stack,
		version:        "2022-09-13",
		namespaceName:  namespaceName,
		campaignModels: campaignModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "SerialKey_CurrentCampaignMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::SerialKey::CurrentCampaignMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	campaignModels := make([]map[string]interface{}, len(p.campaignModels))
	for i, item := range p.campaignModels {
		campaignModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":        p.version,
			"CampaignModels": campaignModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack       *Stack
	name        string
	description *string
	logSetting  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:       stack,
		name:        name,
		description: description,
		logSetting:  logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "SerialKey_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::SerialKey::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
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
	campaignModels []CampaignModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		campaignModels,
	).AddDependsOn(
		p,
	)
	return p
}

type CampaignModelMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	description        *string
	metadata           *string
	enableCampaignCode bool
}

func NewCampaignModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	enableCampaignCode bool,
	description *string,
	metadata *string,
) *CampaignModelMaster {

	self := CampaignModelMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		description:        description,
		metadata:           metadata,
		enableCampaignCode: enableCampaignCode,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *CampaignModelMaster) ResourceName() string {
	return "SerialKey_CampaignModelMaster_" + p.name
}

func (p *CampaignModelMaster) ResourceType() string {
	return "GS2::SerialKey::CampaignModelMaster"
}

func (p *CampaignModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["EnableCampaignCode"] = p.enableCampaignCode
	return properties
}

func (p *CampaignModelMaster) Ref(
	namespaceName string,
) CampaignModelMasterRef {
	return CampaignModelMasterRef{
		namespaceName:     namespaceName,
		campaignModelName: p.name,
	}
}

func (p *CampaignModelMaster) GetAttrCampaignId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.CampaignId",
	)
}
