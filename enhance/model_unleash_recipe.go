package enhance

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

type UnleashRecipe struct {
	Name            string
	Metadata        *string
	TargetGroupKeys []string
	Materials       []UnleashMaterial
}

type UnleashRecipeOptions struct {
	Metadata        *string
	TargetGroupKeys []string
}

func NewUnleashRecipe(
	name string,
	materials []UnleashMaterial,
	options UnleashRecipeOptions,
) UnleashRecipe {
	_data := UnleashRecipe{
		Name:            name,
		Materials:       materials,
		Metadata:        options.Metadata,
		TargetGroupKeys: options.TargetGroupKeys,
	}
	return _data
}

func (p *UnleashRecipe) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["TargetGroupKeys"] = p.TargetGroupKeys
	{
		values := make([]map[string]interface{}, len(p.Materials))
		for i, element := range p.Materials {
			values[i] = element.Properties()
		}
		properties["Materials"] = values
	}
	return properties
}
