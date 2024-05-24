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

type Member struct {
	UserId   string
	RoleName string
	JoinedAt int64
}

type MemberOptions struct {
}

func NewMember(
	userId string,
	roleName string,
	joinedAt int64,
	options MemberOptions,
) Member {
	data := Member{
		UserId:   userId,
		RoleName: roleName,
		JoinedAt: joinedAt,
	}
	return data
}

func (p *Member) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	properties["RoleName"] = p.RoleName
	properties["JoinedAt"] = p.JoinedAt
	return properties
}
