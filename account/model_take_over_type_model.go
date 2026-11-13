package account

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

type TakeOverTypeModelTakeOverMethod string

const TakeOverTypeModelTakeOverMethodOpenIdConnect = TakeOverTypeModelTakeOverMethod("openIdConnect")
const TakeOverTypeModelTakeOverMethodIssuedCode = TakeOverTypeModelTakeOverMethod("issuedCode")

func (p TakeOverTypeModelTakeOverMethod) Pointer() *TakeOverTypeModelTakeOverMethod {
	return &p
}

type TakeOverTypeModel struct {
	Type                 int32
	Metadata             *string
	TakeOverMethod       TakeOverTypeModelTakeOverMethod
	OpenIdConnectSetting *OpenIdConnectSetting
	IssuedCodeSetting    *IssuedCodeSetting
}

type TakeOverTypeModelOptions struct {
	Metadata             *string
	OpenIdConnectSetting *OpenIdConnectSetting
	IssuedCodeSetting    *IssuedCodeSetting
}

func NewTakeOverTypeModel(
	type_ int32,
	takeOverMethod TakeOverTypeModelTakeOverMethod,
	options TakeOverTypeModelOptions,
) TakeOverTypeModel {
	_data := TakeOverTypeModel{
		Type:                 type_,
		TakeOverMethod:       takeOverMethod,
		Metadata:             options.Metadata,
		OpenIdConnectSetting: options.OpenIdConnectSetting,
		IssuedCodeSetting:    options.IssuedCodeSetting,
	}
	return _data
}

type TakeOverTypeModelTakeOverMethodIsOpenIdConnectOptions struct {
	Metadata *string
}

func NewTakeOverTypeModelTakeOverMethodIsOpenIdConnect(
	type_ int32,
	openIdConnectSetting OpenIdConnectSetting,
	options TakeOverTypeModelTakeOverMethodIsOpenIdConnectOptions,
) TakeOverTypeModel {
	return NewTakeOverTypeModel(
		type_,
		TakeOverTypeModelTakeOverMethodOpenIdConnect,
		TakeOverTypeModelOptions{
			Metadata:             options.Metadata,
			OpenIdConnectSetting: &openIdConnectSetting,
		},
	)
}

type TakeOverTypeModelTakeOverMethodIsIssuedCodeOptions struct {
	Metadata *string
}

func NewTakeOverTypeModelTakeOverMethodIsIssuedCode(
	type_ int32,
	issuedCodeSetting IssuedCodeSetting,
	options TakeOverTypeModelTakeOverMethodIsIssuedCodeOptions,
) TakeOverTypeModel {
	return NewTakeOverTypeModel(
		type_,
		TakeOverTypeModelTakeOverMethodIssuedCode,
		TakeOverTypeModelOptions{
			Metadata:          options.Metadata,
			IssuedCodeSetting: &issuedCodeSetting,
		},
	)
}

func (p *TakeOverTypeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Type"] = p.Type
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["TakeOverMethod"] = p.TakeOverMethod
	if p.OpenIdConnectSetting != nil {
		properties["OpenIdConnectSetting"] = p.OpenIdConnectSetting.Properties()
	}
	if p.IssuedCodeSetting != nil {
		properties["IssuedCodeSetting"] = p.IssuedCodeSetting.Properties()
	}
	return properties
}
