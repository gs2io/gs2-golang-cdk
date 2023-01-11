package matchmaking

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

type SignedBallot struct {
	Body      string
	Signature string
}

type SignedBallotOptions struct {
}

func NewSignedBallot(
	body string,
	signature string,
	options SignedBallotOptions,
) SignedBallot {
	data := SignedBallot{
		Body:      body,
		Signature: signature,
	}
	return data
}

func (p *SignedBallot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Body"] = p.Body
	properties["Signature"] = p.Signature
	return properties
}
