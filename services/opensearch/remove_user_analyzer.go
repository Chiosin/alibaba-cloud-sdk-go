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

// RemoveUserAnalyzer invokes the opensearch.RemoveUserAnalyzer API synchronously
func (client *Client) RemoveUserAnalyzer(request *RemoveUserAnalyzerRequest) (response *RemoveUserAnalyzerResponse, err error) {
	response = CreateRemoveUserAnalyzerResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveUserAnalyzerWithChan invokes the opensearch.RemoveUserAnalyzer API asynchronously
func (client *Client) RemoveUserAnalyzerWithChan(request *RemoveUserAnalyzerRequest) (<-chan *RemoveUserAnalyzerResponse, <-chan error) {
	responseChan := make(chan *RemoveUserAnalyzerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveUserAnalyzer(request)
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

// RemoveUserAnalyzerWithCallback invokes the opensearch.RemoveUserAnalyzer API asynchronously
func (client *Client) RemoveUserAnalyzerWithCallback(request *RemoveUserAnalyzerRequest, callback func(response *RemoveUserAnalyzerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveUserAnalyzerResponse
		var err error
		defer close(result)
		response, err = client.RemoveUserAnalyzer(request)
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

// RemoveUserAnalyzerRequest is the request struct for api RemoveUserAnalyzer
type RemoveUserAnalyzerRequest struct {
	*requests.RoaRequest
	Name string `position:"Path" name:"name"`
}

// RemoveUserAnalyzerResponse is the response struct for api RemoveUserAnalyzer
type RemoveUserAnalyzerResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"requestId" xml:"requestId"`
}

// CreateRemoveUserAnalyzerRequest creates a request to invoke RemoveUserAnalyzer API
func CreateRemoveUserAnalyzerRequest() (request *RemoveUserAnalyzerRequest) {
	request = &RemoveUserAnalyzerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RemoveUserAnalyzer", "/v4/openapi/user-analyzers/[name]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateRemoveUserAnalyzerResponse creates a response to parse from RemoveUserAnalyzer response
func CreateRemoveUserAnalyzerResponse() (response *RemoveUserAnalyzerResponse) {
	response = &RemoveUserAnalyzerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
