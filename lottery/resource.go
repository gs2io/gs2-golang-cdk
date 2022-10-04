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
	name                     string
	metadata                 *string
	mode                     LotteryModelMode
	method                   LotteryModelMethod
	prizeTableName           *string
	choicePrizeTableScriptId *string
}

func NewLotteryModel(
	name string,
	metadata *string,
	mode LotteryModelMode,
	method LotteryModelMethod,
	prizeTableName *string,
	choicePrizeTableScriptId *string,
) *LotteryModel {
	return &LotteryModel{
		name:                     name,
		metadata:                 metadata,
		mode:                     mode,
		method:                   method,
		prizeTableName:           prizeTableName,
		choicePrizeTableScriptId: choicePrizeTableScriptId,
	}
}

func (p *LotteryModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Mode"] = p.mode
	properties["Method"] = p.method
	if p.prizeTableName != nil {
		properties["PrizeTableName"] = p.prizeTableName
	}
	if p.choicePrizeTableScriptId != nil {
		properties["ChoicePrizeTableScriptId"] = p.choicePrizeTableScriptId
	}
	return properties
}

func (p *LotteryModel) Ref(
	namespaceName string,
) LotteryModelRef {
	return LotteryModelRef{
		namespaceName: namespaceName,
		lotteryName:   p.name,
	}
}

func NewPrizeTableLotteryModel(
	name string,
	mode LotteryModelMode,
	prizeTableName string,
	metadata *string,
) LotteryModel {
	return LotteryModel{
		method:         LotteryModelMethodPrizeTable,
		name:           name,
		metadata:       metadata,
		mode:           mode,
		prizeTableName: &prizeTableName,
	}
}

func NewScriptLotteryModel(
	name string,
	mode LotteryModelMode,
	choicePrizeTableScriptId string,
	metadata *string,
) LotteryModel {
	return LotteryModel{
		method:                   LotteryModelMethodScript,
		name:                     name,
		metadata:                 metadata,
		mode:                     mode,
		choicePrizeTableScriptId: &choicePrizeTableScriptId,
	}
}

type PrizeTable struct {
	name     string
	metadata *string
	prizes   []Prize
}

func NewPrizeTable(
	name string,
	metadata *string,
	prizes []Prize,
) *PrizeTable {
	return &PrizeTable{
		name:     name,
		metadata: metadata,
		prizes:   prizes,
	}
}

func (p *PrizeTable) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.prizes))
		for i, element := range p.prizes {
			values[i] = element.Properties()
		}
		properties["Prizes"] = values
	}
	return properties
}

func (p *PrizeTable) Ref(
	namespaceName string,
) PrizeTableRef {
	return PrizeTableRef{
		namespaceName:  namespaceName,
		prizeTableName: p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	lotteryModels []LotteryModel
	prizeTables   []PrizeTable
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	lotteryModels []LotteryModel,
	prizeTables []PrizeTable,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-02-21",
		namespaceName: namespaceName,
		lotteryModels: lotteryModels,
		prizeTables:   prizeTables,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Lottery_CurrentLotteryMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Lottery::CurrentLotteryMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	lotteryModels := make([]map[string]interface{}, len(p.lotteryModels))
	for i, item := range p.lotteryModels {
		lotteryModels[i] = item.Properties()
	}
	prizeTables := make([]map[string]interface{}, len(p.prizeTables))
	for i, item := range p.prizeTables {
		prizeTables[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":       p.version,
			"LotteryModels": lotteryModels,
			"PrizeTables":   prizeTables,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                    *Stack
	name                     string
	description              *string
	transactionSetting       TransactionSetting
	lotteryTriggerScriptId   *string
	choicePrizeTableScriptId *string
	logSetting               *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	transactionSetting TransactionSetting,
	description *string,
	lotteryTriggerScriptId *string,
	choicePrizeTableScriptId *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                    stack,
		name:                     name,
		description:              description,
		transactionSetting:       transactionSetting,
		lotteryTriggerScriptId:   lotteryTriggerScriptId,
		choicePrizeTableScriptId: choicePrizeTableScriptId,
		logSetting:               logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Lottery_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Lottery::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.lotteryTriggerScriptId != nil {
		properties["LotteryTriggerScriptId"] = p.lotteryTriggerScriptId
	}
	if p.choicePrizeTableScriptId != nil {
		properties["ChoicePrizeTableScriptId"] = p.choicePrizeTableScriptId
	}
	if p.logSetting != nil {
		properties["LogSetting"] = p.logSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		namespaceName: p.name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}

func (p *Namespace) MasterData(
	lotteryModels []LotteryModel,
	prizeTables []PrizeTable,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		lotteryModels,
		prizeTables,
	).AddDependsOn(
		p,
	)
	return p
}

type LotteryModelMaster struct {
	CdkResource
	stack                    *Stack
	namespaceName            string
	name                     string
	description              *string
	metadata                 *string
	mode                     string
	method                   string
	prizeTableName           *string
	choicePrizeTableScriptId *string
}

func NewLotteryModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	mode string,
	method string,
	description *string,
	metadata *string,
	prizeTableName *string,
	choicePrizeTableScriptId *string,
) *LotteryModelMaster {

	self := LotteryModelMaster{
		stack:                    stack,
		namespaceName:            namespaceName,
		name:                     name,
		description:              description,
		metadata:                 metadata,
		mode:                     mode,
		method:                   method,
		prizeTableName:           prizeTableName,
		choicePrizeTableScriptId: choicePrizeTableScriptId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *LotteryModelMaster) ResourceName() string {
	return "Lottery_LotteryModelMaster_" + p.name
}

func (p *LotteryModelMaster) ResourceType() string {
	return "GS2::Lottery::LotteryModelMaster"
}

func (p *LotteryModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["Mode"] = p.mode
	properties["Method"] = p.method
	if p.prizeTableName != nil {
		properties["PrizeTableName"] = p.prizeTableName
	}
	if p.choicePrizeTableScriptId != nil {
		properties["ChoicePrizeTableScriptId"] = p.choicePrizeTableScriptId
	}
	return properties
}

func (p *LotteryModelMaster) Ref(
	namespaceName string,
) LotteryModelMasterRef {
	return LotteryModelMasterRef{
		namespaceName: namespaceName,
		lotteryName:   p.name,
	}
}

func (p *LotteryModelMaster) GetAttrLotteryModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.LotteryModelId",
	)
}

type PrizeTableMaster struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	metadata      *string
	prizes        []Prize
}

func NewPrizeTableMaster(
	stack *Stack,
	namespaceName string,
	name string,
	prizes []Prize,
	description *string,
	metadata *string,
) *PrizeTableMaster {

	self := PrizeTableMaster{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		metadata:      metadata,
		prizes:        prizes,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *PrizeTableMaster) ResourceName() string {
	return "Lottery_PrizeTableMaster_" + p.name
}

func (p *PrizeTableMaster) ResourceType() string {
	return "GS2::Lottery::PrizeTableMaster"
}

func (p *PrizeTableMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.prizes))
		for i, element := range p.prizes {
			values[i] = element.Properties()
		}
		properties["Prizes"] = values
	}
	return properties
}

func (p *PrizeTableMaster) Ref(
	namespaceName string,
) PrizeTableMasterRef {
	return PrizeTableMasterRef{
		namespaceName:  namespaceName,
		prizeTableName: p.name,
	}
}

func (p *PrizeTableMaster) GetAttrPrizeTableId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.PrizeTableId",
	)
}
