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

// ListQuotaReviewTasks invokes the opensearch.ListQuotaReviewTasks API synchronously
func (client *Client) ListQuotaReviewTasks(request *ListQuotaReviewTasksRequest) (response *ListQuotaReviewTasksResponse, err error) {
	response = CreateListQuotaReviewTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListQuotaReviewTasksWithChan invokes the opensearch.ListQuotaReviewTasks API asynchronously
func (client *Client) ListQuotaReviewTasksWithChan(request *ListQuotaReviewTasksRequest) (<-chan *ListQuotaReviewTasksResponse, <-chan error) {
	responseChan := make(chan *ListQuotaReviewTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQuotaReviewTasks(request)
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

// ListQuotaReviewTasksWithCallback invokes the opensearch.ListQuotaReviewTasks API asynchronously
func (client *Client) ListQuotaReviewTasksWithCallback(request *ListQuotaReviewTasksRequest, callback func(response *ListQuotaReviewTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQuotaReviewTasksResponse
		var err error
		defer close(result)
		response, err = client.ListQuotaReviewTasks(request)
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

// ListQuotaReviewTasksRequest is the request struct for api ListQuotaReviewTasks
type ListQuotaReviewTasksRequest struct {
	*requests.RoaRequest
	PageSize         requests.Integer `position:"Query" name:"pageSize"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
	PageNumber       requests.Integer `position:"Query" name:"pageNumber"`
}

// ListQuotaReviewTasksResponse is the response struct for api ListQuotaReviewTasks
type ListQuotaReviewTasksResponse struct {
	*responses.BaseResponse
	TotalCount int                                `json:"totalCount" xml:"totalCount"`
	RequestId  string                             `json:"requestId" xml:"requestId"`
	Result     []ResultItemInListQuotaReviewTasks `json:"result" xml:"result"`
}

// CreateListQuotaReviewTasksRequest creates a request to invoke ListQuotaReviewTasks API
func CreateListQuotaReviewTasksRequest() (request *ListQuotaReviewTasksRequest) {
	request = &ListQuotaReviewTasksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListQuotaReviewTasks", "/v4/openapi/app-groups/[appGroupIdentity]/quota-review-tasks", "", "")
	request.Method = requests.GET
	return
}

// CreateListQuotaReviewTasksResponse creates a response to parse from ListQuotaReviewTasks response
func CreateListQuotaReviewTasksResponse() (response *ListQuotaReviewTasksResponse) {
	response = &ListQuotaReviewTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
