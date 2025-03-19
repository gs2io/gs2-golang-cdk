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

type RefundHistory struct {
	TransactionId string
	Year          int32
	Month         int32
	Day           int32
	UserId        *string
	Detail        RefundEvent
	CreatedAt     int64
}

type RefundHistoryOptions struct {
	UserId *string
}

func NewRefundHistory(
	transactionId string,
	year int32,
	month int32,
	day int32,
	detail RefundEvent,
	createdAt int64,
	options RefundHistoryOptions,
) RefundHistory {
	_data := RefundHistory{
		TransactionId: transactionId,
		Year:          year,
		Month:         month,
		Day:           day,
		Detail:        detail,
		CreatedAt:     createdAt,
		UserId:        options.UserId,
	}
	return _data
}

func (p *RefundHistory) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TransactionId"] = p.TransactionId
	properties["Year"] = p.Year
	properties["Month"] = p.Month
	properties["Day"] = p.Day
	if p.UserId != nil {
		properties["UserId"] = p.UserId
	}
	properties["Detail"] = p.Detail.Properties()
	properties["CreatedAt"] = p.CreatedAt
	return properties
}
