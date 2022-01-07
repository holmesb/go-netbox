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

// DcimFrontPortTemplatesUpdateReader is a Reader for the DcimFrontPortTemplatesUpdate structure.
type DcimFrontPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesUpdateOK creates a DcimFrontPortTemplatesUpdateOK with default headers values
func NewDcimFrontPortTemplatesUpdateOK() *DcimFrontPortTemplatesUpdateOK {
	return &DcimFrontPortTemplatesUpdateOK{}
}

/* DcimFrontPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesUpdateOK dcim front port templates update o k
*/
type DcimFrontPortTemplatesUpdateOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesUpdateDefault creates a DcimFrontPortTemplatesUpdateDefault with default headers values
func NewDcimFrontPortTemplatesUpdateDefault(code int) *DcimFrontPortTemplatesUpdateDefault {
	return &DcimFrontPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortTemplatesUpdateDefault dcim front port templates update default
*/
type DcimFrontPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates update default response
func (o *DcimFrontPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
