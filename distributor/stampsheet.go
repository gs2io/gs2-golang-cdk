package distributor

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

func AndExpressionByUserId(
	namespaceName string,
	actions *[]VerifyAction,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if actions != nil {
		properties["actions"] = actions
	}
	return VerifyAction{
		Action:  "Gs2Distributor:AndExpressionByUserId",
		Request: properties,
	}
}

func OrExpressionByUserId(
	namespaceName string,
	actions *[]VerifyAction,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if actions != nil {
		properties["actions"] = actions
	}
	return VerifyAction{
		Action:  "Gs2Distributor:OrExpressionByUserId",
		Request: properties,
	}
}

func IfExpressionByUserId(
	namespaceName string,
	condition VerifyAction,
	trueActions *[]ConsumeAction,
	falseActions *[]ConsumeAction,
	multiplyValueSpecifyingQuantity *bool,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["condition"] = condition
	if trueActions != nil {
		properties["trueActions"] = trueActions
	}
	if falseActions != nil {
		properties["falseActions"] = falseActions
	}
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return ConsumeAction{
		Action:  "Gs2Distributor:IfExpressionByUserId",
		Request: properties,
	}
}
