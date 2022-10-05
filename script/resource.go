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

type Namespace struct {
	CdkResource
	stack       *Stack
	name        string
	description *string
	logSetting  *LogSetting
}

func NewNamespace(
	stack *Stack,
	name string,
	description *string,
	logSetting *LogSetting,
) *Namespace {

	self := Namespace{
		stack:       stack,
		name:        name,
		description: description,
		logSetting:  logSetting,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Namespace) ResourceName() string {
	return "Script_Namespace_" + p.name
}

func (p *Namespace) ResourceType() string {
	return "GS2::Script::Namespace"
}

func (p *Namespace) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	if p.logSetting != nil {
		properties["LogSetting"] = p.logSetting.Properties()
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

type Script struct {
	CdkResource
	stack         *Stack
	namespaceName string
	name          string
	description   *string
	script        string
}

func NewScript(
	stack *Stack,
	namespaceName string,
	name string,
	script string,
	description *string,
) *Script {

	self := Script{
		stack:         stack,
		namespaceName: namespaceName,
		name:          name,
		description:   description,
		script:        script,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)

	return &self
}

func (p *Script) ResourceName() string {
	return "Script_Script_" + p.name
}

func (p *Script) ResourceType() string {
	return "GS2::Script::Script"
}

func (p *Script) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.namespaceName
	properties["Name"] = p.name
	if p.description != nil {
		properties["Description"] = p.description
	}
	properties["Script"] = p.script
	return properties
}

func (p *Script) Ref(
	namespaceName string,
) ScriptRef {
	return ScriptRef{
		namespaceName: namespaceName,
		scriptName:    p.name,
	}
}

func (p *Script) GetAttrScriptId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.ScriptId",
	)
}
