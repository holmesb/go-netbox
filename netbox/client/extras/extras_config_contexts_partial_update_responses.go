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

// ExtrasConfigContextsPartialUpdateReader is a Reader for the ExtrasConfigContextsPartialUpdate structure.
type ExtrasConfigContextsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsPartialUpdateOK creates a ExtrasConfigContextsPartialUpdateOK with default headers values
func NewExtrasConfigContextsPartialUpdateOK() *ExtrasConfigContextsPartialUpdateOK {
	return &ExtrasConfigContextsPartialUpdateOK{}
}

/* ExtrasConfigContextsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsPartialUpdateOK extras config contexts partial update o k
*/
type ExtrasConfigContextsPartialUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/{id}/][%d] extrasConfigContextsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsPartialUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsPartialUpdateDefault creates a ExtrasConfigContextsPartialUpdateDefault with default headers values
func NewExtrasConfigContextsPartialUpdateDefault(code int) *ExtrasConfigContextsPartialUpdateDefault {
	return &ExtrasConfigContextsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasConfigContextsPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsPartialUpdateDefault extras config contexts partial update default
*/
type ExtrasConfigContextsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts partial update default response
func (o *ExtrasConfigContextsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigContextsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/{id}/][%d] extras_config-contexts_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasConfigContextsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
