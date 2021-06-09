package iot

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

// CreateStudioAppDomainOpen invokes the iot.CreateStudioAppDomainOpen API synchronously
func (client *Client) CreateStudioAppDomainOpen(request *CreateStudioAppDomainOpenRequest) (response *CreateStudioAppDomainOpenResponse, err error) {
	response = CreateCreateStudioAppDomainOpenResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStudioAppDomainOpenWithChan invokes the iot.CreateStudioAppDomainOpen API asynchronously
func (client *Client) CreateStudioAppDomainOpenWithChan(request *CreateStudioAppDomainOpenRequest) (<-chan *CreateStudioAppDomainOpenResponse, <-chan error) {
	responseChan := make(chan *CreateStudioAppDomainOpenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStudioAppDomainOpen(request)
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

// CreateStudioAppDomainOpenWithCallback invokes the iot.CreateStudioAppDomainOpen API asynchronously
func (client *Client) CreateStudioAppDomainOpenWithCallback(request *CreateStudioAppDomainOpenRequest, callback func(response *CreateStudioAppDomainOpenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStudioAppDomainOpenResponse
		var err error
		defer close(result)
		response, err = client.CreateStudioAppDomainOpen(request)
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

// CreateStudioAppDomainOpenRequest is the request struct for api CreateStudioAppDomainOpen
type CreateStudioAppDomainOpenRequest struct {
	*requests.RpcRequest
	Protocol      string `position:"Body" name:"Protocol"`
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	Host          string `position:"Body" name:"Host"`
	ProjectId     string `position:"Body" name:"ProjectId"`
	AppId         string `position:"Body" name:"AppId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// CreateStudioAppDomainOpenResponse is the response struct for api CreateStudioAppDomainOpen
type CreateStudioAppDomainOpenResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateStudioAppDomainOpenRequest creates a request to invoke CreateStudioAppDomainOpen API
func CreateCreateStudioAppDomainOpenRequest() (request *CreateStudioAppDomainOpenRequest) {
	request = &CreateStudioAppDomainOpenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateStudioAppDomainOpen", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStudioAppDomainOpenResponse creates a response to parse from CreateStudioAppDomainOpen response
func CreateCreateStudioAppDomainOpenResponse() (response *CreateStudioAppDomainOpenResponse) {
	response = &CreateStudioAppDomainOpenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}