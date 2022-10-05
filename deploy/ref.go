package deploy

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

type StackRef struct {
	stackName string
}

func (p *StackRef) Event(
	eventName string,
) *EventRef {
	return &EventRef{
		stackName: p.stackName,
		eventName: eventName,
	}
}

func (p *StackRef) Output(
	outputName string,
) *OutputRef {
	return &OutputRef{
		stackName:  p.stackName,
		outputName: outputName,
	}
}

func (p *StackRef) Resource(
	resourceName string,
) *ResourceRef {
	return &ResourceRef{
		stackName:    p.stackName,
		resourceName: resourceName,
	}
}

func (p *StackRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stack",
			p.stackName,
		},
	).String()
}

func (p *StackRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type ResourceRef struct {
	stackName    string
	resourceName string
}

func (p *ResourceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stack",
			p.stackName,
			"resource",
			p.resourceName,
		},
	).String()
}

func (p *ResourceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type WorkingStackRef struct {
	stackName string
}

func (p *WorkingStackRef) WorkingResource(
	resourceName string,
) *WorkingResourceRef {
	return &WorkingResourceRef{
		stackName:    p.stackName,
		resourceName: resourceName,
	}
}

func (p *WorkingStackRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"deploy",
			p.stackName,
		},
	).String()
}

func (p *WorkingStackRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type WorkingResourceRef struct {
	stackName    string
	resourceName string
}

func (p *WorkingResourceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"deploy",
			p.stackName,
			"workingResource",
			p.resourceName,
		},
	).String()
}

func (p *WorkingResourceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type EventRef struct {
	stackName string
	eventName string
}

func (p *EventRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stack",
			p.stackName,
			"event",
			p.eventName,
		},
	).String()
}

func (p *EventRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type OutputRef struct {
	stackName  string
	outputName string
}

func (p *OutputRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"stack",
			p.stackName,
			"output",
			p.outputName,
		},
	).String()
}

func (p *OutputRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
