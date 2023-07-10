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

type Content struct {
	Section     string
	Content     string
	FrontMatter string
}

type ContentOptions struct {
}

func NewContent(
	section string,
	content string,
	frontMatter string,
	options ContentOptions,
) Content {
	data := Content{
		Section:     section,
		Content:     content,
		FrontMatter: frontMatter,
	}
	return data
}

func (p *Content) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Section"] = p.Section
	properties["Content"] = p.Content
	properties["FrontMatter"] = p.FrontMatter
	return properties
}
