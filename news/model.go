package news

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

type News struct {
	section         string
	content         string
	title           string
	scheduleEventId *string
	timestamp       int64
	frontMatter     string
}

func NewNews(
	section string,
	content string,
	title string,
	scheduleEventId *string,
	timestamp int64,
	frontMatter string,
) News {
	data := News{
		section:         section,
		content:         content,
		title:           title,
		scheduleEventId: scheduleEventId,
		timestamp:       timestamp,
		frontMatter:     frontMatter,
	}
	return data
}

func (p *News) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Section"] = p.section
	properties["Content"] = p.content
	properties["Title"] = p.title
	if p.scheduleEventId != nil {
		properties["ScheduleEventId"] = p.scheduleEventId
	}
	properties["Timestamp"] = p.timestamp
	properties["FrontMatter"] = p.frontMatter
	return properties
}

type SetCookieRequestEntry struct {
	key   string
	value string
}

func NewSetCookieRequestEntry(
	key string,
	value string,
) SetCookieRequestEntry {
	data := SetCookieRequestEntry{
		key:   key,
		value: value,
	}
	return data
}

func (p *SetCookieRequestEntry) Properties() map[string]interface{} {
	properties := map[string]interface{}{}
	properties["Key"] = p.key
	properties["Value"] = p.value
	return properties
}
