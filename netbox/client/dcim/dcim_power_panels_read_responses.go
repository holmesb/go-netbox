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

// DcimPowerPanelsReadReader is a Reader for the DcimPowerPanelsRead structure.
type DcimPowerPanelsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsReadOK creates a DcimPowerPanelsReadOK with default headers values
func NewDcimPowerPanelsReadOK() *DcimPowerPanelsReadOK {
	return &DcimPowerPanelsReadOK{}
}

/* DcimPowerPanelsReadOK describes a response with status code 200, with default header values.

DcimPowerPanelsReadOK dcim power panels read o k
*/
type DcimPowerPanelsReadOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcimPowerPanelsReadOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPanelsReadOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsReadDefault creates a DcimPowerPanelsReadDefault with default headers values
func NewDcimPowerPanelsReadDefault(code int) *DcimPowerPanelsReadDefault {
	return &DcimPowerPanelsReadDefault{
		_statusCode: code,
	}
}

/* DcimPowerPanelsReadDefault describes a response with status code -1, with default header values.

DcimPowerPanelsReadDefault dcim power panels read default
*/
type DcimPowerPanelsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels read default response
func (o *DcimPowerPanelsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPanelsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcim_power-panels_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerPanelsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
