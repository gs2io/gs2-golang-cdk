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

type Version struct {
	major int32
	minor int32
	micro int32
}

func NewVersion(
	major int32,
	minor int32,
	micro int32,
) Version {
	data := Version{
		major: major,
		minor: minor,
		micro: micro,
	}
	return data
}

func (p *Version) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Major"] = p.major
	properties["Minor"] = p.minor
	properties["Micro"] = p.micro
	return properties
}

type Status struct {
	versionModel   VersionModel
	currentVersion *Version
}

func NewStatus(
	versionModel VersionModel,
	currentVersion *Version,
) Status {
	data := Status{
		versionModel:   versionModel,
		currentVersion: currentVersion,
	}
	return data
}

func (p *Status) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["VersionModel"] = p.versionModel.Properties()
	if p.currentVersion != nil {
		properties["CurrentVersion"] = p.currentVersion.Properties()
	}
	return properties
}

type TargetVersion struct {
	versionName string
	version     Version
	body        *string
	signature   *string
}

func NewTargetVersion(
	versionName string,
	version Version,
	body *string,
	signature *string,
) TargetVersion {
	data := TargetVersion{
		versionName: versionName,
		version:     version,
		body:        body,
		signature:   signature,
	}
	return data
}

func (p *TargetVersion) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["VersionName"] = p.versionName
	properties["Version"] = p.version.Properties()
	if p.body != nil {
		properties["Body"] = p.body
	}
	if p.signature != nil {
		properties["Signature"] = p.signature
	}
	return properties
}

type SignTargetVersion struct {
	region        string
	ownerId       string
	namespaceName string
	versionName   string
	version       Version
}

func NewSignTargetVersion(
	region string,
	ownerId string,
	namespaceName string,
	versionName string,
	version Version,
) SignTargetVersion {
	data := SignTargetVersion{
		region:        region,
		ownerId:       ownerId,
		namespaceName: namespaceName,
		versionName:   versionName,
		version:       version,
	}
	return data
}

func (p *SignTargetVersion) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Region"] = p.region
	properties["OwnerId"] = p.ownerId
	properties["NamespaceName"] = p.namespaceName
	properties["VersionName"] = p.versionName
	properties["Version"] = p.version.Properties()
	return properties
}
