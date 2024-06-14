package matchmaking

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

func VerifyIncludeParticipantByUserId(
	namespaceName string,
	seasonName string,
	season int64,
	tier int64,
	seasonGatheringName string,
	verifyType string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["seasonName"] = seasonName
	properties["season"] = season
	properties["tier"] = tier
	properties["seasonGatheringName"] = seasonGatheringName
	properties["verifyType"] = verifyType
	return ConsumeAction{
		Action:  "Gs2Matchmaking:VerifyIncludeParticipantByUserId",
		Request: properties,
	}
}
