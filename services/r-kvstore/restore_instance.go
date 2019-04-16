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

// RestoreInstance invokes the r_kvstore.RestoreInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoreinstance.html
func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
	response = CreateRestoreInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestoreInstanceWithChan invokes the r_kvstore.RestoreInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoreinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreInstanceWithChan(request *RestoreInstanceRequest) (<-chan *RestoreInstanceResponse, <-chan error) {
	responseChan := make(chan *RestoreInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestoreInstance(request)
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

// RestoreInstanceWithCallback invokes the r_kvstore.RestoreInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restoreinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreInstanceWithCallback(request *RestoreInstanceRequest, callback func(response *RestoreInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestoreInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestoreInstance(request)
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

// RestoreInstanceRequest is the request struct for api RestoreInstance
type RestoreInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RestoreInstanceResponse is the response struct for api RestoreInstance
type RestoreInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestoreInstanceRequest creates a request to invoke RestoreInstance API
func CreateRestoreInstanceRequest() (request *RestoreInstanceRequest) {
	request = &RestoreInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreInstance", "R-kvstore", "openAPI")
	return
}

// CreateRestoreInstanceResponse creates a response to parse from RestoreInstance response
func CreateRestoreInstanceResponse() (response *RestoreInstanceResponse) {
	response = &RestoreInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
