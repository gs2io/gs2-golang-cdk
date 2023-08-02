package experience

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

type ExperienceModel struct {
	Name               string
	Metadata           *string
	DefaultExperience  int64
	DefaultRankCap     int64
	MaxRankCap         int64
	RankThreshold      Threshold
	AcquireActionRates []AcquireActionRate
}

type ExperienceModelOptions struct {
	Metadata           *string
	AcquireActionRates []AcquireActionRate
}

func NewExperienceModel(
	name string,
	defaultExperience int64,
	defaultRankCap int64,
	maxRankCap int64,
	rankThreshold Threshold,
	options ExperienceModelOptions,
) ExperienceModel {
	data := ExperienceModel{
		Name:               name,
		DefaultExperience:  defaultExperience,
		DefaultRankCap:     defaultRankCap,
		MaxRankCap:         maxRankCap,
		RankThreshold:      rankThreshold,
		Metadata:           options.Metadata,
		AcquireActionRates: options.AcquireActionRates,
	}
	return data
}

func (p *ExperienceModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["DefaultExperience"] = p.DefaultExperience
	properties["DefaultRankCap"] = p.DefaultRankCap
	properties["MaxRankCap"] = p.MaxRankCap
	properties["RankThreshold"] = p.RankThreshold.Properties()
	{
		values := make([]map[string]interface{}, len(p.AcquireActionRates))
		for i, element := range p.AcquireActionRates {
			values[i] = element.Properties()
		}
		properties["AcquireActionRates"] = values
	}
	return properties
}
