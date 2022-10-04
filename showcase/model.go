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

type DisplayItemMasterType string

const DisplayItemMasterTypeSalesItem = DisplayItemMasterType("salesItem")
const DisplayItemMasterTypeSalesItemGroup = DisplayItemMasterType("salesItemGroup")

func (p DisplayItemMasterType) Pointer() *DisplayItemMasterType {
	return &p
}

type DisplayItemMaster struct {
	displayItemId      string
	Type               DisplayItemMasterType
	salesItemName      *string
	salesItemGroupName *string
	salesPeriodEventId *string
}

func NewDisplayItemMaster(
	displayItemId string,
	Type DisplayItemMasterType,
	salesItemName *string,
	salesItemGroupName *string,
	salesPeriodEventId *string,
) DisplayItemMaster {
	data := DisplayItemMaster{
		displayItemId:      displayItemId,
		Type:               Type,
		salesItemName:      salesItemName,
		salesItemGroupName: salesItemGroupName,
		salesPeriodEventId: salesPeriodEventId,
	}
	return data
}

func NewSalesItemDisplayItemMaster(
	displayItemId string,
	salesItemName string,
	salesPeriodEventId *string,
) DisplayItemMaster {
	return DisplayItemMaster{
		Type:               DisplayItemMasterTypeSalesItem,
		displayItemId:      displayItemId,
		salesItemName:      &salesItemName,
		salesPeriodEventId: salesPeriodEventId,
	}
}

func NewSalesItemGroupDisplayItemMaster(
	displayItemId string,
	salesItemGroupName string,
	salesPeriodEventId *string,
) DisplayItemMaster {
	return DisplayItemMaster{
		Type:               DisplayItemMasterTypeSalesItemGroup,
		displayItemId:      displayItemId,
		salesItemGroupName: &salesItemGroupName,
		salesPeriodEventId: salesPeriodEventId,
	}
}

func (p *DisplayItemMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["DisplayItemId"] = p.displayItemId
	properties["Type"] = p.Type
	if p.salesItemName != nil {
		properties["SalesItemName"] = p.salesItemName
	}
	if p.salesItemGroupName != nil {
		properties["SalesItemGroupName"] = p.salesItemGroupName
	}
	if p.salesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.salesPeriodEventId
	}
	return properties
}
