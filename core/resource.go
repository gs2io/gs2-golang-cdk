package core

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
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

type ICdkResource interface {
	ResourceName() string
	ResourceType() string
	Properties() map[string]interface{}
}

type CdkResource struct {
	resource  ICdkResource
	dependsOn []string
}

func NewCdkResource(
	resource ICdkResource,
) CdkResource {
	return CdkResource{
		resource:  resource,
		dependsOn: make([]string, 0),
	}
}

func (p *CdkResource) AddDependsOn(
	resource ICdkResource,
) *CdkResource {
	p.dependsOn = append(p.dependsOn, resource.ResourceName())
	return p
}

func (p *CdkResource) Template() map[string]interface{} {
	return map[string]interface{}{
		"Type":       p.resource.ResourceType(),
		"Properties": p.resource.Properties(),
		"DependsOn":  p.dependsOn,
	}
}

type Stack struct {
	Resources []*CdkResource
	Outputs   map[string]string
}

func NewStack() Stack {
	return Stack{
		Resources: make([]*CdkResource, 0),
		Outputs:   map[string]string{},
	}
}

func (p *Stack) AddResource(
	resource *CdkResource,
) {
	p.Resources = append(p.Resources, resource)
}

func (p *Stack) Output(
	name string,
	path string,
) {
	p.Outputs[name] = path
}

func (p *Stack) Template() map[string]interface{} {
	properties := make([]map[string]interface{}, len(p.Resources))
	for i, resource := range p.Resources {
		properties[i] = resource.Template()
	}
	return map[string]interface{}{
		"GS2TemplateFormatVersion": "2019-05-01",
		"Resources":                properties,
		"Outputs":                  p.Outputs,
	}
}

func (p *Stack) Yaml() (string, error) {
	data, err := yaml.Marshal(p.Template())
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile("'!(.*) (.*)'")
	data = re.ReplaceAll([]byte(strings.ReplaceAll(string(data), "''", "'")), []byte("!$1 $2"))
	return string(data), nil
}
