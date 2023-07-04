package rds

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

// ModifySQLCollectorRetention invokes the rds.ModifySQLCollectorRetention API synchronously
func (client *Client) ModifySQLCollectorRetention(request *ModifySQLCollectorRetentionRequest) (response *ModifySQLCollectorRetentionResponse, err error) {
	response = CreateModifySQLCollectorRetentionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySQLCollectorRetentionWithChan invokes the rds.ModifySQLCollectorRetention API asynchronously
func (client *Client) ModifySQLCollectorRetentionWithChan(request *ModifySQLCollectorRetentionRequest) (<-chan *ModifySQLCollectorRetentionResponse, <-chan error) {
	responseChan := make(chan *ModifySQLCollectorRetentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySQLCollectorRetention(request)
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

// ModifySQLCollectorRetentionWithCallback invokes the rds.ModifySQLCollectorRetention API asynchronously
func (client *Client) ModifySQLCollectorRetentionWithCallback(request *ModifySQLCollectorRetentionRequest, callback func(response *ModifySQLCollectorRetentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySQLCollectorRetentionResponse
		var err error
		defer close(result)
		response, err = client.ModifySQLCollectorRetention(request)
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

// ModifySQLCollectorRetentionRequest is the request struct for api ModifySQLCollectorRetention
type ModifySQLCollectorRetentionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ConfigValue          string           `position:"Query" name:"ConfigValue"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifySQLCollectorRetentionResponse is the response struct for api ModifySQLCollectorRetention
type ModifySQLCollectorRetentionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySQLCollectorRetentionRequest creates a request to invoke ModifySQLCollectorRetention API
func CreateModifySQLCollectorRetentionRequest() (request *ModifySQLCollectorRetentionRequest) {
	request = &ModifySQLCollectorRetentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifySQLCollectorRetention", "", "")
	request.Method = requests.POST
	return
}

// CreateModifySQLCollectorRetentionResponse creates a response to parse from ModifySQLCollectorRetention response
func CreateModifySQLCollectorRetentionResponse() (response *ModifySQLCollectorRetentionResponse) {
	response = &ModifySQLCollectorRetentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
