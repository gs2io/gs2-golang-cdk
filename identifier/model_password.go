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
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type PasswordEnableTwoFactorAuthentication string

const PasswordEnableTwoFactorAuthenticationRfc6238 = PasswordEnableTwoFactorAuthentication("RFC6238")
const PasswordEnableTwoFactorAuthenticationDisable = PasswordEnableTwoFactorAuthentication("Disable")

func (p PasswordEnableTwoFactorAuthentication) Pointer() *PasswordEnableTwoFactorAuthentication {
	return &p
}

type Password struct {
	CdkResource
	stack    *Stack
	UserName string
	Password string
}

type PasswordOptions struct {
}

func NewPassword(
	stack *Stack,
	userName string,
	password string,
	options PasswordOptions,
) *Password {
	data := Password{
		stack:    stack,
		UserName: userName,
		Password: password,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Password) ResourceName() string {
	return "Identifier_Password_"
}

func (p *Password) ResourceType() string {
	return "GS2::Identifier::Password"
}

func (p *Password) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.UserName
	properties["Password"] = p.Password
	return properties
}

func (p *Password) Ref(
	userName string,
) PasswordRef {
	return PasswordRef{
		UserName: userName,
	}
}

func (p *Password) GetAttrPasswordId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.PasswordId",
	)
}
