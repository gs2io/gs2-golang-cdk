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

type OpenIdConnectSetting struct {
	ConfigurationPath  string
	ClientId           string
	ClientSecret       *string
	AppleTeamId        *string
	AppleKeyId         *string
	ApplePrivateKeyPem *string
	DoneEndpointUrl    *string
}

type OpenIdConnectSettingOptions struct {
	ClientSecret       *string
	AppleTeamId        *string
	AppleKeyId         *string
	ApplePrivateKeyPem *string
	DoneEndpointUrl    *string
}

func NewOpenIdConnectSetting(
	configurationPath string,
	clientId string,
	options OpenIdConnectSettingOptions,
) OpenIdConnectSetting {
	data := OpenIdConnectSetting{
		ConfigurationPath:  configurationPath,
		ClientId:           clientId,
		ClientSecret:       options.ClientSecret,
		AppleTeamId:        options.AppleTeamId,
		AppleKeyId:         options.AppleKeyId,
		ApplePrivateKeyPem: options.ApplePrivateKeyPem,
		DoneEndpointUrl:    options.DoneEndpointUrl,
	}
	return data
}

func (p *OpenIdConnectSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ConfigurationPath"] = p.ConfigurationPath
	properties["ClientId"] = p.ClientId
	if p.ClientSecret != nil {
		properties["ClientSecret"] = p.ClientSecret
	}
	if p.AppleTeamId != nil {
		properties["AppleTeamId"] = p.AppleTeamId
	}
	if p.AppleKeyId != nil {
		properties["AppleKeyId"] = p.AppleKeyId
	}
	if p.ApplePrivateKeyPem != nil {
		properties["ApplePrivateKeyPem"] = p.ApplePrivateKeyPem
	}
	if p.DoneEndpointUrl != nil {
		properties["DoneEndpointUrl"] = p.DoneEndpointUrl
	}
	return properties
}
