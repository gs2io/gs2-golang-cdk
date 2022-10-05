package formation

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

func (p *NamespaceRef) CurrentFormMaster() *CurrentFormMasterRef {
	return &CurrentFormMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) FormModel(
	formModelName string,
) *FormModelRef {
	return &FormModelRef{
		namespaceName: p.namespaceName,
		formModelName: formModelName,
	}
}

func (p *NamespaceRef) MoldModel(
	moldName string,
) *MoldModelRef {
	return &MoldModelRef{
		namespaceName: p.namespaceName,
		moldName:      moldName,
	}
}

func (p *NamespaceRef) FormModelMaster(
	formModelName string,
) *FormModelMasterRef {
	return &FormModelMasterRef{
		namespaceName: p.namespaceName,
		formModelName: formModelName,
	}
}

func (p *NamespaceRef) MoldModelMaster(
	moldName string,
) *MoldModelMasterRef {
	return &MoldModelMasterRef{
		namespaceName: p.namespaceName,
		moldName:      moldName,
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
			"formation",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type FormModelRef struct {
	namespaceName string
	formModelName string
}

func (p *FormModelRef) AcquireActionsToFormProperties(
	moldName string,
	index int32,
	acquireAction AcquireAction,
	config *[]AcquireActionConfig,
) AcquireAction {
	return AcquireActionsToFormProperties(
		p.namespaceName,
		moldName,
		index,
		acquireAction,
		config,
	)
}

func (p *FormModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.namespaceName,
			"model",
			"form",
			p.formModelName,
		},
	).String()
}

func (p *FormModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type FormModelMasterRef struct {
	namespaceName string
	formModelName string
}

func (p *FormModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.namespaceName,
			"model",
			"form",
			p.formModelName,
		},
	).String()
}

func (p *FormModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MoldModelRef struct {
	namespaceName string
	moldName      string
}

func (p *MoldModelRef) AddMoldCapacity(
	capacity int32,
) AcquireAction {
	return AddMoldCapacityByUserId(
		p.namespaceName,
		p.moldName,
		capacity,
	)
}

func (p *MoldModelRef) SetMoldCapacity(
	capacity int32,
) AcquireAction {
	return SetMoldCapacityByUserId(
		p.namespaceName,
		p.moldName,
		capacity,
	)
}

func (p *MoldModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.namespaceName,
			"model",
			"mold",
			p.moldName,
		},
	).String()
}

func (p *MoldModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MoldModelMasterRef struct {
	namespaceName string
	moldName      string
}

func (p *MoldModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.namespaceName,
			"model",
			"mold",
			p.moldName,
		},
	).String()
}

func (p *MoldModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentFormMasterRef struct {
	namespaceName string
}

func (p *CurrentFormMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentFormMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type FormRef struct {
	namespaceName string
	moldName      string
	index         string
}

func (p *FormRef) AcquireActionsToFormProperties(
	acquireAction AcquireAction,
	config *[]AcquireActionConfig,
) AcquireAction {
	return AcquireActionsToFormProperties(
		p.namespaceName,
		p.moldName,
		p.index,
		acquireAction,
		config,
	)
}
