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

// DescribeModel invokes the opensearch.DescribeModel API synchronously
func (client *Client) DescribeModel(request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
	response = CreateDescribeModelResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeModelWithChan invokes the opensearch.DescribeModel API asynchronously
func (client *Client) DescribeModelWithChan(request *DescribeModelRequest) (<-chan *DescribeModelResponse, <-chan error) {
	responseChan := make(chan *DescribeModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeModel(request)
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

// DescribeModelWithCallback invokes the opensearch.DescribeModel API asynchronously
func (client *Client) DescribeModelWithCallback(request *DescribeModelRequest, callback func(response *DescribeModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeModelResponse
		var err error
		defer close(result)
		response, err = client.DescribeModel(request)
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

// DescribeModelRequest is the request struct for api DescribeModel
type DescribeModelRequest struct {
	*requests.RoaRequest
	ModelName        string `position:"Path" name:"modelName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DescribeModelResponse is the response struct for api DescribeModel
type DescribeModelResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"requestId" xml:"requestId"`
}

// CreateDescribeModelRequest creates a request to invoke DescribeModel API
func CreateDescribeModelRequest() (request *DescribeModelRequest) {
	request = &DescribeModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeModel", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/models/[modelName]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeModelResponse creates a response to parse from DescribeModel response
func CreateDescribeModelResponse() (response *DescribeModelResponse) {
	response = &DescribeModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
