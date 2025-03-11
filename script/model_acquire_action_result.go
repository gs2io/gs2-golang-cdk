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

type AcquireActionResultAction string

func (p AcquireActionResultAction) Pointer() *AcquireActionResultAction {
	return &p
}

type AcquireActionResult struct {
	Action         string
	AcquireRequest string
	StatusCode     *int32
	AcquireResult  *string
}

type AcquireActionResultOptions struct {
	StatusCode    *int32
	AcquireResult *string
}

func NewAcquireActionResult(
	action string,
	acquireRequest string,
	options AcquireActionResultOptions,
) AcquireActionResult {
	_data := AcquireActionResult{
		Action:         action,
		AcquireRequest: acquireRequest,
		StatusCode:     options.StatusCode,
		AcquireResult:  options.AcquireResult,
	}
	return _data
}

func (p *AcquireActionResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Action"] = p.Action
	properties["AcquireRequest"] = p.AcquireRequest
	if p.StatusCode != nil {
		properties["StatusCode"] = p.StatusCode
	}
	if p.AcquireResult != nil {
		properties["AcquireResult"] = p.AcquireResult
	}
	return properties
}
