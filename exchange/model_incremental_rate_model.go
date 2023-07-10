package exchange

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

type IncrementalRateModelCalculateType string

const IncrementalRateModelCalculateTypeLinear = IncrementalRateModelCalculateType("linear")
const IncrementalRateModelCalculateTypePower = IncrementalRateModelCalculateType("power")
const IncrementalRateModelCalculateTypeGs2Script = IncrementalRateModelCalculateType("gs2_script")

func (p IncrementalRateModelCalculateType) Pointer() *IncrementalRateModelCalculateType {
	return &p
}

type IncrementalRateModel struct {
	Name              string
	Metadata          *string
	ConsumeAction     ConsumeAction
	CalculateType     IncrementalRateModelCalculateType
	BaseValue         *int64
	CoefficientValue  *int64
	CalculateScriptId *string
	ExchangeCountId   string
	AcquireActions    []AcquireAction
}

type IncrementalRateModelOptions struct {
	Metadata          *string
	BaseValue         *int64
	CoefficientValue  *int64
	CalculateScriptId *string
	AcquireActions    []AcquireAction
}

func NewIncrementalRateModel(
	name string,
	consumeAction ConsumeAction,
	calculateType IncrementalRateModelCalculateType,
	exchangeCountId string,
	options IncrementalRateModelOptions,
) IncrementalRateModel {
	data := IncrementalRateModel{
		Name:              name,
		ConsumeAction:     consumeAction,
		CalculateType:     calculateType,
		ExchangeCountId:   exchangeCountId,
		Metadata:          options.Metadata,
		BaseValue:         options.BaseValue,
		CoefficientValue:  options.CoefficientValue,
		CalculateScriptId: options.CalculateScriptId,
		AcquireActions:    options.AcquireActions,
	}
	return data
}

type IncrementalRateModelCalculateTypeIsLinearOptions struct {
	Metadata       *string
	AcquireActions []AcquireAction
}

func NewIncrementalRateModelCalculateTypeIsLinear(
	name string,
	consumeAction ConsumeAction,
	exchangeCountId string,
	baseValue int64,
	coefficientValue int64,
	options IncrementalRateModelCalculateTypeIsLinearOptions,
) IncrementalRateModel {
	return NewIncrementalRateModel(
		name,
		consumeAction,
		IncrementalRateModelCalculateTypeLinear,
		exchangeCountId,
		IncrementalRateModelOptions{
			Metadata:         options.Metadata,
			BaseValue:        &baseValue,
			CoefficientValue: &coefficientValue,
			AcquireActions:   options.AcquireActions,
		},
	)
}

type IncrementalRateModelCalculateTypeIsPowerOptions struct {
	Metadata       *string
	AcquireActions []AcquireAction
}

func NewIncrementalRateModelCalculateTypeIsPower(
	name string,
	consumeAction ConsumeAction,
	exchangeCountId string,
	coefficientValue int64,
	options IncrementalRateModelCalculateTypeIsPowerOptions,
) IncrementalRateModel {
	return NewIncrementalRateModel(
		name,
		consumeAction,
		IncrementalRateModelCalculateTypePower,
		exchangeCountId,
		IncrementalRateModelOptions{
			Metadata:         options.Metadata,
			CoefficientValue: &coefficientValue,
			AcquireActions:   options.AcquireActions,
		},
	)
}

type IncrementalRateModelCalculateTypeIsGs2ScriptOptions struct {
	Metadata       *string
	AcquireActions []AcquireAction
}

func NewIncrementalRateModelCalculateTypeIsGs2Script(
	name string,
	consumeAction ConsumeAction,
	exchangeCountId string,
	calculateScriptId string,
	options IncrementalRateModelCalculateTypeIsGs2ScriptOptions,
) IncrementalRateModel {
	return NewIncrementalRateModel(
		name,
		consumeAction,
		IncrementalRateModelCalculateTypeGs2Script,
		exchangeCountId,
		IncrementalRateModelOptions{
			Metadata:          options.Metadata,
			CalculateScriptId: &calculateScriptId,
			AcquireActions:    options.AcquireActions,
		},
	)
}

func (p *IncrementalRateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["ConsumeAction"] = p.ConsumeAction.Properties()
	properties["CalculateType"] = p.CalculateType
	if p.BaseValue != nil {
		properties["BaseValue"] = p.BaseValue
	}
	if p.CoefficientValue != nil {
		properties["CoefficientValue"] = p.CoefficientValue
	}
	if p.CalculateScriptId != nil {
		properties["CalculateScriptId"] = p.CalculateScriptId
	}
	properties["ExchangeCountId"] = p.ExchangeCountId
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}
