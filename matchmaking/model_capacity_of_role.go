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

type CapacityOfRole struct {
	RoleName     string
	RoleAliases  []string
	Capacity     int32
	Participants []Player
}

type CapacityOfRoleOptions struct {
	RoleAliases  []string
	Participants []Player
}

func NewCapacityOfRole(
	roleName string,
	capacity int32,
	options CapacityOfRoleOptions,
) CapacityOfRole {
	_data := CapacityOfRole{
		RoleName:     roleName,
		Capacity:     capacity,
		RoleAliases:  options.RoleAliases,
		Participants: options.Participants,
	}
	return _data
}

func (p *CapacityOfRole) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["RoleName"] = p.RoleName
	properties["RoleAliases"] = p.RoleAliases
	properties["Capacity"] = p.Capacity
	{
		values := make([]map[string]interface{}, len(p.Participants))
		for i, element := range p.Participants {
			values[i] = element.Properties()
		}
		properties["Participants"] = values
	}
	return properties
}
