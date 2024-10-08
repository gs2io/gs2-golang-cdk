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

type IgnoreUsers struct {
	GuildModelName string
	Users          []IgnoreUser
	CreatedAt      int64
	UpdatedAt      int64
	Revision       *int64
}

type IgnoreUsersOptions struct {
	Users    []IgnoreUser
	Revision *int64
}

func NewIgnoreUsers(
	guildModelName string,
	createdAt int64,
	updatedAt int64,
	options IgnoreUsersOptions,
) IgnoreUsers {
	data := IgnoreUsers{
		GuildModelName: guildModelName,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Users:          options.Users,
		Revision:       options.Revision,
	}
	return data
}

func (p *IgnoreUsers) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GuildModelName"] = p.GuildModelName
	{
		values := make([]map[string]interface{}, len(p.Users))
		for i, element := range p.Users {
			values[i] = element.Properties()
		}
		properties["Users"] = values
	}
	properties["CreatedAt"] = p.CreatedAt
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
