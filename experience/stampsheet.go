package experience

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

func SubExperienceByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	experienceValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["experienceValue"] = experienceValue
	return ConsumeAction{
		Action:  "Gs2Experience:SubExperienceByUserId",
		Request: properties,
	}
}

func SubRankCapByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	rankCapValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["rankCapValue"] = rankCapValue
	return ConsumeAction{
		Action:  "Gs2Experience:SubRankCapByUserId",
		Request: properties,
	}
}

func VerifyRankByUserId(
	namespaceName string,
	experienceName string,
	verifyType string,
	propertyId string,
	rankValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["verifyType"] = verifyType
	properties["propertyId"] = propertyId
	properties["rankValue"] = rankValue
	return ConsumeAction{
		Action:  "Gs2Experience:VerifyRankByUserId",
		Request: properties,
	}
}

func VerifyRankCapByUserId(
	namespaceName string,
	experienceName string,
	verifyType string,
	propertyId string,
	rankCapValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["verifyType"] = verifyType
	properties["propertyId"] = propertyId
	properties["rankCapValue"] = rankCapValue
	return ConsumeAction{
		Action:  "Gs2Experience:VerifyRankCapByUserId",
		Request: properties,
	}
}

func AddExperienceByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	experienceValue int64,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["experienceValue"] = experienceValue
	return AcquireAction{
		Action:  "Gs2Experience:AddExperienceByUserId",
		Request: properties,
	}
}

func AddRankCapByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	rankCapValue int64,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["rankCapValue"] = rankCapValue
	return AcquireAction{
		Action:  "Gs2Experience:AddRankCapByUserId",
		Request: properties,
	}
}

func SetRankCapByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	rankCapValue int64,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["rankCapValue"] = rankCapValue
	return AcquireAction{
		Action:  "Gs2Experience:SetRankCapByUserId",
		Request: properties,
	}
}

func MultiplyAcquireActionsByUserId(
	namespaceName string,
	experienceName string,
	propertyId string,
	rateName string,
	acquireActions *[]AcquireAction,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["experienceName"] = experienceName
	properties["propertyId"] = propertyId
	properties["rateName"] = rateName
	if acquireActions != nil {
		properties["acquireActions"] = acquireActions
	}
	return AcquireAction{
		Action:  "Gs2Experience:MultiplyAcquireActionsByUserId",
		Request: properties,
	}
}
