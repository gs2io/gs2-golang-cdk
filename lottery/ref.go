package lottery

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

func (p *NamespaceRef) CurrentLotteryMaster() *CurrentLotteryMasterRef {
	return &CurrentLotteryMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) PrizeTable(
	prizeTableName string,
) *PrizeTableRef {
	return &PrizeTableRef{
		namespaceName:  p.namespaceName,
		prizeTableName: prizeTableName,
	}
}

func (p *NamespaceRef) LotteryModel(
	lotteryName string,
) *LotteryModelRef {
	return &LotteryModelRef{
		namespaceName: p.namespaceName,
		lotteryName:   lotteryName,
	}
}

func (p *NamespaceRef) PrizeTableMaster(
	prizeTableName string,
) *PrizeTableMasterRef {
	return &PrizeTableMasterRef{
		namespaceName:  p.namespaceName,
		prizeTableName: prizeTableName,
	}
}

func (p *NamespaceRef) LotteryModelMaster(
	lotteryName string,
) *LotteryModelMasterRef {
	return &LotteryModelMasterRef{
		namespaceName: p.namespaceName,
		lotteryName:   lotteryName,
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
			"lottery",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LotteryModelMasterRef struct {
	namespaceName string
	lotteryName   string
}

func (p *LotteryModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
			"lotteryModel",
			p.lotteryName,
		},
	).String()
}

func (p *LotteryModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type PrizeTableMasterRef struct {
	namespaceName  string
	prizeTableName string
}

func (p *PrizeTableMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
			"table",
			p.prizeTableName,
		},
	).String()
}

func (p *PrizeTableMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LotteryModelRef struct {
	namespaceName string
	lotteryName   string
}

func (p *LotteryModelRef) Draw(
	count int32,
	config *[]Config,
) AcquireAction {
	return DrawByUserId(
		p.namespaceName,
		p.lotteryName,
		count,
		config,
	)
}

func (p *LotteryModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
			"lotteryModel",
			p.lotteryName,
		},
	).String()
}

func (p *LotteryModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type PrizeTableRef struct {
	namespaceName  string
	prizeTableName string
}

func (p *PrizeTableRef) PrizeLimit(
	prizeId string,
) *PrizeLimitRef {
	return &PrizeLimitRef{
		namespaceName:  p.namespaceName,
		prizeTableName: p.prizeTableName,
		prizeId:        prizeId,
	}
}

func (p *PrizeTableRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
			"table",
			p.prizeTableName,
		},
	).String()
}

func (p *PrizeTableRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentLotteryMasterRef struct {
	namespaceName string
}

func (p *CurrentLotteryMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentLotteryMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type PrizeLimitRef struct {
	namespaceName  string
	prizeTableName string
	prizeId        string
}

func (p *PrizeLimitRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"lottery",
			p.namespaceName,
			"table",
			p.prizeTableName,
			"prize",
			p.prizeId,
		},
	).String()
}

func (p *PrizeLimitRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
