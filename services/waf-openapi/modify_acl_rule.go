package waf_openapi

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

// ModifyAclRule invokes the waf_openapi.ModifyAclRule API synchronously
func (client *Client) ModifyAclRule(request *ModifyAclRuleRequest) (response *ModifyAclRuleResponse, err error) {
	response = CreateModifyAclRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAclRuleWithChan invokes the waf_openapi.ModifyAclRule API asynchronously
func (client *Client) ModifyAclRuleWithChan(request *ModifyAclRuleRequest) (<-chan *ModifyAclRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyAclRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAclRule(request)
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

// ModifyAclRuleWithCallback invokes the waf_openapi.ModifyAclRule API asynchronously
func (client *Client) ModifyAclRuleWithCallback(request *ModifyAclRuleRequest, callback func(response *ModifyAclRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAclRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyAclRule(request)
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

// ModifyAclRuleRequest is the request struct for api ModifyAclRule
type ModifyAclRuleRequest struct {
	*requests.RpcRequest
	Rules      string `position:"Query" name:"Rules"`
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Domain     string `position:"Query" name:"Domain"`
	Lang       string `position:"Query" name:"Lang"`
	Region     string `position:"Query" name:"Region"`
}

// ModifyAclRuleResponse is the response struct for api ModifyAclRule
type ModifyAclRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateModifyAclRuleRequest creates a request to invoke ModifyAclRule API
func CreateModifyAclRuleRequest() (request *ModifyAclRuleRequest) {
	request = &ModifyAclRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "ModifyAclRule", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyAclRuleResponse creates a response to parse from ModifyAclRule response
func CreateModifyAclRuleResponse() (response *ModifyAclRuleResponse) {
	response = &ModifyAclRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
