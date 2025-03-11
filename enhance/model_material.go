package enhance

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

type Material struct {
	MaterialItemSetId string
	Count             int32
}

type MaterialOptions struct {
}

func NewMaterial(
	materialItemSetId string,
	count int32,
	options MaterialOptions,
) Material {
	_data := Material{
		MaterialItemSetId: materialItemSetId,
		Count:             count,
	}
	return _data
}

func (p *Material) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["MaterialItemSetId"] = p.MaterialItemSetId
	properties["Count"] = p.Count
	return properties
}
