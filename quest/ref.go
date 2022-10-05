package quest

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

type NamespaceRef struct {
	namespaceName string
}

func (p *NamespaceRef) CurrentQuestMaster() *CurrentQuestMasterRef {
	return &CurrentQuestMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) QuestGroupModel(
	questGroupName string,
) *QuestGroupModelRef {
	return &QuestGroupModelRef{
		namespaceName:  p.namespaceName,
		questGroupName: questGroupName,
	}
}

func (p *NamespaceRef) QuestGroupModelMaster(
	questGroupName string,
) *QuestGroupModelMasterRef {
	return &QuestGroupModelMasterRef{
		namespaceName:  p.namespaceName,
		questGroupName: questGroupName,
	}
}

func (p *NamespaceRef) CreateProgress(
	questModelId string,
	force bool,
	config *[]Config,
) AcquireAction {
	return CreateProgressByUserId(
		p.namespaceName,
		questModelId,
		force,
		config,
	)
}

func (p *NamespaceRef) DeleteProgress() ConsumeAction {
	return DeleteProgressByUserId(
		p.namespaceName,
	)
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type QuestGroupModelMasterRef struct {
	namespaceName  string
	questGroupName string
}

func (p *QuestGroupModelMasterRef) QuestModelMaster(
	questName string,
) *QuestModelMasterRef {
	return &QuestModelMasterRef{
		namespaceName:  p.namespaceName,
		questGroupName: p.questGroupName,
		questName:      questName,
	}
}

func (p *QuestGroupModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
			"group",
			p.questGroupName,
		},
	).String()
}

func (p *QuestGroupModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type QuestModelMasterRef struct {
	namespaceName  string
	questGroupName string
	questName      string
}

func (p *QuestModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
			"group",
			p.questGroupName,
			"quest",
			p.questName,
		},
	).String()
}

func (p *QuestModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentQuestMasterRef struct {
	namespaceName string
}

func (p *CurrentQuestMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentQuestMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type QuestGroupModelRef struct {
	namespaceName  string
	questGroupName string
}

func (p *QuestGroupModelRef) QuestModel(
	questName string,
) *QuestModelRef {
	return &QuestModelRef{
		namespaceName:  p.namespaceName,
		questGroupName: p.questGroupName,
		questName:      questName,
	}
}

func (p *QuestGroupModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
			"group",
			p.questGroupName,
		},
	).String()
}

func (p *QuestGroupModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type QuestModelRef struct {
	namespaceName  string
	questGroupName string
	questName      string
}

func (p *QuestModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"quest",
			p.namespaceName,
			"group",
			p.questGroupName,
			"quest",
			p.questName,
		},
	).String()
}

func (p *QuestModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
