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

type Showcase struct {
	Name               string
	Metadata           *string
	SalesPeriodEventId *string
	DisplayItems       []DisplayItem
}

type ShowcaseOptions struct {
	Metadata           *string
	SalesPeriodEventId *string
}

func NewShowcase(
	name string,
	displayItems []DisplayItem,
	options ShowcaseOptions,
) Showcase {
	data := Showcase{
		Name:               name,
		DisplayItems:       displayItems,
		Metadata:           options.Metadata,
		SalesPeriodEventId: options.SalesPeriodEventId,
	}
	return data
}

func (p *Showcase) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	if p.SalesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.SalesPeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.DisplayItems))
		for i, element := range p.DisplayItems {
			values[i] = element.Properties()
		}
		properties["DisplayItems"] = values
	}
	return properties
}
