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

// NewDcimRackRolesDeleteParams creates a new DcimRackRolesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesDeleteParams() *DcimRackRolesDeleteParams {
	return &DcimRackRolesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesDeleteParamsWithTimeout creates a new DcimRackRolesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesDeleteParamsWithTimeout(timeout time.Duration) *DcimRackRolesDeleteParams {
	return &DcimRackRolesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesDeleteParamsWithContext creates a new DcimRackRolesDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackRolesDeleteParamsWithContext(ctx context.Context) *DcimRackRolesDeleteParams {
	return &DcimRackRolesDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackRolesDeleteParamsWithHTTPClient creates a new DcimRackRolesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesDeleteParamsWithHTTPClient(client *http.Client) *DcimRackRolesDeleteParams {
	return &DcimRackRolesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackRolesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack roles delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackRolesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this rack role.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesDeleteParams) WithDefaults() *DcimRackRolesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) WithTimeout(timeout time.Duration) *DcimRackRolesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) WithContext(ctx context.Context) *DcimRackRolesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) WithHTTPClient(client *http.Client) *DcimRackRolesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) WithID(id int64) *DcimRackRolesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack roles delete params
func (o *DcimRackRolesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
