// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamVlanGroupsBulkPartialUpdateReader is a Reader for the IpamVlanGroupsBulkPartialUpdate structure.
type IpamVlanGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsBulkPartialUpdateOK creates a IpamVlanGroupsBulkPartialUpdateOK with default headers values
func NewIpamVlanGroupsBulkPartialUpdateOK() *IpamVlanGroupsBulkPartialUpdateOK {
	return &IpamVlanGroupsBulkPartialUpdateOK{}
}

/* IpamVlanGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkPartialUpdateOK ipam vlan groups bulk partial update o k
*/
type IpamVlanGroupsBulkPartialUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/][%d] ipamVlanGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlanGroupsBulkPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsBulkPartialUpdateDefault creates a IpamVlanGroupsBulkPartialUpdateDefault with default headers values
func NewIpamVlanGroupsBulkPartialUpdateDefault(code int) *IpamVlanGroupsBulkPartialUpdateDefault {
	return &IpamVlanGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* IpamVlanGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamVlanGroupsBulkPartialUpdateDefault ipam vlan groups bulk partial update default
*/
type IpamVlanGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups bulk partial update default response
func (o *IpamVlanGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlanGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlanGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
