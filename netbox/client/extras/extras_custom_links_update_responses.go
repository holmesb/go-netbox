// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomLinksUpdateReader is a Reader for the ExtrasCustomLinksUpdate structure.
type ExtrasCustomLinksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksUpdateOK creates a ExtrasCustomLinksUpdateOK with default headers values
func NewExtrasCustomLinksUpdateOK() *ExtrasCustomLinksUpdateOK {
	return &ExtrasCustomLinksUpdateOK{}
}

/* ExtrasCustomLinksUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksUpdateOK extras custom links update o k
*/
type ExtrasCustomLinksUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/{id}/][%d] extrasCustomLinksUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksUpdateDefault creates a ExtrasCustomLinksUpdateDefault with default headers values
func NewExtrasCustomLinksUpdateDefault(code int) *ExtrasCustomLinksUpdateDefault {
	return &ExtrasCustomLinksUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasCustomLinksUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksUpdateDefault extras custom links update default
*/
type ExtrasCustomLinksUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links update default response
func (o *ExtrasCustomLinksUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomLinksUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/{id}/][%d] extras_custom-links_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasCustomLinksUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
