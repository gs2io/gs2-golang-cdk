package guard

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

type BlockingPolicyModelDefaultRestriction string

const BlockingPolicyModelDefaultRestrictionAllow = BlockingPolicyModelDefaultRestriction("Allow")
const BlockingPolicyModelDefaultRestrictionDeny = BlockingPolicyModelDefaultRestriction("Deny")

func (p BlockingPolicyModelDefaultRestriction) Pointer() *BlockingPolicyModelDefaultRestriction {
	return &p
}

type BlockingPolicyModelLocationDetection string

const BlockingPolicyModelLocationDetectionEnable = BlockingPolicyModelLocationDetection("Enable")
const BlockingPolicyModelLocationDetectionDisable = BlockingPolicyModelLocationDetection("Disable")

func (p BlockingPolicyModelLocationDetection) Pointer() *BlockingPolicyModelLocationDetection {
	return &p
}

type BlockingPolicyModelLocationRestriction string

const BlockingPolicyModelLocationRestrictionAllow = BlockingPolicyModelLocationRestriction("Allow")
const BlockingPolicyModelLocationRestrictionDeny = BlockingPolicyModelLocationRestriction("Deny")

func (p BlockingPolicyModelLocationRestriction) Pointer() *BlockingPolicyModelLocationRestriction {
	return &p
}

type BlockingPolicyModelAnonymousIpDetection string

const BlockingPolicyModelAnonymousIpDetectionEnable = BlockingPolicyModelAnonymousIpDetection("Enable")
const BlockingPolicyModelAnonymousIpDetectionDisable = BlockingPolicyModelAnonymousIpDetection("Disable")

func (p BlockingPolicyModelAnonymousIpDetection) Pointer() *BlockingPolicyModelAnonymousIpDetection {
	return &p
}

type BlockingPolicyModelAnonymousIpRestriction string

const BlockingPolicyModelAnonymousIpRestrictionDeny = BlockingPolicyModelAnonymousIpRestriction("Deny")

func (p BlockingPolicyModelAnonymousIpRestriction) Pointer() *BlockingPolicyModelAnonymousIpRestriction {
	return &p
}

type BlockingPolicyModelHostingProviderIpDetection string

const BlockingPolicyModelHostingProviderIpDetectionEnable = BlockingPolicyModelHostingProviderIpDetection("Enable")
const BlockingPolicyModelHostingProviderIpDetectionDisable = BlockingPolicyModelHostingProviderIpDetection("Disable")

func (p BlockingPolicyModelHostingProviderIpDetection) Pointer() *BlockingPolicyModelHostingProviderIpDetection {
	return &p
}

type BlockingPolicyModelHostingProviderIpRestriction string

const BlockingPolicyModelHostingProviderIpRestrictionDeny = BlockingPolicyModelHostingProviderIpRestriction("Deny")

func (p BlockingPolicyModelHostingProviderIpRestriction) Pointer() *BlockingPolicyModelHostingProviderIpRestriction {
	return &p
}

type BlockingPolicyModelReputationIpDetection string

const BlockingPolicyModelReputationIpDetectionEnable = BlockingPolicyModelReputationIpDetection("Enable")
const BlockingPolicyModelReputationIpDetectionDisable = BlockingPolicyModelReputationIpDetection("Disable")

func (p BlockingPolicyModelReputationIpDetection) Pointer() *BlockingPolicyModelReputationIpDetection {
	return &p
}

type BlockingPolicyModelReputationIpRestriction string

const BlockingPolicyModelReputationIpRestrictionDeny = BlockingPolicyModelReputationIpRestriction("Deny")

func (p BlockingPolicyModelReputationIpRestriction) Pointer() *BlockingPolicyModelReputationIpRestriction {
	return &p
}

type BlockingPolicyModelIpAddressesDetection string

const BlockingPolicyModelIpAddressesDetectionEnable = BlockingPolicyModelIpAddressesDetection("Enable")
const BlockingPolicyModelIpAddressesDetectionDisable = BlockingPolicyModelIpAddressesDetection("Disable")

func (p BlockingPolicyModelIpAddressesDetection) Pointer() *BlockingPolicyModelIpAddressesDetection {
	return &p
}

type BlockingPolicyModelIpAddressRestriction string

const BlockingPolicyModelIpAddressRestrictionAllow = BlockingPolicyModelIpAddressRestriction("Allow")
const BlockingPolicyModelIpAddressRestrictionDeny = BlockingPolicyModelIpAddressRestriction("Deny")

func (p BlockingPolicyModelIpAddressRestriction) Pointer() *BlockingPolicyModelIpAddressRestriction {
	return &p
}

type BlockingPolicyModel struct {
	PassServices                 []string
	DefaultRestriction           BlockingPolicyModelDefaultRestriction
	LocationDetection            BlockingPolicyModelLocationDetection
	Locations                    []string
	LocationRestriction          *BlockingPolicyModelLocationRestriction
	AnonymousIpDetection         BlockingPolicyModelAnonymousIpDetection
	AnonymousIpRestriction       *BlockingPolicyModelAnonymousIpRestriction
	HostingProviderIpDetection   BlockingPolicyModelHostingProviderIpDetection
	HostingProviderIpRestriction *BlockingPolicyModelHostingProviderIpRestriction
	ReputationIpDetection        BlockingPolicyModelReputationIpDetection
	ReputationIpRestriction      *BlockingPolicyModelReputationIpRestriction
	IpAddressesDetection         BlockingPolicyModelIpAddressesDetection
	IpAddresses                  []string
	IpAddressRestriction         *BlockingPolicyModelIpAddressRestriction
}

type BlockingPolicyModelOptions struct {
	Locations                    []string
	LocationRestriction          *BlockingPolicyModelLocationRestriction
	AnonymousIpRestriction       *BlockingPolicyModelAnonymousIpRestriction
	HostingProviderIpRestriction *BlockingPolicyModelHostingProviderIpRestriction
	ReputationIpRestriction      *BlockingPolicyModelReputationIpRestriction
	IpAddresses                  []string
	IpAddressRestriction         *BlockingPolicyModelIpAddressRestriction
}

func NewBlockingPolicyModel(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	options BlockingPolicyModelOptions,
) BlockingPolicyModel {
	_data := BlockingPolicyModel{
		PassServices:                 passServices,
		DefaultRestriction:           defaultRestriction,
		LocationDetection:            locationDetection,
		AnonymousIpDetection:         anonymousIpDetection,
		HostingProviderIpDetection:   hostingProviderIpDetection,
		ReputationIpDetection:        reputationIpDetection,
		IpAddressesDetection:         ipAddressesDetection,
		Locations:                    options.Locations,
		LocationRestriction:          options.LocationRestriction,
		AnonymousIpRestriction:       options.AnonymousIpRestriction,
		HostingProviderIpRestriction: options.HostingProviderIpRestriction,
		ReputationIpRestriction:      options.ReputationIpRestriction,
		IpAddresses:                  options.IpAddresses,
		IpAddressRestriction:         options.IpAddressRestriction,
	}
	return _data
}

type BlockingPolicyModelLocationDetectionIsEnableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelLocationDetectionIsEnable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	locations []string,
	locationRestriction BlockingPolicyModelLocationRestriction,
	options BlockingPolicyModelLocationDetectionIsEnableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		BlockingPolicyModelLocationDetectionEnable,
		anonymousIpDetection,
		hostingProviderIpDetection,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			Locations:           locations,
			LocationRestriction: &locationRestriction,
			IpAddresses:         options.IpAddresses,
		},
	)
}

type BlockingPolicyModelLocationDetectionIsDisableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelLocationDetectionIsDisable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	options BlockingPolicyModelLocationDetectionIsDisableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		BlockingPolicyModelLocationDetectionDisable,
		anonymousIpDetection,
		hostingProviderIpDetection,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			IpAddresses: options.IpAddresses,
		},
	)
}

type BlockingPolicyModelAnonymousIpDetectionIsEnableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelAnonymousIpDetectionIsEnable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	anonymousIpRestriction BlockingPolicyModelAnonymousIpRestriction,
	options BlockingPolicyModelAnonymousIpDetectionIsEnableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		BlockingPolicyModelAnonymousIpDetectionEnable,
		hostingProviderIpDetection,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			AnonymousIpRestriction: &anonymousIpRestriction,
			IpAddresses:            options.IpAddresses,
		},
	)
}

type BlockingPolicyModelAnonymousIpDetectionIsDisableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelAnonymousIpDetectionIsDisable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	options BlockingPolicyModelAnonymousIpDetectionIsDisableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		BlockingPolicyModelAnonymousIpDetectionDisable,
		hostingProviderIpDetection,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			IpAddresses: options.IpAddresses,
		},
	)
}

type BlockingPolicyModelHostingProviderIpDetectionIsEnableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelHostingProviderIpDetectionIsEnable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	hostingProviderIpRestriction BlockingPolicyModelHostingProviderIpRestriction,
	options BlockingPolicyModelHostingProviderIpDetectionIsEnableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		BlockingPolicyModelHostingProviderIpDetectionEnable,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			HostingProviderIpRestriction: &hostingProviderIpRestriction,
			IpAddresses:                  options.IpAddresses,
		},
	)
}

type BlockingPolicyModelHostingProviderIpDetectionIsDisableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelHostingProviderIpDetectionIsDisable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	options BlockingPolicyModelHostingProviderIpDetectionIsDisableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		BlockingPolicyModelHostingProviderIpDetectionDisable,
		reputationIpDetection,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			IpAddresses: options.IpAddresses,
		},
	)
}

type BlockingPolicyModelReputationIpDetectionIsEnableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelReputationIpDetectionIsEnable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	reputationIpRestriction BlockingPolicyModelReputationIpRestriction,
	options BlockingPolicyModelReputationIpDetectionIsEnableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		hostingProviderIpDetection,
		BlockingPolicyModelReputationIpDetectionEnable,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			ReputationIpRestriction: &reputationIpRestriction,
			IpAddresses:             options.IpAddresses,
		},
	)
}

type BlockingPolicyModelReputationIpDetectionIsDisableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelReputationIpDetectionIsDisable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	ipAddressesDetection BlockingPolicyModelIpAddressesDetection,
	options BlockingPolicyModelReputationIpDetectionIsDisableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		hostingProviderIpDetection,
		BlockingPolicyModelReputationIpDetectionDisable,
		ipAddressesDetection,
		BlockingPolicyModelOptions{
			IpAddresses: options.IpAddresses,
		},
	)
}

type BlockingPolicyModelIpAddressesDetectionIsEnableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelIpAddressesDetectionIsEnable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	ipAddressRestriction BlockingPolicyModelIpAddressRestriction,
	options BlockingPolicyModelIpAddressesDetectionIsEnableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		hostingProviderIpDetection,
		reputationIpDetection,
		BlockingPolicyModelIpAddressesDetectionEnable,
		BlockingPolicyModelOptions{
			IpAddresses:          options.IpAddresses,
			IpAddressRestriction: &ipAddressRestriction,
		},
	)
}

type BlockingPolicyModelIpAddressesDetectionIsDisableOptions struct {
	IpAddresses []string
}

func NewBlockingPolicyModelIpAddressesDetectionIsDisable(
	passServices []string,
	defaultRestriction BlockingPolicyModelDefaultRestriction,
	locationDetection BlockingPolicyModelLocationDetection,
	anonymousIpDetection BlockingPolicyModelAnonymousIpDetection,
	hostingProviderIpDetection BlockingPolicyModelHostingProviderIpDetection,
	reputationIpDetection BlockingPolicyModelReputationIpDetection,
	options BlockingPolicyModelIpAddressesDetectionIsDisableOptions,
) BlockingPolicyModel {
	return NewBlockingPolicyModel(
		passServices,
		defaultRestriction,
		locationDetection,
		anonymousIpDetection,
		hostingProviderIpDetection,
		reputationIpDetection,
		BlockingPolicyModelIpAddressesDetectionDisable,
		BlockingPolicyModelOptions{
			IpAddresses: options.IpAddresses,
		},
	)
}

func (p *BlockingPolicyModel) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["PassServices"] = p.PassServices
	properties["DefaultRestriction"] = p.DefaultRestriction
	properties["LocationDetection"] = p.LocationDetection
	properties["Locations"] = p.Locations
	if p.LocationRestriction != nil {
		properties["LocationRestriction"] = p.LocationRestriction
	}
	properties["AnonymousIpDetection"] = p.AnonymousIpDetection
	if p.AnonymousIpRestriction != nil {
		properties["AnonymousIpRestriction"] = p.AnonymousIpRestriction
	}
	properties["HostingProviderIpDetection"] = p.HostingProviderIpDetection
	if p.HostingProviderIpRestriction != nil {
		properties["HostingProviderIpRestriction"] = p.HostingProviderIpRestriction
	}
	properties["ReputationIpDetection"] = p.ReputationIpDetection
	if p.ReputationIpRestriction != nil {
		properties["ReputationIpRestriction"] = p.ReputationIpRestriction
	}
	properties["IpAddressesDetection"] = p.IpAddressesDetection
	properties["IpAddresses"] = p.IpAddresses
	if p.IpAddressRestriction != nil {
		properties["IpAddressRestriction"] = p.IpAddressRestriction
	}
	return properties
}
