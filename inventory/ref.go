package inventory

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

func (p *NamespaceRef) CurrentItemModelMaster() *CurrentItemModelMasterRef {
	return &CurrentItemModelMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) InventoryModel(
	inventoryName string,
) *InventoryModelRef {
	return &InventoryModelRef{
		namespaceName: p.namespaceName,
		inventoryName: inventoryName,
	}
}

func (p *NamespaceRef) InventoryModelMaster(
	inventoryName string,
) *InventoryModelMasterRef {
	return &InventoryModelMasterRef{
		namespaceName: p.namespaceName,
		inventoryName: inventoryName,
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
			"inventory",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type InventoryModelMasterRef struct {
	namespaceName string
	inventoryName string
}

func (p *InventoryModelMasterRef) ItemModelMaster(
	itemName string,
) *ItemModelMasterRef {
	return &ItemModelMasterRef{
		namespaceName: p.namespaceName,
		inventoryName: p.inventoryName,
		itemName:      itemName,
	}
}

func (p *InventoryModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.namespaceName,
			"model",
			p.inventoryName,
		},
	).String()
}

func (p *InventoryModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type InventoryModelRef struct {
	namespaceName string
	inventoryName string
}

func (p *InventoryModelRef) ItemModel(
	itemName string,
) *ItemModelRef {
	return &ItemModelRef{
		namespaceName: p.namespaceName,
		inventoryName: p.inventoryName,
		itemName:      itemName,
	}
}

func (p *InventoryModelRef) AddCapacity(
	addCapacityValue int32,
) AcquireAction {
	return AddCapacityByUserId(
		p.namespaceName,
		p.inventoryName,
		addCapacityValue,
	)
}

func (p *InventoryModelRef) SetCapacity(
	newCapacityValue int32,
) AcquireAction {
	return SetCapacityByUserId(
		p.namespaceName,
		p.inventoryName,
		newCapacityValue,
	)
}

func (p *InventoryModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.namespaceName,
			"model",
			p.inventoryName,
		},
	).String()
}

func (p *InventoryModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ItemModelMasterRef struct {
	namespaceName string
	inventoryName string
	itemName      string
}

func (p *ItemModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.namespaceName,
			"model",
			p.inventoryName,
			"item",
			p.itemName,
		},
	).String()
}

func (p *ItemModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ItemModelRef struct {
	namespaceName string
	inventoryName string
	itemName      string
}

func (p *ItemModelRef) AcquireItemSet(
	acquireCount int64,
	expiresAt int64,
	createNewItemSet bool,
	itemSetName *string,
) AcquireAction {
	return AcquireItemSetByUserId(
		p.namespaceName,
		p.inventoryName,
		p.itemName,
		acquireCount,
		expiresAt,
		createNewItemSet,
		itemSetName,
	)
}

func (p *ItemModelRef) AddReferenceOf(
	itemSetName string,
	referenceOf string,
) AcquireAction {
	return AddReferenceOfByUserId(
		p.namespaceName,
		p.inventoryName,
		p.itemName,
		itemSetName,
		referenceOf,
	)
}

func (p *ItemModelRef) ConsumeItemSet(
	consumeCount int64,
	itemSetName *string,
) ConsumeAction {
	return ConsumeItemSetByUserId(
		p.namespaceName,
		p.inventoryName,
		p.itemName,
		consumeCount,
		itemSetName,
	)
}

func (p *ItemModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.namespaceName,
			"model",
			p.inventoryName,
			"item",
			p.itemName,
		},
	).String()
}

func (p *ItemModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentItemModelMasterRef struct {
	namespaceName string
}

func (p *CurrentItemModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentItemModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ReferenceOfRef struct {
	namespaceName string
	inventoryName string
	itemName      string
	itemSetName   string
	referenceOf   string
}

func (p *ReferenceOfRef) DeleteReferenceOf() AcquireAction {
	return DeleteReferenceOfByUserId(
		p.namespaceName,
		p.inventoryName,
		p.itemName,
		p.itemSetName,
		p.referenceOf,
	)
}

func (p *ReferenceOfRef) VerifyReferenceOf(
	verifyType string,
) ConsumeAction {
	return VerifyReferenceOfByUserId(
		p.namespaceName,
		p.inventoryName,
		p.itemName,
		p.itemSetName,
		p.referenceOf,
		verifyType,
	)
}
