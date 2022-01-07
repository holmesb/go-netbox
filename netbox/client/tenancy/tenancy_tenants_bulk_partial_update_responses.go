// Code generated by go-swagger; DO NOT EDIT.

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyTenantsBulkPartialUpdateReader is a Reader for the TenancyTenantsBulkPartialUpdate structure.
type TenancyTenantsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsBulkPartialUpdateOK creates a TenancyTenantsBulkPartialUpdateOK with default headers values
func NewTenancyTenantsBulkPartialUpdateOK() *TenancyTenantsBulkPartialUpdateOK {
	return &TenancyTenantsBulkPartialUpdateOK{}
}

/* TenancyTenantsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsBulkPartialUpdateOK tenancy tenants bulk partial update o k
*/
type TenancyTenantsBulkPartialUpdateOK struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/][%d] tenancyTenantsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantsBulkPartialUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsBulkPartialUpdateDefault creates a TenancyTenantsBulkPartialUpdateDefault with default headers values
func NewTenancyTenantsBulkPartialUpdateDefault(code int) *TenancyTenantsBulkPartialUpdateDefault {
	return &TenancyTenantsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyTenantsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantsBulkPartialUpdateDefault tenancy tenants bulk partial update default
*/
type TenancyTenantsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenants bulk partial update default response
func (o *TenancyTenantsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/][%d] tenancy_tenants_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyTenantsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
