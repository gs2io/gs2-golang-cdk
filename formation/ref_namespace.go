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

type NamespaceRef struct {
	NamespaceName string
}

func (p *NamespaceRef) MoldModel(
	moldModelName string,
) *MoldModelRef {
	return &MoldModelRef{
		NamespaceName: p.NamespaceName,
		MoldModelName: moldModelName,
	}
}

func (p *NamespaceRef) PropertyFormModel(
	propertyFormModelName string,
) *PropertyFormModelRef {
	return &PropertyFormModelRef{
		NamespaceName:         p.NamespaceName,
		PropertyFormModelName: propertyFormModelName,
	}
}

func (p *NamespaceRef) AddMoldCapacity(
	moldName string,
	capacity int32,
) AcquireAction {
	return AddMoldCapacityByUserId(
		p.NamespaceName,
		moldName,
		capacity,
	)
}

func (p *NamespaceRef) SetMoldCapacity(
	moldName string,
	capacity int32,
) AcquireAction {
	return SetMoldCapacityByUserId(
		p.NamespaceName,
		moldName,
		capacity,
	)
}

func (p *NamespaceRef) SubMoldCapacity(
	moldName string,
	capacity int32,
) ConsumeAction {
	return SubMoldCapacityByUserId(
		p.NamespaceName,
		moldName,
		capacity,
	)
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"formation",
			p.NamespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
