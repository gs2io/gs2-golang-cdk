package seasonRating

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

type SeasonModel struct {
	Name                   string
	Metadata               *string
	Tiers                  []TierModel
	ExperienceModelId      string
	ChallengePeriodEventId *string
}

type SeasonModelOptions struct {
	Metadata               *string
	ChallengePeriodEventId *string
}

func NewSeasonModel(
	name string,
	tiers []TierModel,
	experienceModelId string,
	options SeasonModelOptions,
) SeasonModel {
	_data := SeasonModel{
		Name:                   name,
		Tiers:                  tiers,
		ExperienceModelId:      experienceModelId,
		Metadata:               options.Metadata,
		ChallengePeriodEventId: options.ChallengePeriodEventId,
	}
	return _data
}

func (p *SeasonModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.Tiers))
		for i, element := range p.Tiers {
			values[i] = element.Properties()
		}
		properties["Tiers"] = values
	}
	properties["ExperienceModelId"] = p.ExperienceModelId
	if p.ChallengePeriodEventId != nil {
		properties["ChallengePeriodEventId"] = p.ChallengePeriodEventId
	}
	return properties
}
