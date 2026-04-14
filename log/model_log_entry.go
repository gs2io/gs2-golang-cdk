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

type LogEntryStatus string

const LogEntryStatusOk = LogEntryStatus("ok")
const LogEntryStatusInfo = LogEntryStatus("info")
const LogEntryStatusNotice = LogEntryStatus("notice")
const LogEntryStatusError = LogEntryStatus("error")
const LogEntryStatusWarn = LogEntryStatus("warn")
const LogEntryStatusEmag = LogEntryStatus("emag")

func (p LogEntryStatus) Pointer() *LogEntryStatus {
	return &p
}

type LogEntry struct {
	Timestamp int64
	Status    LogEntryStatus
	Duration  int64
	Line      string
	Labels    []Label
}

type LogEntryOptions struct {
	Labels []Label
}

func NewLogEntry(
	timestamp int64,
	status LogEntryStatus,
	duration int64,
	line string,
	options LogEntryOptions,
) LogEntry {
	_data := LogEntry{
		Timestamp: timestamp,
		Status:    status,
		Duration:  duration,
		Line:      line,
		Labels:    options.Labels,
	}
	return _data
}

func (p *LogEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.Timestamp
	properties["Status"] = p.Status
	properties["Duration"] = p.Duration
	properties["Line"] = p.Line
	{
		values := make([]map[string]interface{}, len(p.Labels))
		for i, element := range p.Labels {
			values[i] = element.Properties()
		}
		properties["Labels"] = values
	}
	return properties
}
