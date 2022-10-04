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
	scriptId    string
	args        string
	maxTryCount int32
}

func NewJobEntry(
	scriptId string,
	args string,
	maxTryCount int32,
) JobEntry {
	data := JobEntry{
		scriptId:    scriptId,
		args:        args,
		maxTryCount: maxTryCount,
	}
	return data
}

func (p *JobEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["ScriptId"] = p.scriptId
	properties["Args"] = p.args
	properties["MaxTryCount"] = p.maxTryCount
	return properties
}

type JobResultBody struct {
	tryNumber  int32
	statusCode int32
	result     string
	tryAt      int64
}

func NewJobResultBody(
	tryNumber int32,
	statusCode int32,
	result string,
	tryAt int64,
) JobResultBody {
	data := JobResultBody{
		tryNumber:  tryNumber,
		statusCode: statusCode,
		result:     result,
		tryAt:      tryAt,
	}
	return data
}

func (p *JobResultBody) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TryNumber"] = p.tryNumber
	properties["StatusCode"] = p.statusCode
	properties["Result"] = p.result
	properties["TryAt"] = p.tryAt
	return properties
}
