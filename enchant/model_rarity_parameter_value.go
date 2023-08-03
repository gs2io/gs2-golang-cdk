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

type RarityParameterValue struct {
	Name          string
	ResourceName  string
	ResourceValue int64
}

type RarityParameterValueOptions struct {
}

func NewRarityParameterValue(
	name string,
	resourceName string,
	resourceValue int64,
	options RarityParameterValueOptions,
) RarityParameterValue {
	data := RarityParameterValue{
		Name:          name,
		ResourceName:  resourceName,
		ResourceValue: resourceValue,
	}
	return data
}

func (p *RarityParameterValue) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["ResourceName"] = p.ResourceName
	properties["ResourceValue"] = p.ResourceValue
	return properties
}
