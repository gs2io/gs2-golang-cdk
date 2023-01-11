package formation

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

type FormModel struct {
	Name     string
	Metadata *string
	Slots    []SlotModel
}

type FormModelOptions struct {
	Metadata *string
}

func NewFormModel(
	name string,
	slots []SlotModel,
	options FormModelOptions,
) FormModel {
	data := FormModel{
		Name:     name,
		Slots:    slots,
		Metadata: options.Metadata,
	}
	return data
}

func (p *FormModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.Slots))
		for i, element := range p.Slots {
			values[i] = element.Properties()
		}
		properties["Slots"] = values
	}
	return properties
}
