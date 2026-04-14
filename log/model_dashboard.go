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

deny overwrite
*/

import (
	. "github.com/gs2io/gs2-golang-cdk/core"
)

var _ = AcquireAction{}

type Dashboard struct {
	CdkResource
	stack         *Stack
	NamespaceName string
	DisplayName   string
	Description   *string
}

type DashboardOptions struct {
	Description *string
}

func NewDashboard(
	stack *Stack,
	namespaceName string,
	displayName string,
	options DashboardOptions,
) *Dashboard {
	data := Dashboard{
		stack:         stack,
		NamespaceName: namespaceName,
		DisplayName:   displayName,
		Description:   options.Description,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *Dashboard) ResourceName() string {
	return "Log_Dashboard_" + p.DisplayName
}

func (p *Dashboard) ResourceType() string {
	return "GS2::Log::Dashboard"
}

func (p *Dashboard) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["DisplayName"] = p.DisplayName
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	return properties
}

func (p *Dashboard) Ref(
	namespaceName string,
) DashboardRef {
	return DashboardRef{
		NamespaceName: namespaceName,
		DashboardName: p.DisplayName,
	}
}

func (p *Dashboard) GetAttrDashboardId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.DashboardId",
	)
}
