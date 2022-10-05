package megaField

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
	namespaceName string
}

func (p *NamespaceRef) CurrentFieldMaster() *CurrentFieldMasterRef {
	return &CurrentFieldMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) AreaModel(
	areaModelName string,
) *AreaModelRef {
	return &AreaModelRef{
		namespaceName: p.namespaceName,
		areaModelName: areaModelName,
	}
}

func (p *NamespaceRef) Node(
	nodeName string,
) *NodeRef {
	return &NodeRef{
		namespaceName: p.namespaceName,
		nodeName:      nodeName,
	}
}

func (p *NamespaceRef) Layer(
	areaModelName string,
	layerModelName string,
) *LayerRef {
	return &LayerRef{
		namespaceName:  p.namespaceName,
		areaModelName:  areaModelName,
		layerModelName: layerModelName,
	}
}

func (p *NamespaceRef) AreaModelMaster(
	areaModelName string,
) *AreaModelMasterRef {
	return &AreaModelMasterRef{
		namespaceName: p.namespaceName,
		areaModelName: areaModelName,
	}
}

func (p *NamespaceRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type AreaModelRef struct {
	namespaceName string
	areaModelName string
}

func (p *AreaModelRef) LayerModel(
	layerModelName string,
) *LayerModelRef {
	return &LayerModelRef{
		namespaceName:  p.namespaceName,
		areaModelName:  p.areaModelName,
		layerModelName: layerModelName,
	}
}

func (p *AreaModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"model",
			"area",
			p.areaModelName,
		},
	).String()
}

func (p *AreaModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type AreaModelMasterRef struct {
	namespaceName string
	areaModelName string
}

func (p *AreaModelMasterRef) LayerModelMaster(
	layerModelName string,
) *LayerModelMasterRef {
	return &LayerModelMasterRef{
		namespaceName:  p.namespaceName,
		areaModelName:  p.areaModelName,
		layerModelName: layerModelName,
	}
}

func (p *AreaModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"master",
			"area",
			p.areaModelName,
		},
	).String()
}

func (p *AreaModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LayerModelRef struct {
	namespaceName  string
	areaModelName  string
	layerModelName string
}

func (p *LayerModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"model",
			"area",
			p.areaModelName,
			"layer",
			p.layerModelName,
		},
	).String()
}

func (p *LayerModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LayerModelMasterRef struct {
	namespaceName  string
	areaModelName  string
	layerModelName string
}

func (p *LayerModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"model",
			"area",
			p.areaModelName,
			"layer",
			p.layerModelName,
		},
	).String()
}

func (p *LayerModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentFieldMasterRef struct {
	namespaceName string
}

func (p *CurrentFieldMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentFieldMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type NodeRef struct {
	namespaceName string
	nodeName      string
}

func (p *NodeRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"node",
			p.nodeName,
		},
	).String()
}

func (p *NodeRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LayerRef struct {
	namespaceName  string
	areaModelName  string
	layerModelName string
}

func (p *LayerRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"megaField",
			p.namespaceName,
			"layer",
			p.areaModelName,
			p.layerModelName,
		},
	).String()
}

func (p *LayerRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
