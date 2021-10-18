package sddp

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

// DescribeInstancePortrait invokes the sddp.DescribeInstancePortrait API synchronously
func (client *Client) DescribeInstancePortrait(request *DescribeInstancePortraitRequest) (response *DescribeInstancePortraitResponse, err error) {
	response = CreateDescribeInstancePortraitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancePortraitWithChan invokes the sddp.DescribeInstancePortrait API asynchronously
func (client *Client) DescribeInstancePortraitWithChan(request *DescribeInstancePortraitRequest) (<-chan *DescribeInstancePortraitResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancePortraitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstancePortrait(request)
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

// DescribeInstancePortraitWithCallback invokes the sddp.DescribeInstancePortrait API asynchronously
func (client *Client) DescribeInstancePortraitWithCallback(request *DescribeInstancePortraitRequest, callback func(response *DescribeInstancePortraitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancePortraitResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstancePortrait(request)
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

// DescribeInstancePortraitRequest is the request struct for api DescribeInstancePortrait
type DescribeInstancePortraitRequest struct {
	*requests.RpcRequest
	ProductId  requests.Integer `position:"Query" name:"ProductId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	InstanceId requests.Integer `position:"Query" name:"InstanceId"`
	ItemKeys   string           `position:"Query" name:"ItemKeys"`
}

// DescribeInstancePortraitResponse is the response struct for api DescribeInstancePortrait
type DescribeInstancePortraitResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Content   []InstancePortrait `json:"Content" xml:"Content"`
}

// CreateDescribeInstancePortraitRequest creates a request to invoke DescribeInstancePortrait API
func CreateDescribeInstancePortraitRequest() (request *DescribeInstancePortraitRequest) {
	request = &DescribeInstancePortraitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeInstancePortrait", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeInstancePortraitResponse creates a response to parse from DescribeInstancePortrait response
func CreateDescribeInstancePortraitResponse() (response *DescribeInstancePortraitResponse) {
	response = &DescribeInstancePortraitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}