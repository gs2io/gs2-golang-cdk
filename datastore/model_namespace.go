package datastore

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

type Namespace struct {
	CdkResource
	stack              *Stack
	Name               string
	Description        *string
	TransactionSetting *TransactionSetting
	LogSetting         *LogSetting
	DoneUploadScript   *ScriptSetting
}

type NamespaceOptions struct {
	Description        *string
	TransactionSetting *TransactionSetting
	LogSetting         *LogSetting
	DoneUploadScript   *ScriptSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:              stack,
		Name:               name,
		Description:        options.Description,
		TransactionSetting: options.TransactionSetting,
		LogSetting:         options.LogSetting,
		DoneUploadScript:   options.DoneUploadScript,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Datastore_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Datastore::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.TransactionSetting != nil {
		properties["TransactionSetting"] = p.TransactionSetting.Properties()
	}
	if p.LogSetting != nil {
		properties["LogSetting"] = p.LogSetting.Properties()
	}
	if p.DoneUploadScript != nil {
		properties["DoneUploadScript"] = p.DoneUploadScript.Properties()
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
