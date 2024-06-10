package ranking2

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

type ClusterRankingModelClusterType string

const ClusterRankingModelClusterTypeRaw = ClusterRankingModelClusterType("Raw")
const ClusterRankingModelClusterTypeGs2Guild_guild = ClusterRankingModelClusterType("Gs2Guild::Guild")

func (p ClusterRankingModelClusterType) Pointer() *ClusterRankingModelClusterType {
	return &p
}

type ClusterRankingModelOrderDirection string

const ClusterRankingModelOrderDirectionAsc = ClusterRankingModelOrderDirection("asc")
const ClusterRankingModelOrderDirectionDesc = ClusterRankingModelOrderDirection("desc")

func (p ClusterRankingModelOrderDirection) Pointer() *ClusterRankingModelOrderDirection {
	return &p
}

type ClusterRankingModel struct {
	Name                string
	Metadata            *string
	ClusterType         ClusterRankingModelClusterType
	MinimumValue        *int64
	MaximumValue        *int64
	Sum                 bool
	OrderDirection      ClusterRankingModelOrderDirection
	EntryPeriodEventId  *string
	RankingRewards      []RankingReward
	AccessPeriodEventId *string
}

type ClusterRankingModelOptions struct {
	Metadata            *string
	MinimumValue        *int64
	MaximumValue        *int64
	EntryPeriodEventId  *string
	RankingRewards      []RankingReward
	AccessPeriodEventId *string
}

func NewClusterRankingModel(
	name string,
	clusterType ClusterRankingModelClusterType,
	sum bool,
	orderDirection ClusterRankingModelOrderDirection,
	options ClusterRankingModelOptions,
) ClusterRankingModel {
	data := ClusterRankingModel{
		Name:                name,
		ClusterType:         clusterType,
		Sum:                 sum,
		OrderDirection:      orderDirection,
		Metadata:            options.Metadata,
		MinimumValue:        options.MinimumValue,
		MaximumValue:        options.MaximumValue,
		EntryPeriodEventId:  options.EntryPeriodEventId,
		RankingRewards:      options.RankingRewards,
		AccessPeriodEventId: options.AccessPeriodEventId,
	}
	return data
}

func (p *ClusterRankingModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["ClusterType"] = p.ClusterType
	if p.MinimumValue != nil {
		properties["MinimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		properties["MaximumValue"] = p.MaximumValue
	}
	properties["Sum"] = p.Sum
	properties["OrderDirection"] = p.OrderDirection
	if p.EntryPeriodEventId != nil {
		properties["EntryPeriodEventId"] = p.EntryPeriodEventId
	}
	{
		values := make([]map[string]interface{}, len(p.RankingRewards))
		for i, element := range p.RankingRewards {
			values[i] = element.Properties()
		}
		properties["RankingRewards"] = values
	}
	if p.AccessPeriodEventId != nil {
		properties["AccessPeriodEventId"] = p.AccessPeriodEventId
	}
	return properties
}
