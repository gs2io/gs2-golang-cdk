package money2

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

type StoreSubscriptionContentModelTriggerExtendMode string

const StoreSubscriptionContentModelTriggerExtendModeJust = StoreSubscriptionContentModelTriggerExtendMode("just")
const StoreSubscriptionContentModelTriggerExtendModeRollupHour = StoreSubscriptionContentModelTriggerExtendMode("rollupHour")

func (p StoreSubscriptionContentModelTriggerExtendMode) Pointer() *StoreSubscriptionContentModelTriggerExtendMode {
	return &p
}

type StoreSubscriptionContentModel struct {
	Name                string
	Metadata            *string
	ScheduleNamespaceId string
	TriggerName         string
	TriggerExtendMode   StoreSubscriptionContentModelTriggerExtendMode
	RollupHour          *int32
	ReallocateSpanDays  int32
	AppleAppStore       *AppleAppStoreSubscriptionContent
	GooglePlay          *GooglePlaySubscriptionContent
}

type StoreSubscriptionContentModelOptions struct {
	Metadata      *string
	RollupHour    *int32
	AppleAppStore *AppleAppStoreSubscriptionContent
	GooglePlay    *GooglePlaySubscriptionContent
}

func NewStoreSubscriptionContentModel(
	name string,
	scheduleNamespaceId string,
	triggerName string,
	triggerExtendMode StoreSubscriptionContentModelTriggerExtendMode,
	reallocateSpanDays int32,
	options StoreSubscriptionContentModelOptions,
) StoreSubscriptionContentModel {
	_data := StoreSubscriptionContentModel{
		Name:                name,
		ScheduleNamespaceId: scheduleNamespaceId,
		TriggerName:         triggerName,
		TriggerExtendMode:   triggerExtendMode,
		ReallocateSpanDays:  reallocateSpanDays,
		Metadata:            options.Metadata,
		RollupHour:          options.RollupHour,
		AppleAppStore:       options.AppleAppStore,
		GooglePlay:          options.GooglePlay,
	}
	return _data
}

type StoreSubscriptionContentModelTriggerExtendModeIsJustOptions struct {
	Metadata      *string
	AppleAppStore *AppleAppStoreSubscriptionContent
	GooglePlay    *GooglePlaySubscriptionContent
}

func NewStoreSubscriptionContentModelTriggerExtendModeIsJust(
	name string,
	scheduleNamespaceId string,
	triggerName string,
	reallocateSpanDays int32,
	options StoreSubscriptionContentModelTriggerExtendModeIsJustOptions,
) StoreSubscriptionContentModel {
	return NewStoreSubscriptionContentModel(
		name,
		scheduleNamespaceId,
		triggerName,
		StoreSubscriptionContentModelTriggerExtendModeJust,
		reallocateSpanDays,
		StoreSubscriptionContentModelOptions{
			Metadata:      options.Metadata,
			AppleAppStore: options.AppleAppStore,
			GooglePlay:    options.GooglePlay,
		},
	)
}

type StoreSubscriptionContentModelTriggerExtendModeIsRollupHourOptions struct {
	Metadata      *string
	AppleAppStore *AppleAppStoreSubscriptionContent
	GooglePlay    *GooglePlaySubscriptionContent
}

func NewStoreSubscriptionContentModelTriggerExtendModeIsRollupHour(
	name string,
	scheduleNamespaceId string,
	triggerName string,
	reallocateSpanDays int32,
	rollupHour int32,
	options StoreSubscriptionContentModelTriggerExtendModeIsRollupHourOptions,
) StoreSubscriptionContentModel {
	return NewStoreSubscriptionContentModel(
		name,
		scheduleNamespaceId,
		triggerName,
		StoreSubscriptionContentModelTriggerExtendModeRollupHour,
		reallocateSpanDays,
		StoreSubscriptionContentModelOptions{
			Metadata:      options.Metadata,
			RollupHour:    &rollupHour,
			AppleAppStore: options.AppleAppStore,
			GooglePlay:    options.GooglePlay,
		},
	)
}

func (p *StoreSubscriptionContentModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["ScheduleNamespaceId"] = p.ScheduleNamespaceId
	properties["TriggerName"] = p.TriggerName
	properties["TriggerExtendMode"] = p.TriggerExtendMode
	if p.RollupHour != nil {
		properties["RollupHour"] = p.RollupHour
	}
	properties["ReallocateSpanDays"] = p.ReallocateSpanDays
	if p.AppleAppStore != nil {
		properties["AppleAppStore"] = p.AppleAppStore.Properties()
	}
	if p.GooglePlay != nil {
		properties["GooglePlay"] = p.GooglePlay.Properties()
	}
	return properties
}
