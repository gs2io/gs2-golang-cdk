package log

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

type AggregationConfigType string

const AggregationConfigTypeCount = AggregationConfigType("count")
const AggregationConfigTypeUnique = AggregationConfigType("unique")
const AggregationConfigTypeSum = AggregationConfigType("sum")
const AggregationConfigTypeAvg = AggregationConfigType("avg")
const AggregationConfigTypeMax = AggregationConfigType("max")
const AggregationConfigTypeMin = AggregationConfigType("min")
const AggregationConfigTypeP90 = AggregationConfigType("p90")
const AggregationConfigTypeP95 = AggregationConfigType("p95")
const AggregationConfigTypeP99 = AggregationConfigType("p99")

func (p AggregationConfigType) Pointer() *AggregationConfigType {
	return &p
}

type AggregationConfig struct {
	Type  *AggregationConfigType
	Field *string
}

type AggregationConfigOptions struct {
	Type  *AggregationConfigType
	Field *string
}

func NewAggregationConfig(
	options AggregationConfigOptions,
) AggregationConfig {
	_data := AggregationConfig{
		Type:  options.Type,
		Field: options.Field,
	}
	return _data
}

func (p *AggregationConfig) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.Type != nil {
		properties["Type"] = p.Type
	}
	if p.Field != nil {
		properties["Field"] = p.Field
	}
	return properties
}
