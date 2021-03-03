package dms_enterprise

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

// ListDBTaskSQLJobDetail invokes the dms_enterprise.ListDBTaskSQLJobDetail API synchronously
func (client *Client) ListDBTaskSQLJobDetail(request *ListDBTaskSQLJobDetailRequest) (response *ListDBTaskSQLJobDetailResponse, err error) {
	response = CreateListDBTaskSQLJobDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ListDBTaskSQLJobDetailWithChan invokes the dms_enterprise.ListDBTaskSQLJobDetail API asynchronously
func (client *Client) ListDBTaskSQLJobDetailWithChan(request *ListDBTaskSQLJobDetailRequest) (<-chan *ListDBTaskSQLJobDetailResponse, <-chan error) {
	responseChan := make(chan *ListDBTaskSQLJobDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDBTaskSQLJobDetail(request)
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

// ListDBTaskSQLJobDetailWithCallback invokes the dms_enterprise.ListDBTaskSQLJobDetail API asynchronously
func (client *Client) ListDBTaskSQLJobDetailWithCallback(request *ListDBTaskSQLJobDetailRequest, callback func(response *ListDBTaskSQLJobDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDBTaskSQLJobDetailResponse
		var err error
		defer close(result)
		response, err = client.ListDBTaskSQLJobDetail(request)
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

// ListDBTaskSQLJobDetailRequest is the request struct for api ListDBTaskSQLJobDetail
type ListDBTaskSQLJobDetailRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	JobId      requests.Integer `position:"Query" name:"JobId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListDBTaskSQLJobDetailResponse is the response struct for api ListDBTaskSQLJobDetail
type ListDBTaskSQLJobDetailResponse struct {
	*responses.BaseResponse
	TotalCount             int64                `json:"TotalCount" xml:"TotalCount"`
	RequestId              string               `json:"RequestId" xml:"RequestId"`
	ErrorCode              string               `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage           string               `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                bool                 `json:"Success" xml:"Success"`
	DBTaskSQLJobDetailList []DBTaskSQLJobDetail `json:"DBTaskSQLJobDetailList" xml:"DBTaskSQLJobDetailList"`
}

// CreateListDBTaskSQLJobDetailRequest creates a request to invoke ListDBTaskSQLJobDetail API
func CreateListDBTaskSQLJobDetailRequest() (request *ListDBTaskSQLJobDetailRequest) {
	request = &ListDBTaskSQLJobDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDBTaskSQLJobDetail", "dmsenterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDBTaskSQLJobDetailResponse creates a response to parse from ListDBTaskSQLJobDetail response
func CreateListDBTaskSQLJobDetailResponse() (response *ListDBTaskSQLJobDetailResponse) {
	response = &ListDBTaskSQLJobDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
