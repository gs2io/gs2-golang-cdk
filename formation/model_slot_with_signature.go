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
	Body         string
	Signature    string
	Metadata     *string
}

type SlotWithSignatureOptions struct {
	Metadata *string
}

func NewSlotWithSignature(
	name string,
	propertyType SlotWithSignaturePropertyType,
	body string,
	signature string,
	options SlotWithSignatureOptions,
) SlotWithSignature {
	data := SlotWithSignature{
		Name:         name,
		PropertyType: propertyType,
		Body:         body,
		Signature:    signature,
		Metadata:     options.Metadata,
	}
	return data
}

func (p *SlotWithSignature) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	properties["PropertyType"] = p.PropertyType
	properties["Body"] = p.Body
	properties["Signature"] = p.Signature
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	return properties
}
