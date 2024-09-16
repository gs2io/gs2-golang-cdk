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

type CounterModelRef struct {
	NamespaceName string
	CounterName   string
}

func (p *CounterModelRef) IncreaseCounter(
	value int64,
) AcquireAction {
	return IncreaseCounterByUserId(
		p.NamespaceName,
		p.CounterName,
		value,
	)
}

func (p *CounterModelRef) SetCounter(
	values *[]ScopedValue,
) AcquireAction {
	return SetCounterByUserId(
		p.NamespaceName,
		p.CounterName,
		values,
	)
}

func (p *CounterModelRef) DecreaseCounter(
	value int64,
) ConsumeAction {
	return DecreaseCounterByUserId(
		p.NamespaceName,
		p.CounterName,
		value,
	)
}

func (p *CounterModelRef) VerifyCounterValue(
	verifyType string,
	scopeType *string,
	resetType *string,
	conditionName *string,
	value *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	return VerifyCounterValueByUserId(
		p.NamespaceName,
		p.CounterName,
		verifyType,
		scopeType,
		resetType,
		conditionName,
		value,
		multiplyValueSpecifyingQuantity,
	)
}

func (p *CounterModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.NamespaceName,
			"counter",
			p.CounterName,
		},
	).String()
}

func (p *CounterModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
