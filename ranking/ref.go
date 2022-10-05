package ranking

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

func (p *NamespaceRef) CurrentRankingMaster() *CurrentRankingMasterRef {
	return &CurrentRankingMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) CategoryModel(
	categoryName string,
) *CategoryModelRef {
	return &CategoryModelRef{
		namespaceName: p.namespaceName,
		categoryName:  categoryName,
	}
}

func (p *NamespaceRef) CategoryModelMaster(
	categoryName string,
) *CategoryModelMasterRef {
	return &CategoryModelMasterRef{
		namespaceName: p.namespaceName,
		categoryName:  categoryName,
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
			"ranking",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CategoryModelRef struct {
	namespaceName string
	categoryName  string
}

func (p *CategoryModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"ranking",
			p.namespaceName,
			"categoryModel",
			p.categoryName,
		},
	).String()
}

func (p *CategoryModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CategoryModelMasterRef struct {
	namespaceName string
	categoryName  string
}

func (p *CategoryModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"ranking",
			p.namespaceName,
			"categoryModelMaster",
			p.categoryName,
		},
	).String()
}

func (p *CategoryModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentRankingMasterRef struct {
	namespaceName string
}

func (p *CurrentRankingMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"ranking",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentRankingMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
