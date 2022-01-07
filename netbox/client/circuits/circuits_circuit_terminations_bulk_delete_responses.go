// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitTerminationsBulkDeleteReader is a Reader for the CircuitsCircuitTerminationsBulkDelete structure.
type CircuitsCircuitTerminationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitTerminationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsBulkDeleteNoContent creates a CircuitsCircuitTerminationsBulkDeleteNoContent with default headers values
func NewCircuitsCircuitTerminationsBulkDeleteNoContent() *CircuitsCircuitTerminationsBulkDeleteNoContent {
	return &CircuitsCircuitTerminationsBulkDeleteNoContent{}
}

/* CircuitsCircuitTerminationsBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitTerminationsBulkDeleteNoContent circuits circuit terminations bulk delete no content
*/
type CircuitsCircuitTerminationsBulkDeleteNoContent struct {
}

func (o *CircuitsCircuitTerminationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTerminationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCircuitsCircuitTerminationsBulkDeleteDefault creates a CircuitsCircuitTerminationsBulkDeleteDefault with default headers values
func NewCircuitsCircuitTerminationsBulkDeleteDefault(code int) *CircuitsCircuitTerminationsBulkDeleteDefault {
	return &CircuitsCircuitTerminationsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitTerminationsBulkDeleteDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsBulkDeleteDefault circuits circuit terminations bulk delete default
*/
type CircuitsCircuitTerminationsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations bulk delete default response
func (o *CircuitsCircuitTerminationsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-terminations/][%d] circuits_circuit-terminations_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
