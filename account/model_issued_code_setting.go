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

type IssuedCodeSetting struct {
	ExpiresDays       int32
	MaxFailedAttempts *int32
	DeleteOnUse       bool
}

type IssuedCodeSettingOptions struct {
	MaxFailedAttempts *int32
}

func NewIssuedCodeSetting(
	expiresDays int32,
	deleteOnUse bool,
	options IssuedCodeSettingOptions,
) IssuedCodeSetting {
	_data := IssuedCodeSetting{
		ExpiresDays:       expiresDays,
		DeleteOnUse:       deleteOnUse,
		MaxFailedAttempts: options.MaxFailedAttempts,
	}
	return _data
}

func (p *IssuedCodeSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ExpiresDays"] = p.ExpiresDays
	if p.MaxFailedAttempts != nil {
		properties["MaxFailedAttempts"] = p.MaxFailedAttempts
	}
	properties["DeleteOnUse"] = p.DeleteOnUse
	return properties
}
