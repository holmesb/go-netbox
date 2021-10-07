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

package dcim

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
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimConsolePortTemplatesUpdateParams creates a new DcimConsolePortTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesUpdateParams() *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithTimeout creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithContext creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesUpdateParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesUpdateParamsWithHTTPClient creates a new DcimConsolePortTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesUpdateParams {
	return &DcimConsolePortTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console port templates update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableConsolePortTemplate

	/* ID.

	   A unique integer value identifying this console port template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesUpdateParams) WithDefaults() *DcimConsolePortTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithData(data *models.WritableConsolePortTemplate) *DcimConsolePortTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) WithID(id int64) *DcimConsolePortTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates update params
func (o *DcimConsolePortTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
