package serialKey

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

func VerifyCodeByUserId(
	namespaceName string,
	code string,
	verifyType string,
	campaignModelName *string,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["code"] = code
	if campaignModelName != nil {
		properties["campaignModelName"] = campaignModelName
	}
	properties["verifyType"] = verifyType
	return VerifyAction{
		Action:  "Gs2SerialKey:VerifyCodeByUserId",
		Request: properties,
	}
}

func UseByUserId(
	namespaceName string,
	code string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["code"] = code
	return ConsumeAction{
		Action:  "Gs2SerialKey:UseByUserId",
		Request: properties,
	}
}

func RevertUseByUserId(
	namespaceName string,
	code string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["code"] = code
	return AcquireAction{
		Action:  "Gs2SerialKey:RevertUseByUserId",
		Request: properties,
	}
}

func IssueOnce(
	namespaceName string,
	campaignModelName string,
	metadata *string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["campaignModelName"] = campaignModelName
	if metadata != nil {
		properties["metadata"] = metadata
	}
	return AcquireAction{
		Action:  "Gs2SerialKey:IssueOnce",
		Request: properties,
	}
}
