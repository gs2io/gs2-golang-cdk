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

type BonusRate struct {
	rate   float32
	weight int32
}

func NewBonusRate(
	rate float32,
	weight int32,
) BonusRate {
	data := BonusRate{
		rate:   rate,
		weight: weight,
	}
	return data
}

func (p *BonusRate) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Rate"] = p.rate
	properties["Weight"] = p.weight
	return properties
}

type Material struct {
	materialItemSetId string
	count             int32
}

func NewMaterial(
	materialItemSetId string,
	count int32,
) Material {
	data := Material{
		materialItemSetId: materialItemSetId,
		count:             count,
	}
	return data
}

func (p *Material) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["MaterialItemSetId"] = p.materialItemSetId
	properties["Count"] = p.count
	return properties
}
