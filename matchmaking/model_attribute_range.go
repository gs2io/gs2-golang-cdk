package matchmaking

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

type AttributeRange struct {
	Name string
	Min  int32
	Max  int32
}

type AttributeRangeOptions struct {
}

func NewAttributeRange(
	name string,
	min int32,
	max int32,
	options AttributeRangeOptions,
) AttributeRange {
	data := AttributeRange{
		Name: name,
		Min:  min,
		Max:  max,
	}
	return data
}

func (p *AttributeRange) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Min"] = p.Min
	properties["Max"] = p.Max
	return properties
}
