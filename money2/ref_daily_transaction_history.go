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

type DailyTransactionHistoryRef struct {
	NamespaceName string
	Year          string
	Month         string
	Day           string
	Currency      string
}

func (p *DailyTransactionHistoryRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"money2",
			p.NamespaceName,
			"transaction",
			"history",
			"daily",
			p.Year,
			p.Month,
			p.Day,
			"currency",
			p.Currency,
		},
	).String()
}

func (p *DailyTransactionHistoryRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
