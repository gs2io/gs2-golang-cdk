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
	name                       string
	description                *string
	followScript               *ScriptSetting
	unfollowScript             *ScriptSetting
	sendRequestScript          *ScriptSetting
	cancelRequestScript        *ScriptSetting
	acceptRequestScript        *ScriptSetting
	rejectRequestScript        *ScriptSetting
	deleteFriendScript         *ScriptSetting
	updateProfileScript        *ScriptSetting
	followNotification         *NotificationSetting
	receiveRequestNotification *NotificationSetting
	acceptRequestNotification  *NotificationSetting
	logSetting                 *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	followScript *ScriptSetting,
	unfollowScript *ScriptSetting,
	sendRequestScript *ScriptSetting,
	cancelRequestScript *ScriptSetting,
	acceptRequestScript *ScriptSetting,
	rejectRequestScript *ScriptSetting,
	deleteFriendScript *ScriptSetting,
	updateProfileScript *ScriptSetting,
	followNotification *NotificationSetting,
	receiveRequestNotification *NotificationSetting,
	acceptRequestNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                      stack,
		name:                       name,
		description:                description,
		followScript:               followScript,
		unfollowScript:             unfollowScript,
		sendRequestScript:          sendRequestScript,
		cancelRequestScript:        cancelRequestScript,
		acceptRequestScript:        acceptRequestScript,
		rejectRequestScript:        rejectRequestScript,
		deleteFriendScript:         deleteFriendScript,
		updateProfileScript:        updateProfileScript,
		followNotification:         followNotification,
		receiveRequestNotification: receiveRequestNotification,
		acceptRequestNotification:  acceptRequestNotification,
		logSetting:                 logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Friend_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Friend::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.followScript != nil {
		properties["FollowScript"] = p.followScript.Properties()
	}
	if p.unfollowScript != nil {
		properties["UnfollowScript"] = p.unfollowScript.Properties()
	}
	if p.sendRequestScript != nil {
		properties["SendRequestScript"] = p.sendRequestScript.Properties()
	}
	if p.cancelRequestScript != nil {
		properties["CancelRequestScript"] = p.cancelRequestScript.Properties()
	}
	if p.acceptRequestScript != nil {
		properties["AcceptRequestScript"] = p.acceptRequestScript.Properties()
	}
	if p.rejectRequestScript != nil {
		properties["RejectRequestScript"] = p.rejectRequestScript.Properties()
	}
	if p.deleteFriendScript != nil {
		properties["DeleteFriendScript"] = p.deleteFriendScript.Properties()
	}
	if p.updateProfileScript != nil {
		properties["UpdateProfileScript"] = p.updateProfileScript.Properties()
	}
	if p.followNotification != nil {
		properties["FollowNotification"] = p.followNotification.Properties()
	}
	if p.receiveRequestNotification != nil {
		properties["ReceiveRequestNotification"] = p.receiveRequestNotification.Properties()
	}
	if p.acceptRequestNotification != nil {
		properties["AcceptRequestNotification"] = p.acceptRequestNotification.Properties()
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
