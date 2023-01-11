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

type Player struct {
	UserId      string
	Attributes  []Attribute
	RoleName    string
	DenyUserIds []string
}

type PlayerOptions struct {
	Attributes  []Attribute
	DenyUserIds []string
}

func NewPlayer(
	userId string,
	roleName string,
	options PlayerOptions,
) Player {
	data := Player{
		UserId:      userId,
		RoleName:    roleName,
		Attributes:  options.Attributes,
		DenyUserIds: options.DenyUserIds,
	}
	return data
}

func (p *Player) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	{
		values := make([]map[string]interface{}, len(p.Attributes))
		for i, element := range p.Attributes {
			values[i] = element.Properties()
		}
		properties["Attributes"] = values
	}
	properties["RoleName"] = p.RoleName
	properties["DenyUserIds"] = p.DenyUserIds
	return properties
}
