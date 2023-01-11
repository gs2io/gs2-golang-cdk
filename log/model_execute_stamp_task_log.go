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

type ExecuteStampTaskLog struct {
	Timestamp int64
	TaskId    string
	Service   string
	Method    string
	UserId    string
	Action    string
	Args      string
}

type ExecuteStampTaskLogOptions struct {
}

func NewExecuteStampTaskLog(
	timestamp int64,
	taskId string,
	service string,
	method string,
	userId string,
	action string,
	args string,
	options ExecuteStampTaskLogOptions,
) ExecuteStampTaskLog {
	data := ExecuteStampTaskLog{
		Timestamp: timestamp,
		TaskId:    taskId,
		Service:   service,
		Method:    method,
		UserId:    userId,
		Action:    action,
		Args:      args,
	}
	return data
}

func (p *ExecuteStampTaskLog) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.Timestamp
	properties["TaskId"] = p.TaskId
	properties["Service"] = p.Service
	properties["Method"] = p.Method
	properties["UserId"] = p.UserId
	properties["Action"] = p.Action
	properties["Args"] = p.Args
	return properties
}
