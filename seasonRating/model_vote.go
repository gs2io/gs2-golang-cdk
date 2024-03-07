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

type Vote struct {
	SeasonName     string
	SessionName    string
	WrittenBallots []WrittenBallot
	CreatedAt      int64
	UpdatedAt      int64
	Revision       *int64
}

type VoteOptions struct {
	WrittenBallots []WrittenBallot
	Revision       *int64
}

func NewVote(
	seasonName string,
	sessionName string,
	createdAt int64,
	updatedAt int64,
	options VoteOptions,
) Vote {
	data := Vote{
		SeasonName:     seasonName,
		SessionName:    sessionName,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		WrittenBallots: options.WrittenBallots,
		Revision:       options.Revision,
	}
	return data
}

func (p *Vote) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["SeasonName"] = p.SeasonName
	properties["SessionName"] = p.SessionName
	{
		values := make([]map[string]interface{}, len(p.WrittenBallots))
		for i, element := range p.WrittenBallots {
			values[i] = element.Properties()
		}
		properties["WrittenBallots"] = values
	}
	properties["CreatedAt"] = p.CreatedAt
	properties["UpdatedAt"] = p.UpdatedAt
	if p.Revision != nil {
		properties["Revision"] = p.Revision
	}
	return properties
}
