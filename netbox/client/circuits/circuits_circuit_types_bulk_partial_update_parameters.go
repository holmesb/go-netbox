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

// NewCircuitsCircuitTypesBulkPartialUpdateParams creates a new CircuitsCircuitTypesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesBulkPartialUpdateParams() *CircuitsCircuitTypesBulkPartialUpdateParams {
	return &CircuitsCircuitTypesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesBulkPartialUpdateParamsWithTimeout creates a new CircuitsCircuitTypesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesBulkPartialUpdateParams {
	return &CircuitsCircuitTypesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesBulkPartialUpdateParamsWithContext creates a new CircuitsCircuitTypesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesBulkPartialUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesBulkPartialUpdateParams {
	return &CircuitsCircuitTypesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesBulkPartialUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesBulkPartialUpdateParams {
	return &CircuitsCircuitTypesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTypesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuit types bulk partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesBulkPartialUpdateParams struct {

	// Data.
	Data *models.CircuitType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WithDefaults() *CircuitsCircuitTypesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types bulk partial update params
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
