package enhance

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

func DirectEnhanceByUserId(
	namespaceName string,
	rateName string,
	targetItemSetId string,
	materials []Material,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	properties["targetItemSetId"] = targetItemSetId
	properties["materials"] = materials
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Enhance:DirectEnhanceByUserId",
		Request: properties,
	}
}

func UnleashByUserId(
	namespaceName string,
	rateName string,
	targetItemSetId string,
	materials []string,
	config *[]Config,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	properties["targetItemSetId"] = targetItemSetId
	properties["materials"] = materials
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Enhance:UnleashByUserId",
		Request: properties,
	}
}

func DeleteProgressByUserId(
	namespaceName string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	return ConsumeAction{
		Action:  "Gs2Enhance:DeleteProgressByUserId",
		Request: properties,
	}
}

func CreateProgressByUserId(
	namespaceName string,
	rateName string,
	targetItemSetId string,
	force bool,
	materials *[]Material,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rateName"] = rateName
	properties["targetItemSetId"] = targetItemSetId
	if materials != nil {
		properties["materials"] = materials
	}
	properties["force"] = force
	return AcquireAction{
		Action:  "Gs2Enhance:CreateProgressByUserId",
		Request: properties,
	}
}
