package waf_openapi

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

// DescribeDomainConfigStatus invokes the waf_openapi.DescribeDomainConfigStatus API synchronously
func (client *Client) DescribeDomainConfigStatus(request *DescribeDomainConfigStatusRequest) (response *DescribeDomainConfigStatusResponse, err error) {
	response = CreateDescribeDomainConfigStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainConfigStatusWithChan invokes the waf_openapi.DescribeDomainConfigStatus API asynchronously
func (client *Client) DescribeDomainConfigStatusWithChan(request *DescribeDomainConfigStatusRequest) (<-chan *DescribeDomainConfigStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainConfigStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainConfigStatus(request)
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

// DescribeDomainConfigStatusWithCallback invokes the waf_openapi.DescribeDomainConfigStatus API asynchronously
func (client *Client) DescribeDomainConfigStatusWithCallback(request *DescribeDomainConfigStatusRequest, callback func(response *DescribeDomainConfigStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainConfigStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainConfigStatus(request)
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

// DescribeDomainConfigStatusRequest is the request struct for api DescribeDomainConfigStatus
type DescribeDomainConfigStatusRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Domain     string `position:"Query" name:"Domain"`
	Lang       string `position:"Query" name:"Lang"`
	Region     string `position:"Query" name:"Region"`
}

// DescribeDomainConfigStatusResponse is the response struct for api DescribeDomainConfigStatus
type DescribeDomainConfigStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeDomainConfigStatusRequest creates a request to invoke DescribeDomainConfigStatus API
func CreateDescribeDomainConfigStatusRequest() (request *DescribeDomainConfigStatusRequest) {
	request = &DescribeDomainConfigStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "DescribeDomainConfigStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainConfigStatusResponse creates a response to parse from DescribeDomainConfigStatus response
func CreateDescribeDomainConfigStatusResponse() (response *DescribeDomainConfigStatusResponse) {
	response = &DescribeDomainConfigStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}