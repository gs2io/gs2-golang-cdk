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

type Trace struct {
	TraceId   string
	Spans     []LogEntry
	Truncated bool
}

type TraceOptions struct {
	Spans []LogEntry
}

func NewTrace(
	traceId string,
	truncated bool,
	options TraceOptions,
) Trace {
	_data := Trace{
		TraceId:   traceId,
		Truncated: truncated,
		Spans:     options.Spans,
	}
	return _data
}

func (p *Trace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TraceId"] = p.TraceId
	{
		values := make([]map[string]interface{}, len(p.Spans))
		for i, element := range p.Spans {
			values[i] = element.Properties()
		}
		properties["Spans"] = values
	}
	properties["Truncated"] = p.Truncated
	return properties
}
