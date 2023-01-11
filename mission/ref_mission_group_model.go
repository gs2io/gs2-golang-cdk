package mission

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

type MissionGroupModelRef struct {
	NamespaceName    string
	MissionGroupName string
}

func (p *MissionGroupModelRef) MissionTaskModel(
	missionTaskName string,
) *MissionTaskModelRef {
	return &MissionTaskModelRef{
		NamespaceName:    p.NamespaceName,
		MissionGroupName: p.MissionGroupName,
		MissionTaskName:  missionTaskName,
	}
}

func (p *MissionGroupModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.NamespaceName,
			"group",
			p.MissionGroupName,
		},
	).String()
}

func (p *MissionGroupModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
