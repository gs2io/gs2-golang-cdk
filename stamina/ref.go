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

type NamespaceRef struct {
	namespaceName string
}

func (p *NamespaceRef) CurrentStaminaMaster() *CurrentStaminaMasterRef {
	return &CurrentStaminaMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) MaxStaminaTable(
	maxStaminaTableName string,
) *MaxStaminaTableRef {
	return &MaxStaminaTableRef{
		namespaceName:       p.namespaceName,
		maxStaminaTableName: maxStaminaTableName,
	}
}

func (p *NamespaceRef) RecoverIntervalTable(
	recoverIntervalTableName string,
) *RecoverIntervalTableRef {
	return &RecoverIntervalTableRef{
		namespaceName:            p.namespaceName,
		recoverIntervalTableName: recoverIntervalTableName,
	}
}

func (p *NamespaceRef) RecoverValueTable(
	recoverValueTableName string,
) *RecoverValueTableRef {
	return &RecoverValueTableRef{
		namespaceName:         p.namespaceName,
		recoverValueTableName: recoverValueTableName,
	}
}

func (p *NamespaceRef) StaminaModel(
	staminaName string,
) *StaminaModelRef {
	return &StaminaModelRef{
		namespaceName: p.namespaceName,
		staminaName:   staminaName,
	}
}

func (p *NamespaceRef) RecoverIntervalTableMaster(
	recoverIntervalTableName string,
) *RecoverIntervalTableMasterRef {
	return &RecoverIntervalTableMasterRef{
		namespaceName:            p.namespaceName,
		recoverIntervalTableName: recoverIntervalTableName,
	}
}

func (p *NamespaceRef) MaxStaminaTableMaster(
	maxStaminaTableName string,
) *MaxStaminaTableMasterRef {
	return &MaxStaminaTableMasterRef{
		namespaceName:       p.namespaceName,
		maxStaminaTableName: maxStaminaTableName,
	}
}

func (p *NamespaceRef) RecoverValueTableMaster(
	recoverValueTableName string,
) *RecoverValueTableMasterRef {
	return &RecoverValueTableMasterRef{
		namespaceName:         p.namespaceName,
		recoverValueTableName: recoverValueTableName,
	}
}

func (p *NamespaceRef) StaminaModelMaster(
	staminaName string,
) *StaminaModelMasterRef {
	return &StaminaModelMasterRef{
		namespaceName: p.namespaceName,
		staminaName:   staminaName,
	}
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type StaminaModelMasterRef struct {
	namespaceName string
	staminaName   string
}

func (p *StaminaModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"model",
			p.staminaName,
		},
	).String()
}

func (p *StaminaModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MaxStaminaTableMasterRef struct {
	namespaceName       string
	maxStaminaTableName string
}

func (p *MaxStaminaTableMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"maxStaminaTable",
			p.maxStaminaTableName,
		},
	).String()
}

func (p *MaxStaminaTableMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RecoverIntervalTableMasterRef struct {
	namespaceName            string
	recoverIntervalTableName string
}

func (p *RecoverIntervalTableMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"recoverIntervalTable",
			p.recoverIntervalTableName,
		},
	).String()
}

func (p *RecoverIntervalTableMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RecoverValueTableMasterRef struct {
	namespaceName         string
	recoverValueTableName string
}

func (p *RecoverValueTableMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"recoverValueTable",
			p.recoverValueTableName,
		},
	).String()
}

func (p *RecoverValueTableMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentStaminaMasterRef struct {
	namespaceName string
}

func (p *CurrentStaminaMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentStaminaMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type StaminaModelRef struct {
	namespaceName string
	staminaName   string
}

func (p *StaminaModelRef) RecoverStamina(
	recoverValue int32,
) AcquireAction {
	return RecoverStaminaByUserId(
		p.namespaceName,
		p.staminaName,
		recoverValue,
	)
}

func (p *StaminaModelRef) RaiseMaxValue(
	raiseValue int32,
) AcquireAction {
	return RaiseMaxValueByUserId(
		p.namespaceName,
		p.staminaName,
		raiseValue,
	)
}

func (p *StaminaModelRef) SetMaxValue(
	maxValue int32,
) AcquireAction {
	return SetMaxValueByUserId(
		p.namespaceName,
		p.staminaName,
		maxValue,
	)
}

func (p *StaminaModelRef) SetRecoverInterval(
	recoverIntervalMinutes int32,
) AcquireAction {
	return SetRecoverIntervalByUserId(
		p.namespaceName,
		p.staminaName,
		recoverIntervalMinutes,
	)
}

func (p *StaminaModelRef) SetRecoverValue(
	recoverValue int32,
) AcquireAction {
	return SetRecoverValueByUserId(
		p.namespaceName,
		p.staminaName,
		recoverValue,
	)
}

func (p *StaminaModelRef) ConsumeStamina(
	consumeValue int32,
) ConsumeAction {
	return ConsumeStaminaByUserId(
		p.namespaceName,
		p.staminaName,
		consumeValue,
	)
}

func (p *StaminaModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"model",
			p.staminaName,
		},
	).String()
}

func (p *StaminaModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type MaxStaminaTableRef struct {
	namespaceName       string
	maxStaminaTableName string
}

func (p *MaxStaminaTableRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"maxStaminaTable",
			p.maxStaminaTableName,
		},
	).String()
}

func (p *MaxStaminaTableRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RecoverIntervalTableRef struct {
	namespaceName            string
	recoverIntervalTableName string
}

func (p *RecoverIntervalTableRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"recoverIntervalTable",
			p.recoverIntervalTableName,
		},
	).String()
}

func (p *RecoverIntervalTableRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RecoverValueTableRef struct {
	namespaceName         string
	recoverValueTableName string
}

func (p *RecoverValueTableRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stamina",
			p.namespaceName,
			"recoverValueTable",
			p.recoverValueTableName,
		},
	).String()
}

func (p *RecoverValueTableRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
