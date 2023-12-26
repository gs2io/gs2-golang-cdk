package grade

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

func SubGradeByUserId(
	namespaceName string,
	gradeName string,
	propertyId string,
	gradeValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["propertyId"] = propertyId
	properties["gradeValue"] = gradeValue
	return ConsumeAction{
		Action:  "Gs2Grade:SubGradeByUserId",
		Request: properties,
	}
}

func VerifyGradeByUserId(
	namespaceName string,
	gradeName string,
	verifyType string,
	propertyId string,
	gradeValue int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["verifyType"] = verifyType
	properties["propertyId"] = propertyId
	properties["gradeValue"] = gradeValue
	return ConsumeAction{
		Action:  "Gs2Grade:VerifyGradeByUserId",
		Request: properties,
	}
}

func VerifyGradeUpMaterialByUserId(
	namespaceName string,
	gradeName string,
	verifyType string,
	propertyId string,
	materialPropertyId string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["verifyType"] = verifyType
	properties["propertyId"] = propertyId
	properties["materialPropertyId"] = materialPropertyId
	return ConsumeAction{
		Action:  "Gs2Grade:VerifyGradeUpMaterialByUserId",
		Request: properties,
	}
}

func AddGradeByUserId(
	namespaceName string,
	gradeName string,
	propertyId string,
	gradeValue int64,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["propertyId"] = propertyId
	properties["gradeValue"] = gradeValue
	return AcquireAction{
		Action:  "Gs2Grade:AddGradeByUserId",
		Request: properties,
	}
}

func ApplyRankCapByUserId(
	namespaceName string,
	gradeName string,
	propertyId string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["propertyId"] = propertyId
	return AcquireAction{
		Action:  "Gs2Grade:ApplyRankCapByUserId",
		Request: properties,
	}
}

func MultiplyAcquireActionsByUserId(
	namespaceName string,
	gradeName string,
	propertyId string,
	rateName string,
	acquireActions *[]AcquireAction,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["gradeName"] = gradeName
	properties["propertyId"] = propertyId
	properties["rateName"] = rateName
	if acquireActions != nil {
		properties["acquireActions"] = acquireActions
	}
	return AcquireAction{
		Action:  "Gs2Grade:MultiplyAcquireActionsByUserId",
		Request: properties,
	}
}
