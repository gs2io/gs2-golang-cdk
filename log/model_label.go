package log

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

type Label struct {
	Key   string
	Value string
}

type LabelOptions struct {
}

func NewLabel(
	key string,
	value string,
	options LabelOptions,
) Label {
	_data := Label{
		Key:   key,
		Value: value,
	}
	return _data
}

func (p *Label) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Key"] = p.Key
	properties["Value"] = p.Value
	return properties
}
