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

type ItemModelRef struct {
	NamespaceName string
	InventoryName string
	ItemName      string
}

func (p *ItemModelRef) AcquireItemSet(
	acquireCount int64,
	expiresAt int64,
	createNewItemSet bool,
	itemSetName *string,
) AcquireAction {
	return AcquireItemSetByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
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
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		referenceOf,
		itemSetName,
	)
}

func (p *ItemModelRef) ConsumeItemSet(
	consumeCount int64,
	itemSetName *string,
) ConsumeAction {
	return ConsumeItemSetByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		consumeCount,
		itemSetName,
	)
}

func (p *ItemModelRef) VerifyItemSet(
	verifyType string,
	count int64,
	itemSetName *string,
) ConsumeAction {
	return VerifyItemSetByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		verifyType,
		count,
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
			p.NamespaceName,
			"model",
			p.InventoryName,
			"item",
			p.ItemName,
		},
	).String()
}

func (p *ItemModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
