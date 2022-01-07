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

// IpamServicesBulkDeleteReader is a Reader for the IpamServicesBulkDelete structure.
type IpamServicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamServicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesBulkDeleteNoContent creates a IpamServicesBulkDeleteNoContent with default headers values
func NewIpamServicesBulkDeleteNoContent() *IpamServicesBulkDeleteNoContent {
	return &IpamServicesBulkDeleteNoContent{}
}

/* IpamServicesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamServicesBulkDeleteNoContent ipam services bulk delete no content
*/
type IpamServicesBulkDeleteNoContent struct {
}

func (o *IpamServicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/services/][%d] ipamServicesBulkDeleteNoContent ", 204)
}

func (o *IpamServicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamServicesBulkDeleteDefault creates a IpamServicesBulkDeleteDefault with default headers values
func NewIpamServicesBulkDeleteDefault(code int) *IpamServicesBulkDeleteDefault {
	return &IpamServicesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamServicesBulkDeleteDefault describes a response with status code -1, with default header values.

IpamServicesBulkDeleteDefault ipam services bulk delete default
*/
type IpamServicesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam services bulk delete default response
func (o *IpamServicesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/services/][%d] ipam_services_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamServicesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
