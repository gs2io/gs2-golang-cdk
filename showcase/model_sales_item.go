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

type SalesItem struct {
	Name           string
	Metadata       *string
	VerifyActions  []VerifyAction
	ConsumeActions []ConsumeAction
	AcquireActions []AcquireAction
}

type SalesItemOptions struct {
	Metadata       *string
	VerifyActions  []VerifyAction
	ConsumeActions []ConsumeAction
}

func NewSalesItem(
	name string,
	acquireActions []AcquireAction,
	options SalesItemOptions,
) SalesItem {
	data := SalesItem{
		Name:           name,
		AcquireActions: acquireActions,
		Metadata:       options.Metadata,
		VerifyActions:  options.VerifyActions,
		ConsumeActions: options.ConsumeActions,
	}
	return data
}

func (p *SalesItem) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.VerifyActions))
		for i, element := range p.VerifyActions {
			values[i] = element.Properties()
		}
		properties["VerifyActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ConsumeActions))
		for i, element := range p.ConsumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}
