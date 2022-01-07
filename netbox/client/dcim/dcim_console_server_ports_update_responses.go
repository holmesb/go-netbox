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

// DcimConsoleServerPortsUpdateReader is a Reader for the DcimConsoleServerPortsUpdate structure.
type DcimConsoleServerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsUpdateOK creates a DcimConsoleServerPortsUpdateOK with default headers values
func NewDcimConsoleServerPortsUpdateOK() *DcimConsoleServerPortsUpdateOK {
	return &DcimConsoleServerPortsUpdateOK{}
}

/* DcimConsoleServerPortsUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsUpdateOK dcim console server ports update o k
*/
type DcimConsoleServerPortsUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortsUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsUpdateDefault creates a DcimConsoleServerPortsUpdateDefault with default headers values
func NewDcimConsoleServerPortsUpdateDefault(code int) *DcimConsoleServerPortsUpdateDefault {
	return &DcimConsoleServerPortsUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortsUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsUpdateDefault dcim console server ports update default
*/
type DcimConsoleServerPortsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports update default response
func (o *DcimConsoleServerPortsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
