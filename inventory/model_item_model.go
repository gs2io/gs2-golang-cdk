package inventory

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

type ItemModel struct {
	Name                string
	Metadata            *string
	StackingLimit       int64
	AllowMultipleStacks bool
	SortValue           int32
}

type ItemModelOptions struct {
	Metadata *string
}

func NewItemModel(
	name string,
	stackingLimit int64,
	allowMultipleStacks bool,
	sortValue int32,
	options ItemModelOptions,
) ItemModel {
	data := ItemModel{
		Name:                name,
		StackingLimit:       stackingLimit,
		AllowMultipleStacks: allowMultipleStacks,
		SortValue:           sortValue,
		Metadata:            options.Metadata,
	}
	return data
}

func (p *ItemModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["StackingLimit"] = p.StackingLimit
	properties["AllowMultipleStacks"] = p.AllowMultipleStacks
	properties["SortValue"] = p.SortValue
	return properties
}
