// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// NewCircuitsProviderNetworksBulkPartialUpdateParams creates a new CircuitsProviderNetworksBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProviderNetworksBulkPartialUpdateParams() *CircuitsProviderNetworksBulkPartialUpdateParams {
	return &CircuitsProviderNetworksBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProviderNetworksBulkPartialUpdateParamsWithTimeout creates a new CircuitsProviderNetworksBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProviderNetworksBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProviderNetworksBulkPartialUpdateParams {
	return &CircuitsProviderNetworksBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProviderNetworksBulkPartialUpdateParamsWithContext creates a new CircuitsProviderNetworksBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProviderNetworksBulkPartialUpdateParamsWithContext(ctx context.Context) *CircuitsProviderNetworksBulkPartialUpdateParams {
	return &CircuitsProviderNetworksBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProviderNetworksBulkPartialUpdateParamsWithHTTPClient creates a new CircuitsProviderNetworksBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProviderNetworksBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProviderNetworksBulkPartialUpdateParams {
	return &CircuitsProviderNetworksBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsProviderNetworksBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits provider networks bulk partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsProviderNetworksBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableProviderNetwork

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits provider networks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WithDefaults() *CircuitsProviderNetworksBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits provider networks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProviderNetworksBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WithContext(ctx context.Context) *CircuitsProviderNetworksBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProviderNetworksBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WithData(data *models.WritableProviderNetwork) *CircuitsProviderNetworksBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits provider networks bulk partial update params
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) SetData(data *models.WritableProviderNetwork) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProviderNetworksBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
