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

type Namespace struct {
	CdkResource
	stack               *Stack
	name                string
	description         *string
	Type                string
	gcpCredentialJson   *string
	bigQueryDatasetName *string
	logExpireDays       *int32
	awsRegion           *string
	awsAccessKeyId      *string
	awsSecretAccessKey  *string
	firehoseStreamName  *string
}

func NewNamespace(
	stack *Stack,
	name string,
	Type string,
	description *string,
	gcpCredentialJson *string,
	bigQueryDatasetName *string,
	logExpireDays *int32,
	awsRegion *string,
	awsAccessKeyId *string,
	awsSecretAccessKey *string,
	firehoseStreamName *string,
) *Namespace {

	self := Namespace{
		stack:               stack,
		name:                name,
		description:         description,
		Type:                Type,
		gcpCredentialJson:   gcpCredentialJson,
		bigQueryDatasetName: bigQueryDatasetName,
		logExpireDays:       logExpireDays,
		awsRegion:           awsRegion,
		awsAccessKeyId:      awsAccessKeyId,
		awsSecretAccessKey:  awsSecretAccessKey,
		firehoseStreamName:  firehoseStreamName,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Log_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Log::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["Type"] = p.Type
	if p.gcpCredentialJson != nil {
		properties["GcpCredentialJson"] = p.gcpCredentialJson
	}
	if p.bigQueryDatasetName != nil {
		properties["BigQueryDatasetName"] = p.bigQueryDatasetName
	}
	if p.logExpireDays != nil {
		properties["LogExpireDays"] = p.logExpireDays
	}
	if p.awsRegion != nil {
		properties["AwsRegion"] = p.awsRegion
	}
	if p.awsAccessKeyId != nil {
		properties["AwsAccessKeyId"] = p.awsAccessKeyId
	}
	if p.awsSecretAccessKey != nil {
		properties["AwsSecretAccessKey"] = p.awsSecretAccessKey
	}
	if p.firehoseStreamName != nil {
		properties["FirehoseStreamName"] = p.firehoseStreamName
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		namespaceName: p.name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}
