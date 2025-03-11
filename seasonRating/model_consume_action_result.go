package seasonRating

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

type ConsumeActionResultAction string

func (p ConsumeActionResultAction) Pointer() *ConsumeActionResultAction {
	return &p
}

type ConsumeActionResult struct {
	Action         string
	ConsumeRequest string
	StatusCode     *int32
	ConsumeResult  *string
}

type ConsumeActionResultOptions struct {
	StatusCode    *int32
	ConsumeResult *string
}

func NewConsumeActionResult(
	action string,
	consumeRequest string,
	options ConsumeActionResultOptions,
) ConsumeActionResult {
	_data := ConsumeActionResult{
		Action:         action,
		ConsumeRequest: consumeRequest,
		StatusCode:     options.StatusCode,
		ConsumeResult:  options.ConsumeResult,
	}
	return _data
}

func (p *ConsumeActionResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Action"] = p.Action
	properties["ConsumeRequest"] = p.ConsumeRequest
	if p.StatusCode != nil {
		properties["StatusCode"] = p.StatusCode
	}
	if p.ConsumeResult != nil {
		properties["ConsumeResult"] = p.ConsumeResult
	}
	return properties
}
