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

type JobResultBody struct {
	TryNumber  int32
	StatusCode int32
	Result     string
	TryAt      int64
}

type JobResultBodyOptions struct {
}

func NewJobResultBody(
	tryNumber int32,
	statusCode int32,
	result string,
	tryAt int64,
	options JobResultBodyOptions,
) JobResultBody {
	_data := JobResultBody{
		TryNumber:  tryNumber,
		StatusCode: statusCode,
		Result:     result,
		TryAt:      tryAt,
	}
	return _data
}

func (p *JobResultBody) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["TryNumber"] = p.TryNumber
	properties["StatusCode"] = p.StatusCode
	properties["Result"] = p.Result
	properties["TryAt"] = p.TryAt
	return properties
}
