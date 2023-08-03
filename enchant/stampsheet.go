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

func ReDrawBalanceParameterStatusByUserId(
	namespaceName string,
	parameterName string,
	propertyId string,
	fixedParameterNames *[]string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["parameterName"] = parameterName
	properties["propertyId"] = propertyId
	if fixedParameterNames != nil {
		properties["fixedParameterNames"] = fixedParameterNames
	}
	return AcquireAction{
		Action:  "Gs2Enchant:ReDrawBalanceParameterStatusByUserId",
		Request: properties,
	}
}

func VerifyRarityParameterStatusByUserId(
	namespaceName string,
	parameterName string,
	propertyId string,
	verifyType string,
	parameterValueName string,
	parameterCount int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["parameterName"] = parameterName
	properties["propertyId"] = propertyId
	properties["verifyType"] = verifyType
	properties["parameterValueName"] = parameterValueName
	properties["parameterCount"] = parameterCount
	return ConsumeAction{
		Action:  "Gs2Enchant:VerifyRarityParameterStatusByUserId",
		Request: properties,
	}
}

func ReDrawRarityParameterStatusByUserId(
	namespaceName string,
	parameterName string,
	propertyId string,
	fixedParameterNames *[]string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["parameterName"] = parameterName
	properties["propertyId"] = propertyId
	if fixedParameterNames != nil {
		properties["fixedParameterNames"] = fixedParameterNames
	}
	return AcquireAction{
		Action:  "Gs2Enchant:ReDrawRarityParameterStatusByUserId",
		Request: properties,
	}
}

func AddRarityParameterStatusByUserId(
	namespaceName string,
	parameterName string,
	propertyId string,
	count int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["parameterName"] = parameterName
	properties["propertyId"] = propertyId
	properties["count"] = count
	return AcquireAction{
		Action:  "Gs2Enchant:AddRarityParameterStatusByUserId",
		Request: properties,
	}
}
