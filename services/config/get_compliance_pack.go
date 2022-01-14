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

// GetCompliancePack invokes the config.GetCompliancePack API synchronously
func (client *Client) GetCompliancePack(request *GetCompliancePackRequest) (response *GetCompliancePackResponse, err error) {
	response = CreateGetCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// GetCompliancePackWithChan invokes the config.GetCompliancePack API asynchronously
func (client *Client) GetCompliancePackWithChan(request *GetCompliancePackRequest) (<-chan *GetCompliancePackResponse, <-chan error) {
	responseChan := make(chan *GetCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCompliancePack(request)
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

// GetCompliancePackWithCallback invokes the config.GetCompliancePack API asynchronously
func (client *Client) GetCompliancePackWithCallback(request *GetCompliancePackRequest, callback func(response *GetCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.GetCompliancePack(request)
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

// GetCompliancePackRequest is the request struct for api GetCompliancePack
type GetCompliancePackRequest struct {
	*requests.RpcRequest
	CompliancePackId string `position:"Query" name:"CompliancePackId"`
}

// GetCompliancePackResponse is the response struct for api GetCompliancePack
type GetCompliancePackResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	CompliancePack CompliancePack `json:"CompliancePack" xml:"CompliancePack"`
}

// CreateGetCompliancePackRequest creates a request to invoke GetCompliancePack API
func CreateGetCompliancePackRequest() (request *GetCompliancePackRequest) {
	request = &GetCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetCompliancePack", "", "")
	request.Method = requests.GET
	return
}

// CreateGetCompliancePackResponse creates a response to parse from GetCompliancePack response
func CreateGetCompliancePackResponse() (response *GetCompliancePackResponse) {
	response = &GetCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}