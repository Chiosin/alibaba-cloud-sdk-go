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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetValidationReport invokes the opensearch.GetValidationReport API synchronously
func (client *Client) GetValidationReport(request *GetValidationReportRequest) (response *GetValidationReportResponse, err error) {
	response = CreateGetValidationReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetValidationReportWithChan invokes the opensearch.GetValidationReport API asynchronously
func (client *Client) GetValidationReportWithChan(request *GetValidationReportRequest) (<-chan *GetValidationReportResponse, <-chan error) {
	responseChan := make(chan *GetValidationReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetValidationReport(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetValidationReportWithCallback invokes the opensearch.GetValidationReport API asynchronously
func (client *Client) GetValidationReportWithCallback(request *GetValidationReportRequest, callback func(response *GetValidationReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetValidationReportResponse
		var err error
		defer close(result)
		response, err = client.GetValidationReport(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetValidationReportRequest is the request struct for api GetValidationReport
type GetValidationReportRequest struct {
	*requests.RoaRequest
	Type             string `position:"Query" name:"type"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetValidationReportResponse is the response struct for api GetValidationReport
type GetValidationReportResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"requestId" xml:"requestId"`
	Result    []map[string]interface{} `json:"result" xml:"result"`
}

// CreateGetValidationReportRequest creates a request to invoke GetValidationReport API
func CreateGetValidationReportRequest() (request *GetValidationReportRequest) {
	request = &GetValidationReportRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetValidationReport", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/data/validation-report", "", "")
	request.Method = requests.GET
	return
}

// CreateGetValidationReportResponse creates a response to parse from GetValidationReport response
func CreateGetValidationReportResponse() (response *GetValidationReportResponse) {
	response = &GetValidationReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
