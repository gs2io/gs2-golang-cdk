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

type GuildRef struct {
	NamespaceName  string
	GuildModelName string
	GuildName      string
}

func (p *GuildRef) IncreaseMaximumCurrentMaximumMemberCountByGuildName(
	value *int32,
) AcquireAction {
	return IncreaseMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		p.GuildName,
		value,
	)
}

func (p *GuildRef) SetMaximumCurrentMaximumMemberCountByGuildName(
	value *int32,
) AcquireAction {
	return SetMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildName,
		p.GuildModelName,
		value,
	)
}

func (p *GuildRef) DecreaseMaximumCurrentMaximumMemberCountByGuildName(
	value *int32,
) ConsumeAction {
	return DecreaseMaximumCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		p.GuildName,
		value,
	)
}

func (p *GuildRef) VerifyCurrentMaximumMemberCountByGuildName(
	verifyType string,
	value *int32,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	return VerifyCurrentMaximumMemberCountByGuildName(
		p.NamespaceName,
		p.GuildModelName,
		p.GuildName,
		verifyType,
		value,
		multiplyValueSpecifyingQuantity,
	)
}

func (p *GuildRef) VerifyIncludeMember(
	verifyType string,
) VerifyAction {
	return VerifyIncludeMemberByUserId(
		p.NamespaceName,
		p.GuildModelName,
		verifyType,
		&p.GuildName,
	)
}

func (p *GuildRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"guild",
			p.NamespaceName,
			"guild",
			p.GuildModelName,
			p.GuildName,
		},
	).String()
}

func (p *GuildRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
