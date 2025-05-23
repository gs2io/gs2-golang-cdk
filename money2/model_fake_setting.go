package money2

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

type FakeSettingAcceptFakeReceipt string

const FakeSettingAcceptFakeReceiptAccept = FakeSettingAcceptFakeReceipt("Accept")
const FakeSettingAcceptFakeReceiptReject = FakeSettingAcceptFakeReceipt("Reject")

func (p FakeSettingAcceptFakeReceipt) Pointer() *FakeSettingAcceptFakeReceipt {
	return &p
}

type FakeSetting struct {
	AcceptFakeReceipt FakeSettingAcceptFakeReceipt
}

type FakeSettingOptions struct {
}

func NewFakeSetting(
	acceptFakeReceipt FakeSettingAcceptFakeReceipt,
	options FakeSettingOptions,
) FakeSetting {
	_data := FakeSetting{
		AcceptFakeReceipt: acceptFakeReceipt,
	}
	return _data
}

func (p *FakeSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["AcceptFakeReceipt"] = p.AcceptFakeReceipt
	return properties
}
