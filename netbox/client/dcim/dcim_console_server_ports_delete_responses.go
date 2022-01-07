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

// DcimConsoleServerPortsDeleteReader is a Reader for the DcimConsoleServerPortsDelete structure.
type DcimConsoleServerPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsDeleteNoContent creates a DcimConsoleServerPortsDeleteNoContent with default headers values
func NewDcimConsoleServerPortsDeleteNoContent() *DcimConsoleServerPortsDeleteNoContent {
	return &DcimConsoleServerPortsDeleteNoContent{}
}

/* DcimConsoleServerPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortsDeleteNoContent dcim console server ports delete no content
*/
type DcimConsoleServerPortsDeleteNoContent struct {
}

func (o *DcimConsoleServerPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsoleServerPortsDeleteDefault creates a DcimConsoleServerPortsDeleteDefault with default headers values
func NewDcimConsoleServerPortsDeleteDefault(code int) *DcimConsoleServerPortsDeleteDefault {
	return &DcimConsoleServerPortsDeleteDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortsDeleteDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsDeleteDefault dcim console server ports delete default
*/
type DcimConsoleServerPortsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports delete default response
func (o *DcimConsoleServerPortsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
