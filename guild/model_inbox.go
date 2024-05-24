package guild

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

type Inbox struct {
	GuildName   string
	FromUserIds []string
	CreatedAt   int64
	UpdatedAt   int64
	Revision    *int64
}

type InboxOptions struct {
	FromUserIds []string
	Revision    *int64
}

func NewInbox(
	guildName string,
	createdAt int64,
	updatedAt int64,
	options InboxOptions,
) Inbox {
	data := Inbox{
		GuildName:   guildName,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		FromUserIds: options.FromUserIds,
		Revision:    options.Revision,
	}
	return data
}

func (p *Inbox) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["GuildName"] = p.GuildName
	properties["FromUserIds"] = p.FromUserIds
	properties["CreatedAt"] = p.CreatedAt
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
