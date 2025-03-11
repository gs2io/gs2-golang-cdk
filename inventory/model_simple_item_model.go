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

type SimpleItemModel struct {
	Name     string
	Metadata *string
}

type SimpleItemModelOptions struct {
	Metadata *string
}

func NewSimpleItemModel(
	name string,
	options SimpleItemModelOptions,
) SimpleItemModel {
	_data := SimpleItemModel{
		Name:     name,
		Metadata: options.Metadata,
	}
	return _data
}

func (p *SimpleItemModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	return properties
}
