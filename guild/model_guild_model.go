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

type GuildModel struct {
	Name                          string
	Metadata                      *string
	DefaultMaximumMemberCount     int32
	MaximumMemberCount            int32
	InactivityPeriodDays          int32
	Roles                         []RoleModel
	GuildMasterRole               string
	GuildMemberDefaultRole        string
	RejoinCoolTimeMinutes         int32
	MaxConcurrentJoinGuilds       *int32
	MaxConcurrentGuildMasterCount *int32
}

type GuildModelOptions struct {
	Metadata                      *string
	MaxConcurrentJoinGuilds       *int32
	MaxConcurrentGuildMasterCount *int32
}

func NewGuildModel(
	name string,
	defaultMaximumMemberCount int32,
	maximumMemberCount int32,
	inactivityPeriodDays int32,
	roles []RoleModel,
	guildMasterRole string,
	guildMemberDefaultRole string,
	rejoinCoolTimeMinutes int32,
	options GuildModelOptions,
) GuildModel {
	_data := GuildModel{
		Name:                          name,
		DefaultMaximumMemberCount:     defaultMaximumMemberCount,
		MaximumMemberCount:            maximumMemberCount,
		InactivityPeriodDays:          inactivityPeriodDays,
		Roles:                         roles,
		GuildMasterRole:               guildMasterRole,
		GuildMemberDefaultRole:        guildMemberDefaultRole,
		RejoinCoolTimeMinutes:         rejoinCoolTimeMinutes,
		Metadata:                      options.Metadata,
		MaxConcurrentJoinGuilds:       options.MaxConcurrentJoinGuilds,
		MaxConcurrentGuildMasterCount: options.MaxConcurrentGuildMasterCount,
	}
	return _data
}

func (p *GuildModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["DefaultMaximumMemberCount"] = p.DefaultMaximumMemberCount
	properties["MaximumMemberCount"] = p.MaximumMemberCount
	properties["InactivityPeriodDays"] = p.InactivityPeriodDays
	{
		values := make([]map[string]interface{}, len(p.Roles))
		for i, element := range p.Roles {
			values[i] = element.Properties()
		}
		properties["Roles"] = values
	}
	properties["GuildMasterRole"] = p.GuildMasterRole
	properties["GuildMemberDefaultRole"] = p.GuildMemberDefaultRole
	properties["RejoinCoolTimeMinutes"] = p.RejoinCoolTimeMinutes
	if p.MaxConcurrentJoinGuilds != nil {
		properties["MaxConcurrentJoinGuilds"] = p.MaxConcurrentJoinGuilds
	}
	if p.MaxConcurrentGuildMasterCount != nil {
		properties["MaxConcurrentGuildMasterCount"] = p.MaxConcurrentGuildMasterCount
	}
	return properties
}
