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

type StaminaModelRef struct {
	NamespaceName string
	StaminaName   string
}

func (p *StaminaModelRef) RecoverStamina(
	recoverValue int32,
) AcquireAction {
	return RecoverStaminaByUserId(
		p.NamespaceName,
		p.StaminaName,
		recoverValue,
	)
}

func (p *StaminaModelRef) RaiseMaxValue(
	raiseValue int32,
) AcquireAction {
	return RaiseMaxValueByUserId(
		p.NamespaceName,
		p.StaminaName,
		raiseValue,
	)
}

func (p *StaminaModelRef) SetMaxValue(
	maxValue int32,
) AcquireAction {
	return SetMaxValueByUserId(
		p.NamespaceName,
		p.StaminaName,
		maxValue,
	)
}

func (p *StaminaModelRef) SetRecoverInterval(
	recoverIntervalMinutes int32,
) AcquireAction {
	return SetRecoverIntervalByUserId(
		p.NamespaceName,
		p.StaminaName,
		recoverIntervalMinutes,
	)
}

func (p *StaminaModelRef) SetRecoverValue(
	recoverValue int32,
) AcquireAction {
	return SetRecoverValueByUserId(
		p.NamespaceName,
		p.StaminaName,
		recoverValue,
	)
}

func (p *StaminaModelRef) DecreaseMaxValue(
	decreaseValue int32,
) ConsumeAction {
	return DecreaseMaxValueByUserId(
		p.NamespaceName,
		p.StaminaName,
		decreaseValue,
	)
}

func (p *StaminaModelRef) ConsumeStamina(
	consumeValue int32,
) ConsumeAction {
	return ConsumeStaminaByUserId(
		p.NamespaceName,
		p.StaminaName,
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
			p.NamespaceName,
			"model",
			p.StaminaName,
		},
	).String()
}

func (p *StaminaModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
