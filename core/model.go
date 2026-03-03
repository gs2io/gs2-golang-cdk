package core

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

type AcquireAction struct {
	Action  string
	Request map[string]interface{}
}

func (p AcquireAction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["action"] = p.Action
	if p.Request != nil {
		properties["request"] = p.Request
	}
	return properties
}

type ConsumeAction struct {
	Action  string
	Request map[string]interface{}
}

func (p ConsumeAction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["action"] = p.Action
	if p.Request != nil {
		properties["request"] = p.Request
	}
	return properties
}

type VerifyAction struct {
	Action  string
	Request map[string]interface{}
}

func (p VerifyAction) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["action"] = p.Action
	if p.Request != nil {
		properties["request"] = p.Request
	}
	return properties
}

type Config struct {
	Key   string
	Value string
}

func (p Config) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["key"] = p.Key
	properties["value"] = p.Value
	return properties
}

type LogSetting struct {
	LoggingNamespaceId string
}

func NewLogSetting(
	loggingNamespaceId string,
) LogSetting {
	return LogSetting{
		LoggingNamespaceId: loggingNamespaceId,
	}
}

func (p LogSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["LoggingNamespaceId"] = p.LoggingNamespaceId
	return properties
}

type NotificationSettingEnable string

const NotificationSettingEnableEnabled = NotificationSettingEnable("Enabled")
const NotificationSettingEnableDisabled = NotificationSettingEnable("Disabled")

func (p NotificationSettingEnable) Pointer() *NotificationSettingEnable {
	return &p
}

type NotificationSettingOptions struct {
	GatewayNamespaceId               *string
	EnableTransferMobileNotification *bool
	Sound                            *string
	Enable                           *NotificationSettingEnable
}

type NotificationSetting struct {
	GatewayNamespaceId               *string
	EnableTransferMobileNotification *bool
	Sound                            *string
	Enable                           *NotificationSettingEnable
}

func NewNotificationSetting(
	options NotificationSettingOptions,
) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               options.GatewayNamespaceId,
		EnableTransferMobileNotification: options.EnableTransferMobileNotification,
		Sound:                            options.Sound,
		Enable:                           options.Enable,
	}
}

func (p NotificationSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.GatewayNamespaceId != nil {
		properties["GatewayNamespaceId"] = *p.GatewayNamespaceId
	}
	if p.EnableTransferMobileNotification != nil {
		properties["EnableTransferMobileNotification"] = *p.EnableTransferMobileNotification
	}
	if p.Sound != nil {
		properties["Sound"] = *p.Sound
	}
	if p.Enable != nil {
		properties["Enable"] = *p.Enable
	}
	return properties
}

type ScriptSettingDoneTriggerTargetType string

const ScriptSettingDoneTriggerTargetTypeNone = ScriptSettingDoneTriggerTargetType("none")
const ScriptSettingDoneTriggerTargetTypeGs2Script = ScriptSettingDoneTriggerTargetType("gs2_script")
const ScriptSettingDoneTriggerTargetTypeAws = ScriptSettingDoneTriggerTargetType("aws")

func (p ScriptSettingDoneTriggerTargetType) Pointer() *ScriptSettingDoneTriggerTargetType {
	return &p
}

type ScriptSettingOptions struct {
	TriggerScriptId             *string
	DoneTriggerTargetType       *ScriptSettingDoneTriggerTargetType
	DoneTriggerScriptId         *string
	DoneTriggerQueueNamespaceId *string
}

type ScriptSetting struct {
	TriggerScriptId             *string
	DoneTriggerTargetType       *ScriptSettingDoneTriggerTargetType
	DoneTriggerScriptId         *string
	DoneTriggerQueueNamespaceId *string
}

func NewScriptSetting(
	options ScriptSettingOptions,
) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId:             options.TriggerScriptId,
		DoneTriggerTargetType:       options.DoneTriggerTargetType,
		DoneTriggerScriptId:         options.DoneTriggerScriptId,
		DoneTriggerQueueNamespaceId: options.DoneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		properties["TriggerScriptId"] = *p.TriggerScriptId
	}
	if p.DoneTriggerTargetType != nil {
		properties["DoneTriggerTargetType"] = *p.DoneTriggerTargetType
	}
	if p.DoneTriggerScriptId != nil {
		properties["DoneTriggerScriptId"] = *p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		properties["DoneTriggerQueueNamespaceId"] = *p.DoneTriggerQueueNamespaceId
	}
	return properties
}

type TransactionSettingOptions struct {
	EnableAtomicCommit        *bool
	TransactionUseDistributor *bool
	AcquireActionUseJobQueue  *bool
	DistributorNamespaceId    *string
	QueueNamespaceId          *string
}

type TransactionSetting struct {
	EnableAtomicCommit        *bool
	TransactionUseDistributor *bool
	AcquireActionUseJobQueue  *bool
	DistributorNamespaceId    *string
	QueueNamespaceId          *string
}

func NewTransactionSetting(
	options TransactionSettingOptions,
) TransactionSetting {
	return TransactionSetting{
		EnableAtomicCommit:        options.EnableAtomicCommit,
		TransactionUseDistributor: options.TransactionUseDistributor,
		AcquireActionUseJobQueue:  options.AcquireActionUseJobQueue,
		DistributorNamespaceId:    options.DistributorNamespaceId,
		QueueNamespaceId:          options.QueueNamespaceId,
	}
}

func (p TransactionSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["EnableAutoRun"] = true
	if p.EnableAtomicCommit != nil {
		properties["EnableAtomicCommit"] = *p.EnableAtomicCommit
	}
	if p.TransactionUseDistributor != nil {
		properties["TransactionUseDistributor"] = *p.TransactionUseDistributor
	}
	if p.AcquireActionUseJobQueue != nil {
		properties["AcquireActionUseJobQueue"] = *p.AcquireActionUseJobQueue
	}
	if p.DistributorNamespaceId != nil {
		properties["DistributorNamespaceId"] = *p.DistributorNamespaceId
	}
	if p.QueueNamespaceId != nil {
		properties["QueueNamespaceId"] = *p.QueueNamespaceId
	}
	return properties
}
