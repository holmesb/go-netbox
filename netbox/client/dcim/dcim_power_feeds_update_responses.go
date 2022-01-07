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

// DcimPowerFeedsUpdateReader is a Reader for the DcimPowerFeedsUpdate structure.
type DcimPowerFeedsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsUpdateOK creates a DcimPowerFeedsUpdateOK with default headers values
func NewDcimPowerFeedsUpdateOK() *DcimPowerFeedsUpdateOK {
	return &DcimPowerFeedsUpdateOK{}
}

/* DcimPowerFeedsUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsUpdateOK dcim power feeds update o k
*/
type DcimPowerFeedsUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/{id}/][%d] dcimPowerFeedsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsUpdateDefault creates a DcimPowerFeedsUpdateDefault with default headers values
func NewDcimPowerFeedsUpdateDefault(code int) *DcimPowerFeedsUpdateDefault {
	return &DcimPowerFeedsUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPowerFeedsUpdateDefault describes a response with status code -1, with default header values.

DcimPowerFeedsUpdateDefault dcim power feeds update default
*/
type DcimPowerFeedsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds update default response
func (o *DcimPowerFeedsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/{id}/][%d] dcim_power-feeds_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerFeedsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
