package formation

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

type AcquireActionConfig struct {
	Name   string
	Config []Config
}

type AcquireActionConfigOptions struct {
	Config []Config
}

func NewAcquireActionConfig(
	name string,
	options AcquireActionConfigOptions,
) AcquireActionConfig {
	data := AcquireActionConfig{
		Name:   name,
		Config: options.Config,
	}
	return data
}

func (p *AcquireActionConfig) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	{
		values := make([]map[string]interface{}, len(p.Config))
		for i, element := range p.Config {
			values[i] = element.Properties()
		}
		properties["Config"] = values
	}
	return properties
}
