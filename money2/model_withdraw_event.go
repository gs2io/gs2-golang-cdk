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

type WithdrawEvent struct {
	Slot            int32
	WithdrawDetails []DepositTransaction
	Status          WalletSummary
}

type WithdrawEventOptions struct {
	WithdrawDetails []DepositTransaction
}

func NewWithdrawEvent(
	slot int32,
	status WalletSummary,
	options WithdrawEventOptions,
) WithdrawEvent {
	_data := WithdrawEvent{
		Slot:            slot,
		Status:          status,
		WithdrawDetails: options.WithdrawDetails,
	}
	return _data
}

func (p *WithdrawEvent) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Slot"] = p.Slot
	{
		values := make([]map[string]interface{}, len(p.WithdrawDetails))
		for i, element := range p.WithdrawDetails {
			values[i] = element.Properties()
		}
		properties["WithdrawDetails"] = values
	}
	properties["Status"] = p.Status.Properties()
	return properties
}
