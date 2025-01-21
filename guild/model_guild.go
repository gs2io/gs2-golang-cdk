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

type GuildJoinPolicy string

const GuildJoinPolicyAnybody = GuildJoinPolicy("anybody")
const GuildJoinPolicyApproval = GuildJoinPolicy("approval")

func (p GuildJoinPolicy) Pointer() *GuildJoinPolicy {
	return &p
}

type Guild struct {
	CdkResource
	stack                  *Stack
	NamespaceName          string
	UserId                 string
	GuildModelName         string
	DisplayName            string
	Attribute1             *int32
	Attribute2             *int32
	Attribute3             *int32
	Attribute4             *int32
	Attribute5             *int32
	Metadata               *string
	MemberMetadata         *string
	JoinPolicy             GuildJoinPolicy
	CustomRoles            []RoleModel
	GuildMemberDefaultRole *string
	TimeOffsetToken        *string
}

type GuildOptions struct {
	Attribute1             *int32
	Attribute2             *int32
	Attribute3             *int32
	Attribute4             *int32
	Attribute5             *int32
	Metadata               *string
	MemberMetadata         *string
	CustomRoles            []RoleModel
	GuildMemberDefaultRole *string
	TimeOffsetToken        *string
}

func NewGuild(
	stack *Stack,
	namespaceName string,
	userId string,
	guildModelName string,
	displayName string,
	joinPolicy GuildJoinPolicy,
	options GuildOptions,
) *Guild {
	data := Guild{
		stack:                  stack,
		NamespaceName:          namespaceName,
		UserId:                 userId,
		GuildModelName:         guildModelName,
		DisplayName:            displayName,
		JoinPolicy:             joinPolicy,
		Attribute1:             options.Attribute1,
		Attribute2:             options.Attribute2,
		Attribute3:             options.Attribute3,
		Attribute4:             options.Attribute4,
		Attribute5:             options.Attribute5,
		Metadata:               options.Metadata,
		MemberMetadata:         options.MemberMetadata,
		CustomRoles:            options.CustomRoles,
		GuildMemberDefaultRole: options.GuildMemberDefaultRole,
		TimeOffsetToken:        options.TimeOffsetToken,
	}
	return &data
}

func (p *Guild) ResourceName() string {
	return "Guild_Guild_" + p.GuildModelName + ':' + p.Name
}

func (p *Guild) ResourceType() string {
	return "GS2::Guild::Guild"
}

func (p *Guild) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["UserId"] = p.UserId
	properties["GuildModelName"] = p.GuildModelName
	properties["DisplayName"] = p.DisplayName
	if p.Attribute1 != nil {
		properties["Attribute1"] = p.Attribute1
	}
	if p.Attribute2 != nil {
		properties["Attribute2"] = p.Attribute2
	}
	if p.Attribute3 != nil {
		properties["Attribute3"] = p.Attribute3
	}
	if p.Attribute4 != nil {
		properties["Attribute4"] = p.Attribute4
	}
	if p.Attribute5 != nil {
		properties["Attribute5"] = p.Attribute5
	}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	if p.MemberMetadata != nil {
		properties["MemberMetadata"] = p.MemberMetadata
	}
	properties["JoinPolicy"] = p.JoinPolicy
	{
		values := make([]map[string]interface{}, len(p.CustomRoles))
		for i, element := range p.CustomRoles {
			values[i] = element.Properties()
		}
		properties["CustomRoles"] = values
	}
	if p.GuildMemberDefaultRole != nil {
		properties["GuildMemberDefaultRole"] = p.GuildMemberDefaultRole
	}
	if p.TimeOffsetToken != nil {
		properties["TimeOffsetToken"] = p.TimeOffsetToken
	}
	return properties
}

func (p *Guild) Ref(
	namespaceName string,
) GuildRef {
	return GuildRef{
		NamespaceName:  namespaceName,
		GuildModelName: p.GuildModelName,
		GuildName:      p.Name,
	}
}

func (p *Guild) GetAttrGuildId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.GuildId",
	)
}
