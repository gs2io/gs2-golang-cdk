package formation

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

func SubMoldCapacityByUserId(
	namespaceName string,
	moldName string,
	capacity int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["moldName"] = moldName
	properties["capacity"] = capacity
	return ConsumeAction{
		Action:  "Gs2Formation:SubMoldCapacityByUserId",
		Request: properties,
	}
}

func AddMoldCapacityByUserId(
	namespaceName string,
	moldName string,
	capacity int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["moldName"] = moldName
	properties["capacity"] = capacity
	return AcquireAction{
		Action:  "Gs2Formation:AddMoldCapacityByUserId",
		Request: properties,
	}
}

func SetMoldCapacityByUserId(
	namespaceName string,
	moldName string,
	capacity int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["moldName"] = moldName
	properties["capacity"] = capacity
	return AcquireAction{
		Action:  "Gs2Formation:SetMoldCapacityByUserId",
		Request: properties,
	}
}

func AcquireActionsToFormProperties(
	namespaceName string,
	moldName string,
	index int32,
	acquireAction AcquireAction,
	config *[]AcquireActionConfig,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["moldName"] = moldName
	properties["index"] = index
	properties["acquireAction"] = acquireAction
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Formation:AcquireActionsToFormProperties",
		Request: properties,
	}
}

func AcquireActionsToPropertyFormProperties(
	namespaceName string,
	formModelName string,
	propertyId string,
	acquireAction AcquireAction,
	config *[]AcquireActionConfig,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["formModelName"] = formModelName
	properties["propertyId"] = propertyId
	properties["acquireAction"] = acquireAction
	if config != nil {
		properties["config"] = config
	}
	return AcquireAction{
		Action:  "Gs2Formation:AcquireActionsToPropertyFormProperties",
		Request: properties,
	}
}
