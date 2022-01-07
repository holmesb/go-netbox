// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIpamIPRangesUpdateParams creates a new IpamIPRangesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPRangesUpdateParams() *IpamIPRangesUpdateParams {
	return &IpamIPRangesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPRangesUpdateParamsWithTimeout creates a new IpamIPRangesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamIPRangesUpdateParamsWithTimeout(timeout time.Duration) *IpamIPRangesUpdateParams {
	return &IpamIPRangesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamIPRangesUpdateParamsWithContext creates a new IpamIPRangesUpdateParams object
// with the ability to set a context for a request.
func NewIpamIPRangesUpdateParamsWithContext(ctx context.Context) *IpamIPRangesUpdateParams {
	return &IpamIPRangesUpdateParams{
		Context: ctx,
	}
}

// NewIpamIPRangesUpdateParamsWithHTTPClient creates a new IpamIPRangesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPRangesUpdateParamsWithHTTPClient(client *http.Client) *IpamIPRangesUpdateParams {
	return &IpamIPRangesUpdateParams{
		HTTPClient: client,
	}
}

/* IpamIPRangesUpdateParams contains all the parameters to send to the API endpoint
   for the ipam ip ranges update operation.

   Typically these are written to a http.Request.
*/
type IpamIPRangesUpdateParams struct {

	// Data.
	Data *models.WritableIPRange

	/* ID.

	   A unique integer value identifying this IP range.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip ranges update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesUpdateParams) WithDefaults() *IpamIPRangesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip ranges update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) WithTimeout(timeout time.Duration) *IpamIPRangesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) WithContext(ctx context.Context) *IpamIPRangesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) WithHTTPClient(client *http.Client) *IpamIPRangesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) WithData(data *models.WritableIPRange) *IpamIPRangesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) SetData(data *models.WritableIPRange) {
	o.Data = data
}

// WithID adds the id to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) WithID(id int64) *IpamIPRangesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip ranges update params
func (o *IpamIPRangesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPRangesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
