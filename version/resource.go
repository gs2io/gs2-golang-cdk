package version

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

type VersionModelScope string

const VersionModelScopePassive = VersionModelScope("passive")
const VersionModelScopeActive = VersionModelScope("active")

func (p VersionModelScope) Pointer() *VersionModelScope {
	return &p
}

type VersionModel struct {
	name           string
	metadata       *string
	warningVersion Version
	errorVersion   Version
	scope          VersionModelScope
	currentVersion *Version
	needSignature  *bool
	signatureKeyId *string
}

func NewVersionModel(
	name string,
	metadata *string,
	warningVersion Version,
	errorVersion Version,
	scope VersionModelScope,
	currentVersion *Version,
	needSignature *bool,
	signatureKeyId *string,
) *VersionModel {
	return &VersionModel{
		name:           name,
		metadata:       metadata,
		warningVersion: warningVersion,
		errorVersion:   errorVersion,
		scope:          scope,
		currentVersion: currentVersion,
		needSignature:  needSignature,
		signatureKeyId: signatureKeyId,
	}
}

func (p *VersionModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["WarningVersion"] = p.warningVersion.Properties()
	properties["ErrorVersion"] = p.errorVersion.Properties()
	properties["Scope"] = p.scope
	if p.currentVersion != nil {
		properties["CurrentVersion"] = p.currentVersion.Properties()
	}
	if p.needSignature != nil {
		properties["NeedSignature"] = p.needSignature
	}
	if p.signatureKeyId != nil {
		properties["SignatureKeyId"] = p.signatureKeyId
	}
	return properties
}

func (p *VersionModel) Ref(
	namespaceName string,
) VersionModelRef {
	return VersionModelRef{
		namespaceName: namespaceName,
		versionName:   p.name,
	}
}

func NewPassiveVersionModel(
	name string,
	warningVersion Version,
	errorVersion Version,
	needSignature bool,
	metadata *string,
	signatureKeyId *string,
) VersionModel {
	return VersionModel{
		scope:          VersionModelScopePassive,
		name:           name,
		metadata:       metadata,
		warningVersion: warningVersion,
		errorVersion:   errorVersion,
		needSignature:  &needSignature,
		signatureKeyId: signatureKeyId,
	}
}

func NewActiveVersionModel(
	name string,
	warningVersion Version,
	errorVersion Version,
	currentVersion Version,
	metadata *string,
	signatureKeyId *string,
) VersionModel {
	return VersionModel{
		scope:          VersionModelScopeActive,
		name:           name,
		metadata:       metadata,
		warningVersion: warningVersion,
		errorVersion:   errorVersion,
		currentVersion: &currentVersion,
		signatureKeyId: signatureKeyId,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	versionModels []VersionModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	versionModels []VersionModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-10-09",
		namespaceName: namespaceName,
		versionModels: versionModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Version_CurrentVersionMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Version::CurrentVersionMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	versionModels := make([]map[string]interface{}, len(p.versionModels))
	for i, item := range p.versionModels {
		versionModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":       p.version,
			"VersionModels": versionModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                       *Stack
	name                        string
	description                 *string
	assumeUserId                string
	acceptVersionScript         *ScriptSetting
	checkVersionTriggerScriptId *string
	logSetting                  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	assumeUserId string,
	description *string,
	acceptVersionScript *ScriptSetting,
	checkVersionTriggerScriptId *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                       stack,
		name:                        name,
		description:                 description,
		assumeUserId:                assumeUserId,
		acceptVersionScript:         acceptVersionScript,
		checkVersionTriggerScriptId: checkVersionTriggerScriptId,
		logSetting:                  logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Version_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Version::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["AssumeUserId"] = p.assumeUserId
	if p.acceptVersionScript != nil {
		properties["AcceptVersionScript"] = p.acceptVersionScript.Properties()
	}
	if p.checkVersionTriggerScriptId != nil {
		properties["CheckVersionTriggerScriptId"] = p.checkVersionTriggerScriptId
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
	versionModels []VersionModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		versionModels,
	).AddDependsOn(
		p,
	)
	return p
}

type VersionModelMaster struct {
	CdkResource
	stack          *Stack
	namespaceName  string
	name           string
	description    *string
	metadata       *string
	warningVersion Version
	errorVersion   Version
	scope          string
	currentVersion *Version
	needSignature  *bool
	signatureKeyId *string
}

func NewVersionModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	warningVersion Version,
	errorVersion Version,
	scope string,
	description *string,
	metadata *string,
	currentVersion *Version,
	needSignature *bool,
	signatureKeyId *string,
) *VersionModelMaster {

	self := VersionModelMaster{
		stack:          stack,
		namespaceName:  namespaceName,
		name:           name,
		description:    description,
		metadata:       metadata,
		warningVersion: warningVersion,
		errorVersion:   errorVersion,
		scope:          scope,
		currentVersion: currentVersion,
		needSignature:  needSignature,
		signatureKeyId: signatureKeyId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *VersionModelMaster) ResourceName() string {
	return "Version_VersionModelMaster_" + p.name
}

func (p *VersionModelMaster) ResourceType() string {
	return "GS2::Version::VersionModelMaster"
}

func (p *VersionModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["WarningVersion"] = p.warningVersion.Properties()
	properties["ErrorVersion"] = p.errorVersion.Properties()
	properties["Scope"] = p.scope
	if p.currentVersion != nil {
		properties["CurrentVersion"] = p.currentVersion.Properties()
	}
	if p.needSignature != nil {
		properties["NeedSignature"] = p.needSignature
	}
	if p.signatureKeyId != nil {
		properties["SignatureKeyId"] = p.signatureKeyId
	}
	return properties
}

func (p *VersionModelMaster) Ref(
	namespaceName string,
) VersionModelMasterRef {
	return VersionModelMasterRef{
		namespaceName: namespaceName,
		versionName:   p.name,
	}
}

func (p *VersionModelMaster) GetAttrVersionModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.VersionModelId",
	)
}
