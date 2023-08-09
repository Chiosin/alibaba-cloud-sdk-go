package config

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

// ListRemediationExecutions invokes the config.ListRemediationExecutions API synchronously
func (client *Client) ListRemediationExecutions(request *ListRemediationExecutionsRequest) (response *ListRemediationExecutionsResponse, err error) {
	response = CreateListRemediationExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRemediationExecutionsWithChan invokes the config.ListRemediationExecutions API asynchronously
func (client *Client) ListRemediationExecutionsWithChan(request *ListRemediationExecutionsRequest) (<-chan *ListRemediationExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListRemediationExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRemediationExecutions(request)
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

// ListRemediationExecutionsWithCallback invokes the config.ListRemediationExecutions API asynchronously
func (client *Client) ListRemediationExecutionsWithCallback(request *ListRemediationExecutionsRequest, callback func(response *ListRemediationExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRemediationExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListRemediationExecutions(request)
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

// ListRemediationExecutionsRequest is the request struct for api ListRemediationExecutions
type ListRemediationExecutionsRequest struct {
	*requests.RpcRequest
	ConfigRuleId    string           `position:"Query" name:"ConfigRuleId"`
	ExecutionStatus string           `position:"Query" name:"ExecutionStatus"`
	NextToken       string           `position:"Query" name:"NextToken"`
	MaxResults      requests.Integer `position:"Query" name:"MaxResults"`
}

// ListRemediationExecutionsResponse is the response struct for api ListRemediationExecutions
type ListRemediationExecutionsResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	RemediationExecutionData RemediationExecutionData `json:"RemediationExecutionData" xml:"RemediationExecutionData"`
}

// CreateListRemediationExecutionsRequest creates a request to invoke ListRemediationExecutions API
func CreateListRemediationExecutionsRequest() (request *ListRemediationExecutionsRequest) {
	request = &ListRemediationExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListRemediationExecutions", "", "")
	request.Method = requests.POST
	return
}

// CreateListRemediationExecutionsResponse creates a response to parse from ListRemediationExecutions response
func CreateListRemediationExecutionsResponse() (response *ListRemediationExecutionsResponse) {
	response = &ListRemediationExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
