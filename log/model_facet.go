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

type Facet struct {
	Field       string
	Values      []FacetValueCount
	Range       *NumericRange
	GlobalRange *NumericRange
}

type FacetOptions struct {
	Values      []FacetValueCount
	Range       *NumericRange
	GlobalRange *NumericRange
}

func NewFacet(
	field string,
	options FacetOptions,
) Facet {
	_data := Facet{
		Field:       field,
		Values:      options.Values,
		Range:       options.Range,
		GlobalRange: options.GlobalRange,
	}
	return _data
}

func (p *Facet) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Field"] = p.Field
	{
		values := make([]map[string]interface{}, len(p.Values))
		for i, element := range p.Values {
			values[i] = element.Properties()
		}
		properties["Values"] = values
	}
	if p.Range != nil {
		properties["Range"] = p.Range.Properties()
	}
	if p.GlobalRange != nil {
		properties["GlobalRange"] = p.GlobalRange.Properties()
	}
	return properties
}
