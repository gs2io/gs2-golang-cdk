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
	namespaceName string
}

func (p *NamespaceRef) CurrentCampaignMaster() *CurrentCampaignMasterRef {
	return &CurrentCampaignMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) CampaignModel(
	campaignModelName string,
) *CampaignModelRef {
	return &CampaignModelRef{
		namespaceName:     p.namespaceName,
		campaignModelName: campaignModelName,
	}
}

func (p *NamespaceRef) SerialKey(
	serialKeyCode string,
) *SerialKeyRef {
	return &SerialKeyRef{
		namespaceName: p.namespaceName,
		serialKeyCode: serialKeyCode,
	}
}

func (p *NamespaceRef) CampaignModelMaster(
	campaignModelName string,
) *CampaignModelMasterRef {
	return &CampaignModelMasterRef{
		namespaceName:     p.namespaceName,
		campaignModelName: campaignModelName,
	}
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
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type IssueJobRef struct {
	namespaceName     string
	campaignModelName string
	issueJobName      string
}

func (p *IssueJobRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.namespaceName,
			"master",
			"campaign",
			p.campaignModelName,
			"issue",
			"job",
			p.issueJobName,
		},
	).String()
}

func (p *IssueJobRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type SerialKeyRef struct {
	namespaceName string
	serialKeyCode string
}

func (p *SerialKeyRef) Use(
	code string,
) ConsumeAction {
	return UseByUserId(
		p.namespaceName,
		code,
	)
}

func (p *SerialKeyRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.namespaceName,
			"serialKey",
			p.serialKeyCode,
		},
	).String()
}

func (p *SerialKeyRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CampaignModelRef struct {
	namespaceName     string
	campaignModelName string
}

func (p *CampaignModelRef) IssueJob(
	issueJobName string,
) *IssueJobRef {
	return &IssueJobRef{
		namespaceName:     p.namespaceName,
		campaignModelName: p.campaignModelName,
		issueJobName:      issueJobName,
	}
}

func (p *CampaignModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.namespaceName,
			"master",
			"campaign",
			p.campaignModelName,
		},
	).String()
}

func (p *CampaignModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CampaignModelMasterRef struct {
	namespaceName     string
	campaignModelName string
}

func (p *CampaignModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.namespaceName,
			"master",
			"campaign",
			p.campaignModelName,
		},
	).String()
}

func (p *CampaignModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentCampaignMasterRef struct {
	namespaceName string
}

func (p *CurrentCampaignMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"serialKey",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentCampaignMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
