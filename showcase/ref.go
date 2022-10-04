package showcase

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

func (p *NamespaceRef) CurrentShowcaseMaster() *CurrentShowcaseMasterRef {
	return &CurrentShowcaseMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) DisplayItem() *DisplayItemRef {
	return &DisplayItemRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) SalesItemMaster(
	salesItemName string,
) *SalesItemMasterRef {
	return &SalesItemMasterRef{
		namespaceName: p.namespaceName,
		salesItemName: salesItemName,
	}
}

func (p *NamespaceRef) SalesItemGroupMaster(
	salesItemGroupName string,
) *SalesItemGroupMasterRef {
	return &SalesItemGroupMasterRef{
		namespaceName:      p.namespaceName,
		salesItemGroupName: salesItemGroupName,
	}
}

func (p *NamespaceRef) ShowcaseMaster(
	showcaseName string,
) *ShowcaseMasterRef {
	return &ShowcaseMasterRef{
		namespaceName: p.namespaceName,
		showcaseName:  showcaseName,
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
			"showcase",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type SalesItemMasterRef struct {
	namespaceName string
	salesItemName string
}

func (p *SalesItemMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"showcase",
			p.namespaceName,
			"salesItem",
			p.salesItemName,
		},
	).String()
}

func (p *SalesItemMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type SalesItemGroupMasterRef struct {
	namespaceName      string
	salesItemGroupName string
}

func (p *SalesItemGroupMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"showcase",
			p.namespaceName,
			"salesItemGroup",
			p.salesItemGroupName,
		},
	).String()
}

func (p *SalesItemGroupMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ShowcaseMasterRef struct {
	namespaceName string
	showcaseName  string
}

func (p *ShowcaseMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"showcase",
			p.namespaceName,
			"showcase",
			p.showcaseName,
		},
	).String()
}

func (p *ShowcaseMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentShowcaseMasterRef struct {
	namespaceName string
}

func (p *CurrentShowcaseMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"showcase",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentShowcaseMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ShowcaseRef struct {
	namespaceName string
	showcaseName  string
}

func (p *ShowcaseRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"showcase",
			p.namespaceName,
			"showcase",
			p.showcaseName,
		},
	).String()
}

func (p *ShowcaseRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type DisplayItemRef struct {
	namespaceName string
}
