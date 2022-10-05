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

type AttributeRange struct {
	name string
	min  int32
	max  int32
}

func NewAttributeRange(
	name string,
	min int32,
	max int32,
) AttributeRange {
	data := AttributeRange{
		name: name,
		min:  min,
		max:  max,
	}
	return data
}

func (p *AttributeRange) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["Min"] = p.min
	properties["Max"] = p.max
	return properties
}

type CapacityOfRole struct {
	roleName     string
	roleAliases  []string
	capacity     int32
	participants []Player
}

func NewCapacityOfRole(
	roleName string,
	roleAliases []string,
	capacity int32,
	participants []Player,
) CapacityOfRole {
	data := CapacityOfRole{
		roleName:     roleName,
		roleAliases:  roleAliases,
		capacity:     capacity,
		participants: participants,
	}
	return data
}

func (p *CapacityOfRole) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["RoleName"] = p.roleName
	properties["RoleAliases"] = p.roleAliases
	properties["Capacity"] = p.capacity
	{
		values := make([]map[string]interface{}, len(p.participants))
		for i, element := range p.participants {
			values[i] = element.Properties()
		}
		properties["Participants"] = values
	}
	return properties
}

type Attribute struct {
	name  string
	value int32
}

func NewAttribute(
	name string,
	value int32,
) Attribute {
	data := Attribute{
		name:  name,
		value: value,
	}
	return data
}

func (p *Attribute) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Name"] = p.name
	properties["Value"] = p.value
	return properties
}

type Player struct {
	userId      string
	attributes  []Attribute
	roleName    string
	denyUserIds []string
}

func NewPlayer(
	userId string,
	attributes []Attribute,
	roleName string,
	denyUserIds []string,
) Player {
	data := Player{
		userId:      userId,
		attributes:  attributes,
		roleName:    roleName,
		denyUserIds: denyUserIds,
	}
	return data
}

func (p *Player) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.userId
	{
		values := make([]map[string]interface{}, len(p.attributes))
		for i, element := range p.attributes {
			values[i] = element.Properties()
		}
		properties["Attributes"] = values
	}
	properties["RoleName"] = p.roleName
	properties["DenyUserIds"] = p.denyUserIds
	return properties
}

type GameResult struct {
	rank   int32
	userId string
}

func NewGameResult(
	rank int32,
	userId string,
) GameResult {
	data := GameResult{
		rank:   rank,
		userId: userId,
	}
	return data
}

func (p *GameResult) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Rank"] = p.rank
	properties["UserId"] = p.userId
	return properties
}

type Ballot struct {
	userId         string
	ratingName     string
	gatheringName  string
	numberOfPlayer int32
}

func NewBallot(
	userId string,
	ratingName string,
	gatheringName string,
	numberOfPlayer int32,
) Ballot {
	data := Ballot{
		userId:         userId,
		ratingName:     ratingName,
		gatheringName:  gatheringName,
		numberOfPlayer: numberOfPlayer,
	}
	return data
}

func (p *Ballot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["UserId"] = p.userId
	properties["RatingName"] = p.ratingName
	properties["GatheringName"] = p.gatheringName
	properties["NumberOfPlayer"] = p.numberOfPlayer
	return properties
}

type SignedBallot struct {
	body      string
	signature string
}

func NewSignedBallot(
	body string,
	signature string,
) SignedBallot {
	data := SignedBallot{
		body:      body,
		signature: signature,
	}
	return data
}

func (p *SignedBallot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Body"] = p.body
	properties["Signature"] = p.signature
	return properties
}

type WrittenBallot struct {
	ballot      Ballot
	gameResults []GameResult
}

func NewWrittenBallot(
	ballot Ballot,
	gameResults []GameResult,
) WrittenBallot {
	data := WrittenBallot{
		ballot:      ballot,
		gameResults: gameResults,
	}
	return data
}

func (p *WrittenBallot) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Ballot"] = p.ballot.Properties()
	{
		values := make([]map[string]interface{}, len(p.gameResults))
		for i, element := range p.gameResults {
			values[i] = element.Properties()
		}
		properties["GameResults"] = values
	}
	return properties
}

type TimeSpan struct {
	days    int32
	hours   int32
	minutes int32
}

func NewTimeSpan(
	days int32,
	hours int32,
	minutes int32,
) TimeSpan {
	data := TimeSpan{
		days:    days,
		hours:   hours,
		minutes: minutes,
	}
	return data
}

func (p *TimeSpan) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Days"] = p.days
	properties["Hours"] = p.hours
	properties["Minutes"] = p.minutes
	return properties
}
