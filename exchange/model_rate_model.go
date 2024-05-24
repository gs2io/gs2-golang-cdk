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

type RateModelTimingType string

const RateModelTimingTypeImmediate = RateModelTimingType("immediate")
const RateModelTimingTypeAwait = RateModelTimingType("await")

func (p RateModelTimingType) Pointer() *RateModelTimingType {
	return &p
}

type RateModel struct {
	Name           string
	Metadata       *string
	ConsumeActions []ConsumeAction
	TimingType     RateModelTimingType
	LockTime       *int32
	AcquireActions []AcquireAction
}

type RateModelOptions struct {
	Metadata       *string
	ConsumeActions []ConsumeAction
	LockTime       *int32
	AcquireActions []AcquireAction
}

func NewRateModel(
	name string,
	timingType RateModelTimingType,
	options RateModelOptions,
) RateModel {
	data := RateModel{
		Name:           name,
		TimingType:     timingType,
		Metadata:       options.Metadata,
		ConsumeActions: options.ConsumeActions,
		LockTime:       options.LockTime,
		AcquireActions: options.AcquireActions,
	}
	return data
}

type RateModelTimingTypeIsImmediateOptions struct {
	Metadata       *string
	ConsumeActions []ConsumeAction
	AcquireActions []AcquireAction
}

func NewRateModelTimingTypeIsImmediate(
	name string,
	options RateModelTimingTypeIsImmediateOptions,
) RateModel {
	return NewRateModel(
		name,
		RateModelTimingTypeImmediate,
		RateModelOptions{
			Metadata:       options.Metadata,
			ConsumeActions: options.ConsumeActions,
			AcquireActions: options.AcquireActions,
		},
	)
}

type RateModelTimingTypeIsAwaitOptions struct {
	Metadata       *string
	ConsumeActions []ConsumeAction
	AcquireActions []AcquireAction
}

func NewRateModelTimingTypeIsAwait(
	name string,
	lockTime int32,
	options RateModelTimingTypeIsAwaitOptions,
) RateModel {
	return NewRateModel(
		name,
		RateModelTimingTypeAwait,
		RateModelOptions{
			Metadata:       options.Metadata,
			ConsumeActions: options.ConsumeActions,
			LockTime:       &lockTime,
			AcquireActions: options.AcquireActions,
		},
	)
}

func (p *RateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.ConsumeActions))
		for i, element := range p.ConsumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	properties["TimingType"] = p.TimingType
	if p.LockTime != nil {
		properties["LockTime"] = p.LockTime
	}
	{
		values := make([]map[string]interface{}, len(p.AcquireActions))
		for i, element := range p.AcquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}
