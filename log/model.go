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

type AccessLog struct {
	timestamp int64
	requestId string
	service   string
	method    string
	userId    *string
	request   string
	result    string
}

func NewAccessLog(
	timestamp int64,
	requestId string,
	service string,
	method string,
	userId *string,
	request string,
	result string,
) AccessLog {
	data := AccessLog{
		timestamp: timestamp,
		requestId: requestId,
		service:   service,
		method:    method,
		userId:    userId,
		request:   request,
		result:    result,
	}
	return data
}

func (p *AccessLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.timestamp
	properties["RequestId"] = p.requestId
	properties["Service"] = p.service
	properties["Method"] = p.method
	if p.userId != nil {
		properties["UserId"] = p.userId
	}
	properties["Request"] = p.request
	properties["Result"] = p.result
	return properties
}

type AccessLogCount struct {
	service *string
	method  *string
	userId  *string
	count   int64
}

func NewAccessLogCount(
	service *string,
	method *string,
	userId *string,
	count int64,
) AccessLogCount {
	data := AccessLogCount{
		service: service,
		method:  method,
		userId:  userId,
		count:   count,
	}
	return data
}

func (p *AccessLogCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.service != nil {
		properties["Service"] = p.service
	}
	if p.method != nil {
		properties["Method"] = p.method
	}
	if p.userId != nil {
		properties["UserId"] = p.userId
	}
	properties["Count"] = p.count
	return properties
}

type IssueStampSheetLog struct {
	timestamp     int64
	transactionId string
	service       string
	method        string
	userId        string
	action        string
	args          string
	tasks         []string
}

func NewIssueStampSheetLog(
	timestamp int64,
	transactionId string,
	service string,
	method string,
	userId string,
	action string,
	args string,
	tasks []string,
) IssueStampSheetLog {
	data := IssueStampSheetLog{
		timestamp:     timestamp,
		transactionId: transactionId,
		service:       service,
		method:        method,
		userId:        userId,
		action:        action,
		args:          args,
		tasks:         tasks,
	}
	return data
}

func (p *IssueStampSheetLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.timestamp
	properties["TransactionId"] = p.transactionId
	properties["Service"] = p.service
	properties["Method"] = p.method
	properties["UserId"] = p.userId
	properties["Action"] = p.action
	properties["Args"] = p.args
	properties["Tasks"] = p.tasks
	return properties
}

type IssueStampSheetLogCount struct {
	service *string
	method  *string
	userId  *string
	action  *string
	count   int64
}

func NewIssueStampSheetLogCount(
	service *string,
	method *string,
	userId *string,
	action *string,
	count int64,
) IssueStampSheetLogCount {
	data := IssueStampSheetLogCount{
		service: service,
		method:  method,
		userId:  userId,
		action:  action,
		count:   count,
	}
	return data
}

func (p *IssueStampSheetLogCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.service != nil {
		properties["Service"] = p.service
	}
	if p.method != nil {
		properties["Method"] = p.method
	}
	if p.userId != nil {
		properties["UserId"] = p.userId
	}
	if p.action != nil {
		properties["Action"] = p.action
	}
	properties["Count"] = p.count
	return properties
}

type ExecuteStampSheetLog struct {
	timestamp     int64
	transactionId string
	service       string
	method        string
	userId        string
	action        string
	args          string
}

func NewExecuteStampSheetLog(
	timestamp int64,
	transactionId string,
	service string,
	method string,
	userId string,
	action string,
	args string,
) ExecuteStampSheetLog {
	data := ExecuteStampSheetLog{
		timestamp:     timestamp,
		transactionId: transactionId,
		service:       service,
		method:        method,
		userId:        userId,
		action:        action,
		args:          args,
	}
	return data
}

func (p *ExecuteStampSheetLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.timestamp
	properties["TransactionId"] = p.transactionId
	properties["Service"] = p.service
	properties["Method"] = p.method
	properties["UserId"] = p.userId
	properties["Action"] = p.action
	properties["Args"] = p.args
	return properties
}

type ExecuteStampSheetLogCount struct {
	service *string
	method  *string
	userId  *string
	action  *string
	count   int64
}

func NewExecuteStampSheetLogCount(
	service *string,
	method *string,
	userId *string,
	action *string,
	count int64,
) ExecuteStampSheetLogCount {
	data := ExecuteStampSheetLogCount{
		service: service,
		method:  method,
		userId:  userId,
		action:  action,
		count:   count,
	}
	return data
}

func (p *ExecuteStampSheetLogCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.service != nil {
		properties["Service"] = p.service
	}
	if p.method != nil {
		properties["Method"] = p.method
	}
	if p.userId != nil {
		properties["UserId"] = p.userId
	}
	if p.action != nil {
		properties["Action"] = p.action
	}
	properties["Count"] = p.count
	return properties
}

type ExecuteStampTaskLog struct {
	timestamp int64
	taskId    string
	service   string
	method    string
	userId    string
	action    string
	args      string
}

func NewExecuteStampTaskLog(
	timestamp int64,
	taskId string,
	service string,
	method string,
	userId string,
	action string,
	args string,
) ExecuteStampTaskLog {
	data := ExecuteStampTaskLog{
		timestamp: timestamp,
		taskId:    taskId,
		service:   service,
		method:    method,
		userId:    userId,
		action:    action,
		args:      args,
	}
	return data
}

func (p *ExecuteStampTaskLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.timestamp
	properties["TaskId"] = p.taskId
	properties["Service"] = p.service
	properties["Method"] = p.method
	properties["UserId"] = p.userId
	properties["Action"] = p.action
	properties["Args"] = p.args
	return properties
}

type ExecuteStampTaskLogCount struct {
	service *string
	method  *string
	userId  *string
	action  *string
	count   int64
}

func NewExecuteStampTaskLogCount(
	service *string,
	method *string,
	userId *string,
	action *string,
	count int64,
) ExecuteStampTaskLogCount {
	data := ExecuteStampTaskLogCount{
		service: service,
		method:  method,
		userId:  userId,
		action:  action,
		count:   count,
	}
	return data
}

func (p *ExecuteStampTaskLogCount) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	if p.service != nil {
		properties["Service"] = p.service
	}
	if p.method != nil {
		properties["Method"] = p.method
	}
	if p.userId != nil {
		properties["UserId"] = p.userId
	}
	if p.action != nil {
		properties["Action"] = p.action
	}
	properties["Count"] = p.count
	return properties
}
