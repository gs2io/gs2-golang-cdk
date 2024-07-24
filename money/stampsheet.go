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

func WithdrawByUserId(
	namespaceName string,
	slot int32,
	count int32,
	paidOnly *bool,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["slot"] = slot
	properties["count"] = count
	if paidOnly != nil {
		properties["paidOnly"] = paidOnly
	}
	return ConsumeAction{
		Action:  "Gs2Money:WithdrawByUserId",
		Request: properties,
	}
}

func DepositByUserId(
	namespaceName string,
	slot int32,
	price float32,
	count int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["slot"] = slot
	properties["price"] = price
	properties["count"] = count
	return AcquireAction{
		Action:  "Gs2Money:DepositByUserId",
		Request: properties,
	}
}

func RecordReceipt(
	namespaceName string,
	contentsId string,
	receipt string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["contentsId"] = contentsId
	properties["receipt"] = receipt
	return ConsumeAction{
		Action:  "Gs2Money:RecordReceipt",
		Request: properties,
	}
}

func RevertRecordReceipt(
	namespaceName string,
	receipt string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["receipt"] = receipt
	return AcquireAction{
		Action:  "Gs2Money:RevertRecordReceipt",
		Request: properties,
	}
}
