package realtime

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

type ParticipantSetting struct {
	UserId   string
	ScriptId string
}

type ParticipantSettingOptions struct {
}

func NewParticipantSetting(
	userId string,
	scriptId string,
	options ParticipantSettingOptions,
) ParticipantSetting {
	_data := ParticipantSetting{
		UserId:   userId,
		ScriptId: scriptId,
	}
	return _data
}

func (p *ParticipantSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.UserId
	properties["ScriptId"] = p.ScriptId
	return properties
}
