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

type AcquireActionRateMode string

const AcquireActionRateModeDouble = AcquireActionRateMode("double")
const AcquireActionRateModeBig = AcquireActionRateMode("big")

func (p AcquireActionRateMode) Pointer() *AcquireActionRateMode {
	return &p
}

type AcquireActionRate struct {
	Name     string
	Mode     AcquireActionRateMode
	Rates    []float64
	BigRates []string
}

type AcquireActionRateOptions struct {
	Rates    []float64
	BigRates []string
}

func NewAcquireActionRate(
	name string,
	mode AcquireActionRateMode,
	options AcquireActionRateOptions,
) AcquireActionRate {
	_data := AcquireActionRate{
		Name:     name,
		Mode:     mode,
		Rates:    options.Rates,
		BigRates: options.BigRates,
	}
	return _data
}

type AcquireActionRateModeIsDoubleOptions struct {
}

func NewAcquireActionRateModeIsDouble(
	name string,
	rates []float64,
	options AcquireActionRateModeIsDoubleOptions,
) AcquireActionRate {
	return NewAcquireActionRate(
		name,
		AcquireActionRateModeDouble,
		AcquireActionRateOptions{
			Rates: rates,
		},
	)
}

type AcquireActionRateModeIsBigOptions struct {
}

func NewAcquireActionRateModeIsBig(
	name string,
	bigRates []string,
	options AcquireActionRateModeIsBigOptions,
) AcquireActionRate {
	return NewAcquireActionRate(
		name,
		AcquireActionRateModeBig,
		AcquireActionRateOptions{
			BigRates: bigRates,
		},
	)
}

func (p *AcquireActionRate) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Mode"] = p.Mode
	properties["Rates"] = p.Rates
	properties["BigRates"] = p.BigRates
	return properties
}
