package script

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

func InvokeScript(
	scriptId string,
	args *string,
	randomStatus *RandomStatus,
	forceUseDistributor *bool,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["scriptId"] = scriptId
	if args != nil {
		properties["args"] = args
	}
	if randomStatus != nil {
		properties["randomStatus"] = randomStatus
	}
	if forceUseDistributor != nil {
		properties["forceUseDistributor"] = forceUseDistributor
	}
	return AcquireAction{
		Action:  "Gs2Script:InvokeScript",
		Request: properties,
	}
}
