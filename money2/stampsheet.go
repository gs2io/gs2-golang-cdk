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

func WithdrawByUserId(
	namespaceName string,
	slot int32,
	withdrawCount int32,
	paidOnly bool,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["slot"] = slot
	properties["withdrawCount"] = withdrawCount
	properties["paidOnly"] = paidOnly
	return ConsumeAction{
		Action:  "Gs2Money2:WithdrawByUserId",
		Request: properties,
	}
}

func DepositByUserId(
	namespaceName string,
	slot int32,
	depositTransactions []DepositTransaction,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["slot"] = slot
	properties["depositTransactions"] = depositTransactions
	return AcquireAction{
		Action:  "Gs2Money2:DepositByUserId",
		Request: properties,
	}
}

func VerifyReceiptByUserId(
	namespaceName string,
	contentName string,
	receipt Receipt,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["contentName"] = contentName
	properties["receipt"] = receipt
	return ConsumeAction{
		Action:  "Gs2Money2:VerifyReceiptByUserId",
		Request: properties,
	}
}
