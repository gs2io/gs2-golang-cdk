package realtime

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

type NamespaceServerType string

const NamespaceServerTypeRelay = NamespaceServerType("relay")

func (p NamespaceServerType) Pointer() *NamespaceServerType {
	return &p
}

type NamespaceServerSpec string

const NamespaceServerSpecRealtime1Nano = NamespaceServerSpec("realtime1.nano")

func (p NamespaceServerSpec) Pointer() *NamespaceServerSpec {
	return &p
}

type Namespace struct {
	CdkResource
	stack              *Stack
	Name               string
	Description        *string
	TransactionSetting *TransactionSetting
	ServerType         NamespaceServerType
	ServerSpec         NamespaceServerSpec
	CreateNotification *NotificationSetting
	LogSetting         *LogSetting
}

type NamespaceOptions struct {
	Description        *string
	TransactionSetting *TransactionSetting
	CreateNotification *NotificationSetting
	LogSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	serverType NamespaceServerType,
	serverSpec NamespaceServerSpec,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:              stack,
		Name:               name,
		ServerType:         serverType,
		ServerSpec:         serverSpec,
		Description:        options.Description,
		TransactionSetting: options.TransactionSetting,
		CreateNotification: options.CreateNotification,
		LogSetting:         options.LogSetting,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Realtime_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Realtime::Namespace"
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
	properties["ServerType"] = p.ServerType
	properties["ServerSpec"] = p.ServerSpec
	if p.CreateNotification != nil {
		properties["CreateNotification"] = p.CreateNotification.Properties()
	}
	if p.LogSetting != nil {
		properties["LogSetting"] = p.LogSetting.Properties()
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
