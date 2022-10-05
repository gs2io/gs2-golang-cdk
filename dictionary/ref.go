package dictionary

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

func (p *NamespaceRef) CurrentEntryMaster() *CurrentEntryMasterRef {
	return &CurrentEntryMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) EntryModel(
	entryName string,
) *EntryModelRef {
	return &EntryModelRef{
		namespaceName: p.namespaceName,
		entryName:     entryName,
	}
}

func (p *NamespaceRef) EntryModelMaster(
	entryName string,
) *EntryModelMasterRef {
	return &EntryModelMasterRef{
		namespaceName: p.namespaceName,
		entryName:     entryName,
	}
}

func (p *NamespaceRef) AddEntries(
	entryModelNames *[]string,
) AcquireAction {
	return AddEntriesByUserId(
		p.namespaceName,
		entryModelNames,
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
			"dictionary",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type EntryModelRef struct {
	namespaceName string
	entryName     string
}

func (p *EntryModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"dictionary",
			p.namespaceName,
			"model",
			p.entryName,
		},
	).String()
}

func (p *EntryModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type EntryModelMasterRef struct {
	namespaceName string
	entryName     string
}

func (p *EntryModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"dictionary",
			p.namespaceName,
			"model",
			p.entryName,
		},
	).String()
}

func (p *EntryModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentEntryMasterRef struct {
	namespaceName string
}

func (p *CurrentEntryMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"dictionary",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentEntryMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
