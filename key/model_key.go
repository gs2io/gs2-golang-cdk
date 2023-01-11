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

type Key struct {
	CdkResource
	stack         *Stack
	NamespaceName string
	Name          string
	Description   *string
}

type KeyOptions struct {
	Description *string
}

func NewKey(
	stack *Stack,
	namespaceName string,
	name string,
	options KeyOptions,
) *Key {
	data := Key{
		stack:         stack,
		NamespaceName: namespaceName,
		Name:          name,
		Description:   options.Description,
	}
	return &data
}

func (p *Key) ResourceName() string {
	return "Key_Key_" + p.Name
}

func (p *Key) ResourceType() string {
	return "GS2::Key::Key"
}

func (p *Key) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	return properties
}

func (p *Key) Ref(
	namespaceName string,
) KeyRef {
	return KeyRef{
		NamespaceName: namespaceName,
		KeyName:       p.Name,
	}
}

func (p *Key) GetAttrKeyId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.KeyId",
	)
}
