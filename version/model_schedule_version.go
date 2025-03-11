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

type ScheduleVersion struct {
	CurrentVersion  Version
	WarningVersion  Version
	ErrorVersion    Version
	ScheduleEventId *string
}

type ScheduleVersionOptions struct {
	ScheduleEventId *string
}

func NewScheduleVersion(
	currentVersion Version,
	warningVersion Version,
	errorVersion Version,
	options ScheduleVersionOptions,
) ScheduleVersion {
	_data := ScheduleVersion{
		CurrentVersion:  currentVersion,
		WarningVersion:  warningVersion,
		ErrorVersion:    errorVersion,
		ScheduleEventId: options.ScheduleEventId,
	}
	return _data
}

func (p *ScheduleVersion) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CurrentVersion"] = p.CurrentVersion.Properties()
	properties["WarningVersion"] = p.WarningVersion.Properties()
	properties["ErrorVersion"] = p.ErrorVersion.Properties()
	if p.ScheduleEventId != nil {
		properties["ScheduleEventId"] = p.ScheduleEventId
	}
	return properties
}
