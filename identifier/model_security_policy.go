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

type SecurityPolicy struct {
	CdkResource
	stack       *Stack
	Name        string
	Description *string
	Policy      Policy
}

type SecurityPolicyOptions struct {
	Description *string
}

func NewSecurityPolicy(
	stack *Stack,
	name string,
	policy Policy,
	options SecurityPolicyOptions,
) *SecurityPolicy {
	data := SecurityPolicy{
		stack:       stack,
		Name:        name,
		Policy:      policy,
		Description: options.Description,
	}
	return &data
}

func (p *SecurityPolicy) ResourceName() string {
	return "Identifier_SecurityPolicy_" + p.Name
}

func (p *SecurityPolicy) ResourceType() string {
	return "GS2::Identifier::SecurityPolicy"
}

func (p *SecurityPolicy) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["Policy"] = p.Policy.Properties()
	return properties
}

func (p *SecurityPolicy) Ref() SecurityPolicyRef {
	return SecurityPolicyRef{
		SecurityPolicyName: p.Name,
	}
}

func (p *SecurityPolicy) GetAttrSecurityPolicyId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.SecurityPolicyId",
	)
}
