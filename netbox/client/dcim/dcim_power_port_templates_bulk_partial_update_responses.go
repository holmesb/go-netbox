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

// DcimPowerPortTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerPortTemplatesBulkPartialUpdate structure.
type DcimPowerPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateOK creates a DcimPowerPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesBulkPartialUpdateOK() *DcimPowerPortTemplatesBulkPartialUpdateOK {
	return &DcimPowerPortTemplatesBulkPartialUpdateOK{}
}

/* DcimPowerPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesBulkPartialUpdateOK dcim power port templates bulk partial update o k
*/
type DcimPowerPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesBulkPartialUpdateDefault creates a DcimPowerPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimPowerPortTemplatesBulkPartialUpdateDefault(code int) *DcimPowerPortTemplatesBulkPartialUpdateDefault {
	return &DcimPowerPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPowerPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesBulkPartialUpdateDefault dcim power port templates bulk partial update default
*/
type DcimPowerPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power port templates bulk partial update default response
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/][%d] dcim_power-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
