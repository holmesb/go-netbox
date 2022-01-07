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

// DcimRacksReadReader is a Reader for the DcimRacksRead structure.
type DcimRacksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksReadOK creates a DcimRacksReadOK with default headers values
func NewDcimRacksReadOK() *DcimRacksReadOK {
	return &DcimRacksReadOK{}
}

/* DcimRacksReadOK describes a response with status code 200, with default header values.

DcimRacksReadOK dcim racks read o k
*/
type DcimRacksReadOK struct {
	Payload *models.Rack
}

func (o *DcimRacksReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcimRacksReadOK  %+v", 200, o.Payload)
}
func (o *DcimRacksReadOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksReadDefault creates a DcimRacksReadDefault with default headers values
func NewDcimRacksReadDefault(code int) *DcimRacksReadDefault {
	return &DcimRacksReadDefault{
		_statusCode: code,
	}
}

/* DcimRacksReadDefault describes a response with status code -1, with default header values.

DcimRacksReadDefault dcim racks read default
*/
type DcimRacksReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks read default response
func (o *DcimRacksReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRacksReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcim_racks_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRacksReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
