package eflo

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

// ListErRouteEntries invokes the eflo.ListErRouteEntries API synchronously
func (client *Client) ListErRouteEntries(request *ListErRouteEntriesRequest) (response *ListErRouteEntriesResponse, err error) {
	response = CreateListErRouteEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListErRouteEntriesWithChan invokes the eflo.ListErRouteEntries API asynchronously
func (client *Client) ListErRouteEntriesWithChan(request *ListErRouteEntriesRequest) (<-chan *ListErRouteEntriesResponse, <-chan error) {
	responseChan := make(chan *ListErRouteEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListErRouteEntries(request)
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

// ListErRouteEntriesWithCallback invokes the eflo.ListErRouteEntries API asynchronously
func (client *Client) ListErRouteEntriesWithCallback(request *ListErRouteEntriesRequest, callback func(response *ListErRouteEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListErRouteEntriesResponse
		var err error
		defer close(result)
		response, err = client.ListErRouteEntries(request)
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

// ListErRouteEntriesRequest is the request struct for api ListErRouteEntries
type ListErRouteEntriesRequest struct {
	*requests.RpcRequest
	PageNumber           requests.Integer `position:"Body" name:"PageNumber"`
	RouteType            string           `position:"Body" name:"RouteType"`
	PageSize             requests.Integer `position:"Body" name:"PageSize"`
	NextHopId            string           `position:"Body" name:"NextHopId"`
	NextHopType          string           `position:"Body" name:"NextHopType"`
	DestinationCidrBlock string           `position:"Body" name:"DestinationCidrBlock"`
	ErId                 string           `position:"Body" name:"ErId"`
	InstanceId           string           `position:"Body" name:"InstanceId"`
	EnablePage           requests.Boolean `position:"Body" name:"EnablePage"`
	Status               string           `position:"Body" name:"Status"`
}

// ListErRouteEntriesResponse is the response struct for api ListErRouteEntries
type ListErRouteEntriesResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateListErRouteEntriesRequest creates a request to invoke ListErRouteEntries API
func CreateListErRouteEntriesRequest() (request *ListErRouteEntriesRequest) {
	request = &ListErRouteEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListErRouteEntries", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListErRouteEntriesResponse creates a response to parse from ListErRouteEntries response
func CreateListErRouteEntriesResponse() (response *ListErRouteEntriesResponse) {
	response = &ListErRouteEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
