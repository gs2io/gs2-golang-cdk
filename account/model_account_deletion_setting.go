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

type AccountDeletionSettingLoginDuringGracePeriod string

const AccountDeletionSettingLoginDuringGracePeriodReject = AccountDeletionSettingLoginDuringGracePeriod("reject")
const AccountDeletionSettingLoginDuringGracePeriodCancel = AccountDeletionSettingLoginDuringGracePeriod("cancel")

func (p AccountDeletionSettingLoginDuringGracePeriod) Pointer() *AccountDeletionSettingLoginDuringGracePeriod {
	return &p
}

type AccountDeletionSettingCleanUserDataScope string

const AccountDeletionSettingCleanUserDataScopeAccount = AccountDeletionSettingCleanUserDataScope("account")
const AccountDeletionSettingCleanUserDataScopeProject = AccountDeletionSettingCleanUserDataScope("project")
const AccountDeletionSettingCleanUserDataScopeProjectExceptMoney = AccountDeletionSettingCleanUserDataScope("projectExceptMoney")

func (p AccountDeletionSettingCleanUserDataScope) Pointer() *AccountDeletionSettingCleanUserDataScope {
	return &p
}

type AccountDeletionSetting struct {
	GracePeriodMinutes     int32
	LoginDuringGracePeriod AccountDeletionSettingLoginDuringGracePeriod
	CleanUserDataScope     AccountDeletionSettingCleanUserDataScope
}

type AccountDeletionSettingOptions struct {
}

func NewAccountDeletionSetting(
	gracePeriodMinutes int32,
	loginDuringGracePeriod AccountDeletionSettingLoginDuringGracePeriod,
	cleanUserDataScope AccountDeletionSettingCleanUserDataScope,
	options AccountDeletionSettingOptions,
) AccountDeletionSetting {
	_data := AccountDeletionSetting{
		GracePeriodMinutes:     gracePeriodMinutes,
		LoginDuringGracePeriod: loginDuringGracePeriod,
		CleanUserDataScope:     cleanUserDataScope,
	}
	return _data
}

func (p *AccountDeletionSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GracePeriodMinutes"] = p.GracePeriodMinutes
	properties["LoginDuringGracePeriod"] = p.LoginDuringGracePeriod
	properties["CleanUserDataScope"] = p.CleanUserDataScope
	return properties
}
