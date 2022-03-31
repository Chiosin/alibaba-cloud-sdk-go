package dms_enterprise

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

// DAGInstance is a nested struct in dms_enterprise response
type DAGInstance struct {
	DagId                string `json:"DagId" xml:"DagId"`
	Status               int    `json:"Status" xml:"Status"`
	Message              string `json:"Message" xml:"Message"`
	LatestInstanceStatus int    `json:"LatestInstanceStatus" xml:"LatestInstanceStatus"`
	OwnerName            string `json:"OwnerName" xml:"OwnerName"`
	BusinessTime         string `json:"BusinessTime" xml:"BusinessTime"`
	Id                   int64  `json:"Id" xml:"Id"`
	CreatorNickName      string `json:"CreatorNickName" xml:"CreatorNickName"`
	DeployId             int64  `json:"DeployId" xml:"DeployId"`
	EndTime              string `json:"EndTime" xml:"EndTime"`
	CreatorId            string `json:"CreatorId" xml:"CreatorId"`
	LatestInstanceTime   string `json:"LatestInstanceTime" xml:"LatestInstanceTime"`
	DagOwnerNickName     string `json:"DagOwnerNickName" xml:"DagOwnerNickName"`
	DagName              string `json:"DagName" xml:"DagName"`
	TriggerType          int    `json:"TriggerType" xml:"TriggerType"`
	HistoryDagId         int64  `json:"HistoryDagId" xml:"HistoryDagId"`
}
