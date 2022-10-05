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

type CounterModelMasterRef struct {
	namespaceName string
	counterName   string
}

func (p *CounterModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
			"counter",
			p.counterName,
		},
	).String()
}

func (p *CounterModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MissionGroupModelMasterRef struct {
	namespaceName    string
	missionGroupName string
}

func (p *MissionGroupModelMasterRef) MissionTaskModelMaster(
	missionTaskName string,
) *MissionTaskModelMasterRef {
	return &MissionTaskModelMasterRef{
		namespaceName:    p.namespaceName,
		missionGroupName: p.missionGroupName,
		missionTaskName:  missionTaskName,
	}
}

func (p *MissionGroupModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
			"group",
			p.missionGroupName,
		},
	).String()
}

func (p *MissionGroupModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type NamespaceRef struct {
	namespaceName string
}

func (p *NamespaceRef) CurrentMissionMaster() *CurrentMissionMasterRef {
	return &CurrentMissionMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) MissionGroupModel(
	missionGroupName string,
) *MissionGroupModelRef {
	return &MissionGroupModelRef{
		namespaceName:    p.namespaceName,
		missionGroupName: missionGroupName,
	}
}

func (p *NamespaceRef) CounterModel(
	counterName string,
) *CounterModelRef {
	return &CounterModelRef{
		namespaceName: p.namespaceName,
		counterName:   counterName,
	}
}

func (p *NamespaceRef) MissionGroupModelMaster(
	missionGroupName string,
) *MissionGroupModelMasterRef {
	return &MissionGroupModelMasterRef{
		namespaceName:    p.namespaceName,
		missionGroupName: missionGroupName,
	}
}

func (p *NamespaceRef) CounterModelMaster(
	counterName string,
) *CounterModelMasterRef {
	return &CounterModelMasterRef{
		namespaceName: p.namespaceName,
		counterName:   counterName,
	}
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentMissionMasterRef struct {
	namespaceName string
}

func (p *CurrentMissionMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentMissionMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CounterModelRef struct {
	namespaceName string
	counterName   string
}

func (p *CounterModelRef) IncreaseCounter(
	value int64,
) AcquireAction {
	return IncreaseCounterByUserId(
		p.namespaceName,
		p.counterName,
		value,
	)
}

func (p *CounterModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
			"counter",
			p.counterName,
		},
	).String()
}

func (p *CounterModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MissionGroupModelRef struct {
	namespaceName    string
	missionGroupName string
}

func (p *MissionGroupModelRef) MissionTaskModel(
	missionTaskName string,
) *MissionTaskModelRef {
	return &MissionTaskModelRef{
		namespaceName:    p.namespaceName,
		missionGroupName: p.missionGroupName,
		missionTaskName:  missionTaskName,
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
			p.namespaceName,
			"group",
			p.missionGroupName,
		},
	).String()
}

func (p *MissionGroupModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MissionTaskModelRef struct {
	namespaceName    string
	missionGroupName string
	missionTaskName  string
}

func (p *MissionTaskModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
			"group",
			p.missionGroupName,
			"missionTaskModel",
			p.missionTaskName,
		},
	).String()
}

func (p *MissionTaskModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MissionTaskModelMasterRef struct {
	namespaceName    string
	missionGroupName string
	missionTaskName  string
}

func (p *MissionTaskModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"mission",
			p.namespaceName,
			"group",
			p.missionGroupName,
			"missionTaskModelMaster",
			p.missionTaskName,
		},
	).String()
}

func (p *MissionTaskModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
