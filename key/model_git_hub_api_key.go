package key

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

type GitHubApiKey struct {
	Name              string
	Description       *string
	ApiKey            string
	EncryptionKeyName string
	CreatedAt         int64
	UpdatedAt         int64
}

type GitHubApiKeyOptions struct {
	Description *string
}

func NewGitHubApiKey(
	name string,
	apiKey string,
	encryptionKeyName string,
	createdAt int64,
	updatedAt int64,
	options GitHubApiKeyOptions,
) GitHubApiKey {
	data := GitHubApiKey{
		Name:              name,
		ApiKey:            apiKey,
		EncryptionKeyName: encryptionKeyName,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		Description:       options.Description,
	}
	return data
}

func (p *GitHubApiKey) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.Name
	if p.Description != nil {
		properties["Description"] = p.Description
	}
	properties["ApiKey"] = p.ApiKey
	properties["EncryptionKeyName"] = p.EncryptionKeyName
	properties["CreatedAt"] = p.CreatedAt
	properties["UpdatedAt"] = p.UpdatedAt
	return properties
}
