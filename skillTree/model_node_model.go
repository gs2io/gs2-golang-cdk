package skillTree

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

type NodeModel struct {
	Name                  string
	Metadata              *string
	ReleaseVerifyActions  []VerifyAction
	ReleaseConsumeActions []ConsumeAction
	ReturnAcquireActions  []AcquireAction
	RestrainReturnRate    float32
	PremiseNodeNames      []string
}

type NodeModelOptions struct {
	Metadata             *string
	ReleaseVerifyActions []VerifyAction
	ReturnAcquireActions []AcquireAction
	PremiseNodeNames     []string
}

func NewNodeModel(
	name string,
	releaseConsumeActions []ConsumeAction,
	restrainReturnRate float32,
	options NodeModelOptions,
) NodeModel {
	data := NodeModel{
		Name:                  name,
		ReleaseConsumeActions: releaseConsumeActions,
		RestrainReturnRate:    restrainReturnRate,
		Metadata:              options.Metadata,
		ReleaseVerifyActions:  options.ReleaseVerifyActions,
		ReturnAcquireActions:  options.ReturnAcquireActions,
		PremiseNodeNames:      options.PremiseNodeNames,
	}
	return data
}

func (p *NodeModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	{
		values := make([]map[string]interface{}, len(p.ReleaseVerifyActions))
		for i, element := range p.ReleaseVerifyActions {
			values[i] = element.Properties()
		}
		properties["ReleaseVerifyActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ReleaseConsumeActions))
		for i, element := range p.ReleaseConsumeActions {
			values[i] = element.Properties()
		}
		properties["ReleaseConsumeActions"] = values
	}
	{
		values := make([]map[string]interface{}, len(p.ReturnAcquireActions))
		for i, element := range p.ReturnAcquireActions {
			values[i] = element.Properties()
		}
		properties["ReturnAcquireActions"] = values
	}
	properties["RestrainReturnRate"] = p.RestrainReturnRate
	properties["PremiseNodeNames"] = p.PremiseNodeNames
	return properties
}
