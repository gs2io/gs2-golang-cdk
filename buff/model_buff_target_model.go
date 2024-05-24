package buff

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

type BuffTargetModelTargetModelName string

func (p BuffTargetModelTargetModelName) Pointer() *BuffTargetModelTargetModelName {
	return &p
}

type BuffTargetModel struct {
	TargetModelName string
	TargetFieldName string
	ConditionGrns   []BuffTargetGrn
	Rate            float32
}

type BuffTargetModelOptions struct {
}

func NewBuffTargetModel(
	targetModelName string,
	targetFieldName string,
	conditionGrns []BuffTargetGrn,
	rate float32,
	options BuffTargetModelOptions,
) BuffTargetModel {
	data := BuffTargetModel{
		TargetModelName: targetModelName,
		TargetFieldName: targetFieldName,
		ConditionGrns:   conditionGrns,
		Rate:            rate,
	}
	return data
}

func (p *BuffTargetModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TargetModelName"] = p.TargetModelName
	properties["TargetFieldName"] = p.TargetFieldName
	{
		values := make([]map[string]interface{}, len(p.ConditionGrns))
		for i, element := range p.ConditionGrns {
			values[i] = element.Properties()
		}
		properties["ConditionGrns"] = values
	}
	properties["Rate"] = p.Rate
	return properties
}
