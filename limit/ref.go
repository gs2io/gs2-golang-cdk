package limit

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

func (p *NamespaceRef) CurrentLimitMaster() *CurrentLimitMasterRef {
	return &CurrentLimitMasterRef{
		namespaceName: p.namespaceName,
	}
}

func (p *NamespaceRef) LimitModel(
	limitName string,
) *LimitModelRef {
	return &LimitModelRef{
		namespaceName: p.namespaceName,
		limitName:     limitName,
	}
}

func (p *NamespaceRef) LimitModelMaster(
	limitName string,
) *LimitModelMasterRef {
	return &LimitModelMasterRef{
		namespaceName: p.namespaceName,
		limitName:     limitName,
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
			"limit",
			p.namespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LimitModelMasterRef struct {
	namespaceName string
	limitName     string
}

func (p *LimitModelMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"limit",
			p.namespaceName,
			"limit",
			p.limitName,
		},
	).String()
}

func (p *LimitModelMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type CurrentLimitMasterRef struct {
	namespaceName string
}

func (p *CurrentLimitMasterRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"limit",
			p.namespaceName,
		},
	).String()
}

func (p *CurrentLimitMasterRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}

type LimitModelRef struct {
	namespaceName string
	limitName     string
}

func (p *LimitModelRef) Grn() string {
	return NewJoin(
		":",
		[]string{
			"grn",
			"gs2",
			NewGetAttrRegion().String(),
			NewGetAttrOwnerId().String(),
			"limit",
			p.namespaceName,
			"limit",
			p.limitName,
		},
	).String()
}

func (p *LimitModelRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
