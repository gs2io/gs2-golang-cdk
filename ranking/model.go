package ranking

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

type Ranking struct {
	rank         int64
	index        int64
	categoryName string
	userId       string
	score        int64
	metadata     *string
	createdAt    int64
}

func NewRanking(
	rank int64,
	index int64,
	categoryName string,
	userId string,
	score int64,
	metadata *string,
	createdAt int64,
) Ranking {
	data := Ranking{
		rank:         rank,
		index:        index,
		categoryName: categoryName,
		userId:       userId,
		score:        score,
		metadata:     metadata,
		createdAt:    createdAt,
	}
	return data
}

func (p *Ranking) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Rank"] = p.rank
	properties["Index"] = p.index
	properties["CategoryName"] = p.categoryName
	properties["UserId"] = p.userId
	properties["Score"] = p.score
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	properties["CreatedAt"] = p.createdAt
	return properties
}

type CalculatedAt struct {
	categoryName string
	calculatedAt int64
}

func NewCalculatedAt(
	categoryName string,
	calculatedAt int64,
) CalculatedAt {
	data := CalculatedAt{
		categoryName: categoryName,
		calculatedAt: calculatedAt,
	}
	return data
}

func (p *CalculatedAt) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CategoryName"] = p.categoryName
	properties["CalculatedAt"] = p.calculatedAt
	return properties
}

type SubscribeUser struct {
	categoryName string
	userId       string
	targetUserId string
}

func NewSubscribeUser(
	categoryName string,
	userId string,
	targetUserId string,
) SubscribeUser {
	data := SubscribeUser{
		categoryName: categoryName,
		userId:       userId,
		targetUserId: targetUserId,
	}
	return data
}

func (p *SubscribeUser) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CategoryName"] = p.categoryName
	properties["UserId"] = p.userId
	properties["TargetUserId"] = p.targetUserId
	return properties
}
