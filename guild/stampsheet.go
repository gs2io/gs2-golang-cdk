package guild

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

func DecreaseMaximumCurrentMaximumMemberCountByGuildName(
	namespaceName string,
	guildModelName string,
	guildName string,
	value *int32,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["guildModelName"] = guildModelName
	properties["guildName"] = guildName
	if value != nil {
		properties["value"] = value
	}
	return ConsumeAction{
		Action:  "Gs2Guild:DecreaseMaximumCurrentMaximumMemberCountByGuildName",
		Request: properties,
	}
}

func IncreaseMaximumCurrentMaximumMemberCountByGuildName(
	namespaceName string,
	guildModelName string,
	guildName string,
	value *int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["guildModelName"] = guildModelName
	properties["guildName"] = guildName
	if value != nil {
		properties["value"] = value
	}
	return AcquireAction{
		Action:  "Gs2Guild:IncreaseMaximumCurrentMaximumMemberCountByGuildName",
		Request: properties,
	}
}

func SetMaximumCurrentMaximumMemberCountByGuildName(
	namespaceName string,
	guildName string,
	guildModelName string,
	value *int32,
) AcquireAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["guildName"] = guildName
	properties["guildModelName"] = guildModelName
	if value != nil {
		properties["value"] = value
	}
	return AcquireAction{
		Action:  "Gs2Guild:SetMaximumCurrentMaximumMemberCountByGuildName",
		Request: properties,
	}
}
