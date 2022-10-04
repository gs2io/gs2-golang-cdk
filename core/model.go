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
	loggingNamespaceId string
}

func NewLogSetting(
	loggingNamespaceId string,
) LogSetting {
	return LogSetting{
		loggingNamespaceId: loggingNamespaceId,
	}
}

func (p LogSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["LoggingNamespaceId"] = p.loggingNamespaceId
	return properties
}

type NotificationSetting struct {
	gatewayNamespaceId               string
	enableTransferMobileNotification *bool
	sound                            *string
}

func NewNotificationSetting(
	gatewayNamespaceId string,
	enableTransferMobileNotification *bool,
	sound *string,
) NotificationSetting {
	return NotificationSetting{
		gatewayNamespaceId:               gatewayNamespaceId,
		enableTransferMobileNotification: enableTransferMobileNotification,
		sound:                            sound,
	}
}

func (p NotificationSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GatewayNamespaceId"] = p.gatewayNamespaceId
	if p.enableTransferMobileNotification != nil {
		properties["EnableTransferMobileNotification"] = *p.enableTransferMobileNotification
	}
	if p.sound != nil {
		properties["Sound"] = *p.sound
	}
	return properties
}

type ScriptSettingDoneTriggerTargetType string

const ScriptSettingDoneTriggerTargetTypeNone = ScriptSettingDoneTriggerTargetType("none")
const ScriptSettingDoneTriggerTargetTypeGs2Script = ScriptSettingDoneTriggerTargetType("gs2_script")
const ScriptSettingDoneTriggerTargetTypeAws = ScriptSettingDoneTriggerTargetType("aws")

type ScriptSetting struct {
	triggerScriptId             *string
	doneTriggerTargetType       ScriptSettingDoneTriggerTargetType
	doneTriggerScriptId         *string
	doneTriggerQueueNamespaceId *string
}

func NewScriptSetting(
	triggerScriptId *string,
	doneTriggerTargetType ScriptSettingDoneTriggerTargetType,
	doneTriggerScriptId *string,
	doneTriggerQueueNamespaceId *string,
) ScriptSetting {
	return ScriptSetting{
		triggerScriptId:             triggerScriptId,
		doneTriggerTargetType:       doneTriggerTargetType,
		doneTriggerScriptId:         doneTriggerScriptId,
		doneTriggerQueueNamespaceId: doneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.triggerScriptId != nil {
		properties["TriggerScriptId"] = *p.triggerScriptId
	}
	properties["DoneTriggerTargetType"] = p.doneTriggerTargetType
	if p.doneTriggerScriptId != nil {
		properties["DoneTriggerScriptId"] = *p.doneTriggerScriptId
	}
	if p.doneTriggerQueueNamespaceId != nil {
		properties["DoneTriggerQueueNamespaceId"] = *p.doneTriggerQueueNamespaceId
	}
	return properties
}

type TransactionSetting struct {
	enableAutoRun          bool
	distributorNamespaceId *string
	keyId                  *string
	queueNamespaceId       *string
}

func NewTransactionSetting(
	enableAutoRun bool,
	distributorNamespaceId *string,
	keyId *string,
	queueNamespaceId *string,
) TransactionSetting {
	return TransactionSetting{
		enableAutoRun:          enableAutoRun,
		distributorNamespaceId: distributorNamespaceId,
		keyId:                  keyId,
		queueNamespaceId:       queueNamespaceId,
	}
}

func (p TransactionSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["EnableAutoRun"] = p.enableAutoRun
	if p.distributorNamespaceId != nil {
		properties["DistributorNamespaceId"] = *p.distributorNamespaceId
	}
	if p.keyId != nil {
		properties["KeyId"] = *p.keyId
	}
	if p.queueNamespaceId != nil {
		properties["QueueNamespaceId"] = *p.queueNamespaceId
	}
	return properties
}
