package enchant

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

type RarityParameterModel struct {
	Name                  string
	Metadata              *string
	MaximumParameterCount int32
	ParameterCounts       []RarityParameterCountModel
	Parameters            []RarityParameterValueModel
}

type RarityParameterModelOptions struct {
	Metadata *string
}

func NewRarityParameterModel(
	name string,
	maximumParameterCount int32,
	parameterCounts []RarityParameterCountModel,
	parameters []RarityParameterValueModel,
	options RarityParameterModelOptions,
) RarityParameterModel {
	_data := RarityParameterModel{
		Name:                  name,
		MaximumParameterCount: maximumParameterCount,
		ParameterCounts:       parameterCounts,
		Parameters:            parameters,
		Metadata:              options.Metadata,
	}
	return _data
}

func (p *RarityParameterModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["MaximumParameterCount"] = p.MaximumParameterCount
	{
		values := make([]map[string]interface{}, len(p.ParameterCounts))
		for i, element := range p.ParameterCounts {
			values[i] = element.Properties()
		}
		properties["ParameterCounts"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.Parameters))
		for i, element := range p.Parameters {
			values[i] = element.Properties()
		}
		properties["Parameters"] = values
	}
	return properties
}
