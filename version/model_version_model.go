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

type VersionModelType string

const VersionModelTypeSimple = VersionModelType("simple")
const VersionModelTypeSchedule = VersionModelType("schedule")

func (p VersionModelType) Pointer() *VersionModelType {
	return &p
}

type VersionModelApproveRequirement string

const VersionModelApproveRequirementRequired = VersionModelApproveRequirement("required")
const VersionModelApproveRequirementOptional = VersionModelApproveRequirement("optional")

func (p VersionModelApproveRequirement) Pointer() *VersionModelApproveRequirement {
	return &p
}

type VersionModel struct {
	Name               string
	Metadata           *string
	Scope              VersionModelScope
	Type               VersionModelType
	CurrentVersion     *Version
	WarningVersion     *Version
	ErrorVersion       *Version
	ScheduleVersions   []ScheduleVersion
	NeedSignature      *bool
	SignatureKeyId     *string
	ApproveRequirement *VersionModelApproveRequirement
}

type VersionModelOptions struct {
	Metadata           *string
	CurrentVersion     *Version
	WarningVersion     *Version
	ErrorVersion       *Version
	ScheduleVersions   []ScheduleVersion
	NeedSignature      *bool
	SignatureKeyId     *string
	ApproveRequirement *VersionModelApproveRequirement
}

func NewVersionModel(
	name string,
	scope VersionModelScope,
	type_ VersionModelType,
	options VersionModelOptions,
) VersionModel {
	_data := VersionModel{
		Name:               name,
		Scope:              scope,
		Type:               type_,
		Metadata:           options.Metadata,
		CurrentVersion:     options.CurrentVersion,
		WarningVersion:     options.WarningVersion,
		ErrorVersion:       options.ErrorVersion,
		ScheduleVersions:   options.ScheduleVersions,
		NeedSignature:      options.NeedSignature,
		SignatureKeyId:     options.SignatureKeyId,
		ApproveRequirement: options.ApproveRequirement,
	}
	return _data
}

type VersionModelTypeIsSimpleOptions struct {
	Metadata         *string
	ScheduleVersions []ScheduleVersion
}

func NewVersionModelTypeIsSimple(
	name string,
	scope VersionModelScope,
	warningVersion Version,
	errorVersion Version,
	options VersionModelTypeIsSimpleOptions,
) VersionModel {
	return NewVersionModel(
		name,
		scope,
		VersionModelTypeSimple,
		VersionModelOptions{
			Metadata:         options.Metadata,
			WarningVersion:   &warningVersion,
			ErrorVersion:     &errorVersion,
			ScheduleVersions: options.ScheduleVersions,
		},
	)
}

type VersionModelTypeIsScheduleOptions struct {
	Metadata         *string
	ScheduleVersions []ScheduleVersion
}

func NewVersionModelTypeIsSchedule(
	name string,
	scope VersionModelScope,
	options VersionModelTypeIsScheduleOptions,
) VersionModel {
	return NewVersionModel(
		name,
		scope,
		VersionModelTypeSchedule,
		VersionModelOptions{
			Metadata:         options.Metadata,
			ScheduleVersions: options.ScheduleVersions,
		},
	)
}

type VersionModelScopeIsPassiveOptions struct {
	Metadata         *string
	ScheduleVersions []ScheduleVersion
}

func NewVersionModelScopeIsPassive(
	name string,
	type_ VersionModelType,
	needSignature bool,
	options VersionModelScopeIsPassiveOptions,
) VersionModel {
	return NewVersionModel(
		name,
		VersionModelScopePassive,
		type_,
		VersionModelOptions{
			Metadata:         options.Metadata,
			ScheduleVersions: options.ScheduleVersions,
			NeedSignature:    &needSignature,
		},
	)
}

type VersionModelScopeIsActiveOptions struct {
	Metadata         *string
	ScheduleVersions []ScheduleVersion
}

func NewVersionModelScopeIsActive(
	name string,
	type_ VersionModelType,
	approveRequirement VersionModelApproveRequirement,
	options VersionModelScopeIsActiveOptions,
) VersionModel {
	return NewVersionModel(
		name,
		VersionModelScopeActive,
		type_,
		VersionModelOptions{
			Metadata:           options.Metadata,
			ScheduleVersions:   options.ScheduleVersions,
			ApproveRequirement: &approveRequirement,
		},
	)
}

func (p *VersionModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["Scope"] = p.Scope
	properties["Type"] = p.Type
	if p.CurrentVersion != nil {
		properties["CurrentVersion"] = p.CurrentVersion.Properties()
	}
	if p.WarningVersion != nil {
		properties["WarningVersion"] = p.WarningVersion.Properties()
	}
	if p.ErrorVersion != nil {
		properties["ErrorVersion"] = p.ErrorVersion.Properties()
	}
	{
		values := make([]map[string]interface{}, len(p.ScheduleVersions))
		for i, element := range p.ScheduleVersions {
			values[i] = element.Properties()
		}
		properties["ScheduleVersions"] = values
	}
	if p.NeedSignature != nil {
		properties["NeedSignature"] = p.NeedSignature
	}
	if p.SignatureKeyId != nil {
		properties["SignatureKeyId"] = p.SignatureKeyId
	}
	if p.ApproveRequirement != nil {
		properties["ApproveRequirement"] = p.ApproveRequirement
	}
	return properties
}
