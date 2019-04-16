package r_kvstore

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

// DescribeStaticVerificationList invokes the r_kvstore.DescribeStaticVerificationList API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describestaticverificationlist.html
func (client *Client) DescribeStaticVerificationList(request *DescribeStaticVerificationListRequest) (response *DescribeStaticVerificationListResponse, err error) {
	response = CreateDescribeStaticVerificationListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStaticVerificationListWithChan invokes the r_kvstore.DescribeStaticVerificationList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describestaticverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStaticVerificationListWithChan(request *DescribeStaticVerificationListRequest) (<-chan *DescribeStaticVerificationListResponse, <-chan error) {
	responseChan := make(chan *DescribeStaticVerificationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStaticVerificationList(request)
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

// DescribeStaticVerificationListWithCallback invokes the r_kvstore.DescribeStaticVerificationList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describestaticverificationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStaticVerificationListWithCallback(request *DescribeStaticVerificationListRequest, callback func(response *DescribeStaticVerificationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStaticVerificationListResponse
		var err error
		defer close(result)
		response, err = client.DescribeStaticVerificationList(request)
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

// DescribeStaticVerificationListRequest is the request struct for api DescribeStaticVerificationList
type DescribeStaticVerificationListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	ReplicaId             string           `position:"Query" name:"ReplicaId"`
	DestinationInstanceId string           `position:"Query" name:"DestinationInstanceId"`
	SourceInstanceId      string           `position:"Query" name:"SourceInstanceId"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStaticVerificationListResponse is the response struct for api DescribeStaticVerificationList
type DescribeStaticVerificationListResponse struct {
	*responses.BaseResponse
	RequestId              string      `json:"RequestId" xml:"RequestId"`
	ReplicaId              string      `json:"ReplicaId" xml:"ReplicaId"`
	SourceInstanceId       string      `json:"SourceInstanceId" xml:"SourceInstanceId"`
	SourceDBNumber         int         `json:"SourceDBNumber" xml:"SourceDBNumber"`
	SourceTableNumber      int         `json:"SourceTableNumber" xml:"SourceTableNumber"`
	SourceCountNumber      int         `json:"SourceCountNumber" xml:"SourceCountNumber"`
	SourceDBSize           int         `json:"SourceDBSize" xml:"SourceDBSize"`
	DestinationInstanceId  string      `json:"DestinationInstanceId" xml:"DestinationInstanceId"`
	DestinationDBNumber    int         `json:"DestinationDBNumber" xml:"DestinationDBNumber"`
	DestinationTableNumber int         `json:"DestinationTableNumber" xml:"DestinationTableNumber"`
	DestinationCountNumber int         `json:"DestinationCountNumber" xml:"DestinationCountNumber"`
	DestinationDBSize      int         `json:"DestinationDBSize" xml:"DestinationDBSize"`
	ConsistencyPercent     string      `json:"ConsistencyPercent" xml:"ConsistencyPercent"`
	Items                  []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeStaticVerificationListRequest creates a request to invoke DescribeStaticVerificationList API
func CreateDescribeStaticVerificationListRequest() (request *DescribeStaticVerificationListRequest) {
	request = &DescribeStaticVerificationListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeStaticVerificationList", "R-kvstore", "openAPI")
	return
}

// CreateDescribeStaticVerificationListResponse creates a response to parse from DescribeStaticVerificationList response
func CreateDescribeStaticVerificationListResponse() (response *DescribeStaticVerificationListResponse) {
	response = &DescribeStaticVerificationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
