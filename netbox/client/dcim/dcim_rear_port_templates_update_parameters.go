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

	"github.com/holmesb/go-netbox/netbox/models"
)

// NewDcimRearPortTemplatesUpdateParams creates a new DcimRearPortTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesUpdateParams() *DcimRearPortTemplatesUpdateParams {
	return &DcimRearPortTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesUpdateParamsWithTimeout creates a new DcimRearPortTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesUpdateParams {
	return &DcimRearPortTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesUpdateParamsWithContext creates a new DcimRearPortTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesUpdateParamsWithContext(ctx context.Context) *DcimRearPortTemplatesUpdateParams {
	return &DcimRearPortTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesUpdateParamsWithHTTPClient creates a new DcimRearPortTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesUpdateParams {
	return &DcimRearPortTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates update operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableRearPortTemplate

	/* ID.

	   A unique integer value identifying this rear port template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesUpdateParams) WithDefaults() *DcimRearPortTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) WithContext(ctx context.Context) *DcimRearPortTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) WithData(data *models.WritableRearPortTemplate) *DcimRearPortTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) SetData(data *models.WritableRearPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) WithID(id int64) *DcimRearPortTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates update params
func (o *DcimRearPortTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
