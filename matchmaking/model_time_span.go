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

type TimeSpan struct {
	Days    int32
	Hours   int32
	Minutes int32
}

type TimeSpanOptions struct {
}

func NewTimeSpan(
	days int32,
	hours int32,
	minutes int32,
	options TimeSpanOptions,
) TimeSpan {
	_data := TimeSpan{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
	}
	return _data
}

func (p *TimeSpan) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Days"] = p.Days
	properties["Hours"] = p.Hours
	properties["Minutes"] = p.Minutes
	return properties
}
