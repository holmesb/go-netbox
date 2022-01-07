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

// TenancyTenantGroupsUpdateReader is a Reader for the TenancyTenantGroupsUpdate structure.
type TenancyTenantGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsUpdateOK creates a TenancyTenantGroupsUpdateOK with default headers values
func NewTenancyTenantGroupsUpdateOK() *TenancyTenantGroupsUpdateOK {
	return &TenancyTenantGroupsUpdateOK{}
}

/* TenancyTenantGroupsUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsUpdateOK tenancy tenant groups update o k
*/
type TenancyTenantGroupsUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantGroupsUpdateDefault creates a TenancyTenantGroupsUpdateDefault with default headers values
func NewTenancyTenantGroupsUpdateDefault(code int) *TenancyTenantGroupsUpdateDefault {
	return &TenancyTenantGroupsUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyTenantGroupsUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantGroupsUpdateDefault tenancy tenant groups update default
*/
type TenancyTenantGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenant groups update default response
func (o *TenancyTenantGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/{id}/][%d] tenancy_tenant-groups_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyTenantGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
