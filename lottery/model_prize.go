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

type PrizeType string

const PrizeTypeAction = PrizeType("action")
const PrizeTypePrizeTable = PrizeType("prize_table")

func (p PrizeType) Pointer() *PrizeType {
	return &p
}

type Prize struct {
	PrizeId              string
	Type                 PrizeType
	AcquireActions       []AcquireAction
	DrawnLimit           *int32
	LimitFailOverPrizeId *string
	PrizeTableName       *string
	Weight               int32
}

type PrizeOptions struct {
	AcquireActions       []AcquireAction
	DrawnLimit           *int32
	LimitFailOverPrizeId *string
	PrizeTableName       *string
}

func NewPrize(
	prizeId string,
	type_ PrizeType,
	weight int32,
	options PrizeOptions,
) Prize {
	_data := Prize{
		PrizeId:              prizeId,
		Type:                 type_,
		Weight:               weight,
		AcquireActions:       options.AcquireActions,
		DrawnLimit:           options.DrawnLimit,
		LimitFailOverPrizeId: options.LimitFailOverPrizeId,
		PrizeTableName:       options.PrizeTableName,
	}
	return _data
}

type PrizeTypeIsActionOptions struct {
	DrawnLimit *int32
}

func NewPrizeTypeIsAction(
	prizeId string,
	weight int32,
	acquireActions []AcquireAction,
	options PrizeTypeIsActionOptions,
) Prize {
	return NewPrize(
		prizeId,
		PrizeTypeAction,
		weight,
		PrizeOptions{
			AcquireActions: acquireActions,
			DrawnLimit:     options.DrawnLimit,
		},
	)
}

type PrizeTypeIsPrizeTableOptions struct {
	DrawnLimit *int32
}

func NewPrizeTypeIsPrizeTable(
	prizeId string,
	weight int32,
	prizeTableName string,
	options PrizeTypeIsPrizeTableOptions,
) Prize {
	return NewPrize(
		prizeId,
		PrizeTypePrizeTable,
		weight,
		PrizeOptions{
			DrawnLimit:     options.DrawnLimit,
			PrizeTableName: &prizeTableName,
		},
	)
}

func (p *Prize) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["PrizeId"] = p.PrizeId
	properties["Type"] = p.Type
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	if p.DrawnLimit != nil {
		properties["DrawnLimit"] = p.DrawnLimit
	}
	if p.LimitFailOverPrizeId != nil {
		properties["LimitFailOverPrizeId"] = p.LimitFailOverPrizeId
	}
	if p.PrizeTableName != nil {
		properties["PrizeTableName"] = p.PrizeTableName
	}
	properties["Weight"] = p.Weight
	return properties
}
