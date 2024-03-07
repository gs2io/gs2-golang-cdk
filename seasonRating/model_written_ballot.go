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

type WrittenBallot struct {
	Ballot      Ballot
	GameResults []GameResult
}

type WrittenBallotOptions struct {
	GameResults []GameResult
}

func NewWrittenBallot(
	ballot Ballot,
	options WrittenBallotOptions,
) WrittenBallot {
	data := WrittenBallot{
		Ballot:      ballot,
		GameResults: options.GameResults,
	}
	return data
}

func (p *WrittenBallot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Ballot"] = p.Ballot.Properties()
	{
		values := make([]map[string]interface{}, len(p.GameResults))
		for i, element := range p.GameResults {
			values[i] = element.Properties()
		}
		properties["GameResults"] = values
	}
	return properties
}
