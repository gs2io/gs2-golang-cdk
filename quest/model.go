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

type Contents struct {
	metadata               *string
	completeAcquireActions []AcquireAction
	weight                 int32
}

func NewContents(
	metadata *string,
	completeAcquireActions []AcquireAction,
	weight int32,
) Contents {
	data := Contents{
		metadata:               metadata,
		completeAcquireActions: completeAcquireActions,
		weight:                 weight,
	}
	return data
}

func (p *Contents) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	{
		values := make([]map[string]interface{}, len(p.completeAcquireActions))
		for i, element := range p.completeAcquireActions {
			values[i] = element.Properties()
		}
		properties["CompleteAcquireActions"] = values
	}
	properties["Weight"] = p.weight
	return properties
}

type RewardAction string

func (p RewardAction) Pointer() *RewardAction {
	return &p
}

type Reward struct {
	action  string
	request string
	itemId  string
	value   int32
}

func NewReward(
	action string,
	request string,
	itemId string,
	value int32,
) Reward {
	data := Reward{
		action:  action,
		request: request,
		itemId:  itemId,
		value:   value,
	}
	return data
}

func (p *Reward) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Action"] = p.action
	properties["Request"] = p.request
	properties["ItemId"] = p.itemId
	properties["Value"] = p.value
	return properties
}
