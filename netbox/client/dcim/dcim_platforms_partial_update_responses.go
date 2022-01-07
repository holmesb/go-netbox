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

// DcimPlatformsPartialUpdateReader is a Reader for the DcimPlatformsPartialUpdate structure.
type DcimPlatformsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsPartialUpdateOK creates a DcimPlatformsPartialUpdateOK with default headers values
func NewDcimPlatformsPartialUpdateOK() *DcimPlatformsPartialUpdateOK {
	return &DcimPlatformsPartialUpdateOK{}
}

/* DcimPlatformsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsPartialUpdateOK dcim platforms partial update o k
*/
type DcimPlatformsPartialUpdateOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsPartialUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsPartialUpdateDefault creates a DcimPlatformsPartialUpdateDefault with default headers values
func NewDcimPlatformsPartialUpdateDefault(code int) *DcimPlatformsPartialUpdateDefault {
	return &DcimPlatformsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPlatformsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPlatformsPartialUpdateDefault dcim platforms partial update default
*/
type DcimPlatformsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms partial update default response
func (o *DcimPlatformsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPlatformsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcim_platforms_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPlatformsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
