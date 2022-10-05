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
	name               string
	metadata           *string
	consumeActions     []ConsumeAction
	timingType         RateModelTimingType
	lockTime           *int32
	enableSkip         *bool
	skipConsumeActions []ConsumeAction
	acquireActions     []AcquireAction
}

func NewRateModel(
	name string,
	metadata *string,
	consumeActions []ConsumeAction,
	timingType RateModelTimingType,
	lockTime *int32,
	enableSkip *bool,
	skipConsumeActions []ConsumeAction,
	acquireActions []AcquireAction,
) *RateModel {
	return &RateModel{
		name:               name,
		metadata:           metadata,
		consumeActions:     consumeActions,
		timingType:         timingType,
		lockTime:           lockTime,
		enableSkip:         enableSkip,
		skipConsumeActions: skipConsumeActions,
		acquireActions:     acquireActions,
	}
}

func (p *RateModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.consumeActions))
		for i, element := range p.consumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	properties["TimingType"] = p.timingType
	if p.lockTime != nil {
		properties["LockTime"] = p.lockTime
	}
	if p.enableSkip != nil {
		properties["EnableSkip"] = p.enableSkip
	}
	{
		values := make([]map[string]interface{}, len(p.skipConsumeActions))
		for i, element := range p.skipConsumeActions {
			values[i] = element.Properties()
		}
		properties["SkipConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}

func (p *RateModel) Ref(
	namespaceName string,
) RateModelRef {
	return RateModelRef{
		namespaceName: namespaceName,
		rateName:      p.name,
	}
}

func NewImmediateRateModel(
	name string,
	consumeActions []ConsumeAction,
	skipConsumeActions []ConsumeAction,
	acquireActions []AcquireAction,
	metadata *string,
) RateModel {
	return RateModel{
		timingType:         RateModelTimingTypeImmediate,
		name:               name,
		metadata:           metadata,
		consumeActions:     consumeActions,
		skipConsumeActions: skipConsumeActions,
		acquireActions:     acquireActions,
	}
}

func NewAwaitRateModel(
	name string,
	consumeActions []ConsumeAction,
	lockTime int32,
	enableSkip bool,
	skipConsumeActions []ConsumeAction,
	acquireActions []AcquireAction,
	metadata *string,
) RateModel {
	return RateModel{
		timingType:         RateModelTimingTypeAwait,
		name:               name,
		metadata:           metadata,
		consumeActions:     consumeActions,
		lockTime:           &lockTime,
		enableSkip:         &enableSkip,
		skipConsumeActions: skipConsumeActions,
		acquireActions:     acquireActions,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	rateModels    []RateModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	rateModels []RateModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-08-19",
		namespaceName: namespaceName,
		rateModels:    rateModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Exchange_CurrentRateMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Exchange::CurrentRateMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	rateModels := make([]map[string]interface{}, len(p.rateModels))
	for i, item := range p.rateModels {
		rateModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":    p.version,
			"RateModels": rateModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                *Stack
	name                 string
	description          *string
	enableAwaitExchange  bool
	enableDirectExchange bool
	transactionSetting   TransactionSetting
	exchangeScript       *ScriptSetting
	logSetting           *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	enableAwaitExchange bool,
	enableDirectExchange bool,
	transactionSetting TransactionSetting,
	description *string,
	exchangeScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                stack,
		name:                 name,
		description:          description,
		enableAwaitExchange:  enableAwaitExchange,
		enableDirectExchange: enableDirectExchange,
		transactionSetting:   transactionSetting,
		exchangeScript:       exchangeScript,
		logSetting:           logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Exchange_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Exchange::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["EnableAwaitExchange"] = p.enableAwaitExchange
	properties["EnableDirectExchange"] = p.enableDirectExchange
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.exchangeScript != nil {
		properties["ExchangeScript"] = p.exchangeScript.Properties()
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
	rateModels []RateModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		rateModels,
	).AddDependsOn(
		p,
	)
	return p
}

type RateModelMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	description        *string
	metadata           *string
	timingType         string
	lockTime           *int32
	enableSkip         *bool
	skipConsumeActions []ConsumeAction
	acquireActions     []AcquireAction
	consumeActions     []ConsumeAction
}

func NewRateModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	timingType string,
	skipConsumeActions []ConsumeAction,
	acquireActions []AcquireAction,
	consumeActions []ConsumeAction,
	description *string,
	metadata *string,
	lockTime *int32,
	enableSkip *bool,
) *RateModelMaster {

	self := RateModelMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		description:        description,
		metadata:           metadata,
		timingType:         timingType,
		lockTime:           lockTime,
		enableSkip:         enableSkip,
		skipConsumeActions: skipConsumeActions,
		acquireActions:     acquireActions,
		consumeActions:     consumeActions,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *RateModelMaster) ResourceName() string {
	return "Exchange_RateModelMaster_" + p.name
}

func (p *RateModelMaster) ResourceType() string {
	return "GS2::Exchange::RateModelMaster"
}

func (p *RateModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["TimingType"] = p.timingType
	if p.lockTime != nil {
		properties["LockTime"] = p.lockTime
	}
	if p.enableSkip != nil {
		properties["EnableSkip"] = p.enableSkip
	}
	{
		values := make([]map[string]interface{}, len(p.skipConsumeActions))
		for i, element := range p.skipConsumeActions {
			values[i] = element.Properties()
		}
		properties["SkipConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.consumeActions))
		for i, element := range p.consumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
	}
	return properties
}

func (p *RateModelMaster) Ref(
	namespaceName string,
) RateModelMasterRef {
	return RateModelMasterRef{
		namespaceName: namespaceName,
		rateName:      p.name,
	}
}

func (p *RateModelMaster) GetAttrRateModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.RateModelId",
	)
}
