package mission

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

type TargetCounterModelResetType string

const TargetCounterModelResetTypeNotReset = TargetCounterModelResetType("notReset")
const TargetCounterModelResetTypeDaily = TargetCounterModelResetType("daily")
const TargetCounterModelResetTypeWeekly = TargetCounterModelResetType("weekly")
const TargetCounterModelResetTypeMonthly = TargetCounterModelResetType("monthly")

func (p TargetCounterModelResetType) Pointer() *TargetCounterModelResetType {
	return &p
}

type TargetCounterModel struct {
	CounterName string
	ResetType   *TargetCounterModelResetType
	Value       int64
}

type TargetCounterModelOptions struct {
	ResetType *TargetCounterModelResetType
}

func NewTargetCounterModel(
	counterName string,
	value int64,
	options TargetCounterModelOptions,
) TargetCounterModel {
	data := TargetCounterModel{
		CounterName: counterName,
		Value:       value,
		ResetType:   options.ResetType,
	}
	return data
}

func (p *TargetCounterModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CounterName"] = p.CounterName
	if p.ResetType != nil {
		properties["ResetType"] = p.ResetType
	}
	properties["Value"] = p.Value
	return properties
}
