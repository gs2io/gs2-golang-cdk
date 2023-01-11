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

type QuestModel struct {
	Name                        string
	Metadata                    *string
	Contents                    []Contents
	ChallengePeriodEventId      *string
	FirstCompleteAcquireActions []AcquireAction
	ConsumeActions              []ConsumeAction
	FailedAcquireActions        []AcquireAction
	PremiseQuestNames           []string
}

type QuestModelOptions struct {
	Metadata                    *string
	ChallengePeriodEventId      *string
	FirstCompleteAcquireActions []AcquireAction
	ConsumeActions              []ConsumeAction
	FailedAcquireActions        []AcquireAction
	PremiseQuestNames           []string
}

func NewQuestModel(
	name string,
	contents []Contents,
	options QuestModelOptions,
) QuestModel {
	data := QuestModel{
		Name:                        name,
		Contents:                    contents,
		Metadata:                    options.Metadata,
		ChallengePeriodEventId:      options.ChallengePeriodEventId,
		FirstCompleteAcquireActions: options.FirstCompleteAcquireActions,
		ConsumeActions:              options.ConsumeActions,
		FailedAcquireActions:        options.FailedAcquireActions,
		PremiseQuestNames:           options.PremiseQuestNames,
	}
	return data
}

func (p *QuestModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.Contents))
		for i, element := range p.Contents {
			values[i] = element.Properties()
		}
		properties["Contents"] = values
	}
	if p.ChallengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.ChallengePeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.FirstCompleteAcquireActions))
		for i, element := range p.FirstCompleteAcquireActions {
			values[i] = element.Properties()
		}
		properties["FirstCompleteAcquireActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ConsumeActions))
		for i, element := range p.ConsumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.FailedAcquireActions))
		for i, element := range p.FailedAcquireActions {
			values[i] = element.Properties()
		}
		properties["FailedAcquireActions"] = values
	}
	properties["PremiseQuestNames"] = p.PremiseQuestNames
	return properties
}
