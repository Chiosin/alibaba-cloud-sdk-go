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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// StartCallV2 invokes the scsp.StartCallV2 API synchronously
func (client *Client) StartCallV2(request *StartCallV2Request) (response *StartCallV2Response, err error) {
	response = CreateStartCallV2Response()
	err = client.DoAction(request, response)
	return
}

// StartCallV2WithChan invokes the scsp.StartCallV2 API asynchronously
func (client *Client) StartCallV2WithChan(request *StartCallV2Request) (<-chan *StartCallV2Response, <-chan error) {
	responseChan := make(chan *StartCallV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartCallV2(request)
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

// StartCallV2WithCallback invokes the scsp.StartCallV2 API asynchronously
func (client *Client) StartCallV2WithCallback(request *StartCallV2Request, callback func(response *StartCallV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartCallV2Response
		var err error
		defer close(result)
		response, err = client.StartCallV2(request)
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

// StartCallV2Request is the request struct for api StartCallV2
type StartCallV2Request struct {
	*requests.RpcRequest
	ClientToken string `position:"Body"`
	InstanceId  string `position:"Body"`
	AccountName string `position:"Body"`
	Caller      string `position:"Body"`
	Callee      string `position:"Body"`
	JsonMsg     string `position:"Body"`
}

// StartCallV2Response is the response struct for api StartCallV2
type StartCallV2Response struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateStartCallV2Request creates a request to invoke StartCallV2 API
func CreateStartCallV2Request() (request *StartCallV2Request) {
	request = &StartCallV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "StartCallV2", "", "")
	request.Method = requests.POST
	return
}

// CreateStartCallV2Response creates a response to parse from StartCallV2 response
func CreateStartCallV2Response() (response *StartCallV2Response) {
	response = &StartCallV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}