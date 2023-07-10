package lottery

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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type LotteryModelMode string

const LotteryModelModeNormal = LotteryModelMode("normal")
const LotteryModelModeBox = LotteryModelMode("box")

func (p LotteryModelMode) Pointer() *LotteryModelMode {
	return &p
}

type LotteryModelMethod string

const LotteryModelMethodPrizeTable = LotteryModelMethod("prize_table")
const LotteryModelMethodScript = LotteryModelMethod("script")

func (p LotteryModelMethod) Pointer() *LotteryModelMethod {
	return &p
}

type LotteryModel struct {
	Name                     string
	Metadata                 *string
	Mode                     LotteryModelMode
	Method                   LotteryModelMethod
	PrizeTableName           *string
	ChoicePrizeTableScriptId *string
}

type LotteryModelOptions struct {
	Metadata                 *string
	PrizeTableName           *string
	ChoicePrizeTableScriptId *string
}

func NewLotteryModel(
	name string,
	mode LotteryModelMode,
	method LotteryModelMethod,
	options LotteryModelOptions,
) LotteryModel {
	data := LotteryModel{
		Name:                     name,
		Mode:                     mode,
		Method:                   method,
		Metadata:                 options.Metadata,
		PrizeTableName:           options.PrizeTableName,
		ChoicePrizeTableScriptId: options.ChoicePrizeTableScriptId,
	}
	return data
}

type LotteryModelMethodIsPrizeTableOptions struct {
	Metadata       *string
	PrizeTableName *string
}

func NewLotteryModelMethodIsPrizeTable(
	name string,
	mode LotteryModelMode,
	options LotteryModelMethodIsPrizeTableOptions,
) LotteryModel {
	return NewLotteryModel(
		name,
		mode,
		LotteryModelMethodPrizeTable,
		LotteryModelOptions{
			Metadata:       options.Metadata,
			PrizeTableName: options.PrizeTableName,
		},
	)
}

type LotteryModelMethodIsScriptOptions struct {
	Metadata                 *string
	ChoicePrizeTableScriptId *string
}

func NewLotteryModelMethodIsScript(
	name string,
	mode LotteryModelMode,
	options LotteryModelMethodIsScriptOptions,
) LotteryModel {
	return NewLotteryModel(
		name,
		mode,
		LotteryModelMethodScript,
		LotteryModelOptions{
			Metadata:                 options.Metadata,
			ChoicePrizeTableScriptId: options.ChoicePrizeTableScriptId,
		},
	)
}

func (p *LotteryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["Mode"] = p.Mode
	properties["Method"] = p.Method
	if p.PrizeTableName != nil {
		properties["PrizeTableName"] = p.PrizeTableName
	}
	if p.ChoicePrizeTableScriptId != nil {
		properties["ChoicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
	}
	return properties
}
