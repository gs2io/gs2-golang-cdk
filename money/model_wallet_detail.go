package money

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

type WalletDetail struct {
	Price float32
	Count int32
}

type WalletDetailOptions struct {
}

func NewWalletDetail(
	price float32,
	count int32,
	options WalletDetailOptions,
) WalletDetail {
	data := WalletDetail{
		Price: price,
		Count: count,
	}
	return data
}

func (p *WalletDetail) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Price"] = p.Price
	properties["Count"] = p.Count
	return properties
}
