package freeze

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

type Output struct {
	Name      string
	Text      string
	CreatedAt int64
	TtlAt     int64
	Revision  *int64
}

type OutputOptions struct {
	Revision *int64
}

func NewOutput(
	name string,
	text string,
	createdAt int64,
	ttlAt int64,
	options OutputOptions,
) Output {
	_data := Output{
		Name:      name,
		Text:      text,
		CreatedAt: createdAt,
		TtlAt:     ttlAt,
		Revision:  options.Revision,
	}
	return _data
}

func (p *Output) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Text"] = p.Text
	properties["CreatedAt"] = p.CreatedAt
	properties["TtlAt"] = p.TtlAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
