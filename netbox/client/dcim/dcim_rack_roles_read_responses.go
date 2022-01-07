// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRackRolesReadReader is a Reader for the DcimRackRolesRead structure.
type DcimRackRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesReadOK creates a DcimRackRolesReadOK with default headers values
func NewDcimRackRolesReadOK() *DcimRackRolesReadOK {
	return &DcimRackRolesReadOK{}
}

/* DcimRackRolesReadOK describes a response with status code 200, with default header values.

DcimRackRolesReadOK dcim rack roles read o k
*/
type DcimRackRolesReadOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcimRackRolesReadOK  %+v", 200, o.Payload)
}
func (o *DcimRackRolesReadOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesReadDefault creates a DcimRackRolesReadDefault with default headers values
func NewDcimRackRolesReadDefault(code int) *DcimRackRolesReadDefault {
	return &DcimRackRolesReadDefault{
		_statusCode: code,
	}
}

/* DcimRackRolesReadDefault describes a response with status code -1, with default header values.

DcimRackRolesReadDefault dcim rack roles read default
*/
type DcimRackRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack roles read default response
func (o *DcimRackRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcim_rack-roles_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRackRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
