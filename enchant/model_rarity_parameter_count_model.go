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

type RarityParameterCountModel struct {
	Count  int32
	Weight int32
}

type RarityParameterCountModelOptions struct {
}

func NewRarityParameterCountModel(
	count int32,
	weight int32,
	options RarityParameterCountModelOptions,
) RarityParameterCountModel {
	_data := RarityParameterCountModel{
		Count:  count,
		Weight: weight,
	}
	return _data
}

func (p *RarityParameterCountModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Count"] = p.Count
	properties["Weight"] = p.Weight
	return properties
}
