package distributor

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

func (p *NamespaceRef) CurrentDistributorMaster() *CurrentDistributorMasterRef {
	return &CurrentDistributorMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) DistributorModel(
	distributorName string,
) *DistributorModelRef {
	return &DistributorModelRef{
		namespaceName:   p.namespaceName,
		distributorName: distributorName,
	}
}

func (p *NamespaceRef) DistributorModelMaster(
	distributorName string,
) *DistributorModelMasterRef {
	return &DistributorModelMasterRef{
		namespaceName:   p.namespaceName,
		distributorName: distributorName,
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
			"distributor",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type DistributorModelMasterRef struct {
	namespaceName   string
	distributorName string
}

func (p *DistributorModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"distributor",
			p.namespaceName,
			"model",
			p.distributorName,
		},
	).String()
}

func (p *DistributorModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type DistributorModelRef struct {
	namespaceName   string
	distributorName string
}

func (p *DistributorModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"distributor",
			p.namespaceName,
			"model",
			p.distributorName,
		},
	).String()
}

func (p *DistributorModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentDistributorMasterRef struct {
	namespaceName string
}

func (p *CurrentDistributorMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"distributor",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentDistributorMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
