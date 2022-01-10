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

package tenancy

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewTenancyTenantsBulkPartialUpdateParams creates a new TenancyTenantsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsBulkPartialUpdateParams() *TenancyTenantsBulkPartialUpdateParams {
	return &TenancyTenantsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsBulkPartialUpdateParamsWithTimeout creates a new TenancyTenantsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantsBulkPartialUpdateParams {
	return &TenancyTenantsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsBulkPartialUpdateParamsWithContext creates a new TenancyTenantsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantsBulkPartialUpdateParamsWithContext(ctx context.Context) *TenancyTenantsBulkPartialUpdateParams {
	return &TenancyTenantsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantsBulkPartialUpdateParamsWithHTTPClient creates a new TenancyTenantsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantsBulkPartialUpdateParams {
	return &TenancyTenantsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenants bulk partial update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableTenant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkPartialUpdateParams) WithDefaults() *TenancyTenantsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) WithContext(ctx context.Context) *TenancyTenantsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) WithData(data *models.WritableTenant) *TenancyTenantsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants bulk partial update params
func (o *TenancyTenantsBulkPartialUpdateParams) SetData(data *models.WritableTenant) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
