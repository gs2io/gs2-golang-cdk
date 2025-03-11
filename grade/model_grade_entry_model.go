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

type GradeEntryModel struct {
	Metadata               *string
	RankCapValue           int64
	PropertyIdRegex        string
	GradeUpPropertyIdRegex string
}

type GradeEntryModelOptions struct {
	Metadata *string
}

func NewGradeEntryModel(
	rankCapValue int64,
	propertyIdRegex string,
	gradeUpPropertyIdRegex string,
	options GradeEntryModelOptions,
) GradeEntryModel {
	_data := GradeEntryModel{
		RankCapValue:           rankCapValue,
		PropertyIdRegex:        propertyIdRegex,
		GradeUpPropertyIdRegex: gradeUpPropertyIdRegex,
		Metadata:               options.Metadata,
	}
	return _data
}

func (p *GradeEntryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["RankCapValue"] = p.RankCapValue
	properties["PropertyIdRegex"] = p.PropertyIdRegex
	properties["GradeUpPropertyIdRegex"] = p.GradeUpPropertyIdRegex
	return properties
}
