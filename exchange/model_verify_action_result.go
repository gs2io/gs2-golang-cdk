package exchange

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

type VerifyActionResultAction string

func (p VerifyActionResultAction) Pointer() *VerifyActionResultAction {
	return &p
}

type VerifyActionResult struct {
	Action        string
	VerifyRequest string
	StatusCode    *int32
	VerifyResult  *string
}

type VerifyActionResultOptions struct {
	StatusCode   *int32
	VerifyResult *string
}

func NewVerifyActionResult(
	action string,
	verifyRequest string,
	options VerifyActionResultOptions,
) VerifyActionResult {
	_data := VerifyActionResult{
		Action:        action,
		VerifyRequest: verifyRequest,
		StatusCode:    options.StatusCode,
		VerifyResult:  options.VerifyResult,
	}
	return _data
}

func (p *VerifyActionResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Action"] = p.Action
	properties["VerifyRequest"] = p.VerifyRequest
	if p.StatusCode != nil {
		properties["StatusCode"] = p.StatusCode
	}
	if p.VerifyResult != nil {
		properties["VerifyResult"] = p.VerifyResult
	}
	return properties
}
