package stateMachine

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

type RandomUsed struct {
	Category int64
	Used     int64
}

type RandomUsedOptions struct {
}

func NewRandomUsed(
	category int64,
	used int64,
	options RandomUsedOptions,
) RandomUsed {
	_data := RandomUsed{
		Category: category,
		Used:     used,
	}
	return _data
}

func (p *RandomUsed) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Category"] = p.Category
	properties["Used"] = p.Used
	return properties
}
