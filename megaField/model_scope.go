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

type Scope struct {
	LayerName string
	R         float32
	Limit     int32
}

type ScopeOptions struct {
}

func NewScope(
	layerName string,
	r float32,
	limit int32,
	options ScopeOptions,
) Scope {
	data := Scope{
		LayerName: layerName,
		R:         r,
		Limit:     limit,
	}
	return data
}

func (p *Scope) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["LayerName"] = p.LayerName
	properties["R"] = p.R
	properties["Limit"] = p.Limit
	return properties
}
