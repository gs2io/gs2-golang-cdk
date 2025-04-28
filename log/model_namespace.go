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

type NamespaceType string

const NamespaceTypeGs2 = NamespaceType("gs2")
const NamespaceTypeBigquery = NamespaceType("bigquery")
const NamespaceTypeFirehose = NamespaceType("firehose")

func (p NamespaceType) Pointer() *NamespaceType {
	return &p
}

type NamespaceFirehoseCompressData string

const NamespaceFirehoseCompressDataNone = NamespaceFirehoseCompressData("none")
const NamespaceFirehoseCompressDataGzip = NamespaceFirehoseCompressData("gzip")

func (p NamespaceFirehoseCompressData) Pointer() *NamespaceFirehoseCompressData {
	return &p
}

type Namespace struct {
	CdkResource
	stack                *Stack
	Name                 string
	Description          *string
	Type_                NamespaceType
	GcpCredentialJson    *string
	BigQueryDatasetName  *string
	LogExpireDays        *int32
	AwsRegion            *string
	AwsAccessKeyId       *string
	AwsSecretAccessKey   *string
	FirehoseStreamName   *string
	FirehoseCompressData *NamespaceFirehoseCompressData
}

type NamespaceOptions struct {
	Description          *string
	Type_                NamespaceType
	GcpCredentialJson    *string
	BigQueryDatasetName  *string
	LogExpireDays        *int32
	AwsRegion            *string
	AwsAccessKeyId       *string
	AwsSecretAccessKey   *string
	FirehoseStreamName   *string
	FirehoseCompressData *NamespaceFirehoseCompressData
}

func NewNamespace(
	stack *Stack,
	name string,
	options NamespaceOptions,
) *Namespace {
	data := Namespace{
		stack:                stack,
		Name:                 name,
		Description:          options.Description,
		Type_:                options.Type_,
		GcpCredentialJson:    options.GcpCredentialJson,
		BigQueryDatasetName:  options.BigQueryDatasetName,
		LogExpireDays:        options.LogExpireDays,
		AwsRegion:            options.AwsRegion,
		AwsAccessKeyId:       options.AwsAccessKeyId,
		AwsSecretAccessKey:   options.AwsSecretAccessKey,
		FirehoseStreamName:   options.FirehoseStreamName,
		FirehoseCompressData: options.FirehoseCompressData,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Namespace) ResourceName() string {
	return "Log_Namespace_" + p.Name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Log::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["Type"] = p.Type_
	if p.GcpCredentialJson != nil {
		properties["GcpCredentialJson"] = p.GcpCredentialJson
	}
	if p.BigQueryDatasetName != nil {
		properties["BigQueryDatasetName"] = p.BigQueryDatasetName
	}
	if p.LogExpireDays != nil {
		properties["LogExpireDays"] = p.LogExpireDays
	}
	if p.AwsRegion != nil {
		properties["AwsRegion"] = p.AwsRegion
	}
	if p.AwsAccessKeyId != nil {
		properties["AwsAccessKeyId"] = p.AwsAccessKeyId
	}
	if p.AwsSecretAccessKey != nil {
		properties["AwsSecretAccessKey"] = p.AwsSecretAccessKey
	}
	if p.FirehoseStreamName != nil {
		properties["FirehoseStreamName"] = p.FirehoseStreamName
	}
	if p.FirehoseCompressData != nil {
		properties["FirehoseCompressData"] = p.FirehoseCompressData
	}
	return properties
}

func (p *Namespace) Ref() NamespaceRef {
	return NamespaceRef{
		NamespaceName: p.Name,
	}
}

func (p *Namespace) GetAttrNamespaceId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.NamespaceId",
	)
}
