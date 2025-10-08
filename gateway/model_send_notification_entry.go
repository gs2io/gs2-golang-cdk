package gateway

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

type SendNotificationEntry struct {
	UserId                           string
	Issuer                           string
	Subject                          string
	Payload                          string
	EnableTransferMobileNotification bool
	Sound                            *string
}

type SendNotificationEntryOptions struct {
	Sound *string
}

func NewSendNotificationEntry(
	userId string,
	issuer string,
	subject string,
	payload string,
	enableTransferMobileNotification bool,
	options SendNotificationEntryOptions,
) SendNotificationEntry {
	_data := SendNotificationEntry{
		UserId:                           userId,
		Issuer:                           issuer,
		Subject:                          subject,
		Payload:                          payload,
		EnableTransferMobileNotification: enableTransferMobileNotification,
		Sound:                            options.Sound,
	}
	return _data
}

func (p *SendNotificationEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	properties["Issuer"] = p.Issuer
	properties["Subject"] = p.Subject
	properties["Payload"] = p.Payload
	properties["EnableTransferMobileNotification"] = p.EnableTransferMobileNotification
	if p.Sound != nil {
		properties["Sound"] = p.Sound
	}
	return properties
}
