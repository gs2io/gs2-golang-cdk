package friend

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
	FollowScript               *ScriptSetting
	UnfollowScript             *ScriptSetting
	SendRequestScript          *ScriptSetting
	CancelRequestScript        *ScriptSetting
	AcceptRequestScript        *ScriptSetting
	RejectRequestScript        *ScriptSetting
	DeleteFriendScript         *ScriptSetting
	UpdateProfileScript        *ScriptSetting
	FollowNotification         *NotificationSetting
	ReceiveRequestNotification NotificationSetting
	AcceptRequestNotification  NotificationSetting
	LogSetting                 *LogSetting
}

type NamespaceOptions struct {
	Description                *string
	FollowScript               *ScriptSetting
	UnfollowScript             *ScriptSetting
	SendRequestScript          *ScriptSetting
	CancelRequestScript        *ScriptSetting
	AcceptRequestScript        *ScriptSetting
	RejectRequestScript        *ScriptSetting
	DeleteFriendScript         *ScriptSetting
	UpdateProfileScript        *ScriptSetting
	FollowNotification         *NotificationSetting
	ReceiveRequestNotification NotificationSetting
	AcceptRequestNotification  NotificationSetting
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
		FollowScript:               options.FollowScript,
		UnfollowScript:             options.UnfollowScript,
		SendRequestScript:          options.SendRequestScript,
		CancelRequestScript:        options.CancelRequestScript,
		AcceptRequestScript:        options.AcceptRequestScript,
		RejectRequestScript:        options.RejectRequestScript,
		DeleteFriendScript:         options.DeleteFriendScript,
		UpdateProfileScript:        options.UpdateProfileScript,
		FollowNotification:         options.FollowNotification,
		ReceiveRequestNotification: options.ReceiveRequestNotification,
		AcceptRequestNotification:  options.AcceptRequestNotification,
		LogSetting:                 options.LogSetting,
	}
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Friend_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Friend::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.FollowScript != nil {
		properties["FollowScript"] = p.FollowScript.Properties()
	}
	if p.UnfollowScript != nil {
		properties["UnfollowScript"] = p.UnfollowScript.Properties()
	}
	if p.SendRequestScript != nil {
		properties["SendRequestScript"] = p.SendRequestScript.Properties()
	}
	if p.CancelRequestScript != nil {
		properties["CancelRequestScript"] = p.CancelRequestScript.Properties()
	}
	if p.AcceptRequestScript != nil {
		properties["AcceptRequestScript"] = p.AcceptRequestScript.Properties()
	}
	if p.RejectRequestScript != nil {
		properties["RejectRequestScript"] = p.RejectRequestScript.Properties()
	}
	if p.DeleteFriendScript != nil {
		properties["DeleteFriendScript"] = p.DeleteFriendScript.Properties()
	}
	if p.UpdateProfileScript != nil {
		properties["UpdateProfileScript"] = p.UpdateProfileScript.Properties()
	}
	if p.FollowNotification != nil {
		properties["FollowNotification"] = p.FollowNotification.Properties()
	}
	properties["ReceiveRequestNotification"] = p.ReceiveRequestNotification.Properties()
	properties["AcceptRequestNotification"] = p.AcceptRequestNotification.Properties()
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
