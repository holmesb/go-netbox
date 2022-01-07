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

// DcimConsoleServerPortTemplatesBulkDeleteReader is a Reader for the DcimConsoleServerPortTemplatesBulkDelete structure.
type DcimConsoleServerPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesBulkDeleteNoContent creates a DcimConsoleServerPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimConsoleServerPortTemplatesBulkDeleteNoContent() *DcimConsoleServerPortTemplatesBulkDeleteNoContent {
	return &DcimConsoleServerPortTemplatesBulkDeleteNoContent{}
}

/* DcimConsoleServerPortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortTemplatesBulkDeleteNoContent dcim console server port templates bulk delete no content
*/
type DcimConsoleServerPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsoleServerPortTemplatesBulkDeleteDefault creates a DcimConsoleServerPortTemplatesBulkDeleteDefault with default headers values
func NewDcimConsoleServerPortTemplatesBulkDeleteDefault(code int) *DcimConsoleServerPortTemplatesBulkDeleteDefault {
	return &DcimConsoleServerPortTemplatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortTemplatesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesBulkDeleteDefault dcim console server port templates bulk delete default
*/
type DcimConsoleServerPortTemplatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates bulk delete default response
func (o *DcimConsoleServerPortTemplatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/][%d] dcim_console-server-port-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
