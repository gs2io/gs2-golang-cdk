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

type Namespace struct {
	CdkResource
	stack              *Stack
	name               string
	description        *string
	serverType         string
	serverSpec         string
	createNotification *NotificationSetting
	logSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	serverType string,
	serverSpec string,
	description *string,
	createNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:              stack,
		name:               name,
		description:        description,
		serverType:         serverType,
		serverSpec:         serverSpec,
		createNotification: createNotification,
		logSetting:         logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Realtime_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Realtime::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["ServerType"] = p.serverType
	properties["ServerSpec"] = p.serverSpec
	if p.createNotification != nil {
		properties["CreateNotification"] = p.createNotification.Properties()
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

type Room struct {
	CdkResource
	stack         *Stack
	ownerId       string
	namespaceName string
	name          string
	ipAddress     *string
	port          *int32
	encryptionKey *string
}

func NewRoom(
	stack *Stack,
	ownerId string,
	namespaceName string,
	name string,
	ipAddress *string,
	port *int32,
	encryptionKey *string,
) *Room {

	self := Room{
		stack:         stack,
		ownerId:       ownerId,
		namespaceName: namespaceName,
		name:          name,
		ipAddress:     ipAddress,
		port:          port,
		encryptionKey: encryptionKey,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Room) ResourceName() string {
	return "Realtime_Room_" + p.name
}

func (p *Room) ResourceType() string {
	return "GS2::Realtime::Room"
}

func (p *Room) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["OwnerId"] = p.ownerId
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.ipAddress != nil {
		properties["IpAddress"] = p.ipAddress
	}
	if p.port != nil {
		properties["Port"] = p.port
	}
	if p.encryptionKey != nil {
		properties["EncryptionKey"] = p.encryptionKey
	}
	return properties
}

func (p *Room) Ref(
	namespaceName string,
) RoomRef {
	return RoomRef{
		namespaceName: namespaceName,
		roomName:      p.name,
	}
}

func (p *Room) GetAttrRoomId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RoomId",
	)
}
