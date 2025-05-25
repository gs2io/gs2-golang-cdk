package freeze

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

type StageStatus string

const StageStatusActive = StageStatus("Active")
const StageStatusUpdating = StageStatus("Updating")
const StageStatusUpdateFailed = StageStatus("UpdateFailed")

func (p StageStatus) Pointer() *StageStatus {
	return &p
}

type Stage struct {
	CdkResource
	stack           *Stack
	Name            string
	SourceStageName *string
	SortNumber      int32
}

type StageOptions struct {
	SourceStageName *string
}

func NewStage(
	stack *Stack,
	name string,
	sortNumber int32,
	options StageOptions,
) *Stage {
	data := Stage{
		stack:           stack,
		Name:            name,
		SortNumber:      sortNumber,
		SourceStageName: options.SourceStageName,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Stage) ResourceName() string {
	return "Freeze_Stage_" + p.Name
}

func (p *Stage) ResourceType() string {
	return "GS2::Freeze::Stage"
}

func (p *Stage) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.SourceStageName != nil {
		properties["SourceStageName"] = p.SourceStageName
	}
	properties["SortNumber"] = p.SortNumber
	return properties
}

func (p *Stage) Ref() StageRef {
	return StageRef{
		StageName: p.Name,
	}
}

func (p *Stage) GetAttrStageId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.StageId",
	)
}
