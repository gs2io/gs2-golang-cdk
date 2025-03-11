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

type TargetCounterModelScopeType string

const TargetCounterModelScopeTypeResetTiming = TargetCounterModelScopeType("resetTiming")
const TargetCounterModelScopeTypeVerifyAction = TargetCounterModelScopeType("verifyAction")

func (p TargetCounterModelScopeType) Pointer() *TargetCounterModelScopeType {
	return &p
}

type TargetCounterModelResetType string

const TargetCounterModelResetTypeNotReset = TargetCounterModelResetType("notReset")
const TargetCounterModelResetTypeDaily = TargetCounterModelResetType("daily")
const TargetCounterModelResetTypeWeekly = TargetCounterModelResetType("weekly")
const TargetCounterModelResetTypeMonthly = TargetCounterModelResetType("monthly")
const TargetCounterModelResetTypeDays = TargetCounterModelResetType("days")

func (p TargetCounterModelResetType) Pointer() *TargetCounterModelResetType {
	return &p
}

type TargetCounterModel struct {
	CounterName   string
	ScopeType     TargetCounterModelScopeType
	ResetType     *TargetCounterModelResetType
	ConditionName *string
	Value         int64
}

type TargetCounterModelOptions struct {
	ResetType     *TargetCounterModelResetType
	ConditionName *string
}

func NewTargetCounterModel(
	counterName string,
	scopeType TargetCounterModelScopeType,
	value int64,
	options TargetCounterModelOptions,
) TargetCounterModel {
	_data := TargetCounterModel{
		CounterName:   counterName,
		ScopeType:     scopeType,
		Value:         value,
		ResetType:     options.ResetType,
		ConditionName: options.ConditionName,
	}
	return _data
}

type TargetCounterModelScopeTypeIsResetTimingOptions struct {
	ResetType *TargetCounterModelResetType
}

func NewTargetCounterModelScopeTypeIsResetTiming(
	counterName string,
	value int64,
	options TargetCounterModelScopeTypeIsResetTimingOptions,
) TargetCounterModel {
	return NewTargetCounterModel(
		counterName,
		TargetCounterModelScopeTypeResetTiming,
		value,
		TargetCounterModelOptions{
			ResetType: options.ResetType,
		},
	)
}

type TargetCounterModelScopeTypeIsVerifyActionOptions struct {
	ResetType *TargetCounterModelResetType
}

func NewTargetCounterModelScopeTypeIsVerifyAction(
	counterName string,
	value int64,
	conditionName string,
	options TargetCounterModelScopeTypeIsVerifyActionOptions,
) TargetCounterModel {
	return NewTargetCounterModel(
		counterName,
		TargetCounterModelScopeTypeVerifyAction,
		value,
		TargetCounterModelOptions{
			ResetType:     options.ResetType,
			ConditionName: &conditionName,
		},
	)
}

func (p *TargetCounterModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CounterName"] = p.CounterName
	properties["ScopeType"] = p.ScopeType
	if p.ResetType != nil {
		properties["ResetType"] = p.ResetType
	}
	if p.ConditionName != nil {
		properties["ConditionName"] = p.ConditionName
	}
	properties["Value"] = p.Value
	return properties
}
