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

type AppleAppStoreVerifyReceiptEventEnvironment string

const AppleAppStoreVerifyReceiptEventEnvironmentSandbox = AppleAppStoreVerifyReceiptEventEnvironment("sandbox")
const AppleAppStoreVerifyReceiptEventEnvironmentProduction = AppleAppStoreVerifyReceiptEventEnvironment("production")

func (p AppleAppStoreVerifyReceiptEventEnvironment) Pointer() *AppleAppStoreVerifyReceiptEventEnvironment {
	return &p
}

type AppleAppStoreVerifyReceiptEvent struct {
	Environment AppleAppStoreVerifyReceiptEventEnvironment
}

type AppleAppStoreVerifyReceiptEventOptions struct {
}

func NewAppleAppStoreVerifyReceiptEvent(
	environment AppleAppStoreVerifyReceiptEventEnvironment,
	options AppleAppStoreVerifyReceiptEventOptions,
) AppleAppStoreVerifyReceiptEvent {
	data := AppleAppStoreVerifyReceiptEvent{
		Environment: environment,
	}
	return data
}

func (p *AppleAppStoreVerifyReceiptEvent) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Environment"] = p.Environment
	return properties
}
