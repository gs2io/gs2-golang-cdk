package distributor

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

type UserDataEntry struct {
	Service       string
	NamespaceName string
	Kind          string
	Payload       string
}

type UserDataEntryOptions struct {
}

func NewUserDataEntry(
	service string,
	namespaceName string,
	kind string,
	payload string,
	options UserDataEntryOptions,
) UserDataEntry {
	_data := UserDataEntry{
		Service:       service,
		NamespaceName: namespaceName,
		Kind:          kind,
		Payload:       payload,
	}
	return _data
}

func (p *UserDataEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Service"] = p.Service
	properties["NamespaceName"] = p.NamespaceName
	properties["Kind"] = p.Kind
	properties["Payload"] = p.Payload
	return properties
}
