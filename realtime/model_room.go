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

type Room struct {
	CdkResource
	stack         *Stack
	NamespaceName string
	Name          string
	IpAddress     *string
	Port          *int32
	EncryptionKey *string
}

type RoomOptions struct {
	IpAddress     *string
	Port          *int32
	EncryptionKey *string
}

func NewRoom(
	stack *Stack,
	namespaceName string,
	name string,
	options RoomOptions,
) *Room {
	data := Room{
		stack:         stack,
		NamespaceName: namespaceName,
		Name:          name,
		IpAddress:     options.IpAddress,
		Port:          options.Port,
		EncryptionKey: options.EncryptionKey,
	}
	return &data
}

func (p *Room) ResourceName() string {
	return "Realtime_Room_" + p.Name
}

func (p *Room) ResourceType() string {
	return "GS2::Realtime::Room"
}

func (p *Room) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["Name"] = p.Name
	if p.IpAddress != nil {
		properties["IpAddress"] = p.IpAddress
	}
	if p.Port != nil {
		properties["Port"] = p.Port
	}
	if p.EncryptionKey != nil {
		properties["EncryptionKey"] = p.EncryptionKey
	}
	return properties
}

func (p *Room) Ref(
	namespaceName string,
) RoomRef {
	return RoomRef{
		NamespaceName: namespaceName,
		RoomName:      p.Name,
	}
}

func (p *Room) GetAttrRoomId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RoomId",
	)
}
