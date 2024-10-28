package exchange

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

func ExchangeByUserId(
	namespaceName string,
	rateName string,
	count int32,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	properties["count"] = count
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Exchange:ExchangeByUserId",
		Request: properties,
	}
}

func IncrementalExchangeByUserId(
	namespaceName string,
	rateName string,
	count int32,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	properties["count"] = count
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Exchange:IncrementalExchangeByUserId",
		Request: properties,
	}
}

func DeleteAwaitByUserId(
	namespaceName string,
	awaitName *string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if awaitName != nil {
		properties["awaitName"] = awaitName
	}
	return ConsumeAction{
		Action:  "Gs2Exchange:DeleteAwaitByUserId",
		Request: properties,
	}
}

func CreateAwaitByUserId(
	namespaceName string,
	rateName string,
	count *int32,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	if count != nil {
		properties["count"] = count
	}
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Exchange:CreateAwaitByUserId",
		Request: properties,
	}
}

func AcquireForceByUserId(
	namespaceName string,
	awaitName *string,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if awaitName != nil {
		properties["awaitName"] = awaitName
	}
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Exchange:AcquireForceByUserId",
		Request: properties,
	}
}

func SkipByUserId(
	namespaceName string,
	awaitName *string,
	skipType *string,
	minutes *int32,
	rate *float32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if awaitName != nil {
		properties["awaitName"] = awaitName
	}
	if skipType != nil {
		properties["skipType"] = skipType
	}
	if minutes != nil {
		properties["minutes"] = minutes
	}
	if rate != nil {
		properties["rate"] = rate
	}
	return AcquireAction{
		Action:  "Gs2Exchange:SkipByUserId",
		Request: properties,
	}
}
