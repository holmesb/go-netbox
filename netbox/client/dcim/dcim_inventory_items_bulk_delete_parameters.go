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
)

// NewDcimInventoryItemsBulkDeleteParams creates a new DcimInventoryItemsBulkDeleteParams object
// with the default values initialized.
func NewDcimInventoryItemsBulkDeleteParams() *DcimInventoryItemsBulkDeleteParams {

	return &DcimInventoryItemsBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsBulkDeleteParamsWithTimeout creates a new DcimInventoryItemsBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInventoryItemsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsBulkDeleteParams {

	return &DcimInventoryItemsBulkDeleteParams{

		timeout: timeout,
	}
}

// NewDcimInventoryItemsBulkDeleteParamsWithContext creates a new DcimInventoryItemsBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInventoryItemsBulkDeleteParamsWithContext(ctx context.Context) *DcimInventoryItemsBulkDeleteParams {

	return &DcimInventoryItemsBulkDeleteParams{

		Context: ctx,
	}
}

// NewDcimInventoryItemsBulkDeleteParamsWithHTTPClient creates a new DcimInventoryItemsBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInventoryItemsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsBulkDeleteParams {

	return &DcimInventoryItemsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*DcimInventoryItemsBulkDeleteParams contains all the parameters to send to the API endpoint
for the dcim inventory items bulk delete operation typically these are written to a http.Request
*/
type DcimInventoryItemsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) WithContext(ctx context.Context) *DcimInventoryItemsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items bulk delete params
func (o *DcimInventoryItemsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}