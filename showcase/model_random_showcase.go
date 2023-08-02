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

type RandomShowcase struct {
	Name                  string
	Metadata              *string
	MaximumNumberOfChoice int32
	DisplayItems          []RandomDisplayItemModel
	BaseTimestamp         int64
	ResetIntervalHours    int32
	SalesPeriodEventId    *string
}

type RandomShowcaseOptions struct {
	Metadata           *string
	SalesPeriodEventId *string
}

func NewRandomShowcase(
	name string,
	maximumNumberOfChoice int32,
	displayItems []RandomDisplayItemModel,
	baseTimestamp int64,
	resetIntervalHours int32,
	options RandomShowcaseOptions,
) RandomShowcase {
	data := RandomShowcase{
		Name:                  name,
		MaximumNumberOfChoice: maximumNumberOfChoice,
		DisplayItems:          displayItems,
		BaseTimestamp:         baseTimestamp,
		ResetIntervalHours:    resetIntervalHours,
		Metadata:              options.Metadata,
		SalesPeriodEventId:    options.SalesPeriodEventId,
	}
	return data
}

func (p *RandomShowcase) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["MaximumNumberOfChoice"] = p.MaximumNumberOfChoice
	{
		values := make([]map[string]interface{}, len(p.DisplayItems))
		for i, element := range p.DisplayItems {
			values[i] = element.Properties()
		}
		properties["DisplayItems"] = values
	}
	properties["BaseTimestamp"] = p.BaseTimestamp
	properties["ResetIntervalHours"] = p.ResetIntervalHours
	if p.SalesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.SalesPeriodEventId
	}
	return properties
}
