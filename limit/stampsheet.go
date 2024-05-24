package limit

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

func CountUpByUserId(
	namespaceName string,
	limitName string,
	counterName string,
	countUpValue int32,
	maxValue *int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["limitName"] = limitName
	properties["counterName"] = counterName
	properties["countUpValue"] = countUpValue
	if maxValue != nil {
		properties["maxValue"] = maxValue
	}
	return ConsumeAction{
		Action:  "Gs2Limit:CountUpByUserId",
		Request: properties,
	}
}

func VerifyCounterByUserId(
	namespaceName string,
	limitName string,
	counterName string,
	verifyType string,
	count int32,
	multiplyValueSpecifyingQuantity bool,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["limitName"] = limitName
	properties["counterName"] = counterName
	properties["verifyType"] = verifyType
	properties["count"] = count
	properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	return ConsumeAction{
		Action:  "Gs2Limit:VerifyCounterByUserId",
		Request: properties,
	}
}

func CountDownByUserId(
	namespaceName string,
	limitName string,
	counterName string,
	countDownValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["limitName"] = limitName
	properties["counterName"] = counterName
	properties["countDownValue"] = countDownValue
	return AcquireAction{
		Action:  "Gs2Limit:CountDownByUserId",
		Request: properties,
	}
}

func DeleteCounterByUserId(
	namespaceName string,
	limitName string,
	counterName string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["limitName"] = limitName
	properties["counterName"] = counterName
	return AcquireAction{
		Action:  "Gs2Limit:DeleteCounterByUserId",
		Request: properties,
	}
}
