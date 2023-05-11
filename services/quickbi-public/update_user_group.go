package quickbi_public

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

// UpdateUserGroup invokes the quickbi_public.UpdateUserGroup API synchronously
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (response *UpdateUserGroupResponse, err error) {
	response = CreateUpdateUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserGroupWithChan invokes the quickbi_public.UpdateUserGroup API asynchronously
func (client *Client) UpdateUserGroupWithChan(request *UpdateUserGroupRequest) (<-chan *UpdateUserGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserGroup(request)
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

// UpdateUserGroupWithCallback invokes the quickbi_public.UpdateUserGroup API asynchronously
func (client *Client) UpdateUserGroupWithCallback(request *UpdateUserGroupRequest, callback func(response *UpdateUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserGroup(request)
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

// UpdateUserGroupRequest is the request struct for api UpdateUserGroup
type UpdateUserGroupRequest struct {
	*requests.RpcRequest
	AccessPoint          string `position:"Query" name:"AccessPoint"`
	UserGroupId          string `position:"Query" name:"UserGroupId"`
	SignType             string `position:"Query" name:"SignType"`
	UserGroupName        string `position:"Query" name:"UserGroupName"`
	UserGroupDescription string `position:"Query" name:"UserGroupDescription"`
}

// UpdateUserGroupResponse is the response struct for api UpdateUserGroup
type UpdateUserGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateUserGroupRequest creates a request to invoke UpdateUserGroup API
func CreateUpdateUserGroupRequest() (request *UpdateUserGroupRequest) {
	request = &UpdateUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateUserGroup", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateUserGroupResponse creates a response to parse from UpdateUserGroup response
func CreateUpdateUserGroupResponse() (response *UpdateUserGroupResponse) {
	response = &UpdateUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
