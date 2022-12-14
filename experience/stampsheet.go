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
