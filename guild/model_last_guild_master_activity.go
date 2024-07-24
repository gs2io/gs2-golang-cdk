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

type LastGuildMasterActivity struct {
	UserId    string
	UpdatedAt int64
}

type LastGuildMasterActivityOptions struct {
}

func NewLastGuildMasterActivity(
	userId string,
	updatedAt int64,
	options LastGuildMasterActivityOptions,
) LastGuildMasterActivity {
	data := LastGuildMasterActivity{
		UserId:    userId,
		UpdatedAt: updatedAt,
	}
	return data
}

func (p *LastGuildMasterActivity) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	properties["UpdatedAt"] = p.UpdatedAt
	return properties
}
