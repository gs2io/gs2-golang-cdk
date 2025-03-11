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

type ReceiptStore string

const ReceiptStoreAppleAppStore = ReceiptStore("AppleAppStore")
const ReceiptStoreGooglePlay = ReceiptStore("GooglePlay")
const ReceiptStoreFake = ReceiptStore("fake")

func (p ReceiptStore) Pointer() *ReceiptStore {
	return &p
}

type Receipt struct {
	Store         ReceiptStore
	TransactionID string
	Payload       string
}

type ReceiptOptions struct {
}

func NewReceipt(
	store ReceiptStore,
	transactionID string,
	payload string,
	options ReceiptOptions,
) Receipt {
	_data := Receipt{
		Store:         store,
		TransactionID: transactionID,
		Payload:       payload,
	}
	return _data
}

func (p *Receipt) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Store"] = p.Store
	properties["TransactionID"] = p.TransactionID
	properties["Payload"] = p.Payload
	return properties
}
