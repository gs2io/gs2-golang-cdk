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

type DailyTransactionHistory struct {
	Year           int32
	Month          int32
	Day            int32
	Currency       string
	DepositAmount  float64
	WithdrawAmount float64
	IssueCount     int64
	ConsumeCount   int64
	UpdatedAt      int64
	Revision       *int64
}

type DailyTransactionHistoryOptions struct {
	Revision *int64
}

func NewDailyTransactionHistory(
	year int32,
	month int32,
	day int32,
	currency string,
	depositAmount float64,
	withdrawAmount float64,
	issueCount int64,
	consumeCount int64,
	updatedAt int64,
	options DailyTransactionHistoryOptions,
) DailyTransactionHistory {
	_data := DailyTransactionHistory{
		Year:           year,
		Month:          month,
		Day:            day,
		Currency:       currency,
		DepositAmount:  depositAmount,
		WithdrawAmount: withdrawAmount,
		IssueCount:     issueCount,
		ConsumeCount:   consumeCount,
		UpdatedAt:      updatedAt,
		Revision:       options.Revision,
	}
	return _data
}

func (p *DailyTransactionHistory) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Year"] = p.Year
	properties["Month"] = p.Month
	properties["Day"] = p.Day
	properties["Currency"] = p.Currency
	properties["DepositAmount"] = p.DepositAmount
	properties["WithdrawAmount"] = p.WithdrawAmount
	properties["IssueCount"] = p.IssueCount
	properties["ConsumeCount"] = p.ConsumeCount
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
