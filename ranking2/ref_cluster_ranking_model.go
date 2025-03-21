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

type ClusterRankingModelRef struct {
	NamespaceName string
	RankingName   string
}

func (p *ClusterRankingModelRef) CreateClusterRankingReceivedReward(
	clusterName string,
	season *int64,
) ConsumeAction {
	return CreateClusterRankingReceivedRewardByUserId(
		p.NamespaceName,
		p.RankingName,
		clusterName,
		season,
	)
}

func (p *ClusterRankingModelRef) VerifyClusterRankingScore(
	clusterName string,
	verifyType string,
	score int64,
	season *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	return VerifyClusterRankingScoreByUserId(
		p.NamespaceName,
		p.RankingName,
		clusterName,
		verifyType,
		score,
		season,
		multiplyValueSpecifyingQuantity,
	)
}

func (p *ClusterRankingModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"ranking2",
			p.NamespaceName,
			"cluster",
			p.RankingName,
		},
	).String()
}

func (p *ClusterRankingModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
