package inbox

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

type GlobalMessage struct {
	name               string
	metadata           string
	readAcquireActions []AcquireAction
	expiresTimeSpan    *TimeSpan
	expiresAt          *int64
}

func NewGlobalMessage(
	name string,
	metadata string,
	readAcquireActions []AcquireAction,
	expiresTimeSpan *TimeSpan,
	expiresAt *int64,
) *GlobalMessage {
	return &GlobalMessage{
		name:               name,
		metadata:           metadata,
		readAcquireActions: readAcquireActions,
		expiresTimeSpan:    expiresTimeSpan,
		expiresAt:          expiresAt,
	}
}

func (p *GlobalMessage) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["Metadata"] = p.metadata
	{
		values := make([]map[string]interface{}, len(p.readAcquireActions))
		for i, element := range p.readAcquireActions {
			values[i] = element.Properties()
		}
		properties["ReadAcquireActions"] = values
	}
	if p.expiresTimeSpan != nil {
		properties["ExpiresTimeSpan"] = p.expiresTimeSpan.Properties()
	}
	if p.expiresAt != nil {
		properties["ExpiresAt"] = p.expiresAt
	}
	return properties
}

func (p *GlobalMessage) Ref(
	namespaceName string,
) GlobalMessageRef {
	return GlobalMessageRef{
		namespaceName:     namespaceName,
		globalMessageName: p.name,
	}
}

type CurrentMasterData struct {
	CdkResource
	stack          *Stack
	version        string
	namespaceName  string
	globalMessages []GlobalMessage
}

func NewCurrentMasterData(
	stack *Stack,
	namespaceName string,
	globalMessages []GlobalMessage,
) *CurrentMasterData {
	self := CurrentMasterData{
		stack:          stack,
		version:        "2020-03-12",
		namespaceName:  namespaceName,
		globalMessages: globalMessages,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *CurrentMasterData) ResourceName() string {
	return "Inbox_CurrentMessageMaster_" + p.namespaceName
}

func (p *CurrentMasterData) ResourceType() string {
	return "GS2::Inbox::CurrentMessageMaster"
}

func (p *CurrentMasterData) Properties() map[string]interface{} {
	globalMessages := make([]map[string]interface{}, len(p.globalMessages))
	for i, item := range p.globalMessages {
		globalMessages[i] = item.Properties()
	}
	return map[string]interface{}{
		"NamespaceName": p.namespaceName,
		"Settings": map[string]interface{}{
			"version":        p.version,
			"GlobalMessages": globalMessages,
		},
	}
}

type Namespace struct {
	CdkResource
	stack                      *Stack
	name                       string
	description                *string
	isAutomaticDeletingEnabled *bool
	transactionSetting         *TransactionSetting
	receiveMessageScript       *ScriptSetting
	readMessageScript          *ScriptSetting
	deleteMessageScript        *ScriptSetting
	receiveNotification        *NotificationSetting
	logSetting                 *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	isAutomaticDeletingEnabled *bool,
	transactionSetting *TransactionSetting,
	receiveMessageScript *ScriptSetting,
	readMessageScript *ScriptSetting,
	deleteMessageScript *ScriptSetting,
	receiveNotification *NotificationSetting,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:                      stack,
		name:                       name,
		description:                description,
		isAutomaticDeletingEnabled: isAutomaticDeletingEnabled,
		transactionSetting:         transactionSetting,
		receiveMessageScript:       receiveMessageScript,
		readMessageScript:          readMessageScript,
		deleteMessageScript:        deleteMessageScript,
		receiveNotification:        receiveNotification,
		logSetting:                 logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Inbox_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Inbox::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.isAutomaticDeletingEnabled != nil {
		properties["IsAutomaticDeletingEnabled"] = p.isAutomaticDeletingEnabled
	}
	if p.transactionSetting != nil {
		properties["TransactionSetting"] = p.transactionSetting.Properties()
	}
	if p.receiveMessageScript != nil {
		properties["ReceiveMessageScript"] = p.receiveMessageScript.Properties()
	}
	if p.readMessageScript != nil {
		properties["ReadMessageScript"] = p.readMessageScript.Properties()
	}
	if p.deleteMessageScript != nil {
		properties["DeleteMessageScript"] = p.deleteMessageScript.Properties()
	}
	if p.receiveNotification != nil {
		properties["ReceiveNotification"] = p.receiveNotification.Properties()
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
	globalMessages []GlobalMessage,

) *Namespace {
	NewCurrentMasterData(
		p.stack,
		p.name,
		globalMessages,
	).AddDependsOn(
		p,
	)
	return p
}

type GlobalMessageMaster struct {
	CdkResource
	stack              *Stack
	namespaceName      string
	name               string
	metadata           string
	readAcquireActions []AcquireAction
	expiresTimeSpan    *TimeSpan
	expiresAt          *int64
}

func NewGlobalMessageMaster(
	stack *Stack,
	namespaceName string,
	name string,
	metadata string,
	readAcquireActions []AcquireAction,
	expiresTimeSpan *TimeSpan,
	expiresAt *int64,
) *GlobalMessageMaster {

	self := GlobalMessageMaster{
		stack:              stack,
		namespaceName:      namespaceName,
		name:               name,
		metadata:           metadata,
		readAcquireActions: readAcquireActions,
		expiresTimeSpan:    expiresTimeSpan,
		expiresAt:          expiresAt,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *GlobalMessageMaster) ResourceName() string {
	return "Inbox_GlobalMessageMaster_" + p.name
}

func (p *GlobalMessageMaster) ResourceType() string {
	return "GS2::Inbox::GlobalMessageMaster"
}

func (p *GlobalMessageMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	properties["Metadata"] = p.metadata
	{
		values := make([]map[string]interface{}, len(p.readAcquireActions))
		for i, element := range p.readAcquireActions {
			values[i] = element.Properties()
		}
		properties["ReadAcquireActions"] = values
	}
	if p.expiresTimeSpan != nil {
		properties["ExpiresTimeSpan"] = p.expiresTimeSpan.Properties()
	}
	if p.expiresAt != nil {
		properties["ExpiresAt"] = p.expiresAt
	}
	return properties
}

func (p *GlobalMessageMaster) Ref(
	namespaceName string,
) GlobalMessageMasterRef {
	return GlobalMessageMasterRef{
		namespaceName:     namespaceName,
		globalMessageName: p.name,
	}
}

func (p *GlobalMessageMaster) GetAttrGlobalMessageId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.GlobalMessageId",
	)
}
