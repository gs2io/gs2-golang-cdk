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

type Layer struct {
	AreaModelName      string
	LayerModelName     string
	Root               *string
	NumberOfMinEntries int32
	NumberOfMaxEntries int32
	Height             int32
	CreatedAt          int64
}

type LayerOptions struct {
	Root *string
}

func NewLayer(
	areaModelName string,
	layerModelName string,
	numberOfMinEntries int32,
	numberOfMaxEntries int32,
	height int32,
	createdAt int64,
	options LayerOptions,
) Layer {
	data := Layer{
		AreaModelName:      areaModelName,
		LayerModelName:     layerModelName,
		NumberOfMinEntries: numberOfMinEntries,
		NumberOfMaxEntries: numberOfMaxEntries,
		Height:             height,
		CreatedAt:          createdAt,
		Root:               options.Root,
	}
	return data
}

func (p *Layer) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["AreaModelName"] = p.AreaModelName
	properties["LayerModelName"] = p.LayerModelName
	if p.Root != nil {
		properties["Root"] = p.Root
	}
	properties["NumberOfMinEntries"] = p.NumberOfMinEntries
	properties["NumberOfMaxEntries"] = p.NumberOfMaxEntries
	properties["Height"] = p.Height
	properties["CreatedAt"] = p.CreatedAt
	return properties
}
