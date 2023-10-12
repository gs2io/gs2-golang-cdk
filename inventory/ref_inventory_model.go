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

type InventoryModelRef struct {
	NamespaceName string
	InventoryName string
}

func (p *InventoryModelRef) ItemModel(
	itemName string,
) *ItemModelRef {
	return &ItemModelRef{
		NamespaceName: p.NamespaceName,
		InventoryName: p.InventoryName,
		ItemName:      itemName,
	}
}

func (p *InventoryModelRef) AddCapacity(
	addCapacityValue int32,
) AcquireAction {
	return AddCapacityByUserId(
		p.NamespaceName,
		p.InventoryName,
		addCapacityValue,
	)
}

func (p *InventoryModelRef) SetCapacity(
	newCapacityValue int32,
) AcquireAction {
	return SetCapacityByUserId(
		p.NamespaceName,
		p.InventoryName,
		newCapacityValue,
	)
}

func (p *InventoryModelRef) VerifyInventoryCurrentMaxCapacity(
	verifyType string,
	currentInventoryMaxCapacity int32,
) ConsumeAction {
	return VerifyInventoryCurrentMaxCapacityByUserId(
		p.NamespaceName,
		p.InventoryName,
		verifyType,
		currentInventoryMaxCapacity,
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
			p.NamespaceName,
			"model",
			p.InventoryName,
		},
	).String()
}

func (p *InventoryModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
