package dictionary

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

func VerifyEntryByUserId(
	namespaceName string,
	entryModelName string,
	verifyType string,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["entryModelName"] = entryModelName
	properties["verifyType"] = verifyType
	return VerifyAction{
		Action:  "Gs2Dictionary:VerifyEntryByUserId",
		Request: properties,
	}
}

func DeleteEntriesByUserId(
	namespaceName string,
	entryModelNames *[]string,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if entryModelNames != nil {
		properties["entryModelNames"] = entryModelNames
	}
	return ConsumeAction{
		Action:  "Gs2Dictionary:DeleteEntriesByUserId",
		Request: properties,
	}
}

func AddEntriesByUserId(
	namespaceName string,
	entryModelNames *[]string,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	if entryModelNames != nil {
		properties["entryModelNames"] = entryModelNames
	}
	return AcquireAction{
		Action:  "Gs2Dictionary:AddEntriesByUserId",
		Request: properties,
	}
}
