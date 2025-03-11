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

type RandomStatus struct {
	Seed int64
	Used []RandomUsed
}

type RandomStatusOptions struct {
	Used []RandomUsed
}

func NewRandomStatus(
	seed int64,
	options RandomStatusOptions,
) RandomStatus {
	_data := RandomStatus{
		Seed: seed,
		Used: options.Used,
	}
	return _data
}

func (p *RandomStatus) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Seed"] = p.Seed
	{
		values := make([]map[string]interface{}, len(p.Used))
		for i, element := range p.Used {
			values[i] = element.Properties()
		}
		properties["Used"] = values
	}
	return properties
}
