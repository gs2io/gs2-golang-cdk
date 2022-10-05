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

type Slot struct {
	name       string
	propertyId string
	metadata   *string
}

func NewSlot(
	name string,
	propertyId string,
	metadata *string,
) Slot {
	data := Slot{
		name:       name,
		propertyId: propertyId,
		metadata:   metadata,
	}
	return data
}

func (p *Slot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["PropertyId"] = p.propertyId
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

type SlotModel struct {
	name          string
	propertyRegex string
	metadata      *string
}

func NewSlotModel(
	name string,
	propertyRegex string,
	metadata *string,
) SlotModel {
	data := SlotModel{
		name:          name,
		propertyRegex: propertyRegex,
		metadata:      metadata,
	}
	return data
}

func (p *SlotModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["PropertyRegex"] = p.propertyRegex
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

type SlotWithSignaturePropertyType string

const SlotWithSignaturePropertyTypeGs2Inventory = SlotWithSignaturePropertyType("gs2_inventory")
const SlotWithSignaturePropertyTypeGs2Dictionary = SlotWithSignaturePropertyType("gs2_dictionary")

func (p SlotWithSignaturePropertyType) Pointer() *SlotWithSignaturePropertyType {
	return &p
}

type SlotWithSignature struct {
	name         string
	propertyType SlotWithSignaturePropertyType
	body         string
	signature    string
	metadata     *string
}

func NewSlotWithSignature(
	name string,
	propertyType SlotWithSignaturePropertyType,
	body string,
	signature string,
	metadata *string,
) SlotWithSignature {
	data := SlotWithSignature{
		name:         name,
		propertyType: propertyType,
		body:         body,
		signature:    signature,
		metadata:     metadata,
	}
	return data
}

func (p *SlotWithSignature) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["PropertyType"] = p.propertyType
	properties["Body"] = p.body
	properties["Signature"] = p.signature
	if p.metadata != nil {
		properties["Metadata"] = p.metadata
	}
	return properties
}

type AcquireActionConfig struct {
	name   string
	config []Config
}

func NewAcquireActionConfig(
	name string,
	config []Config,
) AcquireActionConfig {
	data := AcquireActionConfig{
		name:   name,
		config: config,
	}
	return data
}

func (p *AcquireActionConfig) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	{
		values := make([]map[string]interface{}, len(p.config))
		for i, element := range p.config {
			values[i] = element.Properties()
		}
		properties["Config"] = values
	}
	return properties
}
