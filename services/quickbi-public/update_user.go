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

// UpdateUser invokes the quickbi_public.UpdateUser API synchronously
func (client *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
	response = CreateUpdateUserResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserWithChan invokes the quickbi_public.UpdateUser API asynchronously
func (client *Client) UpdateUserWithChan(request *UpdateUserRequest) (<-chan *UpdateUserResponse, <-chan error) {
	responseChan := make(chan *UpdateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUser(request)
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

// UpdateUserWithCallback invokes the quickbi_public.UpdateUser API asynchronously
func (client *Client) UpdateUserWithCallback(request *UpdateUserRequest, callback func(response *UpdateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserResponse
		var err error
		defer close(result)
		response, err = client.UpdateUser(request)
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

// UpdateUserRequest is the request struct for api UpdateUser
type UpdateUserRequest struct {
	*requests.RpcRequest
	AdminUser     requests.Boolean `position:"Query" name:"AdminUser"`
	AccountType   requests.Integer `position:"Query" name:"AccountType"`
	AccessPoint   string           `position:"Query" name:"AccessPoint"`
	UserType      requests.Integer `position:"Query" name:"UserType"`
	SignType      string           `position:"Query" name:"SignType"`
	UserId        string           `position:"Query" name:"UserId"`
	Phone         string           `position:"Query" name:"Phone"`
	NickName      string           `position:"Query" name:"NickName"`
	AuthAdminUser requests.Boolean `position:"Query" name:"AuthAdminUser"`
	Email         string           `position:"Query" name:"Email"`
}

// UpdateUserResponse is the response struct for api UpdateUser
type UpdateUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateUserRequest creates a request to invoke UpdateUser API
func CreateUpdateUserRequest() (request *UpdateUserRequest) {
	request = &UpdateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateUser", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateUserResponse creates a response to parse from UpdateUser response
func CreateUpdateUserResponse() (response *UpdateUserResponse) {
	response = &UpdateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
