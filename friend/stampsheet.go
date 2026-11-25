package friend

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

func UpdateProfileByUserId(
	namespaceName string,
	publicProfile *string,
	followerProfile *string,
	friendProfile *string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if publicProfile != nil {
		properties["publicProfile"] = publicProfile
	}
	if followerProfile != nil {
		properties["followerProfile"] = followerProfile
	}
	if friendProfile != nil {
		properties["friendProfile"] = friendProfile
	}
	return AcquireAction{
		Action:  "Gs2Friend:UpdateProfileByUserId",
		Request: properties,
	}
}

func VerifyFollowByUserId(
	namespaceName string,
	targetUserId string,
	verifyType string,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["targetUserId"] = targetUserId
	properties["verifyType"] = verifyType
	return VerifyAction{
		Action:  "Gs2Friend:VerifyFollowByUserId",
		Request: properties,
	}
}

func VerifyFriendByUserId(
	namespaceName string,
	targetUserId string,
	verifyType string,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["targetUserId"] = targetUserId
	properties["verifyType"] = verifyType
	return VerifyAction{
		Action:  "Gs2Friend:VerifyFriendByUserId",
		Request: properties,
	}
}

func VerifyFriendCountByUserId(
	namespaceName string,
	verifyType string,
	count int32,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["verifyType"] = verifyType
	properties["count"] = count
	return VerifyAction{
		Action:  "Gs2Friend:VerifyFriendCountByUserId",
		Request: properties,
	}
}
