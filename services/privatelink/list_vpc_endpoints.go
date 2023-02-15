package privatelink

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

// ListVpcEndpoints invokes the privatelink.ListVpcEndpoints API synchronously
func (client *Client) ListVpcEndpoints(request *ListVpcEndpointsRequest) (response *ListVpcEndpointsResponse, err error) {
	response = CreateListVpcEndpointsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcEndpointsWithChan invokes the privatelink.ListVpcEndpoints API asynchronously
func (client *Client) ListVpcEndpointsWithChan(request *ListVpcEndpointsRequest) (<-chan *ListVpcEndpointsResponse, <-chan error) {
	responseChan := make(chan *ListVpcEndpointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcEndpoints(request)
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

// ListVpcEndpointsWithCallback invokes the privatelink.ListVpcEndpoints API asynchronously
func (client *Client) ListVpcEndpointsWithCallback(request *ListVpcEndpointsRequest, callback func(response *ListVpcEndpointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcEndpointsResponse
		var err error
		defer close(result)
		response, err = client.ListVpcEndpoints(request)
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

// ListVpcEndpointsRequest is the request struct for api ListVpcEndpoints
type ListVpcEndpointsRequest struct {
	*requests.RpcRequest
	EndpointId       string                 `position:"Query" name:"EndpointId"`
	EndpointStatus   string                 `position:"Query" name:"EndpointStatus"`
	ResourceGroupId  string                 `position:"Query" name:"ResourceGroupId"`
	NextToken        string                 `position:"Query" name:"NextToken"`
	EndpointType     string                 `position:"Query" name:"EndpointType"`
	ServiceName      string                 `position:"Query" name:"ServiceName"`
	Tag              *[]ListVpcEndpointsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ConnectionStatus string                 `position:"Query" name:"ConnectionStatus"`
	VpcId            string                 `position:"Query" name:"VpcId"`
	EndpointName     string                 `position:"Query" name:"EndpointName"`
	MaxResults       requests.Integer       `position:"Query" name:"MaxResults"`
}

// ListVpcEndpointsTag is a repeated param struct in ListVpcEndpointsRequest
type ListVpcEndpointsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListVpcEndpointsResponse is the response struct for api ListVpcEndpoints
type ListVpcEndpointsResponse struct {
	*responses.BaseResponse
	NextToken  string     `json:"NextToken" xml:"NextToken"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	MaxResults int        `json:"MaxResults" xml:"MaxResults"`
	Endpoints  []Endpoint `json:"Endpoints" xml:"Endpoints"`
}

// CreateListVpcEndpointsRequest creates a request to invoke ListVpcEndpoints API
func CreateListVpcEndpointsRequest() (request *ListVpcEndpointsRequest) {
	request = &ListVpcEndpointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "ListVpcEndpoints", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpcEndpointsResponse creates a response to parse from ListVpcEndpoints response
func CreateListVpcEndpointsResponse() (response *ListVpcEndpointsResponse) {
	response = &ListVpcEndpointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
