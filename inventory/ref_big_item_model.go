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

type BigItemModelRef struct {
	NamespaceName string
	InventoryName string
	ItemName      string
}

func (p *BigItemModelRef) AcquireBigItem(
	acquireCount string,
) AcquireAction {
	return AcquireBigItemByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		acquireCount,
	)
}

func (p *BigItemModelRef) SetBigItem(
	count string,
) AcquireAction {
	return SetBigItemByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		count,
	)
}

func (p *BigItemModelRef) ConsumeBigItem(
	consumeCount string,
) ConsumeAction {
	return ConsumeBigItemByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		consumeCount,
	)
}

func (p *BigItemModelRef) VerifyBigItem(
	verifyType string,
	count string,
) ConsumeAction {
	return VerifyBigItemByUserId(
		p.NamespaceName,
		p.InventoryName,
		p.ItemName,
		verifyType,
		count,
	)
}

func (p *BigItemModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"inventory",
			p.NamespaceName,
			"big",
			"model",
			p.InventoryName,
			"item",
			p.ItemName,
		},
	).String()
}

func (p *BigItemModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
