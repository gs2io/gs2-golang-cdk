package deploy

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

type Stack struct {
	CdkResource
	stack       *Stack
	name        string
	description *string
	template    string
}

func NewStack(
	stack *Stack,
	name string,
	template string,
	description *string,
) *Stack {

	self := Stack{
		stack:       stack,
		name:        name,
		description: description,
		template:    template,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Stack) ResourceName() string {
	return "Deploy_Stack_" + p.name
}

func (p *Stack) ResourceType() string {
	return "GS2::Deploy::Stack"
}

func (p *Stack) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["Template"] = p.template
	return properties
}

func (p *Stack) Ref() StackRef {
	return StackRef{
		stackName: p.name,
	}
}

func (p *Stack) GetAttrStackId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.StackId",
	)
}
