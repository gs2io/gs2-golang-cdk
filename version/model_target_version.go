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

type TargetVersion struct {
	VersionName string
	Body        *string
	Signature   *string
	Version     *Version
}

type TargetVersionOptions struct {
	Body      *string
	Signature *string
	Version   *Version
}

func NewTargetVersion(
	versionName string,
	options TargetVersionOptions,
) TargetVersion {
	data := TargetVersion{
		VersionName: versionName,
		Body:        options.Body,
		Signature:   options.Signature,
		Version:     options.Version,
	}
	return data
}

func (p *TargetVersion) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["VersionName"] = p.VersionName
	if p.Body != nil {
		properties["Body"] = p.Body
	}
	if p.Signature != nil {
		properties["Signature"] = p.Signature
	}
	if p.Version != nil {
		properties["Version"] = p.Version.Properties()
	}
	return properties
}
