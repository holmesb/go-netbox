// Code generated by go-swagger; DO NOT EDIT.

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
)

// NewDcimLocationsDeleteParams creates a new DcimLocationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimLocationsDeleteParams() *DcimLocationsDeleteParams {
	return &DcimLocationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimLocationsDeleteParamsWithTimeout creates a new DcimLocationsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimLocationsDeleteParamsWithTimeout(timeout time.Duration) *DcimLocationsDeleteParams {
	return &DcimLocationsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimLocationsDeleteParamsWithContext creates a new DcimLocationsDeleteParams object
// with the ability to set a context for a request.
func NewDcimLocationsDeleteParamsWithContext(ctx context.Context) *DcimLocationsDeleteParams {
	return &DcimLocationsDeleteParams{
		Context: ctx,
	}
}

// NewDcimLocationsDeleteParamsWithHTTPClient creates a new DcimLocationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimLocationsDeleteParamsWithHTTPClient(client *http.Client) *DcimLocationsDeleteParams {
	return &DcimLocationsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimLocationsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim locations delete operation.

   Typically these are written to a http.Request.
*/
type DcimLocationsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this location.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim locations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimLocationsDeleteParams) WithDefaults() *DcimLocationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim locations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimLocationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim locations delete params
func (o *DcimLocationsDeleteParams) WithTimeout(timeout time.Duration) *DcimLocationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim locations delete params
func (o *DcimLocationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim locations delete params
func (o *DcimLocationsDeleteParams) WithContext(ctx context.Context) *DcimLocationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim locations delete params
func (o *DcimLocationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim locations delete params
func (o *DcimLocationsDeleteParams) WithHTTPClient(client *http.Client) *DcimLocationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim locations delete params
func (o *DcimLocationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim locations delete params
func (o *DcimLocationsDeleteParams) WithID(id int64) *DcimLocationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim locations delete params
func (o *DcimLocationsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimLocationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
