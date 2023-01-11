package quest

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

type RewardAction string

func (p RewardAction) Pointer() *RewardAction {
	return &p
}

type Reward struct {
	Action  string
	Request string
	ItemId  string
	Value   int32
}

type RewardOptions struct {
}

func NewReward(
	action string,
	request string,
	itemId string,
	value int32,
	options RewardOptions,
) Reward {
	data := Reward{
		Action:  action,
		Request: request,
		ItemId:  itemId,
		Value:   value,
	}
	return data
}

func (p *Reward) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Action"] = p.Action
	properties["Request"] = p.Request
	properties["ItemId"] = p.ItemId
	properties["Value"] = p.Value
	return properties
}
