package distributor

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

type BatchRequestPayloadService string

const BatchRequestPayloadServiceAccount = BatchRequestPayloadService("account")
const BatchRequestPayloadServiceAdReward = BatchRequestPayloadService("adReward")
const BatchRequestPayloadServiceAuth = BatchRequestPayloadService("auth")
const BatchRequestPayloadServiceBuff = BatchRequestPayloadService("buff")
const BatchRequestPayloadServiceChat = BatchRequestPayloadService("chat")
const BatchRequestPayloadServiceDatastore = BatchRequestPayloadService("datastore")
const BatchRequestPayloadServiceDeploy = BatchRequestPayloadService("deploy")
const BatchRequestPayloadServiceDictionary = BatchRequestPayloadService("dictionary")
const BatchRequestPayloadServiceDistributor = BatchRequestPayloadService("distributor")
const BatchRequestPayloadServiceEnchant = BatchRequestPayloadService("enchant")
const BatchRequestPayloadServiceEnhance = BatchRequestPayloadService("enhance")
const BatchRequestPayloadServiceExchange = BatchRequestPayloadService("exchange")
const BatchRequestPayloadServiceExperience = BatchRequestPayloadService("experience")
const BatchRequestPayloadServiceFormation = BatchRequestPayloadService("formation")
const BatchRequestPayloadServiceFriend = BatchRequestPayloadService("friend")
const BatchRequestPayloadServiceGateway = BatchRequestPayloadService("gateway")
const BatchRequestPayloadServiceGrade = BatchRequestPayloadService("grade")
const BatchRequestPayloadServiceGuard = BatchRequestPayloadService("guard")
const BatchRequestPayloadServiceGuild = BatchRequestPayloadService("guild")
const BatchRequestPayloadServiceIdentifier = BatchRequestPayloadService("identifier")
const BatchRequestPayloadServiceIdle = BatchRequestPayloadService("idle")
const BatchRequestPayloadServiceInbox = BatchRequestPayloadService("inbox")
const BatchRequestPayloadServiceInventory = BatchRequestPayloadService("inventory")
const BatchRequestPayloadServiceJobQueue = BatchRequestPayloadService("jobQueue")
const BatchRequestPayloadServiceKey = BatchRequestPayloadService("key")
const BatchRequestPayloadServiceLimit = BatchRequestPayloadService("limit")
const BatchRequestPayloadServiceLock = BatchRequestPayloadService("lock")
const BatchRequestPayloadServiceLog = BatchRequestPayloadService("log")
const BatchRequestPayloadServiceLoginReward = BatchRequestPayloadService("loginReward")
const BatchRequestPayloadServiceLottery = BatchRequestPayloadService("lottery")
const BatchRequestPayloadServiceMatchmaking = BatchRequestPayloadService("matchmaking")
const BatchRequestPayloadServiceMegaField = BatchRequestPayloadService("megaField")
const BatchRequestPayloadServiceMission = BatchRequestPayloadService("mission")
const BatchRequestPayloadServiceMoney = BatchRequestPayloadService("money")
const BatchRequestPayloadServiceMoney2 = BatchRequestPayloadService("money2")
const BatchRequestPayloadServiceNews = BatchRequestPayloadService("news")
const BatchRequestPayloadServiceQuest = BatchRequestPayloadService("quest")
const BatchRequestPayloadServiceRanking = BatchRequestPayloadService("ranking")
const BatchRequestPayloadServiceRanking2 = BatchRequestPayloadService("ranking2")
const BatchRequestPayloadServiceRealtime = BatchRequestPayloadService("realtime")
const BatchRequestPayloadServiceSchedule = BatchRequestPayloadService("schedule")
const BatchRequestPayloadServiceScript = BatchRequestPayloadService("script")
const BatchRequestPayloadServiceSeasonRating = BatchRequestPayloadService("seasonRating")
const BatchRequestPayloadServiceSerialKey = BatchRequestPayloadService("serialKey")
const BatchRequestPayloadServiceShowcase = BatchRequestPayloadService("showcase")
const BatchRequestPayloadServiceSkillTree = BatchRequestPayloadService("skillTree")
const BatchRequestPayloadServiceStamina = BatchRequestPayloadService("stamina")
const BatchRequestPayloadServiceStateMachine = BatchRequestPayloadService("stateMachine")
const BatchRequestPayloadServiceVersion = BatchRequestPayloadService("version")

func (p BatchRequestPayloadService) Pointer() *BatchRequestPayloadService {
	return &p
}

type BatchRequestPayload struct {
	RequestId  string
	Service    BatchRequestPayloadService
	MethodName string
	Parameter  string
}

type BatchRequestPayloadOptions struct {
}

func NewBatchRequestPayload(
	requestId string,
	service BatchRequestPayloadService,
	methodName string,
	parameter string,
	options BatchRequestPayloadOptions,
) BatchRequestPayload {
	_data := BatchRequestPayload{
		RequestId:  requestId,
		Service:    service,
		MethodName: methodName,
		Parameter:  parameter,
	}
	return _data
}

func (p *BatchRequestPayload) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["RequestId"] = p.RequestId
	properties["Service"] = p.Service
	properties["MethodName"] = p.MethodName
	properties["Parameter"] = p.Parameter
	return properties
}
