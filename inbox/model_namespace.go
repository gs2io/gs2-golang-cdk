package inbox

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
	stack                      *Stack
	Name                       string
	Description                *string
	IsAutomaticDeletingEnabled bool
	TransactionSetting         *TransactionSetting
	ReceiveMessageScript       *ScriptSetting
	ReadMessageScript          *ScriptSetting
	DeleteMessageScript        *ScriptSetting
	ReceiveNotification        *NotificationSetting
	LogSetting                 *LogSetting
}

type NamespaceOptions struct {
	Description                *string
	IsAutomaticDeletingEnabled bool
	TransactionSetting         *TransactionSetting
	ReceiveMessageScript       *ScriptSetting
	ReadMessageScript          *ScriptSetting
	DeleteMessageScript        *ScriptSetting
	ReceiveNotification        *NotificationSetting
	LogSetting                 *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                      stack,
		Name:                       name,
		Description:                options.Description,
		IsAutomaticDeletingEnabled: options.IsAutomaticDeletingEnabled,
		TransactionSetting:         options.TransactionSetting,
		ReceiveMessageScript:       options.ReceiveMessageScript,
		ReadMessageScript:          options.ReadMessageScript,
		DeleteMessageScript:        options.DeleteMessageScript,
		ReceiveNotification:        options.ReceiveNotification,
		LogSetting:                 options.LogSetting,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Inbox_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Inbox::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["IsAutomaticDeletingEnabled"] = p.IsAutomaticDeletingEnabled
	if p.TransactionSetting != nil {
		properties["TransactionSetting"] = p.TransactionSetting.Properties()
	}
	if p.ReceiveMessageScript != nil {
		properties["ReceiveMessageScript"] = p.ReceiveMessageScript.Properties()
	}
	if p.ReadMessageScript != nil {
		properties["ReadMessageScript"] = p.ReadMessageScript.Properties()
	}
	if p.DeleteMessageScript != nil {
		properties["DeleteMessageScript"] = p.DeleteMessageScript.Properties()
	}
	if p.ReceiveNotification != nil {
		properties["ReceiveNotification"] = p.ReceiveNotification.Properties()
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

func (p *Namespace) MasterData(
	globalMessages []GlobalMessage,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		globalMessages,
	).AddDependsOn(
		p,
	)
	return p
}
