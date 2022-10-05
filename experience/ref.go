package experience

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

func (p *NamespaceRef) CurrentExperienceMaster() *CurrentExperienceMasterRef {
	return &CurrentExperienceMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) ExperienceModel(
	experienceName string,
) *ExperienceModelRef {
	return &ExperienceModelRef{
		namespaceName:  p.namespaceName,
		experienceName: experienceName,
	}
}

func (p *NamespaceRef) ThresholdMaster(
	thresholdName string,
) *ThresholdMasterRef {
	return &ThresholdMasterRef{
		namespaceName: p.namespaceName,
		thresholdName: thresholdName,
	}
}

func (p *NamespaceRef) ExperienceModelMaster(
	experienceName string,
) *ExperienceModelMasterRef {
	return &ExperienceModelMasterRef{
		namespaceName:  p.namespaceName,
		experienceName: experienceName,
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
			"experience",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ExperienceModelMasterRef struct {
	namespaceName  string
	experienceName string
}

func (p *ExperienceModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"experience",
			p.namespaceName,
			"model",
			p.experienceName,
		},
	).String()
}

func (p *ExperienceModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ExperienceModelRef struct {
	namespaceName  string
	experienceName string
}

func (p *ExperienceModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"experience",
			p.namespaceName,
			"model",
			p.experienceName,
		},
	).String()
}

func (p *ExperienceModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ThresholdMasterRef struct {
	namespaceName string
	thresholdName string
}

func (p *ThresholdMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"experience",
			p.namespaceName,
			"threshold",
			p.thresholdName,
		},
	).String()
}

func (p *ThresholdMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentExperienceMasterRef struct {
	namespaceName string
}

func (p *CurrentExperienceMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"experience",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentExperienceMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
