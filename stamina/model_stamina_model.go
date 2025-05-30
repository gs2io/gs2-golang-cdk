package stamina

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

type StaminaModel struct {
	Name                   string
	Metadata               *string
	RecoverIntervalMinutes int32
	RecoverValue           int32
	InitialCapacity        int32
	IsOverflow             bool
	MaxCapacity            *int32
	MaxStaminaTable        *MaxStaminaTable
	RecoverIntervalTable   *RecoverIntervalTable
	RecoverValueTable      *RecoverValueTable
}

type StaminaModelOptions struct {
	Metadata             *string
	MaxCapacity          *int32
	MaxStaminaTable      *MaxStaminaTable
	RecoverIntervalTable *RecoverIntervalTable
	RecoverValueTable    *RecoverValueTable
}

func NewStaminaModel(
	name string,
	recoverIntervalMinutes int32,
	recoverValue int32,
	initialCapacity int32,
	isOverflow bool,
	options StaminaModelOptions,
) StaminaModel {
	_data := StaminaModel{
		Name:                   name,
		RecoverIntervalMinutes: recoverIntervalMinutes,
		RecoverValue:           recoverValue,
		InitialCapacity:        initialCapacity,
		IsOverflow:             isOverflow,
		Metadata:               options.Metadata,
		MaxCapacity:            options.MaxCapacity,
		MaxStaminaTable:        options.MaxStaminaTable,
		RecoverIntervalTable:   options.RecoverIntervalTable,
		RecoverValueTable:      options.RecoverValueTable,
	}
	return _data
}

func (p *StaminaModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["RecoverIntervalMinutes"] = p.RecoverIntervalMinutes
	properties["RecoverValue"] = p.RecoverValue
	properties["InitialCapacity"] = p.InitialCapacity
	properties["IsOverflow"] = p.IsOverflow
	if p.MaxCapacity != nil {
		properties["MaxCapacity"] = p.MaxCapacity
	}
	if p.MaxStaminaTable != nil {
		properties["MaxStaminaTable"] = p.MaxStaminaTable.Properties()
	}
	if p.RecoverIntervalTable != nil {
		properties["RecoverIntervalTable"] = p.RecoverIntervalTable.Properties()
	}
	if p.RecoverValueTable != nil {
		properties["RecoverValueTable"] = p.RecoverValueTable.Properties()
	}
	return properties
}
