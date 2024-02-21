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

type AccessLogWithTelemetryStatus string

const AccessLogWithTelemetryStatusOk = AccessLogWithTelemetryStatus("ok")
const AccessLogWithTelemetryStatusError = AccessLogWithTelemetryStatus("error")

func (p AccessLogWithTelemetryStatus) Pointer() *AccessLogWithTelemetryStatus {
	return &p
}

type AccessLogWithTelemetry struct {
	Timestamp       int64
	SourceRequestId string
	RequestId       string
	Duration        int64
	Service         string
	Method          string
	UserId          *string
	Request         string
	Result          string
	Status          AccessLogWithTelemetryStatus
}

type AccessLogWithTelemetryOptions struct {
	UserId *string
}

func NewAccessLogWithTelemetry(
	timestamp int64,
	sourceRequestId string,
	requestId string,
	duration int64,
	service string,
	method string,
	request string,
	result string,
	status AccessLogWithTelemetryStatus,
	options AccessLogWithTelemetryOptions,
) AccessLogWithTelemetry {
	data := AccessLogWithTelemetry{
		Timestamp:       timestamp,
		SourceRequestId: sourceRequestId,
		RequestId:       requestId,
		Duration:        duration,
		Service:         service,
		Method:          method,
		Request:         request,
		Result:          result,
		Status:          status,
		UserId:          options.UserId,
	}
	return data
}

func (p *AccessLogWithTelemetry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Timestamp"] = p.Timestamp
	properties["SourceRequestId"] = p.SourceRequestId
	properties["RequestId"] = p.RequestId
	properties["Duration"] = p.Duration
	properties["Service"] = p.Service
	properties["Method"] = p.Method
	if p.UserId != nil {
		properties["UserId"] = p.UserId
	}
	properties["Request"] = p.Request
	properties["Result"] = p.Result
	properties["Status"] = p.Status
	return properties
}
