// Copyright 2019 Polyaxon, Inc.
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

package k8_s_config_maps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateK8SConfigMapsParams creates a new CreateK8SConfigMapsParams object
// with the default values initialized.
func NewCreateK8SConfigMapsParams() *CreateK8SConfigMapsParams {
	var ()
	return &CreateK8SConfigMapsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateK8SConfigMapsParamsWithTimeout creates a new CreateK8SConfigMapsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateK8SConfigMapsParamsWithTimeout(timeout time.Duration) *CreateK8SConfigMapsParams {
	var ()
	return &CreateK8SConfigMapsParams{

		timeout: timeout,
	}
}

// NewCreateK8SConfigMapsParamsWithContext creates a new CreateK8SConfigMapsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateK8SConfigMapsParamsWithContext(ctx context.Context) *CreateK8SConfigMapsParams {
	var ()
	return &CreateK8SConfigMapsParams{

		Context: ctx,
	}
}

// NewCreateK8SConfigMapsParamsWithHTTPClient creates a new CreateK8SConfigMapsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateK8SConfigMapsParamsWithHTTPClient(client *http.Client) *CreateK8SConfigMapsParams {
	var ()
	return &CreateK8SConfigMapsParams{
		HTTPClient: client,
	}
}

/*CreateK8SConfigMapsParams contains all the parameters to send to the API endpoint
for the create k8 s config maps operation typically these are written to a http.Request
*/
type CreateK8SConfigMapsParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1K8SResource
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) WithTimeout(timeout time.Duration) *CreateK8SConfigMapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) WithContext(ctx context.Context) *CreateK8SConfigMapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) WithHTTPClient(client *http.Client) *CreateK8SConfigMapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) WithBody(body *service_model.V1K8SResource) *CreateK8SConfigMapsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) SetBody(body *service_model.V1K8SResource) {
	o.Body = body
}

// WithOwner adds the owner to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) WithOwner(owner string) *CreateK8SConfigMapsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create k8 s config maps params
func (o *CreateK8SConfigMapsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateK8SConfigMapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}