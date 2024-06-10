package guild

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

type GuildModelRef struct {
	NamespaceName  string
	GuildModelName string
}

func (p *GuildModelRef) RoleModel() *RoleModelRef {
	return &RoleModelRef{
		NamespaceName:  p.NamespaceName,
		GuildModelName: p.GuildModelName,
	}
}

func (p *GuildModelRef) IncreaseMaximumCurrentMaximumMemberCountByGuildName(
	guildName string,
	value *int32,
) AcquireAction {
	return IncreaseMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		guildName,
		value,
	)
}

func (p *GuildModelRef) SetMaximumCurrentMaximumMemberCountByGuildName(
	guildName string,
	value *int32,
) AcquireAction {
	return SetMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		guildName,
		p.GuildModelName,
		value,
	)
}

func (p *GuildModelRef) DecreaseMaximumCurrentMaximumMemberCountByGuildName(
	guildName string,
	value *int32,
) ConsumeAction {
	return DecreaseMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		guildName,
		value,
	)
}

func (p *GuildModelRef) VerifyCurrentMaximumMemberCountByGuildName(
	guildName string,
	verifyType string,
	multiplyValueSpecifyingQuantity bool,
	value *int32,
) ConsumeAction {
	return VerifyCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		guildName,
		verifyType,
		value,
		multiplyValueSpecifyingQuantity,
	)
}

func (p *GuildModelRef) VerifyIncludeMember(
	guildName string,
	verifyType string,
) ConsumeAction {
	return VerifyIncludeMemberByUserId(
		p.NamespaceName,
		p.GuildModelName,
		verifyType,
		guildName,
	)
}

func (p *GuildModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"guild",
			p.NamespaceName,
			"model",
			p.GuildModelName,
		},
	).String()
}

func (p *GuildModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
