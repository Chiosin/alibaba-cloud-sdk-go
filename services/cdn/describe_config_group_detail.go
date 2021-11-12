package cdn

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

// DescribeConfigGroupDetail invokes the cdn.DescribeConfigGroupDetail API synchronously
func (client *Client) DescribeConfigGroupDetail(request *DescribeConfigGroupDetailRequest) (response *DescribeConfigGroupDetailResponse, err error) {
	response = CreateDescribeConfigGroupDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigGroupDetailWithChan invokes the cdn.DescribeConfigGroupDetail API asynchronously
func (client *Client) DescribeConfigGroupDetailWithChan(request *DescribeConfigGroupDetailRequest) (<-chan *DescribeConfigGroupDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigGroupDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigGroupDetail(request)
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

// DescribeConfigGroupDetailWithCallback invokes the cdn.DescribeConfigGroupDetail API asynchronously
func (client *Client) DescribeConfigGroupDetailWithCallback(request *DescribeConfigGroupDetailRequest, callback func(response *DescribeConfigGroupDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigGroupDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigGroupDetail(request)
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

// DescribeConfigGroupDetailRequest is the request struct for api DescribeConfigGroupDetail
type DescribeConfigGroupDetailRequest struct {
	*requests.RpcRequest
	ConfigGroupName string           `position:"Query" name:"ConfigGroupName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	ConfigGroupId   string           `position:"Query" name:"ConfigGroupId"`
}

// DescribeConfigGroupDetailResponse is the response struct for api DescribeConfigGroupDetail
type DescribeConfigGroupDetailResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ConfigGroupId   string `json:"ConfigGroupId" xml:"ConfigGroupId"`
	ConfigGroupName string `json:"ConfigGroupName" xml:"ConfigGroupName"`
	Description     string `json:"Description" xml:"Description"`
	BizName         string `json:"BizName" xml:"BizName"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	UpdateTime      string `json:"UpdateTime" xml:"UpdateTime"`
}

// CreateDescribeConfigGroupDetailRequest creates a request to invoke DescribeConfigGroupDetail API
func CreateDescribeConfigGroupDetailRequest() (request *DescribeConfigGroupDetailRequest) {
	request = &DescribeConfigGroupDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeConfigGroupDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeConfigGroupDetailResponse creates a response to parse from DescribeConfigGroupDetail response
func CreateDescribeConfigGroupDetailResponse() (response *DescribeConfigGroupDetailResponse) {
	response = &DescribeConfigGroupDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}