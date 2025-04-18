package megaField

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

type MyPosition struct {
	Position Position
	Vector   Vector
	R        float32
}

type MyPositionOptions struct {
}

func NewMyPosition(
	position Position,
	vector Vector,
	r float32,
	options MyPositionOptions,
) MyPosition {
	_data := MyPosition{
		Position: position,
		Vector:   vector,
		R:        r,
	}
	return _data
}

func (p *MyPosition) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Position"] = p.Position.Properties()
	properties["Vector"] = p.Vector.Properties()
	properties["R"] = p.R
	return properties
}
