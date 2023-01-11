package version

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

type SignTargetVersion struct {
	Region        string
	OwnerId       string
	NamespaceName string
	VersionName   string
	Version       Version
}

type SignTargetVersionOptions struct {
}

func NewSignTargetVersion(
	region string,
	ownerId string,
	namespaceName string,
	versionName string,
	version Version,
	options SignTargetVersionOptions,
) SignTargetVersion {
	data := SignTargetVersion{
		Region:        region,
		OwnerId:       ownerId,
		NamespaceName: namespaceName,
		VersionName:   versionName,
		Version:       version,
	}
	return data
}

func (p *SignTargetVersion) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Region"] = p.Region
	properties["OwnerId"] = p.OwnerId
	properties["NamespaceName"] = p.NamespaceName
	properties["VersionName"] = p.VersionName
	properties["Version"] = p.Version.Properties()
	return properties
}
