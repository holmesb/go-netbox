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

// ExtrasTagsUpdateReader is a Reader for the ExtrasTagsUpdate structure.
type ExtrasTagsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsUpdateOK creates a ExtrasTagsUpdateOK with default headers values
func NewExtrasTagsUpdateOK() *ExtrasTagsUpdateOK {
	return &ExtrasTagsUpdateOK{}
}

/* ExtrasTagsUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsUpdateOK extras tags update o k
*/
type ExtrasTagsUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extrasTagsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsUpdateDefault creates a ExtrasTagsUpdateDefault with default headers values
func NewExtrasTagsUpdateDefault(code int) *ExtrasTagsUpdateDefault {
	return &ExtrasTagsUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasTagsUpdateDefault describes a response with status code -1, with default header values.

ExtrasTagsUpdateDefault extras tags update default
*/
type ExtrasTagsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags update default response
func (o *ExtrasTagsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasTagsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/{id}/][%d] extras_tags_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasTagsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
