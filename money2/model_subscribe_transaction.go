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

type SubscribeTransactionStore string

const SubscribeTransactionStoreAppleAppStore = SubscribeTransactionStore("AppleAppStore")
const SubscribeTransactionStoreGooglePlay = SubscribeTransactionStore("GooglePlay")
const SubscribeTransactionStoreFake = SubscribeTransactionStore("fake")

func (p SubscribeTransactionStore) Pointer() *SubscribeTransactionStore {
	return &p
}

type SubscribeTransactionStatusDetail string

const SubscribeTransactionStatusDetailActiveActive = SubscribeTransactionStatusDetail("active@active")
const SubscribeTransactionStatusDetailActiveConvertedFromTrial = SubscribeTransactionStatusDetail("active@converted_from_trial")
const SubscribeTransactionStatusDetailActiveInTrial = SubscribeTransactionStatusDetail("active@in_trial")
const SubscribeTransactionStatusDetailActiveInIntroOffer = SubscribeTransactionStatusDetail("active@in_intro_offer")
const SubscribeTransactionStatusDetailGraceCanceled = SubscribeTransactionStatusDetail("grace@canceled")
const SubscribeTransactionStatusDetailGraceGracePeriod = SubscribeTransactionStatusDetail("grace@grace_period")
const SubscribeTransactionStatusDetailGraceOnHold = SubscribeTransactionStatusDetail("grace@on_hold")
const SubscribeTransactionStatusDetailInactiveExpired = SubscribeTransactionStatusDetail("inactive@expired")
const SubscribeTransactionStatusDetailInactiveRevoked = SubscribeTransactionStatusDetail("inactive@revoked")

func (p SubscribeTransactionStatusDetail) Pointer() *SubscribeTransactionStatusDetail {
	return &p
}

type SubscribeTransaction struct {
	ContentName     string
	TransactionId   string
	Store           SubscribeTransactionStore
	UserId          *string
	StatusDetail    SubscribeTransactionStatusDetail
	ExpiresAt       int64
	LastAllocatedAt *int64
	LastTakeOverAt  *int64
	CreatedAt       int64
	UpdatedAt       int64
	Revision        *int64
}

type SubscribeTransactionOptions struct {
	UserId          *string
	LastAllocatedAt *int64
	LastTakeOverAt  *int64
	Revision        *int64
}

func NewSubscribeTransaction(
	contentName string,
	transactionId string,
	store SubscribeTransactionStore,
	statusDetail SubscribeTransactionStatusDetail,
	expiresAt int64,
	createdAt int64,
	updatedAt int64,
	options SubscribeTransactionOptions,
) SubscribeTransaction {
	_data := SubscribeTransaction{
		ContentName:     contentName,
		TransactionId:   transactionId,
		Store:           store,
		StatusDetail:    statusDetail,
		ExpiresAt:       expiresAt,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		UserId:          options.UserId,
		LastAllocatedAt: options.LastAllocatedAt,
		LastTakeOverAt:  options.LastTakeOverAt,
		Revision:        options.Revision,
	}
	return _data
}

func (p *SubscribeTransaction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ContentName"] = p.ContentName
	properties["TransactionId"] = p.TransactionId
	properties["Store"] = p.Store
	if p.UserId != nil {
		properties["UserId"] = p.UserId
	}
	properties["StatusDetail"] = p.StatusDetail
	properties["ExpiresAt"] = p.ExpiresAt
	if p.LastAllocatedAt != nil {
		properties["LastAllocatedAt"] = p.LastAllocatedAt
	}
	if p.LastTakeOverAt != nil {
		properties["LastTakeOverAt"] = p.LastTakeOverAt
	}
	properties["CreatedAt"] = p.CreatedAt
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
