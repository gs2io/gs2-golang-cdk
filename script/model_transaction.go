package script

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

type Transaction struct {
	TransactionId  *string
	VerifyActions  []VerifyAction
	ConsumeActions []ConsumeAction
	AcquireActions []AcquireAction
}

type TransactionOptions struct {
	TransactionId  *string
	VerifyActions  []VerifyAction
	ConsumeActions []ConsumeAction
	AcquireActions []AcquireAction
}

func NewTransaction(
	options TransactionOptions,
) Transaction {
	data := Transaction{
		TransactionId:  options.TransactionId,
		VerifyActions:  options.VerifyActions,
		ConsumeActions: options.ConsumeActions,
		AcquireActions: options.AcquireActions,
	}
	return data
}

func (p *Transaction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.TransactionId != nil {
		properties["TransactionId"] = p.TransactionId
	}
	{
		values := make([]map[string]interface{}, len(p.VerifyActions))
		for i, element := range p.VerifyActions {
			values[i] = element.Properties()
		}
		properties["VerifyActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ConsumeActions))
		for i, element := range p.ConsumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}
