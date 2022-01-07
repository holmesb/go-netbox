// Code generated by go-swagger; DO NOT EDIT.

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantGroupsBulkDeleteReader is a Reader for the TenancyTenantGroupsBulkDelete structure.
type TenancyTenantGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsBulkDeleteNoContent creates a TenancyTenantGroupsBulkDeleteNoContent with default headers values
func NewTenancyTenantGroupsBulkDeleteNoContent() *TenancyTenantGroupsBulkDeleteNoContent {
	return &TenancyTenantGroupsBulkDeleteNoContent{}
}

/* TenancyTenantGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantGroupsBulkDeleteNoContent tenancy tenant groups bulk delete no content
*/
type TenancyTenantGroupsBulkDeleteNoContent struct {
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyTenantGroupsBulkDeleteDefault creates a TenancyTenantGroupsBulkDeleteDefault with default headers values
func NewTenancyTenantGroupsBulkDeleteDefault(code int) *TenancyTenantGroupsBulkDeleteDefault {
	return &TenancyTenantGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* TenancyTenantGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

TenancyTenantGroupsBulkDeleteDefault tenancy tenant groups bulk delete default
*/
type TenancyTenantGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenant groups bulk delete default response
func (o *TenancyTenantGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancy_tenant-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyTenantGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
