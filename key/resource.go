package key

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
	stack       *Stack
	name        string
	description *string
	logSetting  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:       stack,
		name:        name,
		description: description,
		logSetting:  logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Key_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Key::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
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

type Key struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
}

func NewKey(
	stack *Stack,
	namespaceName string,
	name string,
	description *string,
) *Key {

	self := Key{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Key) ResourceName() string {
	return "Key_Key_" + p.name
}

func (p *Key) ResourceType() string {
	return "GS2::Key::Key"
}

func (p *Key) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	return properties
}

func (p *Key) Ref(
	namespaceName string,
) KeyRef {
	return KeyRef{
		namespaceName: namespaceName,
		keyName:       p.name,
	}
}

func (p *Key) GetAttrKeyId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.KeyId",
	)
}
