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

type Position struct {
	x float32
	y float32
	z float32
}

func NewPosition(
	x float32,
	y float32,
	z float32,
) Position {
	data := Position{
		x: x,
		y: y,
		z: z,
	}
	return data
}

func (p *Position) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["X"] = p.x
	properties["Y"] = p.y
	properties["Z"] = p.z
	return properties
}

type MyPosition struct {
	position Position
	vector   Vector
	r        float32
}

func NewMyPosition(
	position Position,
	vector Vector,
	r float32,
) MyPosition {
	data := MyPosition{
		position: position,
		vector:   vector,
		r:        r,
	}
	return data
}

func (p *MyPosition) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Position"] = p.position.Properties()
	properties["Vector"] = p.vector.Properties()
	properties["R"] = p.r
	return properties
}

type Scope struct {
	layerName string
	r         float32
	limit     int32
}

func NewScope(
	layerName string,
	r float32,
	limit int32,
) Scope {
	data := Scope{
		layerName: layerName,
		r:         r,
		limit:     limit,
	}
	return data
}

func (p *Scope) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["LayerName"] = p.layerName
	properties["R"] = p.r
	properties["Limit"] = p.limit
	return properties
}

type Vector struct {
	x float32
	y float32
	z float32
}

func NewVector(
	x float32,
	y float32,
	z float32,
) Vector {
	data := Vector{
		x: x,
		y: y,
		z: z,
	}
	return data
}

func (p *Vector) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["X"] = p.x
	properties["Y"] = p.y
	properties["Z"] = p.z
	return properties
}
