// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitsPartialUpdateReader is a Reader for the CircuitsCircuitsPartialUpdate structure.
type CircuitsCircuitsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsPartialUpdateOK creates a CircuitsCircuitsPartialUpdateOK with default headers values
func NewCircuitsCircuitsPartialUpdateOK() *CircuitsCircuitsPartialUpdateOK {
	return &CircuitsCircuitsPartialUpdateOK{}
}

/* CircuitsCircuitsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsPartialUpdateOK circuits circuits partial update o k
*/
type CircuitsCircuitsPartialUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/{id}/][%d] circuitsCircuitsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsPartialUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsPartialUpdateDefault creates a CircuitsCircuitsPartialUpdateDefault with default headers values
func NewCircuitsCircuitsPartialUpdateDefault(code int) *CircuitsCircuitsPartialUpdateDefault {
	return &CircuitsCircuitsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitsPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitsPartialUpdateDefault circuits circuits partial update default
*/
type CircuitsCircuitsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuits partial update default response
func (o *CircuitsCircuitsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/{id}/][%d] circuits_circuits_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
