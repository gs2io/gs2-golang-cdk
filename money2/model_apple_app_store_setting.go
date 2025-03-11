package money2

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

type AppleAppStoreSetting struct {
	BundleId        *string
	SharedSecretKey *string
	IssuerId        *string
	KeyId           *string
	PrivateKeyPem   *string
}

type AppleAppStoreSettingOptions struct {
	BundleId        *string
	SharedSecretKey *string
	IssuerId        *string
	KeyId           *string
	PrivateKeyPem   *string
}

func NewAppleAppStoreSetting(
	options AppleAppStoreSettingOptions,
) AppleAppStoreSetting {
	_data := AppleAppStoreSetting{
		BundleId:        options.BundleId,
		SharedSecretKey: options.SharedSecretKey,
		IssuerId:        options.IssuerId,
		KeyId:           options.KeyId,
		PrivateKeyPem:   options.PrivateKeyPem,
	}
	return _data
}

func (p *AppleAppStoreSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.BundleId != nil {
		properties["BundleId"] = p.BundleId
	}
	if p.SharedSecretKey != nil {
		properties["SharedSecretKey"] = p.SharedSecretKey
	}
	if p.IssuerId != nil {
		properties["IssuerId"] = p.IssuerId
	}
	if p.KeyId != nil {
		properties["KeyId"] = p.KeyId
	}
	if p.PrivateKeyPem != nil {
		properties["PrivateKeyPem"] = p.PrivateKeyPem
	}
	return properties
}
