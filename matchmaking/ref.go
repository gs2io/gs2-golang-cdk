package matchmaking

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

func (p *NamespaceRef) CurrentRatingModelMaster() *CurrentRatingModelMasterRef {
	return &CurrentRatingModelMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) RatingModel(
	ratingName string,
) *RatingModelRef {
	return &RatingModelRef{
		namespaceName: p.namespaceName,
		ratingName:    ratingName,
	}
}

func (p *NamespaceRef) Vote(
	ratingName string,
	gatheringName string,
) *VoteRef {
	return &VoteRef{
		namespaceName: p.namespaceName,
		ratingName:    ratingName,
		gatheringName: gatheringName,
	}
}

func (p *NamespaceRef) RatingModelMaster(
	ratingName string,
) *RatingModelMasterRef {
	return &RatingModelMasterRef{
		namespaceName: p.namespaceName,
		ratingName:    ratingName,
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
			"matchmaking",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type GatheringRef struct {
	namespaceName string
	gatheringName string
}

func (p *GatheringRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"matchmaking",
			p.namespaceName,
			"gathering",
			p.gatheringName,
		},
	).String()
}

func (p *GatheringRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RatingModelMasterRef struct {
	namespaceName string
	ratingName    string
}

func (p *RatingModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"matchmaking",
			p.namespaceName,
			"model",
			p.ratingName,
		},
	).String()
}

func (p *RatingModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type RatingModelRef struct {
	namespaceName string
	ratingName    string
}

func (p *RatingModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"matchmaking",
			p.namespaceName,
			"model",
			p.ratingName,
		},
	).String()
}

func (p *RatingModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentRatingModelMasterRef struct {
	namespaceName string
}

func (p *CurrentRatingModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"matchmaking",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentRatingModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type VoteRef struct {
	namespaceName string
	ratingName    string
	gatheringName string
}

func (p *VoteRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"matchmaking",
			p.namespaceName,
			"vote",
			p.ratingName,
			p.gatheringName,
		},
	).String()
}

func (p *VoteRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
