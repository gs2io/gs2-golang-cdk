package buff

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

type BuffEntryModelExpression string

const BuffEntryModelExpressionRateAdd = BuffEntryModelExpression("rate_add")
const BuffEntryModelExpressionMul = BuffEntryModelExpression("mul")
const BuffEntryModelExpressionValueAdd = BuffEntryModelExpression("value_add")

func (p BuffEntryModelExpression) Pointer() *BuffEntryModelExpression {
	return &p
}

type BuffEntryModelTargetType string

const BuffEntryModelTargetTypeModel = BuffEntryModelTargetType("model")
const BuffEntryModelTargetTypeAction = BuffEntryModelTargetType("action")

func (p BuffEntryModelTargetType) Pointer() *BuffEntryModelTargetType {
	return &p
}

type BuffEntryModel struct {
	Name                       string
	Metadata                   *string
	Expression                 BuffEntryModelExpression
	TargetType                 BuffEntryModelTargetType
	TargetModel                *BuffTargetModel
	TargetAction               *BuffTargetAction
	Priority                   int32
	ApplyPeriodScheduleEventId *string
}

type BuffEntryModelOptions struct {
	Metadata                   *string
	TargetModel                *BuffTargetModel
	TargetAction               *BuffTargetAction
	ApplyPeriodScheduleEventId *string
}

func NewBuffEntryModel(
	name string,
	expression BuffEntryModelExpression,
	targetType BuffEntryModelTargetType,
	priority int32,
	options BuffEntryModelOptions,
) BuffEntryModel {
	_data := BuffEntryModel{
		Name:                       name,
		Expression:                 expression,
		TargetType:                 targetType,
		Priority:                   priority,
		Metadata:                   options.Metadata,
		TargetModel:                options.TargetModel,
		TargetAction:               options.TargetAction,
		ApplyPeriodScheduleEventId: options.ApplyPeriodScheduleEventId,
	}
	return _data
}

type BuffEntryModelTargetTypeIsModelOptions struct {
	Metadata                   *string
	ApplyPeriodScheduleEventId *string
}

func NewBuffEntryModelTargetTypeIsModel(
	name string,
	expression BuffEntryModelExpression,
	priority int32,
	targetModel BuffTargetModel,
	options BuffEntryModelTargetTypeIsModelOptions,
) BuffEntryModel {
	return NewBuffEntryModel(
		name,
		expression,
		BuffEntryModelTargetTypeModel,
		priority,
		BuffEntryModelOptions{
			Metadata:                   options.Metadata,
			TargetModel:                &targetModel,
			ApplyPeriodScheduleEventId: options.ApplyPeriodScheduleEventId,
		},
	)
}

type BuffEntryModelTargetTypeIsActionOptions struct {
	Metadata                   *string
	ApplyPeriodScheduleEventId *string
}

func NewBuffEntryModelTargetTypeIsAction(
	name string,
	expression BuffEntryModelExpression,
	priority int32,
	targetAction BuffTargetAction,
	options BuffEntryModelTargetTypeIsActionOptions,
) BuffEntryModel {
	return NewBuffEntryModel(
		name,
		expression,
		BuffEntryModelTargetTypeAction,
		priority,
		BuffEntryModelOptions{
			Metadata:                   options.Metadata,
			TargetAction:               &targetAction,
			ApplyPeriodScheduleEventId: options.ApplyPeriodScheduleEventId,
		},
	)
}

func (p *BuffEntryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["Expression"] = p.Expression
	properties["TargetType"] = p.TargetType
	if p.TargetModel != nil {
		properties["TargetModel"] = p.TargetModel.Properties()
	}
	if p.TargetAction != nil {
		properties["TargetAction"] = p.TargetAction.Properties()
	}
	properties["Priority"] = p.Priority
	if p.ApplyPeriodScheduleEventId != nil {
		properties["ApplyPeriodScheduleEventId"] = p.ApplyPeriodScheduleEventId
	}
	return properties
}
