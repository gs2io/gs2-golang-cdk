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

type VerifyReceiptEventPlatform string

const VerifyReceiptEventPlatformAppleAppStore = VerifyReceiptEventPlatform("AppleAppStore")
const VerifyReceiptEventPlatformGooglePlay = VerifyReceiptEventPlatform("GooglePlay")
const VerifyReceiptEventPlatformFake = VerifyReceiptEventPlatform("fake")

func (p VerifyReceiptEventPlatform) Pointer() *VerifyReceiptEventPlatform {
	return &p
}

type VerifyReceiptEvent struct {
	ContentName                     string
	Platform                        VerifyReceiptEventPlatform
	AppleAppStoreVerifyReceiptEvent *AppleAppStoreVerifyReceiptEvent
	GooglePlayVerifyReceiptEvent    *GooglePlayVerifyReceiptEvent
}

type VerifyReceiptEventOptions struct {
	AppleAppStoreVerifyReceiptEvent *AppleAppStoreVerifyReceiptEvent
	GooglePlayVerifyReceiptEvent    *GooglePlayVerifyReceiptEvent
}

func NewVerifyReceiptEvent(
	contentName string,
	platform VerifyReceiptEventPlatform,
	options VerifyReceiptEventOptions,
) VerifyReceiptEvent {
	_data := VerifyReceiptEvent{
		ContentName:                     contentName,
		Platform:                        platform,
		AppleAppStoreVerifyReceiptEvent: options.AppleAppStoreVerifyReceiptEvent,
		GooglePlayVerifyReceiptEvent:    options.GooglePlayVerifyReceiptEvent,
	}
	return _data
}

func (p *VerifyReceiptEvent) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ContentName"] = p.ContentName
	properties["Platform"] = p.Platform
	if p.AppleAppStoreVerifyReceiptEvent != nil {
		properties["AppleAppStoreVerifyReceiptEvent"] = p.AppleAppStoreVerifyReceiptEvent.Properties()
	}
	if p.GooglePlayVerifyReceiptEvent != nil {
		properties["GooglePlayVerifyReceiptEvent"] = p.GooglePlayVerifyReceiptEvent.Properties()
	}
	return properties
}
