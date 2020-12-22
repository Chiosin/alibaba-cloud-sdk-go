package dataworks_public

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

// Data is a nested struct in dataworks_public response
type Data struct {
	HourSlaDetail       string                `json:"HourSlaDetail" xml:"HourSlaDetail"`
	LastModifyTime      int64                 `json:"LastModifyTime" xml:"LastModifyTime"`
	CreatorName         string                `json:"CreatorName" xml:"CreatorName"`
	ApplicationSecret   string                `json:"ApplicationSecret" xml:"ApplicationSecret"`
	CycTime             int64                 `json:"CycTime" xml:"CycTime"`
	WhereCondition      string                `json:"WhereCondition" xml:"WhereCondition"`
	Property            string                `json:"Property" xml:"Property"`
	BaselineType        string                `json:"BaselineType" xml:"BaselineType"`
	TaskId              string                `json:"TaskId" xml:"TaskId"`
	IsDefault           bool                  `json:"IsDefault" xml:"IsDefault"`
	Content             string                `json:"Content" xml:"Content"`
	EnvType             int                   `json:"EnvType" xml:"EnvType"`
	ProjectName         string                `json:"ProjectName" xml:"ProjectName"`
	Creator             string                `json:"Creator" xml:"Creator"`
	FileName            string                `json:"FileName" xml:"FileName"`
	Type                string                `json:"Type" xml:"Type"`
	TotalCount          int64                 `json:"TotalCount" xml:"TotalCount"`
	FinishTime          int64                 `json:"FinishTime" xml:"FinishTime"`
	Connection          string                `json:"Connection" xml:"Connection"`
	Id                  int64                 `json:"Id" xml:"Id"`
	RuleType            int                   `json:"RuleType" xml:"RuleType"`
	TableName           string                `json:"TableName" xml:"TableName"`
	CategoryId          int64                 `json:"CategoryId" xml:"CategoryId"`
	PredictType         int                   `json:"PredictType" xml:"PredictType"`
	Trend               string                `json:"Trend" xml:"Trend"`
	InstanceId          int64                 `json:"InstanceId" xml:"InstanceId"`
	ExpHour             int                   `json:"ExpHour" xml:"ExpHour"`
	Status              string                `json:"Status" xml:"Status"`
	RemindUnit          string                `json:"RemindUnit" xml:"RemindUnit"`
	Name                string                `json:"Name" xml:"Name"`
	InGroupId           int                   `json:"InGroupId" xml:"InGroupId"`
	ModifyTime          int64                 `json:"ModifyTime" xml:"ModifyTime"`
	ApplicationId       int64                 `json:"ApplicationId" xml:"ApplicationId"`
	FilePropertyContent string                `json:"FilePropertyContent" xml:"FilePropertyContent"`
	ExpTime             int64                 `json:"ExpTime" xml:"ExpTime"`
	TemplateId          int                   `json:"TemplateId" xml:"TemplateId"`
	Repeatability       bool                  `json:"Repeatability" xml:"Repeatability"`
	TenantId            int64                 `json:"TenantId" xml:"TenantId"`
	Detail              string                `json:"Detail" xml:"Detail"`
	PageSize            int                   `json:"PageSize" xml:"PageSize"`
	DataSize            int64                 `json:"DataSize" xml:"DataSize"`
	Comment             string                `json:"Comment" xml:"Comment"`
	RepeatInterval      int64                 `json:"RepeatInterval" xml:"RepeatInterval"`
	Founder             string                `json:"Founder" xml:"Founder"`
	CreateUser          string                `json:"CreateUser" xml:"CreateUser"`
	EndCast             int64                 `json:"EndCast" xml:"EndCast"`
	LastAccessTime      int64                 `json:"LastAccessTime" xml:"LastAccessTime"`
	BaselineName        string                `json:"BaselineName" xml:"BaselineName"`
	Bizdate             int64                 `json:"Bizdate" xml:"Bizdate"`
	ApplicationCode     string                `json:"ApplicationCode" xml:"ApplicationCode"`
	Priority            int                   `json:"Priority" xml:"Priority"`
	TotalColumnCount    int64                 `json:"TotalColumnCount" xml:"TotalColumnCount"`
	Buffer              float64               `json:"Buffer" xml:"Buffer"`
	NextPrimaryKey      string                `json:"NextPrimaryKey" xml:"NextPrimaryKey"`
	CommitUser          string                `json:"CommitUser" xml:"CommitUser"`
	Owner               string                `json:"Owner" xml:"Owner"`
	PageNum             int                   `json:"PageNum" xml:"PageNum"`
	PartitionKeys       string                `json:"PartitionKeys" xml:"PartitionKeys"`
	EntityId            int64                 `json:"EntityId" xml:"EntityId"`
	BeginWaitTimeTime   int64                 `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	OwnerId             string                `json:"OwnerId" xml:"OwnerId"`
	UseFlag             bool                  `json:"UseFlag" xml:"UseFlag"`
	DqcDescription      string                `json:"DqcDescription" xml:"DqcDescription"`
	CheckerName         string                `json:"CheckerName" xml:"CheckerName"`
	RemindType          string                `json:"RemindType" xml:"RemindType"`
	UseType             string                `json:"UseType" xml:"UseType"`
	MethodId            int                   `json:"MethodId" xml:"MethodId"`
	CriticalThreshold   string                `json:"CriticalThreshold" xml:"CriticalThreshold"`
	ApplicationKey      string                `json:"ApplicationKey" xml:"ApplicationKey"`
	DagType             string                `json:"DagType" xml:"DagType"`
	DndStart            string                `json:"DndStart" xml:"DndStart"`
	BeginRunningTime    int64                 `json:"BeginRunningTime" xml:"BeginRunningTime"`
	FixCheck            bool                  `json:"FixCheck" xml:"FixCheck"`
	FileVersion         int                   `json:"FileVersion" xml:"FileVersion"`
	Operator            string                `json:"Operator" xml:"Operator"`
	HasNext             bool                  `json:"HasNext" xml:"HasNext"`
	SlaMinu             int                   `json:"SlaMinu" xml:"SlaMinu"`
	StartTime           int64                 `json:"StartTime" xml:"StartTime"`
	NodeName            string                `json:"NodeName" xml:"NodeName"`
	ProjectNameCn       string                `json:"ProjectNameCn" xml:"ProjectNameCn"`
	Gmtdate             int64                 `json:"Gmtdate" xml:"Gmtdate"`
	FileContent         string                `json:"FileContent" xml:"FileContent"`
	LifeCycle           int                   `json:"LifeCycle" xml:"LifeCycle"`
	HourExpDetail       string                `json:"HourExpDetail" xml:"HourExpDetail"`
	DatabaseName        string                `json:"DatabaseName" xml:"DatabaseName"`
	Checker             int                   `json:"Checker" xml:"Checker"`
	ClusterId           string                `json:"ClusterId" xml:"ClusterId"`
	Version             int64                 `json:"Version" xml:"Version"`
	TopicId             int64                 `json:"TopicId" xml:"TopicId"`
	ApplicationName     string                `json:"ApplicationName" xml:"ApplicationName"`
	OwnerName           string                `json:"OwnerName" xml:"OwnerName"`
	BaselineId          int64                 `json:"BaselineId" xml:"BaselineId"`
	IsCurrentProd       bool                  `json:"IsCurrentProd" xml:"IsCurrentProd"`
	Location            string                `json:"Location" xml:"Location"`
	FinishStatus        string                `json:"FinishStatus" xml:"FinishStatus"`
	Caption             string                `json:"Caption" xml:"Caption"`
	AppGuid             string                `json:"AppGuid" xml:"AppGuid"`
	WarningThreshold    string                `json:"WarningThreshold" xml:"WarningThreshold"`
	NodeId              int64                 `json:"NodeId" xml:"NodeId"`
	RemindName          string                `json:"RemindName" xml:"RemindName"`
	ClusterBizId        string                `json:"ClusterBizId" xml:"ClusterBizId"`
	FolderPath          string                `json:"FolderPath" xml:"FolderPath"`
	NodeContent         string                `json:"NodeContent" xml:"NodeContent"`
	ChangeType          string                `json:"ChangeType" xml:"ChangeType"`
	ParamValues         string                `json:"ParamValues" xml:"ParamValues"`
	MethodName          string                `json:"MethodName" xml:"MethodName"`
	TemplateName        string                `json:"TemplateName" xml:"TemplateName"`
	MaxAlertTimes       int                   `json:"MaxAlertTimes" xml:"MaxAlertTimes"`
	RuleName            string                `json:"RuleName" xml:"RuleName"`
	ErrorMessage        string                `json:"ErrorMessage" xml:"ErrorMessage"`
	BeginWaitResTime    int64                 `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	CreateTime          int64                 `json:"CreateTime" xml:"CreateTime"`
	IsVisible           int                   `json:"IsVisible" xml:"IsVisible"`
	FolderId            string                `json:"FolderId" xml:"FolderId"`
	BlockType           int                   `json:"BlockType" xml:"BlockType"`
	Endpoint            string                `json:"Endpoint" xml:"Endpoint"`
	NextTaskId          string                `json:"NextTaskId" xml:"NextTaskId"`
	SlaTime             int64                 `json:"SlaTime" xml:"SlaTime"`
	PageNumber          int                   `json:"PageNumber" xml:"PageNumber"`
	DqcType             int                   `json:"DqcType" xml:"DqcType"`
	RemindId            int64                 `json:"RemindId" xml:"RemindId"`
	TableGuid           string                `json:"TableGuid" xml:"TableGuid"`
	ExpectValue         string                `json:"ExpectValue" xml:"ExpectValue"`
	ExpMinu             int                   `json:"ExpMinu" xml:"ExpMinu"`
	OnDuty              string                `json:"OnDuty" xml:"OnDuty"`
	DndEnd              string                `json:"DndEnd" xml:"DndEnd"`
	DagId               int64                 `json:"DagId" xml:"DagId"`
	AlertUnit           string                `json:"AlertUnit" xml:"AlertUnit"`
	RelatedFlowId       int64                 `json:"RelatedFlowId" xml:"RelatedFlowId"`
	LastDdlTime         int64                 `json:"LastDdlTime" xml:"LastDdlTime"`
	ModifiedTime        int64                 `json:"ModifiedTime" xml:"ModifiedTime"`
	Useflag             bool                  `json:"Useflag" xml:"Useflag"`
	AlertInterval       int                   `json:"AlertInterval" xml:"AlertInterval"`
	CommitTime          int64                 `json:"CommitTime" xml:"CommitTime"`
	ProjectId           int64                 `json:"ProjectId" xml:"ProjectId"`
	SlaHour             int                   `json:"SlaHour" xml:"SlaHour"`
	AlertTargets        []string              `json:"AlertTargets" xml:"AlertTargets"`
	TableGuidList       []string              `json:"TableGuidList" xml:"TableGuidList"`
	AlertMethods        []string              `json:"AlertMethods" xml:"AlertMethods"`
	Deployment          Deployment            `json:"Deployment" xml:"Deployment"`
	File                File                  `json:"File" xml:"File"`
	NodeConfiguration   NodeConfiguration     `json:"NodeConfiguration" xml:"NodeConfiguration"`
	BlockInstance       BlockInstance         `json:"BlockInstance" xml:"BlockInstance"`
	LastInstance        LastInstance          `json:"LastInstance" xml:"LastInstance"`
	DataEntityList      []DataEntityListItem  `json:"DataEntityList" xml:"DataEntityList"`
	Influences          []InfluencesItem      `json:"Influences" xml:"Influences"`
	Baselines           []BaselinesItem       `json:"Baselines" xml:"Baselines"`
	ThemeList           []ThemeListItem       `json:"ThemeList" xml:"ThemeList"`
	RuleChecks          []RuleChecksItem      `json:"RuleChecks" xml:"RuleChecks"`
	TableEntityList     []TableEntityListItem `json:"TableEntityList" xml:"TableEntityList"`
	Rules               []RulesItem           `json:"Rules" xml:"Rules"`
	ColumnList          []ColumnListItem      `json:"ColumnList" xml:"ColumnList"`
	Robots              []RobotsItem          `json:"Robots" xml:"Robots"`
	BizProcesses        []BizProcessesItem    `json:"BizProcesses" xml:"BizProcesses"`
	Projects            []ProjectsItem        `json:"Projects" xml:"Projects"`
	Nodes               []NodesItem           `json:"Nodes" xml:"Nodes"`
}
