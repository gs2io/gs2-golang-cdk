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
	Name           string
	Metadata       *string
	WarningVersion Version
	ErrorVersion   Version
	Scope          VersionModelScope
	CurrentVersion *Version
	NeedSignature  *bool
	SignatureKeyId *string
}

type VersionModelOptions struct {
	Metadata       *string
	CurrentVersion *Version
	NeedSignature  *bool
	SignatureKeyId *string
}

func NewVersionModel(
	name string,
	warningVersion Version,
	errorVersion Version,
	scope VersionModelScope,
	options VersionModelOptions,
) VersionModel {
	data := VersionModel{
		Name:           name,
		WarningVersion: warningVersion,
		ErrorVersion:   errorVersion,
		Scope:          scope,
		Metadata:       options.Metadata,
		CurrentVersion: options.CurrentVersion,
		NeedSignature:  options.NeedSignature,
		SignatureKeyId: options.SignatureKeyId,
	}
	return data
}

type VersionModelScopeIsPassiveOptions struct {
	Metadata *string
}

func NewVersionModelScopeIsPassive(
	name string,
	warningVersion Version,
	errorVersion Version,
	needSignature bool,
	options VersionModelScopeIsPassiveOptions,
) VersionModel {
	return NewVersionModel(
		name,
		warningVersion,
		errorVersion,
		VersionModelScopePassive,
		VersionModelOptions{
			Metadata:      options.Metadata,
			NeedSignature: &needSignature,
		},
	)
}

type VersionModelScopeIsActiveOptions struct {
	Metadata *string
}

func NewVersionModelScopeIsActive(
	name string,
	warningVersion Version,
	errorVersion Version,
	currentVersion Version,
	options VersionModelScopeIsActiveOptions,
) VersionModel {
	return NewVersionModel(
		name,
		warningVersion,
		errorVersion,
		VersionModelScopeActive,
		VersionModelOptions{
			Metadata:       options.Metadata,
			CurrentVersion: &currentVersion,
		},
	)
}

func (p *VersionModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["WarningVersion"] = p.WarningVersion.Properties()
	properties["ErrorVersion"] = p.ErrorVersion.Properties()
	properties["Scope"] = p.Scope
	if p.CurrentVersion != nil {
		properties["CurrentVersion"] = p.CurrentVersion.Properties()
	}
	if p.NeedSignature != nil {
		properties["NeedSignature"] = p.NeedSignature
	}
	if p.SignatureKeyId != nil {
		properties["SignatureKeyId"] = p.SignatureKeyId
	}
	return properties
}
