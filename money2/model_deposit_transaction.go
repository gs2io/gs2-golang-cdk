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

type DepositTransaction struct {
	Price       float32
	Currency    *string
	Count       int32
	DepositedAt *int64
}

type DepositTransactionOptions struct {
	Currency    *string
	DepositedAt *int64
}

func NewDepositTransaction(
	price float32,
	count int32,
	options DepositTransactionOptions,
) DepositTransaction {
	data := DepositTransaction{
		Price:       price,
		Count:       count,
		Currency:    options.Currency,
		DepositedAt: options.DepositedAt,
	}
	return data
}

func (p *DepositTransaction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Price"] = p.Price
	if p.Currency != nil {
		properties["Currency"] = p.Currency
	}
	properties["Count"] = p.Count
	if p.DepositedAt != nil {
		properties["DepositedAt"] = p.DepositedAt
	}
	return properties
}
