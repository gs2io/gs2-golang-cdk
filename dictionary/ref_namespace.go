package dictionary

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

func (p *NamespaceRef) EntryModel(
	entryModelName string,
) *EntryModelRef {
	return &EntryModelRef{
		NamespaceName:  p.NamespaceName,
		EntryModelName: entryModelName,
	}
}

func (p *NamespaceRef) AddEntries(
	entryModelNames *[]string,
) AcquireAction {
	return AddEntriesByUserId(
		p.NamespaceName,
		entryModelNames,
	)
}

func (p *NamespaceRef) DeleteEntries(
	entryModelNames *[]string,
) ConsumeAction {
	return DeleteEntriesByUserId(
		p.NamespaceName,
		entryModelNames,
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
			"dictionary",
			p.NamespaceName,
		},
	).String()
}

func (p *NamespaceRef) GrnPointer() *string {
	grn := p.Grn()
	return &grn
}
