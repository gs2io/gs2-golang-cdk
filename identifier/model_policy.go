package identifier

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

type Policy struct {
	version    string
	statements []Statement
}

func NewPolicy(
	statements []Statement,
) Policy {
	return Policy{
		version:    "2016-04-01",
		statements: statements,
	}
}

func (p *Policy) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Version"] = p.version
	statements := make([]map[string]interface{}, len(p.statements))
	for i, statement := range p.statements {
		statements[i] = statement.Properties()
	}
	if p.statements != nil {
		properties["Statements"] = statements
	}
	return properties
}
