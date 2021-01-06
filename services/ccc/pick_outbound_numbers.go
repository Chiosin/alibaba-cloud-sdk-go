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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// PickOutboundNumbers invokes the ccc.PickOutboundNumbers API synchronously
func (client *Client) PickOutboundNumbers(request *PickOutboundNumbersRequest) (response *PickOutboundNumbersResponse, err error) {
	response = CreatePickOutboundNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// PickOutboundNumbersWithChan invokes the ccc.PickOutboundNumbers API asynchronously
func (client *Client) PickOutboundNumbersWithChan(request *PickOutboundNumbersRequest) (<-chan *PickOutboundNumbersResponse, <-chan error) {
	responseChan := make(chan *PickOutboundNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PickOutboundNumbers(request)
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

// PickOutboundNumbersWithCallback invokes the ccc.PickOutboundNumbers API asynchronously
func (client *Client) PickOutboundNumbersWithCallback(request *PickOutboundNumbersRequest, callback func(response *PickOutboundNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PickOutboundNumbersResponse
		var err error
		defer close(result)
		response, err = client.PickOutboundNumbers(request)
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

// PickOutboundNumbersRequest is the request struct for api PickOutboundNumbers
type PickOutboundNumbersRequest struct {
	*requests.RpcRequest
	Count            requests.Integer `position:"Query" name:"Count"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
	SkillGroupIdList string           `position:"Query" name:"SkillGroupIdList"`
	CalledNumber     string           `position:"Query" name:"CalledNumber"`
}

// PickOutboundNumbersResponse is the response struct for api PickOutboundNumbers
type PickOutboundNumbersResponse struct {
	*responses.BaseResponse
	Code           string       `json:"Code" xml:"Code"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string       `json:"Message" xml:"Message"`
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Data           []NumberPair `json:"Data" xml:"Data"`
}

// CreatePickOutboundNumbersRequest creates a request to invoke PickOutboundNumbers API
func CreatePickOutboundNumbersRequest() (request *PickOutboundNumbersRequest) {
	request = &PickOutboundNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "PickOutboundNumbers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePickOutboundNumbersResponse creates a response to parse from PickOutboundNumbers response
func CreatePickOutboundNumbersResponse() (response *PickOutboundNumbersResponse) {
	response = &PickOutboundNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
