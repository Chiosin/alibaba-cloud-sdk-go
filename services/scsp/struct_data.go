package scsp

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

// Data is a nested struct in scsp response
type Data struct {
	CreatorName             string                        `json:"CreatorName" xml:"CreatorName"`
	TemplateId              int64                         `json:"TemplateId" xml:"TemplateId"`
	PageSize                int64                         `json:"PageSize" xml:"PageSize"`
	AgentStatus             int                           `json:"AgentStatus" xml:"AgentStatus"`
	Expires                 int64                         `json:"Expires" xml:"Expires"`
	ExtInfo                 string                        `json:"ExtInfo" xml:"ExtInfo"`
	Anonymity               bool                          `json:"Anonymity" xml:"Anonymity"`
	AgentId                 int64                         `json:"AgentId" xml:"AgentId"`
	EntityBizCodeType       string                        `json:"EntityBizCodeType" xml:"EntityBizCodeType"`
	RestType                int                           `json:"RestType" xml:"RestType"`
	Nick                    string                        `json:"Nick" xml:"Nick"`
	MemberId                int64                         `json:"MemberId" xml:"MemberId"`
	FormData                string                        `json:"FormData" xml:"FormData"`
	Token                   string                        `json:"Token" xml:"Token"`
	CategoryId              int64                         `json:"CategoryId" xml:"CategoryId"`
	DepartmentId            string                        `json:"DepartmentId" xml:"DepartmentId"`
	Avatar                  string                        `json:"Avatar" xml:"Avatar"`
	EntityId                string                        `json:"EntityId" xml:"EntityId"`
	FromInfo                string                        `json:"FromInfo" xml:"FromInfo"`
	Gender                  string                        `json:"Gender" xml:"Gender"`
	ServiceId               int64                         `json:"ServiceId" xml:"ServiceId"`
	TicketId                int64                         `json:"TicketId" xml:"TicketId"`
	CreatorId               int64                         `json:"CreatorId" xml:"CreatorId"`
	Assigned                bool                          `json:"Assigned" xml:"Assigned"`
	EntityBizCode           string                        `json:"EntityBizCode" xml:"EntityBizCode"`
	TenantId                int64                         `json:"TenantId" xml:"TenantId"`
	CreateTime              int64                         `json:"CreateTime" xml:"CreateTime"`
	EntityName              string                        `json:"EntityName" xml:"EntityName"`
	MemberName              string                        `json:"MemberName" xml:"MemberName"`
	UserId                  string                        `json:"UserId" xml:"UserId"`
	EntityRelationNumber    string                        `json:"EntityRelationNumber" xml:"EntityRelationNumber"`
	TotalNum                int64                         `json:"TotalNum" xml:"TotalNum"`
	CustomerTypeId          int                           `json:"CustomerTypeId" xml:"CustomerTypeId"`
	Phone                   string                        `json:"Phone" xml:"Phone"`
	Status                  int                           `json:"Status" xml:"Status"`
	ForeignId               string                        `json:"ForeignId" xml:"ForeignId"`
	Priority                int                           `json:"Priority" xml:"Priority"`
	PageNo                  string                        `json:"PageNo" xml:"PageNo"`
	CarbonCopy              string                        `json:"CarbonCopy" xml:"CarbonCopy"`
	CaseStatus              int                           `json:"CaseStatus" xml:"CaseStatus"`
	DisplayName             string                        `json:"DisplayName" xml:"DisplayName"`
	ModifiedTime            int64                         `json:"ModifiedTime" xml:"ModifiedTime"`
	UniqueId                int64                         `json:"UniqueId" xml:"UniqueId"`
	AgentStatusCode         string                        `json:"AgentStatusCode" xml:"AgentStatusCode"`
	AccountName             string                        `json:"AccountName" xml:"AccountName"`
	GroupId                 int64                         `json:"GroupId" xml:"GroupId"`
	CreatorType             int                           `json:"CreatorType" xml:"CreatorType"`
	TicketName              string                        `json:"TicketName" xml:"TicketName"`
	Num                     []NumItem                     `json:"Num" xml:"Num"`
	GroupList               []GroupListItem               `json:"GroupList" xml:"GroupList"`
	NumGroup                []NumGroupItem                `json:"NumGroup" xml:"NumGroup"`
	CallsPerdayResponseList []CallsPerdayResponseListItem `json:"CallsPerdayResponseList" xml:"CallsPerdayResponseList"`
	ActivityRecords         []ActivityRecordsItem         `json:"ActivityRecords" xml:"ActivityRecords"`
	Activities              []ActivitiesItem              `json:"Activities" xml:"Activities"`
}