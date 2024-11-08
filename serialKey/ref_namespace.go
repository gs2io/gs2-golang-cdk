package serialKey

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

type NamespaceRef struct {
	NamespaceName string
}

func (p *NamespaceRef) CampaignModel(
	campaignModelName string,
) *CampaignModelRef {
	return &CampaignModelRef{
		NamespaceName:     p.NamespaceName,
		CampaignModelName: campaignModelName,
	}
}

func (p *NamespaceRef) RevertUse(
	code string,
) AcquireAction {
	return RevertUseByUserId(
		p.NamespaceName,
		code,
	)
}

func (p *NamespaceRef) VerifyCode(
	code string,
	verifyType string,
) VerifyAction {
	return VerifyCodeByUserId(
		p.NamespaceName,
		code,
		verifyType,
	)
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.NamespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
