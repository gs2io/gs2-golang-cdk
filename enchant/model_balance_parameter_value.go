package enchant

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

type BalanceParameterValue struct {
	Name  string
	Value int64
}

type BalanceParameterValueOptions struct {
}

func NewBalanceParameterValue(
	name string,
	value int64,
	options BalanceParameterValueOptions,
) BalanceParameterValue {
	_data := BalanceParameterValue{
		Name:  name,
		Value: value,
	}
	return _data
}

func (p *BalanceParameterValue) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["Value"] = p.Value
	return properties
}
