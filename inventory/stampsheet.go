package inventory

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

func AddCapacityByUserId(
	namespaceName string,
	inventoryName string,
	addCapacityValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["addCapacityValue"] = addCapacityValue
	return AcquireAction{
		Action:  "Gs2Inventory:AddCapacityByUserId",
		Request: properties,
	}
}

func SetCapacityByUserId(
	namespaceName string,
	inventoryName string,
	newCapacityValue int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["newCapacityValue"] = newCapacityValue
	return AcquireAction{
		Action:  "Gs2Inventory:SetCapacityByUserId",
		Request: properties,
	}
}

func ConsumeItemSetByUserId(
	namespaceName string,
	inventoryName string,
	itemName string,
	consumeCount int64,
	itemSetName *string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["itemName"] = itemName
	properties["consumeCount"] = consumeCount
	if itemSetName != nil {
		properties["itemSetName"] = itemSetName
	}
	return ConsumeAction{
		Action:  "Gs2Inventory:ConsumeItemSetByUserId",
		Request: properties,
	}
}

func AcquireItemSetByUserId(
	namespaceName string,
	inventoryName string,
	itemName string,
	acquireCount int64,
	expiresAt int64,
	createNewItemSet bool,
	itemSetName *string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["itemName"] = itemName
	properties["acquireCount"] = acquireCount
	properties["expiresAt"] = expiresAt
	properties["createNewItemSet"] = createNewItemSet
	if itemSetName != nil {
		properties["itemSetName"] = itemSetName
	}
	return AcquireAction{
		Action:  "Gs2Inventory:AcquireItemSetByUserId",
		Request: properties,
	}
}

func VerifyReferenceOfByUserId(
	namespaceName string,
	inventoryName string,
	itemName string,
	itemSetName string,
	referenceOf string,
	verifyType string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["itemName"] = itemName
	properties["itemSetName"] = itemSetName
	properties["referenceOf"] = referenceOf
	properties["verifyType"] = verifyType
	return ConsumeAction{
		Action:  "Gs2Inventory:VerifyReferenceOfByUserId",
		Request: properties,
	}
}

func AddReferenceOfByUserId(
	namespaceName string,
	inventoryName string,
	itemName string,
	itemSetName string,
	referenceOf string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["itemName"] = itemName
	properties["itemSetName"] = itemSetName
	properties["referenceOf"] = referenceOf
	return AcquireAction{
		Action:  "Gs2Inventory:AddReferenceOfByUserId",
		Request: properties,
	}
}

func DeleteReferenceOfByUserId(
	namespaceName string,
	inventoryName string,
	itemName string,
	itemSetName string,
	referenceOf string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["itemName"] = itemName
	properties["itemSetName"] = itemSetName
	properties["referenceOf"] = referenceOf
	return AcquireAction{
		Action:  "Gs2Inventory:DeleteReferenceOfByUserId",
		Request: properties,
	}
}

func ConsumeSimpleItemsByUserId(
	namespaceName string,
	inventoryName string,
	consumeCounts []ConsumeCount,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["consumeCounts"] = consumeCounts
	return ConsumeAction{
		Action:  "Gs2Inventory:ConsumeSimpleItemsByUserId",
		Request: properties,
	}
}

func AcquireSimpleItemsByUserId(
	namespaceName string,
	inventoryName string,
	acquireCounts []AcquireCount,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["inventoryName"] = inventoryName
	properties["acquireCounts"] = acquireCounts
	return AcquireAction{
		Action:  "Gs2Inventory:AcquireSimpleItemsByUserId",
		Request: properties,
	}
}
