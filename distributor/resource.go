package distributor

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

type DistributorModel struct {
	name               string
	metadata           *string
	inboxNamespaceId   *string
	whiteListTargetIds []string
}

func NewDistributorModel(
	name string,
	metadata *string,
	inboxNamespaceId *string,
	whiteListTargetIds []string,
) *DistributorModel {
	return &DistributorModel{
		name:               name,
		metadata:           metadata,
		inboxNamespaceId:   inboxNamespaceId,
		whiteListTargetIds: whiteListTargetIds,
	}
}

func (p *DistributorModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.inboxNamespaceId != nil {
		properties["InboxNamespaceId"] = p.inboxNamespaceId
	}
	properties["WhiteListTargetIds"] = p.whiteListTargetIds
	return properties
}

func (p *DistributorModel) Ref(
	namespaceName string,
) DistributorModelRef {
	return DistributorModelRef{
		namespaceName:   namespaceName,
		distributorName: p.name,
	}
}

type StampSheetResult struct {
	userId            string
	transactionId     string
	taskRequests      []ConsumeAction
	sheetRequest      AcquireAction
	taskResults       []string
	sheetResult       *string
	nextTransactionId *string
	createdAt         int64
	ttlAt             int64
}

func NewStampSheetResult(
	userId string,
	transactionId string,
	taskRequests []ConsumeAction,
	sheetRequest AcquireAction,
	taskResults []string,
	sheetResult *string,
	nextTransactionId *string,
	createdAt int64,
	ttlAt int64,
) *StampSheetResult {
	return &StampSheetResult{
		userId:            userId,
		transactionId:     transactionId,
		taskRequests:      taskRequests,
		sheetRequest:      sheetRequest,
		taskResults:       taskResults,
		sheetResult:       sheetResult,
		nextTransactionId: nextTransactionId,
		createdAt:         createdAt,
		ttlAt:             ttlAt,
	}
}

func (p *StampSheetResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.userId
	properties["TransactionId"] = p.transactionId
	{
		values := make([]map[string]interface{}, len(p.taskRequests))
		for i, element := range p.taskRequests {
			values[i] = element.Properties()
		}
		properties["TaskRequests"] = values
	}
	properties["SheetRequest"] = p.sheetRequest.Properties()
	properties["TaskResults"] = p.taskResults
	if p.sheetResult != nil {
		properties["SheetResult"] = p.sheetResult
	}
	if p.nextTransactionId != nil {
		properties["NextTransactionId"] = p.nextTransactionId
	}
	properties["CreatedAt"] = p.createdAt
	properties["TtlAt"] = p.ttlAt
	return properties
}

type CurrentMasterData struct {
	CdkResource
	stack             *Stack
	version           string
	namespaceName     string
	distributorModels []DistributorModel
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	distributorModels []DistributorModel,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:             stack,
		version:           "2019-03-01",
		namespaceName:     namespaceName,
		distributorModels: distributorModels,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Distributor_CurrentDistributorMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Distributor::CurrentDistributorMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	distributorModels := make([]map[string]interface{}, len(p.distributorModels))
	for i, item := range p.distributorModels {
		distributorModels[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":           p.version,
			"DistributorModels": distributorModels,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                         *Stack
	name                          string
	description                   *string
	assumeUserId                  *string
	autoRunStampSheetNotification *NotificationSetting
	logSetting                    *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	assumeUserId *string,
	autoRunStampSheetNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                         stack,
		name:                          name,
		description:                   description,
		assumeUserId:                  assumeUserId,
		autoRunStampSheetNotification: autoRunStampSheetNotification,
		logSetting:                    logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Distributor_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Distributor::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.assumeUserId != nil {
		properties["AssumeUserId"] = p.assumeUserId
	}
	if p.autoRunStampSheetNotification != nil {
		properties["AutoRunStampSheetNotification"] = p.autoRunStampSheetNotification.Properties()
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
	distributorModels []DistributorModel,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		distributorModels,
	).AddDependsOn(
		p,
	)
	return p
}

type DistributorModelMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	description        *string
	metadata           *string
	inboxNamespaceId   *string
	whiteListTargetIds []string
}

func NewDistributorModelMaster(
	stack *Stack,
	namespaceName string,
	name string,
	whiteListTargetIds []string,
	description *string,
	metadata *string,
	inboxNamespaceId *string,
) *DistributorModelMaster {

	self := DistributorModelMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		description:        description,
		metadata:           metadata,
		inboxNamespaceId:   inboxNamespaceId,
		whiteListTargetIds: whiteListTargetIds,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *DistributorModelMaster) ResourceName() string {
	return "Distributor_DistributorModelMaster_" + p.name
}

func (p *DistributorModelMaster) ResourceType() string {
	return "GS2::Distributor::DistributorModelMaster"
}

func (p *DistributorModelMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	if p.inboxNamespaceId != nil {
		properties["InboxNamespaceId"] = p.inboxNamespaceId
	}
	properties["WhiteListTargetIds"] = p.whiteListTargetIds
	return properties
}

func (p *DistributorModelMaster) Ref(
	namespaceName string,
) DistributorModelMasterRef {
	return DistributorModelMasterRef{
		namespaceName:   namespaceName,
		distributorName: p.name,
	}
}

func (p *DistributorModelMaster) GetAttrDistributorModelId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.DistributorModelId",
	)
}
