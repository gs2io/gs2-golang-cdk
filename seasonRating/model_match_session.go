package seasonRating

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

type MatchSession struct {
	CdkResource
	stack         *Stack
	NamespaceName string
	SessionName   *string
	TtlSeconds    int32
}

type MatchSessionOptions struct {
	SessionName *string
	TtlSeconds  int32
}

func NewMatchSession(
	stack *Stack,
	namespaceName string,
	options MatchSessionOptions,
) *MatchSession {
	data := MatchSession{
		stack:         stack,
		NamespaceName: namespaceName,
		SessionName:   options.SessionName,
		TtlSeconds:    options.TtlSeconds,
	}
	data.CdkResource = NewCdkResource(&data)
	stack.AddResource(&data.CdkResource)
	return &data
}

func (p *MatchSession) ResourceName() string {
	return "SeasonRating_MatchSession_" + p.Name
}

func (p *MatchSession) ResourceType() string {
	return "GS2::SeasonRating::MatchSession"
}

func (p *MatchSession) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["NamespaceName"] = p.NamespaceName
	if p.SessionName != nil {
		properties["SessionName"] = p.SessionName
	}
	properties["TtlSeconds"] = p.TtlSeconds
	return properties
}

func (p *MatchSession) Ref(
	namespaceName string,
) MatchSessionRef {
	return MatchSessionRef{
		NamespaceName: namespaceName,
		SessionName:   p.Name,
	}
}

func (p *MatchSession) GetAttrSessionId() GetAttr {
	return NewGetAttrByResource(
		p,
		"Item.SessionId",
	)
}
