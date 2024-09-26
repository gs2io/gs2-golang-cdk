package script

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

type Script struct {
	CdkResource
	stack                       *Stack
	NamespaceName               string
	Name                        string
	Description                 *string
	Script                      string
	DisableStringNumberToNumber bool
}

type ScriptOptions struct {
	Description                 *string
	DisableStringNumberToNumber bool
}

func NewScript(
	stack *Stack,
	namespaceName string,
	name string,
	script string,
	options ScriptOptions,
) *Script {
	data := Script{
		stack:                       stack,
		NamespaceName:               namespaceName,
		Name:                        name,
		Script:                      script,
		Description:                 options.Description,
		DisableStringNumberToNumber: options.DisableStringNumberToNumber,
	}
	return &data
}

func (p *Script) ResourceName() string {
	return "Script_Script_" + p.Name
}

func (p *Script) ResourceType() string {
	return "GS2::Script::Script"
}

func (p *Script) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["Script"] = p.Script
	properties["DisableStringNumberToNumber"] = p.DisableStringNumberToNumber
	return properties
}

func (p *Script) Ref(
	namespaceName string,
) ScriptRef {
	return ScriptRef{
		NamespaceName: namespaceName,
		ScriptName:    p.Name,
	}
}

func (p *Script) GetAttrScriptId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ScriptId",
	)
}
