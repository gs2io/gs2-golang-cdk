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

type NotificationSetting struct {
	GatewayNamespaceId               string
	EnableTransferMobileNotification *bool
	Sound                            *string
}

func NewNotificationSetting(
	gatewayNamespaceId string,
	enableTransferMobileNotification *bool,
	sound *string,
) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               gatewayNamespaceId,
		EnableTransferMobileNotification: enableTransferMobileNotification,
		Sound:                            sound,
	}
}

func (p NotificationSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GatewayNamespaceId"] = p.GatewayNamespaceId
	if p.EnableTransferMobileNotification != nil {
		properties["EnableTransferMobileNotification"] = *p.EnableTransferMobileNotification
	}
	if p.Sound != nil {
		properties["Sound"] = *p.Sound
	}
	return properties
}

type ScriptSettingDoneTriggerTargetType string

const ScriptSettingDoneTriggerTargetTypeNone = ScriptSettingDoneTriggerTargetType("none")
const ScriptSettingDoneTriggerTargetTypeGs2Script = ScriptSettingDoneTriggerTargetType("gs2_script")
const ScriptSettingDoneTriggerTargetTypeAws = ScriptSettingDoneTriggerTargetType("aws")

type ScriptSetting struct {
	TriggerScriptId             *string
	DoneTriggerTargetType       ScriptSettingDoneTriggerTargetType
	DoneTriggerScriptId         *string
	DoneTriggerQueueNamespaceId *string
}

func NewScriptSetting(
	triggerScriptId *string,
	doneTriggerTargetType ScriptSettingDoneTriggerTargetType,
	doneTriggerScriptId *string,
	doneTriggerQueueNamespaceId *string,
) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId:             triggerScriptId,
		DoneTriggerTargetType:       doneTriggerTargetType,
		DoneTriggerScriptId:         doneTriggerScriptId,
		DoneTriggerQueueNamespaceId: doneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		properties["TriggerScriptId"] = *p.TriggerScriptId
	}
	properties["DoneTriggerTargetType"] = p.DoneTriggerTargetType
	if p.DoneTriggerScriptId != nil {
		properties["DoneTriggerScriptId"] = *p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		properties["DoneTriggerQueueNamespaceId"] = *p.DoneTriggerQueueNamespaceId
	}
	return properties
}

type TransactionSetting struct {
	EnableAutoRun          bool
	DistributorNamespaceId *string
	KeyId                  *string
	QueueNamespaceId       *string
}

func NewTransactionSetting(
	enableAutoRun bool,
	distributorNamespaceId *string,
	keyId *string,
	queueNamespaceId *string,
) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun:          enableAutoRun,
		DistributorNamespaceId: distributorNamespaceId,
		KeyId:                  keyId,
		QueueNamespaceId:       queueNamespaceId,
	}
}

func (p TransactionSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["EnableAutoRun"] = p.EnableAutoRun
	if p.DistributorNamespaceId != nil {
		properties["DistributorNamespaceId"] = *p.DistributorNamespaceId
	}
	if p.KeyId != nil {
		properties["KeyId"] = *p.KeyId
	}
	if p.QueueNamespaceId != nil {
		properties["QueueNamespaceId"] = *p.QueueNamespaceId
	}
	return properties
}
