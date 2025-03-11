package jobQueue

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

type JobEntry struct {
	ScriptId    string
	Args        string
	MaxTryCount int32
}

type JobEntryOptions struct {
}

func NewJobEntry(
	scriptId string,
	args string,
	maxTryCount int32,
	options JobEntryOptions,
) JobEntry {
	_data := JobEntry{
		ScriptId:    scriptId,
		Args:        args,
		MaxTryCount: maxTryCount,
	}
	return _data
}

func (p *JobEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ScriptId"] = p.ScriptId
	properties["Args"] = p.Args
	properties["MaxTryCount"] = p.MaxTryCount
	return properties
}
