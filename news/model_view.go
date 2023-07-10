package news

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

type View struct {
	Contents       []Content
	RemoveContents []Content
}

type ViewOptions struct {
	Contents       []Content
	RemoveContents []Content
}

func NewView(
	options ViewOptions,
) View {
	data := View{
		Contents:       options.Contents,
		RemoveContents: options.RemoveContents,
	}
	return data
}

func (p *View) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	{
		values := make([]map[string]interface{}, len(p.Contents))
		for i, element := range p.Contents {
			values[i] = element.Properties()
		}
		properties["Contents"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.RemoveContents))
		for i, element := range p.RemoveContents {
			values[i] = element.Properties()
		}
		properties["RemoveContents"] = values
	}
	return properties
}
