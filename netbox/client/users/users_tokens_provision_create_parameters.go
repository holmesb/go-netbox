// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersTokensProvisionCreateParams creates a new UsersTokensProvisionCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersTokensProvisionCreateParams() *UsersTokensProvisionCreateParams {
	return &UsersTokensProvisionCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersTokensProvisionCreateParamsWithTimeout creates a new UsersTokensProvisionCreateParams object
// with the ability to set a timeout on a request.
func NewUsersTokensProvisionCreateParamsWithTimeout(timeout time.Duration) *UsersTokensProvisionCreateParams {
	return &UsersTokensProvisionCreateParams{
		timeout: timeout,
	}
}

// NewUsersTokensProvisionCreateParamsWithContext creates a new UsersTokensProvisionCreateParams object
// with the ability to set a context for a request.
func NewUsersTokensProvisionCreateParamsWithContext(ctx context.Context) *UsersTokensProvisionCreateParams {
	return &UsersTokensProvisionCreateParams{
		Context: ctx,
	}
}

// NewUsersTokensProvisionCreateParamsWithHTTPClient creates a new UsersTokensProvisionCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersTokensProvisionCreateParamsWithHTTPClient(client *http.Client) *UsersTokensProvisionCreateParams {
	return &UsersTokensProvisionCreateParams{
		HTTPClient: client,
	}
}

/* UsersTokensProvisionCreateParams contains all the parameters to send to the API endpoint
   for the users tokens provision create operation.

   Typically these are written to a http.Request.
*/
type UsersTokensProvisionCreateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users tokens provision create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersTokensProvisionCreateParams) WithDefaults() *UsersTokensProvisionCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users tokens provision create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersTokensProvisionCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) WithTimeout(timeout time.Duration) *UsersTokensProvisionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) WithContext(ctx context.Context) *UsersTokensProvisionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) WithHTTPClient(client *http.Client) *UsersTokensProvisionCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users tokens provision create params
func (o *UsersTokensProvisionCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UsersTokensProvisionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
