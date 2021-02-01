package polardb

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

// DescribeSQLExplorerRetention invokes the polardb.DescribeSQLExplorerRetention API synchronously
func (client *Client) DescribeSQLExplorerRetention(request *DescribeSQLExplorerRetentionRequest) (response *DescribeSQLExplorerRetentionResponse, err error) {
	response = CreateDescribeSQLExplorerRetentionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLExplorerRetentionWithChan invokes the polardb.DescribeSQLExplorerRetention API asynchronously
func (client *Client) DescribeSQLExplorerRetentionWithChan(request *DescribeSQLExplorerRetentionRequest) (<-chan *DescribeSQLExplorerRetentionResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLExplorerRetentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLExplorerRetention(request)
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

// DescribeSQLExplorerRetentionWithCallback invokes the polardb.DescribeSQLExplorerRetention API asynchronously
func (client *Client) DescribeSQLExplorerRetentionWithCallback(request *DescribeSQLExplorerRetentionRequest, callback func(response *DescribeSQLExplorerRetentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLExplorerRetentionResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLExplorerRetention(request)
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

// DescribeSQLExplorerRetentionRequest is the request struct for api DescribeSQLExplorerRetention
type DescribeSQLExplorerRetentionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSQLExplorerRetentionResponse is the response struct for api DescribeSQLExplorerRetention
type DescribeSQLExplorerRetentionResponse struct {
	*responses.BaseResponse
	ConfigValue    string `json:"ConfigValue" xml:"ConfigValue"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceID   int    `json:"DBInstanceID" xml:"DBInstanceID"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateDescribeSQLExplorerRetentionRequest creates a request to invoke DescribeSQLExplorerRetention API
func CreateDescribeSQLExplorerRetentionRequest() (request *DescribeSQLExplorerRetentionRequest) {
	request = &DescribeSQLExplorerRetentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeSQLExplorerRetention", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLExplorerRetentionResponse creates a response to parse from DescribeSQLExplorerRetention response
func CreateDescribeSQLExplorerRetentionResponse() (response *DescribeSQLExplorerRetentionResponse) {
	response = &DescribeSQLExplorerRetentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}