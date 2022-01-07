// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlansBulkDeleteReader is a Reader for the IpamVlansBulkDelete structure.
type IpamVlansBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlansBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansBulkDeleteNoContent creates a IpamVlansBulkDeleteNoContent with default headers values
func NewIpamVlansBulkDeleteNoContent() *IpamVlansBulkDeleteNoContent {
	return &IpamVlansBulkDeleteNoContent{}
}

/* IpamVlansBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVlansBulkDeleteNoContent ipam vlans bulk delete no content
*/
type IpamVlansBulkDeleteNoContent struct {
}

func (o *IpamVlansBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/][%d] ipamVlansBulkDeleteNoContent ", 204)
}

func (o *IpamVlansBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVlansBulkDeleteDefault creates a IpamVlansBulkDeleteDefault with default headers values
func NewIpamVlansBulkDeleteDefault(code int) *IpamVlansBulkDeleteDefault {
	return &IpamVlansBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamVlansBulkDeleteDefault describes a response with status code -1, with default header values.

IpamVlansBulkDeleteDefault ipam vlans bulk delete default
*/
type IpamVlansBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlans bulk delete default response
func (o *IpamVlansBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/][%d] ipam_vlans_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlansBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
