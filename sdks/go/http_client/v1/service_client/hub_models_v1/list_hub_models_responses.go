// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package hub_models_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListHubModelsReader is a Reader for the ListHubModels structure.
type ListHubModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHubModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListHubModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListHubModelsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListHubModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListHubModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListHubModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHubModelsOK creates a ListHubModelsOK with default headers values
func NewListHubModelsOK() *ListHubModelsOK {
	return &ListHubModelsOK{}
}

/*ListHubModelsOK handles this case with default header values.

A successful response.
*/
type ListHubModelsOK struct {
	Payload *service_model.V1ListHubModelsResponse
}

func (o *ListHubModelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models][%d] listHubModelsOK  %+v", 200, o.Payload)
}

func (o *ListHubModelsOK) GetPayload() *service_model.V1ListHubModelsResponse {
	return o.Payload
}

func (o *ListHubModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListHubModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHubModelsNoContent creates a ListHubModelsNoContent with default headers values
func NewListHubModelsNoContent() *ListHubModelsNoContent {
	return &ListHubModelsNoContent{}
}

/*ListHubModelsNoContent handles this case with default header values.

No content.
*/
type ListHubModelsNoContent struct {
	Payload interface{}
}

func (o *ListHubModelsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models][%d] listHubModelsNoContent  %+v", 204, o.Payload)
}

func (o *ListHubModelsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListHubModelsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHubModelsForbidden creates a ListHubModelsForbidden with default headers values
func NewListHubModelsForbidden() *ListHubModelsForbidden {
	return &ListHubModelsForbidden{}
}

/*ListHubModelsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListHubModelsForbidden struct {
	Payload interface{}
}

func (o *ListHubModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models][%d] listHubModelsForbidden  %+v", 403, o.Payload)
}

func (o *ListHubModelsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListHubModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHubModelsNotFound creates a ListHubModelsNotFound with default headers values
func NewListHubModelsNotFound() *ListHubModelsNotFound {
	return &ListHubModelsNotFound{}
}

/*ListHubModelsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListHubModelsNotFound struct {
	Payload interface{}
}

func (o *ListHubModelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models][%d] listHubModelsNotFound  %+v", 404, o.Payload)
}

func (o *ListHubModelsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListHubModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHubModelsDefault creates a ListHubModelsDefault with default headers values
func NewListHubModelsDefault(code int) *ListHubModelsDefault {
	return &ListHubModelsDefault{
		_statusCode: code,
	}
}

/*ListHubModelsDefault handles this case with default header values.

An unexpected error response
*/
type ListHubModelsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list hub models default response
func (o *ListHubModelsDefault) Code() int {
	return o._statusCode
}

func (o *ListHubModelsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models][%d] ListHubModels default  %+v", o._statusCode, o.Payload)
}

func (o *ListHubModelsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListHubModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}