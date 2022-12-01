package opensearch

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

// ResultItem is a nested struct in opensearch response
type ResultItem struct {
	Date               string                   `json:"date" xml:"date"`
	Traffic            int                      `json:"traffic" xml:"traffic"`
	Service            string                   `json:"service" xml:"service"`
	Created            int                      `json:"created" xml:"created"`
	ZeroHitRate        float64                  `json:"zeroHitRate" xml:"zeroHitRate"`
	IndustryName       string                   `json:"industryName" xml:"industryName"`
	Meta               string                   `json:"meta" xml:"meta"`
	Business           string                   `json:"business" xml:"business"`
	IpvUv              int                      `json:"ipvUv" xml:"ipvUv"`
	Pv                 int                      `json:"pv" xml:"pv"`
	Label              string                   `json:"label" xml:"label"`
	Domain             string                   `json:"domain" xml:"domain"`
	ExperimentName     string                   `json:"experimentName" xml:"experimentName"`
	Name               string                   `json:"name" xml:"name"`
	Assumed            bool                     `json:"assumed" xml:"assumed"`
	IsSys              string                   `json:"isSys" xml:"isSys"`
	Type               string                   `json:"type" xml:"type"`
	Endpoint           string                   `json:"endpoint" xml:"endpoint"`
	ConsoleUrl         string                   `json:"consoleUrl" xml:"consoleUrl"`
	Updated            int                      `json:"updated" xml:"updated"`
	Code               string                   `json:"code" xml:"code"`
	RegionId           string                   `json:"regionId" xml:"regionId"`
	Tag                string                   `json:"tag" xml:"tag"`
	TemplateId         string                   `json:"template_id" xml:"template_id"`
	LocalName          string                   `json:"localName" xml:"localName"`
	Status             int                      `json:"status" xml:"status"`
	Message            string                   `json:"message" xml:"message"`
	Ctr                float64                  `json:"ctr" xml:"ctr"`
	Online             bool                     `json:"online" xml:"online"`
	Description        string                   `json:"description" xml:"description"`
	SundialId          string                   `json:"sundialId" xml:"sundialId"`
	Available          bool                     `json:"available" xml:"available"`
	Params             map[string]interface{}   `json:"params" xml:"params"`
	DataCollectionType string                   `json:"dataCollectionType" xml:"dataCollectionType"`
	Order              int                      `json:"order" xml:"order"`
	Priority           string                   `json:"priority" xml:"priority"`
	IsDefault          string                   `json:"isDefault" xml:"isDefault"`
	Ipv                int                      `json:"ipv" xml:"ipv"`
	Uv                 int                      `json:"uv" xml:"uv"`
	Id                 string                   `json:"id" xml:"id"`
	Active             bool                     `json:"active" xml:"active"`
	Processors         []map[string]interface{} `json:"processors" xml:"processors"`
	Indexes            []string                 `json:"indexes" xml:"indexes"`
	Values             []string                 `json:"values" xml:"values"`
	DataSource         DataSource               `json:"dataSource" xml:"dataSource"`
	Dicts              []DictsItem              `json:"dicts" xml:"dicts"`
}
