package schedule

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

type RepeatSchedule struct {
	RepeatCount          int32
	CurrentRepeatStartAt *int64
	CurrentRepeatEndAt   *int64
	LastRepeatEndAt      *int64
	NextRepeatStartAt    *int64
}

type RepeatScheduleOptions struct {
	CurrentRepeatStartAt *int64
	CurrentRepeatEndAt   *int64
	LastRepeatEndAt      *int64
	NextRepeatStartAt    *int64
}

func NewRepeatSchedule(
	repeatCount int32,
	options RepeatScheduleOptions,
) RepeatSchedule {
	data := RepeatSchedule{
		RepeatCount:          repeatCount,
		CurrentRepeatStartAt: options.CurrentRepeatStartAt,
		CurrentRepeatEndAt:   options.CurrentRepeatEndAt,
		LastRepeatEndAt:      options.LastRepeatEndAt,
		NextRepeatStartAt:    options.NextRepeatStartAt,
	}
	return data
}

func (p *RepeatSchedule) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["RepeatCount"] = p.RepeatCount
	if p.CurrentRepeatStartAt != nil {
		properties["CurrentRepeatStartAt"] = p.CurrentRepeatStartAt
	}
	if p.CurrentRepeatEndAt != nil {
		properties["CurrentRepeatEndAt"] = p.CurrentRepeatEndAt
	}
	if p.LastRepeatEndAt != nil {
		properties["LastRepeatEndAt"] = p.LastRepeatEndAt
	}
	if p.NextRepeatStartAt != nil {
		properties["NextRepeatStartAt"] = p.NextRepeatStartAt
	}
	return properties
}
