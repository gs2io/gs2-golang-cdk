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

type GlobalRankingSetting struct {
	UniqueByUserId           bool
	CalculateIntervalMinutes int32
	CalculateFixedTiming     *FixedTiming
	AdditionalScopes         []Scope
	IgnoreUserIds            []string
	Generation               *string
}

type GlobalRankingSettingOptions struct {
	CalculateFixedTiming *FixedTiming
	AdditionalScopes     []Scope
	IgnoreUserIds        []string
	Generation           *string
}

func NewGlobalRankingSetting(
	uniqueByUserId bool,
	calculateIntervalMinutes int32,
	options GlobalRankingSettingOptions,
) GlobalRankingSetting {
	data := GlobalRankingSetting{
		UniqueByUserId:           uniqueByUserId,
		CalculateIntervalMinutes: calculateIntervalMinutes,
		CalculateFixedTiming:     options.CalculateFixedTiming,
		AdditionalScopes:         options.AdditionalScopes,
		IgnoreUserIds:            options.IgnoreUserIds,
		Generation:               options.Generation,
	}
	return data
}

func (p *GlobalRankingSetting) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UniqueByUserId"] = p.UniqueByUserId
	properties["CalculateIntervalMinutes"] = p.CalculateIntervalMinutes
	if p.CalculateFixedTiming != nil {
		properties["CalculateFixedTiming"] = p.CalculateFixedTiming.Properties()
	}
	{
		values := make([]map[string]interface{}, len(p.AdditionalScopes))
		for i, element := range p.AdditionalScopes {
			values[i] = element.Properties()
		}
		properties["AdditionalScopes"] = values
	}
	properties["IgnoreUserIds"] = p.IgnoreUserIds
	if p.Generation != nil {
		properties["Generation"] = p.Generation
	}
	return properties
}
