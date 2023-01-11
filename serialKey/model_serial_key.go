package serialKey

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

type SerialKeyStatus string

const SerialKeyStatusActive = SerialKeyStatus("ACTIVE")
const SerialKeyStatusUsed = SerialKeyStatus("USED")
const SerialKeyStatusInactive = SerialKeyStatus("INACTIVE")

func (p SerialKeyStatus) Pointer() *SerialKeyStatus {
	return &p
}

type SerialKey struct {
	CampaignModelName string
	Code              string
	Metadata          *string
	Status            SerialKeyStatus
	UsedUserId        *string
	CreatedAt         int64
	UsedAt            *int64
	UpdatedAt         int64
}

type SerialKeyOptions struct {
	Metadata   *string
	UsedUserId *string
	UsedAt     *int64
}

func NewSerialKey(
	campaignModelName string,
	code string,
	status SerialKeyStatus,
	createdAt int64,
	updatedAt int64,
	options SerialKeyOptions,
) SerialKey {
	data := SerialKey{
		CampaignModelName: campaignModelName,
		Code:              code,
		Status:            status,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		Metadata:          options.Metadata,
		UsedUserId:        options.UsedUserId,
		UsedAt:            options.UsedAt,
	}
	return data
}

type SerialKeyStatusIsActiveOptions struct {
	Metadata *string
	UsedAt   *int64
}

func NewSerialKeyStatusIsActive(
	campaignModelName string,
	code string,
	createdAt int64,
	updatedAt int64,
	options SerialKeyStatusIsActiveOptions,
) SerialKey {
	return NewSerialKey(
		campaignModelName,
		code,
		SerialKeyStatusActive,
		createdAt,
		updatedAt,
		SerialKeyOptions{
			Metadata: options.Metadata,
			UsedAt:   options.UsedAt,
		},
	)
}

type SerialKeyStatusIsUsedOptions struct {
	Metadata *string
	UsedAt   *int64
}

func NewSerialKeyStatusIsUsed(
	campaignModelName string,
	code string,
	createdAt int64,
	updatedAt int64,
	usedUserId string,
	options SerialKeyStatusIsUsedOptions,
) SerialKey {
	return NewSerialKey(
		campaignModelName,
		code,
		SerialKeyStatusUsed,
		createdAt,
		updatedAt,
		SerialKeyOptions{
			Metadata:   options.Metadata,
			UsedUserId: &usedUserId,
			UsedAt:     options.UsedAt,
		},
	)
}

type SerialKeyStatusIsInactiveOptions struct {
	Metadata *string
	UsedAt   *int64
}

func NewSerialKeyStatusIsInactive(
	campaignModelName string,
	code string,
	createdAt int64,
	updatedAt int64,
	options SerialKeyStatusIsInactiveOptions,
) SerialKey {
	return NewSerialKey(
		campaignModelName,
		code,
		SerialKeyStatusInactive,
		createdAt,
		updatedAt,
		SerialKeyOptions{
			Metadata: options.Metadata,
			UsedAt:   options.UsedAt,
		},
	)
}

func (p *SerialKey) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["CampaignModelName"] = p.CampaignModelName
	properties["Code"] = p.Code
	if p.Metadata != nil {
		properties["Metadata"] = p.Metadata
	}
	properties["Status"] = p.Status
	if p.UsedUserId != nil {
		properties["UsedUserId"] = p.UsedUserId
	}
	properties["CreatedAt"] = p.CreatedAt
	if p.UsedAt != nil {
		properties["UsedAt"] = p.UsedAt
	}
	properties["UpdatedAt"] = p.UpdatedAt
	return properties
}
