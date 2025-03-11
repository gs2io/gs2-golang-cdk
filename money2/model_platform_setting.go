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

type PlatformSetting struct {
	AppleAppStore *AppleAppStoreSetting
	GooglePlay    *GooglePlaySetting
	Fake          *FakeSetting
}

type PlatformSettingOptions struct {
	AppleAppStore *AppleAppStoreSetting
	GooglePlay    *GooglePlaySetting
	Fake          *FakeSetting
}

func NewPlatformSetting(
	options PlatformSettingOptions,
) PlatformSetting {
	_data := PlatformSetting{
		AppleAppStore: options.AppleAppStore,
		GooglePlay:    options.GooglePlay,
		Fake:          options.Fake,
	}
	return _data
}

func (p *PlatformSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.AppleAppStore != nil {
		properties["AppleAppStore"] = p.AppleAppStore.Properties()
	}
	if p.GooglePlay != nil {
		properties["GooglePlay"] = p.GooglePlay.Properties()
	}
	if p.Fake != nil {
		properties["Fake"] = p.Fake.Properties()
	}
	return properties
}
