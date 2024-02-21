package distributor

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

type StampSheetResult struct {
	UserId            string
	TransactionId     string
	TaskRequests      []ConsumeAction
	SheetRequest      AcquireAction
	TaskResultCodes   []int32
	TaskResults       []string
	SheetResultCode   *int32
	SheetResult       *string
	NextTransactionId *string
	CreatedAt         int64
	TtlAt             int64
	Revision          *int64
}

type StampSheetResultOptions struct {
	TaskRequests      []ConsumeAction
	TaskResultCodes   []int32
	TaskResults       []string
	SheetResultCode   *int32
	SheetResult       *string
	NextTransactionId *string
	Revision          *int64
}

func NewStampSheetResult(
	userId string,
	transactionId string,
	sheetRequest AcquireAction,
	createdAt int64,
	ttlAt int64,
	options StampSheetResultOptions,
) StampSheetResult {
	data := StampSheetResult{
		UserId:            userId,
		TransactionId:     transactionId,
		SheetRequest:      sheetRequest,
		CreatedAt:         createdAt,
		TtlAt:             ttlAt,
		TaskRequests:      options.TaskRequests,
		TaskResultCodes:   options.TaskResultCodes,
		TaskResults:       options.TaskResults,
		SheetResultCode:   options.SheetResultCode,
		SheetResult:       options.SheetResult,
		NextTransactionId: options.NextTransactionId,
		Revision:          options.Revision,
	}
	return data
}

func (p *StampSheetResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	properties["TransactionId"] = p.TransactionId
	{
		values := make([]map[string]interface{}, len(p.TaskRequests))
		for i, element := range p.TaskRequests {
			values[i] = element.Properties()
		}
		properties["TaskRequests"] = values
	}
	properties["SheetRequest"] = p.SheetRequest.Properties()
	properties["TaskResultCodes"] = p.TaskResultCodes
	properties["TaskResults"] = p.TaskResults
	if p.SheetResultCode != nil {
		properties["SheetResultCode"] = p.SheetResultCode
	}
	if p.SheetResult != nil {
		properties["SheetResult"] = p.SheetResult
	}
	if p.NextTransactionId != nil {
		properties["NextTransactionId"] = p.NextTransactionId
	}
	properties["CreatedAt"] = p.CreatedAt
	properties["TtlAt"] = p.TtlAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
