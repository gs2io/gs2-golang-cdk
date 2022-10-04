package showcase

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

type SalesItem struct {
	name           string
	metadata       *string
	consumeActions []ConsumeAction
	acquireActions []AcquireAction
}

func NewSalesItem(
	name string,
	metadata *string,
	consumeActions []ConsumeAction,
	acquireActions []AcquireAction,
) *SalesItem {
	return &SalesItem{
		name:           name,
		metadata:       metadata,
		consumeActions: consumeActions,
		acquireActions: acquireActions,
	}
}

func (p *SalesItem) Properties() map[string]interface{} {
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
	{
		values := make([]map[string]interface{}, len(p.acquireActions))
		for i, element := range p.acquireActions {
			values[i] = element.Properties()
		}
		properties["AcquireActions"] = values
	}
	return properties
}

type SalesItemGroup struct {
	name       string
	metadata   *string
	salesItems []SalesItem
}

func NewSalesItemGroup(
	name string,
	metadata *string,
	salesItems []SalesItem,
) *SalesItemGroup {
	return &SalesItemGroup{
		name:       name,
		metadata:   metadata,
		salesItems: salesItems,
	}
}

func (p *SalesItemGroup) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.salesItems))
		for i, element := range p.salesItems {
			values[i] = element.Properties()
		}
		properties["SalesItems"] = values
	}
	return properties
}

type Showcase struct {
	name               string
	metadata           *string
	salesPeriodEventId *string
	displayItems       []DisplayItem
}

func NewShowcase(
	name string,
	metadata *string,
	salesPeriodEventId *string,
	displayItems []DisplayItem,
) *Showcase {
	return &Showcase{
		name:               name,
		metadata:           metadata,
		salesPeriodEventId: salesPeriodEventId,
		displayItems:       displayItems,
	}
}

func (p *Showcase) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.salesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.salesPeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.displayItems))
		for i, element := range p.displayItems {
			values[i] = element.Properties()
		}
		properties["DisplayItems"] = values
	}
	return properties
}

func (p *Showcase) Ref(
	namespaceName string,
) ShowcaseRef {
	return ShowcaseRef{
		namespaceName: namespaceName,
		showcaseName:  p.name,
	}
}

type DisplayItemType string

const DisplayItemTypeSalesItem = DisplayItemType("salesItem")
const DisplayItemTypeSalesItemGroup = DisplayItemType("salesItemGroup")

func (p DisplayItemType) Pointer() *DisplayItemType {
	return &p
}

type DisplayItem struct {
	displayItemId      string
	Type               DisplayItemType
	salesItem          *SalesItem
	salesItemGroup     *SalesItemGroup
	salesPeriodEventId *string
}

func NewDisplayItem(
	displayItemId string,
	Type DisplayItemType,
	salesItem *SalesItem,
	salesItemGroup *SalesItemGroup,
	salesPeriodEventId *string,
) *DisplayItem {
	return &DisplayItem{
		displayItemId:      displayItemId,
		Type:               Type,
		salesItem:          salesItem,
		salesItemGroup:     salesItemGroup,
		salesPeriodEventId: salesPeriodEventId,
	}
}

func (p *DisplayItem) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["DisplayItemId"] = p.displayItemId
	properties["Type"] = p.Type
	if p.salesItem != nil {
		properties["SalesItem"] = p.salesItem.Properties()
	}
	if p.salesItemGroup != nil {
		properties["SalesItemGroup"] = p.salesItemGroup.Properties()
	}
	if p.salesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.salesPeriodEventId
	}
	return properties
}

func (p *DisplayItem) Ref(
	namespaceName string,
) DisplayItemRef {
	return DisplayItemRef{
		namespaceName: namespaceName,
	}
}

func NewSalesItemDisplayItem(
	displayItemId string,
	salesItem SalesItem,
	salesPeriodEventId *string,
) DisplayItem {
	return DisplayItem{
		Type:               DisplayItemTypeSalesItem,
		displayItemId:      displayItemId,
		salesItem:          &salesItem,
		salesPeriodEventId: salesPeriodEventId,
	}
}

func NewSalesItemGroupDisplayItem(
	displayItemId string,
	salesItemGroup SalesItemGroup,
	salesPeriodEventId *string,
) DisplayItem {
	return DisplayItem{
		Type:               DisplayItemTypeSalesItemGroup,
		displayItemId:      displayItemId,
		salesItemGroup:     &salesItemGroup,
		salesPeriodEventId: salesPeriodEventId,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack         *Stack
	version       string
	namespaceName string
	showcases     []Showcase
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	showcases []Showcase,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:         stack,
		version:       "2019-04-04",
		namespaceName: namespaceName,
		showcases:     showcases,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Showcase_CurrentShowcaseMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Showcase::CurrentShowcaseMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	showcases := make([]map[string]interface{}, len(p.showcases))
	for i, item := range p.showcases {
		showcases[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":   p.version,
			"Showcases": showcases,
		},
	}
}

type Namespace struct {
	CdkResource
	stack              *Stack
	name               string
	description        *string
	transactionSetting TransactionSetting
	buyScript          *ScriptSetting
	logSetting         *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	transactionSetting TransactionSetting,
	description *string,
	buyScript *ScriptSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:              stack,
		name:               name,
		description:        description,
		transactionSetting: transactionSetting,
		buyScript:          buyScript,
		logSetting:         logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Showcase_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Showcase::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["TransactionSetting"] = p.transactionSetting.Properties()
	if p.buyScript != nil {
		properties["BuyScript"] = p.buyScript.Properties()
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
	showcases []Showcase,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		showcases,
	).AddDependsOn(
		p,
	)
	return p
}

type SalesItemMaster struct {
	CdkResource
	stack          *Stack
	namespaceName  string
	name           string
	description    *string
	metadata       *string
	consumeActions []ConsumeAction
	acquireActions []AcquireAction
}

func NewSalesItemMaster(
	stack *Stack,
	namespaceName string,
	name string,
	consumeActions []ConsumeAction,
	acquireActions []AcquireAction,
	description *string,
	metadata *string,
) *SalesItemMaster {

	self := SalesItemMaster{
		stack:          stack,
		namespaceName:  namespaceName,
		name:           name,
		description:    description,
		metadata:       metadata,
		consumeActions: consumeActions,
		acquireActions: acquireActions,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *SalesItemMaster) ResourceName() string {
	return "Showcase_SalesItemMaster_" + p.name
}

func (p *SalesItemMaster) ResourceType() string {
	return "GS2::Showcase::SalesItemMaster"
}

func (p *SalesItemMaster) Properties() map[string]interface{} {
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
		values := make([]map[string]interface{}, len(p.consumeActions))
		for i, element := range p.consumeActions {
			values[i] = element.Properties()
		}
		properties["ConsumeActions"] = values
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

func (p *SalesItemMaster) Ref(
	namespaceName string,
) SalesItemMasterRef {
	return SalesItemMasterRef{
		namespaceName: namespaceName,
		salesItemName: p.name,
	}
}

func (p *SalesItemMaster) GetAttrSalesItemId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.SalesItemId",
	)
}

type SalesItemGroupMaster struct {
	CdkResource
	stack          *Stack
	namespaceName  string
	name           string
	description    *string
	metadata       *string
	salesItemNames []string
}

func NewSalesItemGroupMaster(
	stack *Stack,
	namespaceName string,
	name string,
	salesItemNames []string,
	description *string,
	metadata *string,
) *SalesItemGroupMaster {

	self := SalesItemGroupMaster{
		stack:          stack,
		namespaceName:  namespaceName,
		name:           name,
		description:    description,
		metadata:       metadata,
		salesItemNames: salesItemNames,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *SalesItemGroupMaster) ResourceName() string {
	return "Showcase_SalesItemGroupMaster_" + p.name
}

func (p *SalesItemGroupMaster) ResourceType() string {
	return "GS2::Showcase::SalesItemGroupMaster"
}

func (p *SalesItemGroupMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["SalesItemNames"] = p.salesItemNames
	return properties
}

func (p *SalesItemGroupMaster) Ref(
	namespaceName string,
) SalesItemGroupMasterRef {
	return SalesItemGroupMasterRef{
		namespaceName:      namespaceName,
		salesItemGroupName: p.name,
	}
}

func (p *SalesItemGroupMaster) GetAttrSalesItemGroupId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.SalesItemGroupId",
	)
}

type ShowcaseMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	description        *string
	metadata           *string
	displayItems       []DisplayItemMaster
	salesPeriodEventId *string
}

func NewShowcaseMaster(
	stack *Stack,
	namespaceName string,
	name string,
	displayItems []DisplayItemMaster,
	description *string,
	metadata *string,
	salesPeriodEventId *string,
) *ShowcaseMaster {

	self := ShowcaseMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		description:        description,
		metadata:           metadata,
		displayItems:       displayItems,
		salesPeriodEventId: salesPeriodEventId,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *ShowcaseMaster) ResourceName() string {
	return "Showcase_ShowcaseMaster_" + p.name
}

func (p *ShowcaseMaster) ResourceType() string {
	return "GS2::Showcase::ShowcaseMaster"
}

func (p *ShowcaseMaster) Properties() map[string]interface{} {
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
		values := make([]map[string]interface{}, len(p.displayItems))
		for i, element := range p.displayItems {
			values[i] = element.Properties()
		}
		properties["DisplayItems"] = values
	}
	if p.salesPeriodEventId != nil {
		properties["SalesPeriodEventId"] = p.salesPeriodEventId
	}
	return properties
}

func (p *ShowcaseMaster) Ref(
	namespaceName string,
) ShowcaseMasterRef {
	return ShowcaseMasterRef{
		namespaceName: namespaceName,
		showcaseName:  p.name,
	}
}

func (p *ShowcaseMaster) GetAttrShowcaseId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ShowcaseId",
	)
}
