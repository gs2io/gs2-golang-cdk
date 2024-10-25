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
	"github.com/gs2io/gs2-golang-cdk/guard"
)

var _ = AcquireAction{}

type Identifier struct {
	CdkResource
	stack    *Stack
	UserName string
}

type IdentifierOptions struct {
}

func NewIdentifier(
	stack *Stack,
	userName string,
	options IdentifierOptions,
) *Identifier {
	data := Identifier{
		stack:    stack,
		UserName: userName,
	}
	return &data
}

func (p *Identifier) ResourceName() string {
	return "Identifier_Identifier_" + p.UserName
}

func (p *Identifier) ResourceType() string {
	return "GS2::Identifier::Identifier"
}

func (p *Identifier) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.UserName
	return properties
}

func (p *Identifier) Ref(
	userName string,
	clientId string,
) IdentifierRef {
	return IdentifierRef{
		UserName: userName,
		ClientId: clientId,
	}
}

func (p *Identifier) Attach(
	guardNamespace *guard.Namespace,
) *Identifier {
	policy := NewAttachGuard(
		p.stack,
		p.UserName,
		p.GetAttrClientId().String(),
		guardNamespace.GetAttrNamespaceId().String(),
		AttachGuardOptions{},
	)
	policy.AddDependsOn(
		p,
	).AddDependsOn(
		guardNamespace,
	)
	return p
}

func (p *Identifier) AttachGrn(
	securityPolicyId string,
) *Identifier {
	policy := NewAttachGuard(
		p.stack,
		p.UserName,
		p.GetAttrClientId().String(),
		securityPolicyId,
		AttachGuardOptions{},
	)
	policy.AddDependsOn(
		p,
	)
	return p
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
