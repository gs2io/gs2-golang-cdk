package idle

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

type TransactionResult struct {
	TransactionId  string
	VerifyResults  []VerifyActionResult
	ConsumeResults []ConsumeActionResult
	AcquireResults []AcquireActionResult
}

type TransactionResultOptions struct {
	VerifyResults  []VerifyActionResult
	ConsumeResults []ConsumeActionResult
	AcquireResults []AcquireActionResult
}

func NewTransactionResult(
	transactionId string,
	options TransactionResultOptions,
) TransactionResult {
	_data := TransactionResult{
		TransactionId:  transactionId,
		VerifyResults:  options.VerifyResults,
		ConsumeResults: options.ConsumeResults,
		AcquireResults: options.AcquireResults,
	}
	return _data
}

func (p *TransactionResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TransactionId"] = p.TransactionId
	{
		values := make([]map[string]interface{}, len(p.VerifyResults))
		for i, element := range p.VerifyResults {
			values[i] = element.Properties()
		}
		properties["VerifyResults"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ConsumeResults))
		for i, element := range p.ConsumeResults {
			values[i] = element.Properties()
		}
		properties["ConsumeResults"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.AcquireResults))
		for i, element := range p.AcquireResults {
			values[i] = element.Properties()
		}
		properties["AcquireResults"] = values
	}
	return properties
}
