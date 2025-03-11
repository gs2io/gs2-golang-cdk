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

type RefundEventPlatform string

const RefundEventPlatformAppleAppStore = RefundEventPlatform("AppleAppStore")
const RefundEventPlatformGooglePlay = RefundEventPlatform("GooglePlay")
const RefundEventPlatformFake = RefundEventPlatform("fake")

func (p RefundEventPlatform) Pointer() *RefundEventPlatform {
	return &p
}

type RefundEvent struct {
	ContentName              string
	Platform                 RefundEventPlatform
	AppleAppStoreRefundEvent *AppleAppStoreVerifyReceiptEvent
	GooglePlayRefundEvent    *GooglePlayVerifyReceiptEvent
}

type RefundEventOptions struct {
	AppleAppStoreRefundEvent *AppleAppStoreVerifyReceiptEvent
	GooglePlayRefundEvent    *GooglePlayVerifyReceiptEvent
}

func NewRefundEvent(
	contentName string,
	platform RefundEventPlatform,
	options RefundEventOptions,
) RefundEvent {
	_data := RefundEvent{
		ContentName:              contentName,
		Platform:                 platform,
		AppleAppStoreRefundEvent: options.AppleAppStoreRefundEvent,
		GooglePlayRefundEvent:    options.GooglePlayRefundEvent,
	}
	return _data
}

func (p *RefundEvent) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ContentName"] = p.ContentName
	properties["Platform"] = p.Platform
	if p.AppleAppStoreRefundEvent != nil {
		properties["AppleAppStoreRefundEvent"] = p.AppleAppStoreRefundEvent.Properties()
	}
	if p.GooglePlayRefundEvent != nil {
		properties["GooglePlayRefundEvent"] = p.GooglePlayRefundEvent.Properties()
	}
	return properties
}
