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
	Rate   float32
	Weight int32
}

type BonusRateOptions struct {
}

func NewBonusRate(
	rate float32,
	weight int32,
	options BonusRateOptions,
) BonusRate {
	_data := BonusRate{
		Rate:   rate,
		Weight: weight,
	}
	return _data
}

func (p *BonusRate) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Rate"] = p.Rate
	properties["Weight"] = p.Weight
	return properties
}
