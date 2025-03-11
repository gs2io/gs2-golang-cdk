package inventory

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

type ConsumeCount struct {
	ItemName string
	Count    int64
}

type ConsumeCountOptions struct {
}

func NewConsumeCount(
	itemName string,
	count int64,
	options ConsumeCountOptions,
) ConsumeCount {
	_data := ConsumeCount{
		ItemName: itemName,
		Count:    count,
	}
	return _data
}

func (p *ConsumeCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ItemName"] = p.ItemName
	properties["Count"] = p.Count
	return properties
}
