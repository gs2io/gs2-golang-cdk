package loginReward

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

func MarkReceivedByUserId(
	namespaceName string,
	bonusModelName string,
	stepNumber int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["bonusModelName"] = bonusModelName
	properties["stepNumber"] = stepNumber
	return ConsumeAction{
		Action:  "Gs2LoginReward:MarkReceivedByUserId",
		Request: properties,
	}
}

func DeleteReceiveStatusByUserId(
	namespaceName string,
	bonusModelName string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["bonusModelName"] = bonusModelName
	return AcquireAction{
		Action:  "Gs2LoginReward:DeleteReceiveStatusByUserId",
		Request: properties,
	}
}

func UnmarkReceivedByUserId(
	namespaceName string,
	bonusModelName string,
	stepNumber int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["bonusModelName"] = bonusModelName
	properties["stepNumber"] = stepNumber
	return AcquireAction{
		Action:  "Gs2LoginReward:UnmarkReceivedByUserId",
		Request: properties,
	}
}
