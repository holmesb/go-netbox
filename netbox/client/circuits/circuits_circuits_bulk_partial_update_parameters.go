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

// NewCircuitsCircuitsBulkPartialUpdateParams creates a new CircuitsCircuitsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitsBulkPartialUpdateParams() *CircuitsCircuitsBulkPartialUpdateParams {
	return &CircuitsCircuitsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsBulkPartialUpdateParamsWithTimeout creates a new CircuitsCircuitsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsBulkPartialUpdateParams {
	return &CircuitsCircuitsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitsBulkPartialUpdateParamsWithContext creates a new CircuitsCircuitsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitsBulkPartialUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitsBulkPartialUpdateParams {
	return &CircuitsCircuitsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitsBulkPartialUpdateParamsWithHTTPClient creates a new CircuitsCircuitsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsBulkPartialUpdateParams {
	return &CircuitsCircuitsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuits bulk partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableCircuit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuits bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsBulkPartialUpdateParams) WithDefaults() *CircuitsCircuitsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuits bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) WithData(data *models.WritableCircuit) *CircuitsCircuitsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuits bulk partial update params
func (o *CircuitsCircuitsBulkPartialUpdateParams) SetData(data *models.WritableCircuit) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
