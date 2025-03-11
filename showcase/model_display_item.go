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

type DisplayItemType string

const DisplayItemTypeSalesItem = DisplayItemType("salesItem")
const DisplayItemTypeSalesItemGroup = DisplayItemType("salesItemGroup")

func (p DisplayItemType) Pointer() *DisplayItemType {
	return &p
}

type DisplayItem struct {
	DisplayItemId      string
	Type               DisplayItemType
	SalesItem          *SalesItem
	SalesItemGroup     *SalesItemGroup
	SalesPeriodEventId *string
}

type DisplayItemOptions struct {
	SalesItem          *SalesItem
	SalesItemGroup     *SalesItemGroup
	SalesPeriodEventId *string
}

func NewDisplayItem(
	displayItemId string,
	type_ DisplayItemType,
	options DisplayItemOptions,
) DisplayItem {
	_data := DisplayItem{
		DisplayItemId:      displayItemId,
		Type:               type_,
		SalesItem:          options.SalesItem,
		SalesItemGroup:     options.SalesItemGroup,
		SalesPeriodEventId: options.SalesPeriodEventId,
	}
	return _data
}

type DisplayItemTypeIsSalesItemOptions struct {
	SalesPeriodEventId *string
}

func NewDisplayItemTypeIsSalesItem(
	displayItemId string,
	salesItem SalesItem,
	options DisplayItemTypeIsSalesItemOptions,
) DisplayItem {
	return NewDisplayItem(
		displayItemId,
		DisplayItemTypeSalesItem,
		DisplayItemOptions{
			SalesItem:          &salesItem,
			SalesPeriodEventId: options.SalesPeriodEventId,
		},
	)
}

type DisplayItemTypeIsSalesItemGroupOptions struct {
	SalesPeriodEventId *string
}

func NewDisplayItemTypeIsSalesItemGroup(
	displayItemId string,
	salesItemGroup SalesItemGroup,
	options DisplayItemTypeIsSalesItemGroupOptions,
) DisplayItem {
	return NewDisplayItem(
		displayItemId,
		DisplayItemTypeSalesItemGroup,
		DisplayItemOptions{
			SalesItemGroup:     &salesItemGroup,
			SalesPeriodEventId: options.SalesPeriodEventId,
		},
	)
}

func (p *DisplayItem) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["DisplayItemId"] = p.DisplayItemId
	properties["Type"] = p.Type
	if p.SalesItem != nil {
		properties["SalesItem"] = p.SalesItem.Properties()
	}
	if p.SalesItemGroup != nil {
		properties["SalesItemGroup"] = p.SalesItemGroup.Properties()
	}
	if p.SalesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.SalesPeriodEventId
	}
	return properties
}
