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

type AttachSecurityPolicy struct {
	CdkResource
	stack            *Stack
	userName         string
	securityPolicyId string
}

type AttachSecurityPolicyOptions struct {
}

func NewAttachSecurityPolicy(
	stack *Stack,
	userName string,
	securityPolicyId string,
	options AttachSecurityPolicyOptions,
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
