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

type StoreContentModel struct {
	Name          string
	Metadata      *string
	AppleAppStore *AppleAppStoreContent
	GooglePlay    *GooglePlayContent
}

type StoreContentModelOptions struct {
	Metadata      *string
	AppleAppStore *AppleAppStoreContent
	GooglePlay    *GooglePlayContent
}

func NewStoreContentModel(
	name string,
	options StoreContentModelOptions,
) StoreContentModel {
	data := StoreContentModel{
		Name:          name,
		Metadata:      options.Metadata,
		AppleAppStore: options.AppleAppStore,
		GooglePlay:    options.GooglePlay,
	}
	return data
}

func (p *StoreContentModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	if p.AppleAppStore != nil {
		properties["AppleAppStore"] = p.AppleAppStore.Properties()
	}
	if p.GooglePlay != nil {
		properties["GooglePlay"] = p.GooglePlay.Properties()
	}
	return properties
}
