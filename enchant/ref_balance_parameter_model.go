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

type BalanceParameterModelRef struct {
	NamespaceName string
	ParameterName string
}

func (p *BalanceParameterModelRef) ReDrawBalanceParameterStatus(
	propertyId string,
	fixedParameterNames *[]string,
) AcquireAction {
	return ReDrawBalanceParameterStatusByUserId(
		p.NamespaceName,
		p.ParameterName,
		propertyId,
		fixedParameterNames,
	)
}

func (p *BalanceParameterModelRef) SetBalanceParameterStatus(
	propertyId string,
	parameterValues []BalanceParameterValue,
) AcquireAction {
	return SetBalanceParameterStatusByUserId(
		p.NamespaceName,
		p.ParameterName,
		propertyId,
		parameterValues,
	)
}

func (p *BalanceParameterModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"enchant",
			p.NamespaceName,
			"model",
			"balance",
			p.ParameterName,
		},
	).String()
}

func (p *BalanceParameterModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
