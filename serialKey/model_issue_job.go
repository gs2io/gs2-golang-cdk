package serialKey

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

type IssueJobStatus string

const IssueJobStatusProcessing = IssueJobStatus("PROCESSING")
const IssueJobStatusComplete = IssueJobStatus("COMPLETE")

func (p IssueJobStatus) Pointer() *IssueJobStatus {
	return &p
}

type IssueJob struct {
	Name              string
	Metadata          *string
	IssuedCount       int32
	IssueRequestCount int32
	Status            IssueJobStatus
	CreatedAt         int64
}

type IssueJobOptions struct {
	Metadata *string
}

func NewIssueJob(
	name string,
	issuedCount int32,
	issueRequestCount int32,
	status IssueJobStatus,
	createdAt int64,
	options IssueJobOptions,
) IssueJob {
	data := IssueJob{
		Name:              name,
		IssuedCount:       issuedCount,
		IssueRequestCount: issueRequestCount,
		Status:            status,
		CreatedAt:         createdAt,
		Metadata:          options.Metadata,
	}
	return data
}

func (p *IssueJob) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["IssuedCount"] = p.IssuedCount
	properties["IssueRequestCount"] = p.IssueRequestCount
	properties["Status"] = p.Status
	properties["CreatedAt"] = p.CreatedAt
	return properties
}
