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

// DcimFrontPortsPathsReader is a Reader for the DcimFrontPortsPaths structure.
type DcimFrontPortsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsPathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsPathsOK creates a DcimFrontPortsPathsOK with default headers values
func NewDcimFrontPortsPathsOK() *DcimFrontPortsPathsOK {
	return &DcimFrontPortsPathsOK{}
}

/* DcimFrontPortsPathsOK describes a response with status code 200, with default header values.

DcimFrontPortsPathsOK dcim front ports paths o k
*/
type DcimFrontPortsPathsOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsPathsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcimFrontPortsPathsOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsPathsOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsPathsDefault creates a DcimFrontPortsPathsDefault with default headers values
func NewDcimFrontPortsPathsDefault(code int) *DcimFrontPortsPathsDefault {
	return &DcimFrontPortsPathsDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortsPathsDefault describes a response with status code -1, with default header values.

DcimFrontPortsPathsDefault dcim front ports paths default
*/
type DcimFrontPortsPathsDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports paths default response
func (o *DcimFrontPortsPathsDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsPathsDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcim_front-ports_paths default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortsPathsDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsPathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
