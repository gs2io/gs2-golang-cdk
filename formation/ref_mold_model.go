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

type MoldModelRef struct {
	NamespaceName string
	MoldName      string
}

func (p *MoldModelRef) FormModel(
	formModelName string,
) *FormModelRef {
	return &FormModelRef{
		NamespaceName: p.NamespaceName,
		MoldName:      p.MoldName,
		FormModelName: formModelName,
	}
}

func (p *MoldModelRef) AddMoldCapacity(
	capacity int32,
) AcquireAction {
	return AddMoldCapacityByUserId(
		p.NamespaceName,
		p.MoldName,
		capacity,
	)
}

func (p *MoldModelRef) SetMoldCapacity(
	capacity int32,
) AcquireAction {
	return SetMoldCapacityByUserId(
		p.NamespaceName,
		p.MoldName,
		capacity,
	)
}

func (p *MoldModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.NamespaceName,
			"model",
			"mold",
			p.MoldName,
		},
	).String()
}

func (p *MoldModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
