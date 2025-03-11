package money2

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

type UnusedBalance struct {
	Currency  string
	Balance   float32
	UpdatedAt int64
	Revision  *int64
}

type UnusedBalanceOptions struct {
	Revision *int64
}

func NewUnusedBalance(
	currency string,
	balance float32,
	updatedAt int64,
	options UnusedBalanceOptions,
) UnusedBalance {
	_data := UnusedBalance{
		Currency:  currency,
		Balance:   balance,
		UpdatedAt: updatedAt,
		Revision:  options.Revision,
	}
	return _data
}

func (p *UnusedBalance) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Currency"] = p.Currency
	properties["Balance"] = p.Balance
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
