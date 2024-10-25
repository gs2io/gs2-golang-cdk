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

type AttachGuard struct {
	CdkResource
	stack            *Stack
	userName         string
	clientId         string
	guardNamespaceId string
}

type AttachGuardOptions struct {
}

func NewAttachGuard(
	stack *Stack,
	userName string,
	clientId string,
	guardNamespaceId string,
	options AttachGuardOptions,
) *AttachGuard {
	self := AttachGuard{
		userName:         userName,
		clientId:         clientId,
		guardNamespaceId: guardNamespaceId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *AttachGuard) ResourceName() string {
	return "Identifier_AttachGuard_"
}

func (p *AttachGuard) ResourceType() string {
	return "GS2::Identifier::AttachGuard"
}

func (p *AttachGuard) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserName"] = p.userName
	properties["ClientId"] = p.clientId
	properties["GuardId"] = p.guardNamespaceId
	return properties
}
