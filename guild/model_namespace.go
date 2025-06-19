package guild

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
	ChangeNotification         *NotificationSetting
	JoinNotification           *NotificationSetting
	LeaveNotification          *NotificationSetting
	ChangeMemberNotification   *NotificationSetting
	ReceiveRequestNotification *NotificationSetting
	RemoveRequestNotification  *NotificationSetting
	CreateGuildScript          *ScriptSetting
	UpdateGuildScript          *ScriptSetting
	JoinGuildScript            *ScriptSetting
	LeaveGuildScript           *ScriptSetting
	ChangeRoleScript           *ScriptSetting
	DeleteGuildScript          *ScriptSetting
	LogSetting                 *LogSetting
}

type NamespaceOptions struct {
	Description                *string
	ChangeNotification         *NotificationSetting
	JoinNotification           *NotificationSetting
	LeaveNotification          *NotificationSetting
	ChangeMemberNotification   *NotificationSetting
	ReceiveRequestNotification *NotificationSetting
	RemoveRequestNotification  *NotificationSetting
	CreateGuildScript          *ScriptSetting
	UpdateGuildScript          *ScriptSetting
	JoinGuildScript            *ScriptSetting
	LeaveGuildScript           *ScriptSetting
	ChangeRoleScript           *ScriptSetting
	DeleteGuildScript          *ScriptSetting
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
		ChangeNotification:         options.ChangeNotification,
		JoinNotification:           options.JoinNotification,
		LeaveNotification:          options.LeaveNotification,
		ChangeMemberNotification:   options.ChangeMemberNotification,
		ReceiveRequestNotification: options.ReceiveRequestNotification,
		RemoveRequestNotification:  options.RemoveRequestNotification,
		CreateGuildScript:          options.CreateGuildScript,
		UpdateGuildScript:          options.UpdateGuildScript,
		JoinGuildScript:            options.JoinGuildScript,
		LeaveGuildScript:           options.LeaveGuildScript,
		ChangeRoleScript:           options.ChangeRoleScript,
		DeleteGuildScript:          options.DeleteGuildScript,
		LogSetting:                 options.LogSetting,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Guild_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Guild::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.ChangeNotification != nil {
		properties["ChangeNotification"] = p.ChangeNotification.Properties()
	}
	if p.JoinNotification != nil {
		properties["JoinNotification"] = p.JoinNotification.Properties()
	}
	if p.LeaveNotification != nil {
		properties["LeaveNotification"] = p.LeaveNotification.Properties()
	}
	if p.ChangeMemberNotification != nil {
		properties["ChangeMemberNotification"] = p.ChangeMemberNotification.Properties()
	}
	if p.ReceiveRequestNotification != nil {
		properties["ReceiveRequestNotification"] = p.ReceiveRequestNotification.Properties()
	}
	if p.RemoveRequestNotification != nil {
		properties["RemoveRequestNotification"] = p.RemoveRequestNotification.Properties()
	}
	if p.CreateGuildScript != nil {
		properties["CreateGuildScript"] = p.CreateGuildScript.Properties()
	}
	if p.UpdateGuildScript != nil {
		properties["UpdateGuildScript"] = p.UpdateGuildScript.Properties()
	}
	if p.JoinGuildScript != nil {
		properties["JoinGuildScript"] = p.JoinGuildScript.Properties()
	}
	if p.LeaveGuildScript != nil {
		properties["LeaveGuildScript"] = p.LeaveGuildScript.Properties()
	}
	if p.ChangeRoleScript != nil {
		properties["ChangeRoleScript"] = p.ChangeRoleScript.Properties()
	}
	if p.DeleteGuildScript != nil {
		properties["DeleteGuildScript"] = p.DeleteGuildScript.Properties()
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
	guildModels []GuildModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.Name,
		guildModels,
	).AddDependsOn(
		p,
	)
	return p
}
