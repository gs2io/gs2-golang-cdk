package enchant

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

type BalanceParameterModelInitialValueStrategy string

const BalanceParameterModelInitialValueStrategyAverage = BalanceParameterModelInitialValueStrategy("average")
const BalanceParameterModelInitialValueStrategyLottery = BalanceParameterModelInitialValueStrategy("lottery")

func (p BalanceParameterModelInitialValueStrategy) Pointer() *BalanceParameterModelInitialValueStrategy {
	return &p
}

type BalanceParameterModel struct {
	Name                 string
	Metadata             *string
	TotalValue           int64
	InitialValueStrategy BalanceParameterModelInitialValueStrategy
	Parameters           []BalanceParameterValueModel
}

type BalanceParameterModelOptions struct {
	Metadata *string
}

func NewBalanceParameterModel(
	name string,
	totalValue int64,
	initialValueStrategy BalanceParameterModelInitialValueStrategy,
	parameters []BalanceParameterValueModel,
	options BalanceParameterModelOptions,
) BalanceParameterModel {
	_data := BalanceParameterModel{
		Name:                 name,
		TotalValue:           totalValue,
		InitialValueStrategy: initialValueStrategy,
		Parameters:           parameters,
		Metadata:             options.Metadata,
	}
	return _data
}

func (p *BalanceParameterModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["TotalValue"] = p.TotalValue
	properties["InitialValueStrategy"] = p.InitialValueStrategy
	{
		values := make([]map[string]interface{}, len(p.Parameters))
		for i, element := range p.Parameters {
			values[i] = element.Properties()
		}
		properties["Parameters"] = values
	}
	return properties
}
