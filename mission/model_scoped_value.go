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

type ScopedValueResetType string

const ScopedValueResetTypeNotReset = ScopedValueResetType("notReset")
const ScopedValueResetTypeDaily = ScopedValueResetType("daily")
const ScopedValueResetTypeWeekly = ScopedValueResetType("weekly")
const ScopedValueResetTypeMonthly = ScopedValueResetType("monthly")

func (p ScopedValueResetType) Pointer() *ScopedValueResetType {
	return &p
}

type ScopedValue struct {
	ResetType   ScopedValueResetType
	Value       int64
	NextResetAt *int64
	UpdatedAt   int64
}

type ScopedValueOptions struct {
	NextResetAt *int64
}

func NewScopedValue(
	resetType ScopedValueResetType,
	value int64,
	updatedAt int64,
	options ScopedValueOptions,
) ScopedValue {
	data := ScopedValue{
		ResetType:   resetType,
		Value:       value,
		UpdatedAt:   updatedAt,
		NextResetAt: options.NextResetAt,
	}
	return data
}

func (p *ScopedValue) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ResetType"] = p.ResetType
	properties["Value"] = p.Value
	if p.NextResetAt != nil {
		properties["NextResetAt"] = p.NextResetAt
	}
	properties["UpdatedAt"] = p.UpdatedAt
	return properties
}
