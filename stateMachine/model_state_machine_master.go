package stateMachine

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

type StateMachineMaster struct {
	CdkResource
	stack                *Stack
	version              string
	namespaceName        string
	mainStateMachineName string
	payload              string
}

type StateMachineMasterOptions struct {
}

func NewStateMachineMaster(
	stack *Stack,
	namespaceName string,
	mainStateMachineName string,
	payload string,
	options StateMachineMasterOptions,
) *StateMachineMaster {
	self := StateMachineMaster{
		stack:                stack,
		namespaceName:        namespaceName,
		mainStateMachineName: mainStateMachineName,
		payload:              payload,
	}
	self.CdkResource = NewCdkResource(&self)
	stack.AddResource(&self.CdkResource)
	return &self
}

func (p *StateMachineMaster) AlternateKeys() string {
	return "version"
}

func (p *StateMachineMaster) ResourceName() string {
	return "StateMachine_StateMachineMaster_" + p.namespaceName
}

func (p *StateMachineMaster) ResourceType() string { return "GS2::StateMachine::StateMachineMaster" }

func (p *StateMachineMaster) Properties() map[string]interface{} {
	properties := map[string]interface{}{}

	properties["NamespaceName"] = p.namespaceName

	properties["MainStateMachineName"] = p.mainStateMachineName

	properties["payload"] = p.payload

	return properties
}
