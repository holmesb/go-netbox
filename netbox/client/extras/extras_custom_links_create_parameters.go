// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/holmesb/go-netbox/netbox/models"
)

// NewExtrasCustomLinksCreateParams creates a new ExtrasCustomLinksCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksCreateParams() *ExtrasCustomLinksCreateParams {
	return &ExtrasCustomLinksCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksCreateParamsWithTimeout creates a new ExtrasCustomLinksCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksCreateParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksCreateParams {
	return &ExtrasCustomLinksCreateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksCreateParamsWithContext creates a new ExtrasCustomLinksCreateParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksCreateParamsWithContext(ctx context.Context) *ExtrasCustomLinksCreateParams {
	return &ExtrasCustomLinksCreateParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksCreateParamsWithHTTPClient creates a new ExtrasCustomLinksCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksCreateParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksCreateParams {
	return &ExtrasCustomLinksCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomLinksCreateParams contains all the parameters to send to the API endpoint
   for the extras custom links create operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomLinksCreateParams struct {

	// Data.
	Data *models.CustomLink

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksCreateParams) WithDefaults() *ExtrasCustomLinksCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) WithContext(ctx context.Context) *ExtrasCustomLinksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) WithData(data *models.CustomLink) *ExtrasCustomLinksCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom links create params
func (o *ExtrasCustomLinksCreateParams) SetData(data *models.CustomLink) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
