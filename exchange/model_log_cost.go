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

type LogCost struct {
	Base float64
	Adds []float64
	Subs []float64
}

type LogCostOptions struct {
	Subs []float64
}

func NewLogCost(
	base float64,
	adds []float64,
	options LogCostOptions,
) LogCost {
	_data := LogCost{
		Base: base,
		Adds: adds,
		Subs: options.Subs,
	}
	return _data
}

func (p *LogCost) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Base"] = p.Base
	properties["Adds"] = p.Adds
	properties["Subs"] = p.Subs
	return properties
}
