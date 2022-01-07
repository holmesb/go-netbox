// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRearPortsDeleteReader is a Reader for the DcimRearPortsDelete structure.
type DcimRearPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsDeleteNoContent creates a DcimRearPortsDeleteNoContent with default headers values
func NewDcimRearPortsDeleteNoContent() *DcimRearPortsDeleteNoContent {
	return &DcimRearPortsDeleteNoContent{}
}

/* DcimRearPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortsDeleteNoContent dcim rear ports delete no content
*/
type DcimRearPortsDeleteNoContent struct {
}

func (o *DcimRearPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-ports/{id}/][%d] dcimRearPortsDeleteNoContent ", 204)
}

func (o *DcimRearPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRearPortsDeleteDefault creates a DcimRearPortsDeleteDefault with default headers values
func NewDcimRearPortsDeleteDefault(code int) *DcimRearPortsDeleteDefault {
	return &DcimRearPortsDeleteDefault{
		_statusCode: code,
	}
}

/* DcimRearPortsDeleteDefault describes a response with status code -1, with default header values.

DcimRearPortsDeleteDefault dcim rear ports delete default
*/
type DcimRearPortsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear ports delete default response
func (o *DcimRearPortsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-ports/{id}/][%d] dcim_rear-ports_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRearPortsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
