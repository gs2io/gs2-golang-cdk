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

type IssueStampSheetLogCount struct {
	Service *string
	Method  *string
	UserId  *string
	Action  *string
	Count   int64
}

type IssueStampSheetLogCountOptions struct {
	Service *string
	Method  *string
	UserId  *string
	Action  *string
}

func NewIssueStampSheetLogCount(
	count int64,
	options IssueStampSheetLogCountOptions,
) IssueStampSheetLogCount {
	data := IssueStampSheetLogCount{
		Count:   count,
		Service: options.Service,
		Method:  options.Method,
		UserId:  options.UserId,
		Action:  options.Action,
	}
	return data
}

func (p *IssueStampSheetLogCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.Service != nil {
		properties["Service"] = p.Service
	}
	if p.Method != nil {
		properties["Method"] = p.Method
	}
	if p.UserId != nil {
		properties["UserId"] = p.UserId
	}
	if p.Action != nil {
		properties["Action"] = p.Action
	}
	properties["Count"] = p.Count
	return properties
}
