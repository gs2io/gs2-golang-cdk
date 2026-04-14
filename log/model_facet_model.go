package log

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

type FacetModelType string

const FacetModelTypeString = FacetModelType("string")
const FacetModelTypeDouble = FacetModelType("double")
const FacetModelTypeMeasure = FacetModelType("measure")

func (p FacetModelType) Pointer() *FacetModelType {
	return &p
}

type FacetModel struct {
	CdkResource
	stack         *Stack
	NamespaceName string
	Field         string
	Type_         FacetModelType
	DisplayName   string
	Order         int32
}

type FacetModelOptions struct {
	Order int32
}

func NewFacetModel(
	stack *Stack,
	namespaceName string,
	field string,
	type_ FacetModelType,
	displayName string,
	options FacetModelOptions,
) *FacetModel {
	data := FacetModel{
		stack:         stack,
		NamespaceName: namespaceName,
		Field:         field,
		Type_:         type_,
		DisplayName:   displayName,
		Order:         options.Order,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *FacetModel) ResourceName() string {
	return "Log_FacetModel_" + p.Field
}

func (p *FacetModel) ResourceType() string {
	return "GS2::Log::FacetModel"
}

func (p *FacetModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["Field"] = p.Field
	properties["Type"] = p.Type_
	properties["DisplayName"] = p.DisplayName
	properties["Order"] = p.Order
	return properties
}

func (p *FacetModel) Ref(
	namespaceName string,
) FacetModelRef {
	return FacetModelRef{
		NamespaceName: namespaceName,
		Field:         p.Field,
	}
}

func (p *FacetModel) GetAttrFacetModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.FacetModelId",
	)
}
