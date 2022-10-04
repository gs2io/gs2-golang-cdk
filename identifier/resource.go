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
	name        string
	description *string
}

func NewUser(
	stack *Stack,
	name string,
	description *string,
) *User {

	self := User{
		stack:       stack,
		name:        name,
		description: description,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *User) ResourceName() string {
	return "Identifier_User_" + p.name
}

func (p *User) ResourceType() string {
	return "GS2::Identifier::User"
}

func (p *User) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	return properties
}

func (p *User) GetAttrUserId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.UserId",
	)
}

func (p *User) Attach(
	securityPolicy *SecurityPolicy,
) *User {
	policy := NewAttachSecurityPolicy(
		p.stack,
		p.name,
		securityPolicy.GetAttrSecurityPolicyId().String(),
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
		p.name,
	)
	identifier.AddDependsOn(
		p,
	)
	return identifier
}

type SecurityPolicy struct {
	CdkResource
	stack       *Stack
	name        string
	description *string
	policy      Policy
}

func NewSecurityPolicy(
	stack *Stack,
	name string,
	policy Policy,
	description *string,
) *SecurityPolicy {

	self := SecurityPolicy{
		stack:       stack,
		name:        name,
		description: description,
		policy:      policy,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *SecurityPolicy) ResourceName() string {
	return "Identifier_SecurityPolicy_" + p.name
}

func (p *SecurityPolicy) ResourceType() string {
	return "GS2::Identifier::SecurityPolicy"
}

func (p *SecurityPolicy) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["Policy"] = p.policy.Properties()
	return properties
}

func (p *SecurityPolicy) Ref() SecurityPolicyRef {
	return SecurityPolicyRef{
		securityPolicyName: p.name,
	}
}

func (p *SecurityPolicy) GetAttrSecurityPolicyId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.SecurityPolicyId",
	)
}

type AttachSecurityPolicy struct {
	CdkResource
	stack            *Stack
	userName         string
	securityPolicyId string
}

func NewAttachSecurityPolicy(
	stack *Stack,
	userName string,
	securityPolicyId string,
) *AttachSecurityPolicy {

	self := AttachSecurityPolicy{
		userName:         userName,
		securityPolicyId: securityPolicyId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *AttachSecurityPolicy) ResourceName() string {
	return "Identifier_Identifier_"
}

func (p *AttachSecurityPolicy) ResourceType() string {
	return "GS2::Identifier::Identifier"
}

func (p *AttachSecurityPolicy) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.userName
	properties["SecurityPolicyId"] = p.securityPolicyId
	return properties
}

type Identifier struct {
	CdkResource
	stack    *Stack
	userName string
}

func NewIdentifier(
	stack *Stack,
	userName string,
) *Identifier {

	self := Identifier{
		stack:    stack,
		userName: userName,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Identifier) ResourceName() string {
	return "Identifier_Identifier_"
}

func (p *Identifier) ResourceType() string {
	return "GS2::Identifier::Identifier"
}

func (p *Identifier) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.userName
	return properties
}

func (p *Identifier) Ref(
	userName string,
) IdentifierRef {
	return IdentifierRef{
		userName: userName,
	}
}

func (p *Identifier) GetAttrClientId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ClientId",
	)
}

func (p *Identifier) GetAttrClientSecret() GetAttr {
	return NewGetAttrByResource(
		p,
		"ClientSecret",
	)
}

type Password struct {
	CdkResource
	stack    *Stack
	userName string
	password string
}

func NewPassword(
	stack *Stack,
	userName string,
	password string,
) *Password {

	self := Password{
		stack:    stack,
		userName: userName,
		password: password,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Password) ResourceName() string {
	return "Identifier_Password_" + p.userName
}

func (p *Password) ResourceType() string {
	return "GS2::Identifier::Password"
}

func (p *Password) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.userName
	properties["Password"] = p.password
	return properties
}
