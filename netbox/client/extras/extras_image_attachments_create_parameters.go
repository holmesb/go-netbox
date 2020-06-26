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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewExtrasImageAttachmentsCreateParams creates a new ExtrasImageAttachmentsCreateParams object
// with the default values initialized.
func NewExtrasImageAttachmentsCreateParams() *ExtrasImageAttachmentsCreateParams {
	var ()
	return &ExtrasImageAttachmentsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithTimeout creates a new ExtrasImageAttachmentsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasImageAttachmentsCreateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsCreateParams {
	var ()
	return &ExtrasImageAttachmentsCreateParams{

		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithContext creates a new ExtrasImageAttachmentsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasImageAttachmentsCreateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsCreateParams {
	var ()
	return &ExtrasImageAttachmentsCreateParams{

		Context: ctx,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithHTTPClient creates a new ExtrasImageAttachmentsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasImageAttachmentsCreateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsCreateParams {
	var ()
	return &ExtrasImageAttachmentsCreateParams{
		HTTPClient: client,
	}
}

/*ExtrasImageAttachmentsCreateParams contains all the parameters to send to the API endpoint
for the extras image attachments create operation typically these are written to a http.Request
*/
type ExtrasImageAttachmentsCreateParams struct {

	/*Data*/
	Data *models.ImageAttachment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
