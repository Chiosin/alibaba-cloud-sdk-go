package config

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

// GenerateAggregateResourceInventory invokes the config.GenerateAggregateResourceInventory API synchronously
func (client *Client) GenerateAggregateResourceInventory(request *GenerateAggregateResourceInventoryRequest) (response *GenerateAggregateResourceInventoryResponse, err error) {
	response = CreateGenerateAggregateResourceInventoryResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateAggregateResourceInventoryWithChan invokes the config.GenerateAggregateResourceInventory API asynchronously
func (client *Client) GenerateAggregateResourceInventoryWithChan(request *GenerateAggregateResourceInventoryRequest) (<-chan *GenerateAggregateResourceInventoryResponse, <-chan error) {
	responseChan := make(chan *GenerateAggregateResourceInventoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateAggregateResourceInventory(request)
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

// GenerateAggregateResourceInventoryWithCallback invokes the config.GenerateAggregateResourceInventory API asynchronously
func (client *Client) GenerateAggregateResourceInventoryWithCallback(request *GenerateAggregateResourceInventoryRequest, callback func(response *GenerateAggregateResourceInventoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateAggregateResourceInventoryResponse
		var err error
		defer close(result)
		response, err = client.GenerateAggregateResourceInventory(request)
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

// GenerateAggregateResourceInventoryRequest is the request struct for api GenerateAggregateResourceInventory
type GenerateAggregateResourceInventoryRequest struct {
	*requests.RpcRequest
	Regions       string `position:"Query" name:"Regions"`
	ResourceTypes string `position:"Query" name:"ResourceTypes"`
	AggregatorId  string `position:"Query" name:"AggregatorId"`
	AccountIds    string `position:"Query" name:"AccountIds"`
}

// GenerateAggregateResourceInventoryResponse is the response struct for api GenerateAggregateResourceInventory
type GenerateAggregateResourceInventoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGenerateAggregateResourceInventoryRequest creates a request to invoke GenerateAggregateResourceInventory API
func CreateGenerateAggregateResourceInventoryRequest() (request *GenerateAggregateResourceInventoryRequest) {
	request = &GenerateAggregateResourceInventoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GenerateAggregateResourceInventory", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateAggregateResourceInventoryResponse creates a response to parse from GenerateAggregateResourceInventory response
func CreateGenerateAggregateResourceInventoryResponse() (response *GenerateAggregateResourceInventoryResponse) {
	response = &GenerateAggregateResourceInventoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
