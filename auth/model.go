package auth

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

type AccessToken struct {
	ownerId    string
	token      string
	userId     string
	expire     int64
	timeOffset int32
}

func NewAccessToken(
	ownerId string,
	token string,
	userId string,
	expire int64,
	timeOffset int32,
) AccessToken {
	data := AccessToken{
		ownerId:    ownerId,
		token:      token,
		userId:     userId,
		expire:     expire,
		timeOffset: timeOffset,
	}
	return data
}

func (p *AccessToken) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["OwnerId"] = p.ownerId
	properties["Token"] = p.token
	properties["UserId"] = p.userId
	properties["Expire"] = p.expire
	properties["TimeOffset"] = p.timeOffset
	return properties
}
