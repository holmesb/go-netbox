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

// DcimRearPortsCreateReader is a Reader for the DcimRearPortsCreate structure.
type DcimRearPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsCreateCreated creates a DcimRearPortsCreateCreated with default headers values
func NewDcimRearPortsCreateCreated() *DcimRearPortsCreateCreated {
	return &DcimRearPortsCreateCreated{}
}

/* DcimRearPortsCreateCreated describes a response with status code 201, with default header values.

DcimRearPortsCreateCreated dcim rear ports create created
*/
type DcimRearPortsCreateCreated struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcimRearPortsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRearPortsCreateCreated) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsCreateDefault creates a DcimRearPortsCreateDefault with default headers values
func NewDcimRearPortsCreateDefault(code int) *DcimRearPortsCreateDefault {
	return &DcimRearPortsCreateDefault{
		_statusCode: code,
	}
}

/* DcimRearPortsCreateDefault describes a response with status code -1, with default header values.

DcimRearPortsCreateDefault dcim rear ports create default
*/
type DcimRearPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear ports create default response
func (o *DcimRearPortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcim_rear-ports_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRearPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
