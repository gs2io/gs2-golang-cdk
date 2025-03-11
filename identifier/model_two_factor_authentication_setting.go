package identifier

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

type TwoFactorAuthenticationSettingStatus string

const TwoFactorAuthenticationSettingStatusVerifying = TwoFactorAuthenticationSettingStatus("Verifying")
const TwoFactorAuthenticationSettingStatusEnable = TwoFactorAuthenticationSettingStatus("Enable")

func (p TwoFactorAuthenticationSettingStatus) Pointer() *TwoFactorAuthenticationSettingStatus {
	return &p
}

type TwoFactorAuthenticationSetting struct {
	AuthenticationKey string
	Status            TwoFactorAuthenticationSettingStatus
}

type TwoFactorAuthenticationSettingOptions struct {
}

func NewTwoFactorAuthenticationSetting(
	authenticationKey string,
	status TwoFactorAuthenticationSettingStatus,
	options TwoFactorAuthenticationSettingOptions,
) TwoFactorAuthenticationSetting {
	_data := TwoFactorAuthenticationSetting{
		AuthenticationKey: authenticationKey,
		Status:            status,
	}
	return _data
}

func (p *TwoFactorAuthenticationSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["AuthenticationKey"] = p.AuthenticationKey
	properties["Status"] = p.Status
	return properties
}
