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

type SlotWithSignaturePropertyType string

const SlotWithSignaturePropertyTypeGs2Inventory = SlotWithSignaturePropertyType("gs2_inventory")
const SlotWithSignaturePropertyTypeGs2SimpleInventory = SlotWithSignaturePropertyType("gs2_simple_inventory")
const SlotWithSignaturePropertyTypeGs2Dictionary = SlotWithSignaturePropertyType("gs2_dictionary")

func (p SlotWithSignaturePropertyType) Pointer() *SlotWithSignaturePropertyType {
	return &p
}

type SlotWithSignature struct {
	Name         string
	PropertyType SlotWithSignaturePropertyType
	Body         *string
	Signature    *string
	Metadata     *string
}

type SlotWithSignatureOptions struct {
	Body      *string
	Signature *string
	Metadata  *string
}

func NewSlotWithSignature(
	name string,
	propertyType SlotWithSignaturePropertyType,
	options SlotWithSignatureOptions,
) SlotWithSignature {
	_data := SlotWithSignature{
		Name:         name,
		PropertyType: propertyType,
		Body:         options.Body,
		Signature:    options.Signature,
		Metadata:     options.Metadata,
	}
	return _data
}

func (p *SlotWithSignature) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["PropertyType"] = p.PropertyType
	if p.Body != nil {
		properties["Body"] = p.Body
	}
	if p.Signature != nil {
		properties["Signature"] = p.Signature
	}
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	return properties
}
