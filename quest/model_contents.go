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
	Metadata               *string
	CompleteAcquireActions []AcquireAction
	Weight                 int32
}

type ContentsOptions struct {
	Metadata               *string
	CompleteAcquireActions []AcquireAction
}

func NewContents(
	weight int32,
	options ContentsOptions,
) Contents {
	data := Contents{
		Weight:                 weight,
		Metadata:               options.Metadata,
		CompleteAcquireActions: options.CompleteAcquireActions,
	}
	return data
}

func (p *Contents) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.CompleteAcquireActions))
		for i, element := range p.CompleteAcquireActions {
			values[i] = element.Properties()
		}
		properties["CompleteAcquireActions"] = values
	}
	properties["Weight"] = p.Weight
	return properties
}
