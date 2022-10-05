package chat

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
	stack                 *Stack
	name                  string
	description           *string
	allowCreateRoom       bool
	postMessageScript     *ScriptSetting
	createRoomScript      *ScriptSetting
	deleteRoomScript      *ScriptSetting
	subscribeRoomScript   *ScriptSetting
	unsubscribeRoomScript *ScriptSetting
	postNotification      *NotificationSetting
	logSetting            *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	allowCreateRoom bool,
	description *string,
	postMessageScript *ScriptSetting,
	createRoomScript *ScriptSetting,
	deleteRoomScript *ScriptSetting,
	subscribeRoomScript *ScriptSetting,
	unsubscribeRoomScript *ScriptSetting,
	postNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                 stack,
		name:                  name,
		description:           description,
		allowCreateRoom:       allowCreateRoom,
		postMessageScript:     postMessageScript,
		createRoomScript:      createRoomScript,
		deleteRoomScript:      deleteRoomScript,
		subscribeRoomScript:   subscribeRoomScript,
		unsubscribeRoomScript: unsubscribeRoomScript,
		postNotification:      postNotification,
		logSetting:            logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Chat_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Chat::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["AllowCreateRoom"] = p.allowCreateRoom
	if p.postMessageScript != nil {
		properties["PostMessageScript"] = p.postMessageScript.Properties()
	}
	if p.createRoomScript != nil {
		properties["CreateRoomScript"] = p.createRoomScript.Properties()
	}
	if p.deleteRoomScript != nil {
		properties["DeleteRoomScript"] = p.deleteRoomScript.Properties()
	}
	if p.subscribeRoomScript != nil {
		properties["SubscribeRoomScript"] = p.subscribeRoomScript.Properties()
	}
	if p.unsubscribeRoomScript != nil {
		properties["UnsubscribeRoomScript"] = p.unsubscribeRoomScript.Properties()
	}
	if p.postNotification != nil {
		properties["PostNotification"] = p.postNotification.Properties()
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
