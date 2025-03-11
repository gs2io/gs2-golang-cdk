package ranking2

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

type SubscribeRankingModelOrderDirection string

const SubscribeRankingModelOrderDirectionAsc = SubscribeRankingModelOrderDirection("asc")
const SubscribeRankingModelOrderDirectionDesc = SubscribeRankingModelOrderDirection("desc")

func (p SubscribeRankingModelOrderDirection) Pointer() *SubscribeRankingModelOrderDirection {
	return &p
}

type SubscribeRankingModel struct {
	Name                string
	Metadata            *string
	MinimumValue        *int64
	MaximumValue        *int64
	Sum                 bool
	OrderDirection      SubscribeRankingModelOrderDirection
	EntryPeriodEventId  *string
	AccessPeriodEventId *string
}

type SubscribeRankingModelOptions struct {
	Metadata            *string
	MinimumValue        *int64
	MaximumValue        *int64
	EntryPeriodEventId  *string
	AccessPeriodEventId *string
}

func NewSubscribeRankingModel(
	name string,
	sum bool,
	orderDirection SubscribeRankingModelOrderDirection,
	options SubscribeRankingModelOptions,
) SubscribeRankingModel {
	_data := SubscribeRankingModel{
		Name:                name,
		Sum:                 sum,
		OrderDirection:      orderDirection,
		Metadata:            options.Metadata,
		MinimumValue:        options.MinimumValue,
		MaximumValue:        options.MaximumValue,
		EntryPeriodEventId:  options.EntryPeriodEventId,
		AccessPeriodEventId: options.AccessPeriodEventId,
	}
	return _data
}

func (p *SubscribeRankingModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	if p.MinimumValue != nil {
		properties["MinimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		properties["MaximumValue"] = p.MaximumValue
	}
	properties["Sum"] = p.Sum
	properties["OrderDirection"] = p.OrderDirection
	if p.EntryPeriodEventId != nil {
		properties["EntryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.AccessPeriodEventId != nil {
		properties["AccessPeriodEventId"] = p.AccessPeriodEventId
	}
	return properties
}
