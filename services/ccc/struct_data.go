package ccc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in ccc response
type Data struct {
	PageSize               int          `json:"PageSize" xml:"PageSize"`
	Password               string       `json:"Password" xml:"Password"`
	UserState              string       `json:"UserState" xml:"UserState"`
	InstanceId             string       `json:"InstanceId" xml:"InstanceId"`
	UserId                 string       `json:"UserId" xml:"UserId"`
	TotalCount             int          `json:"TotalCount" xml:"TotalCount"`
	DeviceId               string       `json:"DeviceId" xml:"DeviceId"`
	Reserved               int64        `json:"Reserved" xml:"Reserved"`
	PageNumber             int          `json:"PageNumber" xml:"PageNumber"`
	BreakCode              string       `json:"BreakCode" xml:"BreakCode"`
	Extension              string       `json:"Extension" xml:"Extension"`
	JobId                  string       `json:"JobId" xml:"JobId"`
	Heartbeat              int64        `json:"Heartbeat" xml:"Heartbeat"`
	City                   string       `json:"City" xml:"City"`
	DisplayName            string       `json:"DisplayName" xml:"DisplayName"`
	SipServerUrl           string       `json:"SipServerUrl" xml:"SipServerUrl"`
	Number                 string       `json:"Number" xml:"Number"`
	Province               string       `json:"Province" xml:"Province"`
	WorkMode               string       `json:"WorkMode" xml:"WorkMode"`
	Signature              string       `json:"Signature" xml:"Signature"`
	Mobile                 string       `json:"Mobile" xml:"Mobile"`
	OutboundScenario       bool         `json:"OutboundScenario" xml:"OutboundScenario"`
	UserName               string       `json:"UserName" xml:"UserName"`
	UserKey                string       `json:"UserKey" xml:"UserKey"`
	SignedSkillGroupIdList []string     `json:"SignedSkillGroupIdList" xml:"SignedSkillGroupIdList"`
	UserContext            UserContext  `json:"UserContext" xml:"UserContext"`
	CallContext            CallContext  `json:"CallContext" xml:"CallContext"`
	List                   []SkillGroup `json:"List" xml:"List"`
}
