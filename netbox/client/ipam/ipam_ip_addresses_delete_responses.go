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

// IpamIPAddressesDeleteReader is a Reader for the IpamIPAddressesDelete structure.
type IpamIPAddressesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPAddressesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesDeleteNoContent creates a IpamIPAddressesDeleteNoContent with default headers values
func NewIpamIPAddressesDeleteNoContent() *IpamIPAddressesDeleteNoContent {
	return &IpamIPAddressesDeleteNoContent{}
}

/* IpamIPAddressesDeleteNoContent describes a response with status code 204, with default header values.

IpamIPAddressesDeleteNoContent ipam Ip addresses delete no content
*/
type IpamIPAddressesDeleteNoContent struct {
}

func (o *IpamIPAddressesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/{id}/][%d] ipamIpAddressesDeleteNoContent ", 204)
}

func (o *IpamIPAddressesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamIPAddressesDeleteDefault creates a IpamIPAddressesDeleteDefault with default headers values
func NewIpamIPAddressesDeleteDefault(code int) *IpamIPAddressesDeleteDefault {
	return &IpamIPAddressesDeleteDefault{
		_statusCode: code,
	}
}

/* IpamIPAddressesDeleteDefault describes a response with status code -1, with default header values.

IpamIPAddressesDeleteDefault ipam ip addresses delete default
*/
type IpamIPAddressesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses delete default response
func (o *IpamIPAddressesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPAddressesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamIPAddressesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
