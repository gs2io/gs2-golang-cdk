package account

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

type BanStatus struct {
	Name             string
	Reason           string
	ReleaseTimestamp int64
}

type BanStatusOptions struct {
}

func NewBanStatus(
	name string,
	reason string,
	releaseTimestamp int64,
	options BanStatusOptions,
) BanStatus {
	data := BanStatus{
		Name:             name,
		Reason:           reason,
		ReleaseTimestamp: releaseTimestamp,
	}
	return data
}

func (p *BanStatus) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Reason"] = p.Reason
	properties["ReleaseTimestamp"] = p.ReleaseTimestamp
	return properties
}
