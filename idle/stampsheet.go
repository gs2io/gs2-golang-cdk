package idle

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

func IncreaseMaximumIdleMinutesByUserId(
	namespaceName string,
	categoryName string,
	increaseMinutes *int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["categoryName"] = categoryName
	if increaseMinutes != nil {
		properties["increaseMinutes"] = increaseMinutes
	}
	return AcquireAction{
		Action:  "Gs2Idle:IncreaseMaximumIdleMinutesByUserId",
		Request: properties,
	}
}
