package account

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

type PlatformUser struct {
	Type           int32
	UserIdentifier string
	UserId         string
}

type PlatformUserOptions struct {
}

func NewPlatformUser(
	type_ int32,
	userIdentifier string,
	userId string,
	options PlatformUserOptions,
) PlatformUser {
	_data := PlatformUser{
		Type:           type_,
		UserIdentifier: userIdentifier,
		UserId:         userId,
	}
	return _data
}

func (p *PlatformUser) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Type"] = p.Type
	properties["UserIdentifier"] = p.UserIdentifier
	properties["UserId"] = p.UserId
	return properties
}
