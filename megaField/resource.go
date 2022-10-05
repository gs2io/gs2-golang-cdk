package megaField

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

type AreaModel struct {
	name     string
	metadata *string
}

func NewAreaModel(
	name string,
	metadata *string,
) *AreaModel {
	return &AreaModel{
		name:     name,
		metadata: metadata,
	}
}

func (p *AreaModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

func (p *AreaModel) Ref(
	namespaceName string,
) AreaModelRef {
	return AreaModelRef{
		namespaceName: namespaceName,
		areaModelName: p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	areaModels    []AreaModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	areaModels []AreaModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2022-08-28",
		namespaceName: namespaceName,
		areaModels:    areaModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "MegaField_CurrentFieldMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::MegaField::CurrentFieldMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	areaModels := make([]map[string]interface{}, len(p.areaModels))
	for i, item := range p.areaModels {
		areaModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":    p.version,
			"AreaModels": areaModels,
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
	return "MegaField_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::MegaField::Namespace"
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
	areaModels []AreaModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		areaModels,
	).AddDependsOn(
		p,
	)
	return p
}

type AreaModelMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
}

func NewAreaModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	description *string,
	metadata *string,
) *AreaModelMaster {

	self := AreaModelMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *AreaModelMaster) ResourceName() string {
	return "MegaField_AreaModelMaster_" + p.name
}

func (p *AreaModelMaster) ResourceType() string {
	return "GS2::MegaField::AreaModelMaster"
}

func (p *AreaModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

func (p *AreaModelMaster) Ref(
	namespaceName string,
) AreaModelMasterRef {
	return AreaModelMasterRef{
		namespaceName: namespaceName,
		areaModelName: p.name,
	}
}

func (p *AreaModelMaster) GetAttrAreaModelMasterId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.AreaModelMasterId",
	)
}

type LayerModelMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	areaModelName string
	name          string
	description   *string
	metadata      *string
}

func NewLayerModelMaster(
	stack *Stack,
	namespaceName string,
	areaModelName string,
	name string,
	description *string,
	metadata *string,
) *LayerModelMaster {

	self := LayerModelMaster{
		stack:         stack,
		namespaceName: namespaceName,
		areaModelName: areaModelName,
		name:          name,
		description:   description,
		metadata:      metadata,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *LayerModelMaster) ResourceName() string {
	return "MegaField_LayerModelMaster_" + p.name
}

func (p *LayerModelMaster) ResourceType() string {
	return "GS2::MegaField::LayerModelMaster"
}

func (p *LayerModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["AreaModelName"] = p.areaModelName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

func (p *LayerModelMaster) Ref(
	namespaceName string,
	areaModelName string,
) LayerModelMasterRef {
	return LayerModelMasterRef{
		namespaceName:  namespaceName,
		areaModelName:  areaModelName,
		layerModelName: p.name,
	}
}

func (p *LayerModelMaster) GetAttrLayerModelMasterId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.LayerModelMasterId",
	)
}
