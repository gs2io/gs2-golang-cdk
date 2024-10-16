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

func VerifyGlobalRankingScoreByUserId(
	namespaceName string,
	rankingName string,
	verifyType string,
	score int64,
	season *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rankingName"] = rankingName
	properties["verifyType"] = verifyType
	if season != nil {
		properties["season"] = season
	}
	properties["score"] = score
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Ranking2:VerifyGlobalRankingScoreByUserId",
		Request: properties,
	}
}

func CreateGlobalRankingReceivedRewardByUserId(
	namespaceName string,
	rankingName string,
	season *int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rankingName"] = rankingName
	if season != nil {
		properties["season"] = season
	}
	return ConsumeAction{
		Action:  "Gs2Ranking2:CreateGlobalRankingReceivedRewardByUserId",
		Request: properties,
	}
}

func VerifyClusterRankingScoreByUserId(
	namespaceName string,
	rankingName string,
	clusterName string,
	verifyType string,
	score int64,
	season *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rankingName"] = rankingName
	properties["clusterName"] = clusterName
	properties["verifyType"] = verifyType
	if season != nil {
		properties["season"] = season
	}
	properties["score"] = score
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Ranking2:VerifyClusterRankingScoreByUserId",
		Request: properties,
	}
}

func CreateClusterRankingReceivedRewardByUserId(
	namespaceName string,
	rankingName string,
	clusterName string,
	season *int64,
) ConsumeAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rankingName"] = rankingName
	properties["clusterName"] = clusterName
	if season != nil {
		properties["season"] = season
	}
	return ConsumeAction{
		Action:  "Gs2Ranking2:CreateClusterRankingReceivedRewardByUserId",
		Request: properties,
	}
}

func VerifySubscribeRankingScoreByUserId(
	namespaceName string,
	rankingName string,
	verifyType string,
	score int64,
	season *int64,
	multiplyValueSpecifyingQuantity *bool,
) VerifyAction {
	properties := map[string]interface{}{
		"userId": "#{userId}",
	}
	properties["namespaceName"] = namespaceName
	properties["rankingName"] = rankingName
	properties["verifyType"] = verifyType
	if season != nil {
		properties["season"] = season
	}
	properties["score"] = score
	if multiplyValueSpecifyingQuantity != nil {
		properties["multiplyValueSpecifyingQuantity"] = multiplyValueSpecifyingQuantity
	}
	return VerifyAction{
		Action:  "Gs2Ranking2:VerifySubscribeRankingScoreByUserId",
		Request: properties,
	}
}
