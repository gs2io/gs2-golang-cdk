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
	OwnerId    string
	Token      string
	UserId     string
	Expire     int64
	TimeOffset int32
}

type AccessTokenOptions struct {
}

func NewAccessToken(
	ownerId string,
	token string,
	userId string,
	expire int64,
	timeOffset int32,
	options AccessTokenOptions,
) AccessToken {
	data := AccessToken{
		OwnerId:    ownerId,
		Token:      token,
		UserId:     userId,
		Expire:     expire,
		TimeOffset: timeOffset,
	}
	return data
}

func (p *AccessToken) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["OwnerId"] = p.OwnerId
	properties["Token"] = p.Token
	properties["UserId"] = p.UserId
	properties["Expire"] = p.Expire
	properties["TimeOffset"] = p.TimeOffset
	return properties
}
