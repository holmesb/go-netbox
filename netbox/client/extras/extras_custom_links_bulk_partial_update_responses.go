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

// ExtrasCustomLinksBulkPartialUpdateReader is a Reader for the ExtrasCustomLinksBulkPartialUpdate structure.
type ExtrasCustomLinksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksBulkPartialUpdateOK creates a ExtrasCustomLinksBulkPartialUpdateOK with default headers values
func NewExtrasCustomLinksBulkPartialUpdateOK() *ExtrasCustomLinksBulkPartialUpdateOK {
	return &ExtrasCustomLinksBulkPartialUpdateOK{}
}

/* ExtrasCustomLinksBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksBulkPartialUpdateOK extras custom links bulk partial update o k
*/
type ExtrasCustomLinksBulkPartialUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-links/][%d] extrasCustomLinksBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksBulkPartialUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksBulkPartialUpdateDefault creates a ExtrasCustomLinksBulkPartialUpdateDefault with default headers values
func NewExtrasCustomLinksBulkPartialUpdateDefault(code int) *ExtrasCustomLinksBulkPartialUpdateDefault {
	return &ExtrasCustomLinksBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasCustomLinksBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksBulkPartialUpdateDefault extras custom links bulk partial update default
*/
type ExtrasCustomLinksBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links bulk partial update default response
func (o *ExtrasCustomLinksBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomLinksBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-links/][%d] extras_custom-links_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasCustomLinksBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
