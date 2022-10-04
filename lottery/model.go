package lottery

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

type Probability struct {
	prize DrawnPrize
	rate  float32
}

func NewProbability(
	prize DrawnPrize,
	rate float32,
) Probability {
	data := Probability{
		prize: prize,
		rate:  rate,
	}
	return data
}

func (p *Probability) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Prize"] = p.prize.Properties()
	properties["Rate"] = p.rate
	return properties
}

type PrizeType string

const PrizeTypeAction = PrizeType("action")
const PrizeTypePrizeTable = PrizeType("prize_table")

func (p PrizeType) Pointer() *PrizeType {
	return &p
}

type Prize struct {
	prizeId              string
	Type                 PrizeType
	acquireActions       []AcquireAction
	drawnLimit           *int32
	limitFailOverPrizeId *string
	prizeTableName       *string
	weight               int32
}

func NewPrize(
	prizeId string,
	Type PrizeType,
	acquireActions []AcquireAction,
	drawnLimit *int32,
	limitFailOverPrizeId *string,
	prizeTableName *string,
	weight int32,
) Prize {
	data := Prize{
		prizeId:              prizeId,
		Type:                 Type,
		acquireActions:       acquireActions,
		drawnLimit:           drawnLimit,
		limitFailOverPrizeId: limitFailOverPrizeId,
		prizeTableName:       prizeTableName,
		weight:               weight,
	}
	return data
}

func NewActionPrize(
	prizeId string,
	acquireActions []AcquireAction,
	weight int32,
	drawnLimit *int32,
	limitFailOverPrizeId *string,
) Prize {
	return Prize{
		Type:                 PrizeTypeAction,
		prizeId:              prizeId,
		acquireActions:       acquireActions,
		drawnLimit:           drawnLimit,
		limitFailOverPrizeId: limitFailOverPrizeId,
		weight:               weight,
	}
}

func NewPrizeTablePrize(
	prizeId string,
	acquireActions []AcquireAction,
	prizeTableName string,
	weight int32,
	drawnLimit *int32,
) Prize {
	return Prize{
		Type:           PrizeTypePrizeTable,
		prizeId:        prizeId,
		drawnLimit:     drawnLimit,
		prizeTableName: &prizeTableName,
		weight:         weight,
	}
}

func (p *Prize) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["PrizeId"] = p.prizeId
	properties["Type"] = p.Type
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	if p.drawnLimit != nil {
		properties["DrawnLimit"] = p.drawnLimit
	}
	if p.limitFailOverPrizeId != nil {
		properties["LimitFailOverPrizeId"] = p.limitFailOverPrizeId
	}
	if p.prizeTableName != nil {
		properties["PrizeTableName"] = p.prizeTableName
	}
	properties["Weight"] = p.weight
	return properties
}

type DrawnPrize struct {
	prizeId        string
	acquireActions []AcquireAction
}

func NewDrawnPrize(
	prizeId string,
	acquireActions []AcquireAction,
) DrawnPrize {
	data := DrawnPrize{
		prizeId:        prizeId,
		acquireActions: acquireActions,
	}
	return data
}

func (p *DrawnPrize) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["PrizeId"] = p.prizeId
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}

type BoxItem struct {
	acquireActions []AcquireAction
	remaining      int32
	initial        int32
}

func NewBoxItem(
	acquireActions []AcquireAction,
	remaining int32,
	initial int32,
) BoxItem {
	data := BoxItem{
		acquireActions: acquireActions,
		remaining:      remaining,
		initial:        initial,
	}
	return data
}

func (p *BoxItem) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	properties["Remaining"] = p.remaining
	properties["Initial"] = p.initial
	return properties
}

type BoxItems struct {
	boxId          string
	prizeTableName string
	userId         string
	items          []BoxItem
}

func NewBoxItems(
	boxId string,
	prizeTableName string,
	userId string,
	items []BoxItem,
) BoxItems {
	data := BoxItems{
		boxId:          boxId,
		prizeTableName: prizeTableName,
		userId:         userId,
		items:          items,
	}
	return data
}

func (p *BoxItems) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["BoxId"] = p.boxId
	properties["PrizeTableName"] = p.prizeTableName
	properties["UserId"] = p.userId
	{
		values := make([]map[string]interface{}, len(p.items))
		for i, element := range p.items {
			values[i] = element.Properties()
		}
		properties["Items"] = values
	}
	return properties
}
