package seasonRating

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

type TierModel struct {
	Metadata           *string
	RaiseRankBonus     int32
	EntryFee           int32
	MinimumChangePoint int32
	MaximumChangePoint int32
}

type TierModelOptions struct {
	Metadata *string
}

func NewTierModel(
	raiseRankBonus int32,
	entryFee int32,
	minimumChangePoint int32,
	maximumChangePoint int32,
	options TierModelOptions,
) TierModel {
	_data := TierModel{
		RaiseRankBonus:     raiseRankBonus,
		EntryFee:           entryFee,
		MinimumChangePoint: minimumChangePoint,
		MaximumChangePoint: maximumChangePoint,
		Metadata:           options.Metadata,
	}
	return _data
}

func (p *TierModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["RaiseRankBonus"] = p.RaiseRankBonus
	properties["EntryFee"] = p.EntryFee
	properties["MinimumChangePoint"] = p.MinimumChangePoint
	properties["MaximumChangePoint"] = p.MaximumChangePoint
	return properties
}
