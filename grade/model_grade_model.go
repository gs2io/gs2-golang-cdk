package grade

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

type GradeModel struct {
	Name               string
	Metadata           *string
	DefaultGrades      []DefaultGradeModel
	ExperienceModelId  string
	GradeEntries       []GradeEntryModel
	AcquireActionRates []AcquireActionRate
}

type GradeModelOptions struct {
	Metadata           *string
	DefaultGrades      []DefaultGradeModel
	AcquireActionRates []AcquireActionRate
}

func NewGradeModel(
	name string,
	experienceModelId string,
	gradeEntries []GradeEntryModel,
	options GradeModelOptions,
) GradeModel {
	data := GradeModel{
		Name:               name,
		ExperienceModelId:  experienceModelId,
		GradeEntries:       gradeEntries,
		Metadata:           options.Metadata,
		DefaultGrades:      options.DefaultGrades,
		AcquireActionRates: options.AcquireActionRates,
	}
	return data
}

func (p *GradeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.DefaultGrades))
		for i, element := range p.DefaultGrades {
			values[i] = element.Properties()
		}
		properties["DefaultGrades"] = values
	}
	properties["ExperienceModelId"] = p.ExperienceModelId
	{
		values := make([]map[string]interface{}, len(p.GradeEntries))
		for i, element := range p.GradeEntries {
			values[i] = element.Properties()
		}
		properties["GradeEntries"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.AcquireActionRates))
		for i, element := range p.AcquireActionRates {
			values[i] = element.Properties()
		}
		properties["AcquireActionRates"] = values
	}
	return properties
}
