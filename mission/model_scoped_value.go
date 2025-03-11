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

type ScopedValueScopeType string

const ScopedValueScopeTypeResetTiming = ScopedValueScopeType("resetTiming")
const ScopedValueScopeTypeVerifyAction = ScopedValueScopeType("verifyAction")

func (p ScopedValueScopeType) Pointer() *ScopedValueScopeType {
	return &p
}

type ScopedValueResetType string

const ScopedValueResetTypeNotReset = ScopedValueResetType("notReset")
const ScopedValueResetTypeDaily = ScopedValueResetType("daily")
const ScopedValueResetTypeWeekly = ScopedValueResetType("weekly")
const ScopedValueResetTypeMonthly = ScopedValueResetType("monthly")
const ScopedValueResetTypeDays = ScopedValueResetType("days")

func (p ScopedValueResetType) Pointer() *ScopedValueResetType {
	return &p
}

type ScopedValue struct {
	ScopeType     ScopedValueScopeType
	ResetType     *ScopedValueResetType
	ConditionName *string
	Value         int64
	NextResetAt   *int64
	UpdatedAt     int64
}

type ScopedValueOptions struct {
	ResetType     *ScopedValueResetType
	ConditionName *string
	NextResetAt   *int64
}

func NewScopedValue(
	scopeType ScopedValueScopeType,
	value int64,
	updatedAt int64,
	options ScopedValueOptions,
) ScopedValue {
	_data := ScopedValue{
		ScopeType:     scopeType,
		Value:         value,
		UpdatedAt:     updatedAt,
		ResetType:     options.ResetType,
		ConditionName: options.ConditionName,
		NextResetAt:   options.NextResetAt,
	}
	return _data
}

type ScopedValueScopeTypeIsResetTimingOptions struct {
	NextResetAt *int64
}

func NewScopedValueScopeTypeIsResetTiming(
	value int64,
	updatedAt int64,
	resetType ScopedValueResetType,
	options ScopedValueScopeTypeIsResetTimingOptions,
) ScopedValue {
	return NewScopedValue(
		ScopedValueScopeTypeResetTiming,
		value,
		updatedAt,
		ScopedValueOptions{
			ResetType:   &resetType,
			NextResetAt: options.NextResetAt,
		},
	)
}

type ScopedValueScopeTypeIsVerifyActionOptions struct {
	NextResetAt *int64
}

func NewScopedValueScopeTypeIsVerifyAction(
	value int64,
	updatedAt int64,
	conditionName string,
	options ScopedValueScopeTypeIsVerifyActionOptions,
) ScopedValue {
	return NewScopedValue(
		ScopedValueScopeTypeVerifyAction,
		value,
		updatedAt,
		ScopedValueOptions{
			ConditionName: &conditionName,
			NextResetAt:   options.NextResetAt,
		},
	)
}

func (p *ScopedValue) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ScopeType"] = p.ScopeType
	if p.ResetType != nil {
		properties["ResetType"] = p.ResetType
	}
	if p.ConditionName != nil {
		properties["ConditionName"] = p.ConditionName
	}
	properties["Value"] = p.Value
	if p.NextResetAt != nil {
		properties["NextResetAt"] = p.NextResetAt
	}
	properties["UpdatedAt"] = p.UpdatedAt
	return properties
}
