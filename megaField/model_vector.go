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

type Vector struct {
	X float32
	Y float32
	Z float32
}

type VectorOptions struct {
}

func NewVector(
	x float32,
	y float32,
	z float32,
	options VectorOptions,
) Vector {
	data := Vector{
		X: x,
		Y: y,
		Z: z,
	}
	return data
}

func (p *Vector) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["X"] = p.X
	properties["Y"] = p.Y
	properties["Z"] = p.Z
	return properties
}
