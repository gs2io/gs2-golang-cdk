package identifier

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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type User struct {
	CdkResource
	stack       *Stack
	Name        string
	Description *string
}

type UserOptions struct {
	Description *string
}

func NewUser(
	stack *Stack,
	name string,
	options UserOptions,
) *User {
	data := User{
		stack:       stack,
		Name:        name,
		Description: options.Description,
	}
	return &data
}

func (p *User) ResourceName() string {
	return "Identifier_User_" + p.Name
}

func (p *User) ResourceType() string {
	return "GS2::Identifier::User"
}

func (p *User) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	return properties
}

func (p *User) Ref() UserRef {
	return UserRef{
		UserName: p.Name,
	}
}

func (p *User) Attach(
	securityPolicy *SecurityPolicy,
) *User {
	policy := NewAttachSecurityPolicy(
		p.stack,
		p.Name,
		securityPolicy.GetAttrSecurityPolicyId().String(),
		AttachSecurityPolicyOptions{},
	)
	policy.AddDependsOn(
		p,
	).AddDependsOn(
		securityPolicy,
	)
	return p
}

func (p *User) Identifier() *Identifier {
	identifier := NewIdentifier(
		p.stack,
		p.Name,
		IdentifierOptions{},
	)
	identifier.AddDependsOn(
		p,
	)
	return identifier
}

func (p *User) GetAttrUserId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.UserId",
	)
}
