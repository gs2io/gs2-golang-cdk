package adReward

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

type Namespace struct {
	CdkResource
	stack                   *Stack
	Name                    string
	Description             *string
	TransactionSetting      *TransactionSetting
	Admob                   *AdMob
	UnityAd                 *UnityAd
	AppLovinMaxes           []AppLovinMax
	AcquirePointScript      *ScriptSetting
	ConsumePointScript      *ScriptSetting
	ChangePointNotification NotificationSetting
	LogSetting              *LogSetting
}

type NamespaceOptions struct {
	Description             *string
	TransactionSetting      *TransactionSetting
	Admob                   *AdMob
	UnityAd                 *UnityAd
	AppLovinMaxes           []AppLovinMax
	AcquirePointScript      *ScriptSetting
	ConsumePointScript      *ScriptSetting
	ChangePointNotification NotificationSetting
	LogSetting              *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                   stack,
		Name:                    name,
		Description:             options.Description,
		TransactionSetting:      options.TransactionSetting,
		Admob:                   options.Admob,
		UnityAd:                 options.UnityAd,
		AppLovinMaxes:           options.AppLovinMaxes,
		AcquirePointScript:      options.AcquirePointScript,
		ConsumePointScript:      options.ConsumePointScript,
		ChangePointNotification: options.ChangePointNotification,
		LogSetting:              options.LogSetting,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "AdReward_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::AdReward::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	if p.TransactionSetting != nil {
		properties["TransactionSetting"] = p.TransactionSetting.Properties()
	}
	if p.Admob != nil {
		properties["Admob"] = p.Admob.Properties()
	}
	if p.UnityAd != nil {
		properties["UnityAd"] = p.UnityAd.Properties()
	}
	{
		values := make([]map[string]interface{}, len(p.AppLovinMaxes))
		for i, element := range p.AppLovinMaxes {
			values[i] = element.Properties()
		}
		properties["AppLovinMaxes"] = values
	}
	if p.AcquirePointScript != nil {
		properties["AcquirePointScript"] = p.AcquirePointScript.Properties()
	}
	if p.ConsumePointScript != nil {
		properties["ConsumePointScript"] = p.ConsumePointScript.Properties()
	}
	properties["ChangePointNotification"] = p.ChangePointNotification.Properties()
	if p.LogSetting != nil {
		properties["LogSetting"] = p.LogSetting.Properties()
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		NamespaceName: p.Name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}
