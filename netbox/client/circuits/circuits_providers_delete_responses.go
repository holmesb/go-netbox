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

// CircuitsProvidersDeleteReader is a Reader for the CircuitsProvidersDelete structure.
type CircuitsProvidersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsProvidersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersDeleteNoContent creates a CircuitsProvidersDeleteNoContent with default headers values
func NewCircuitsProvidersDeleteNoContent() *CircuitsProvidersDeleteNoContent {
	return &CircuitsProvidersDeleteNoContent{}
}

/* CircuitsProvidersDeleteNoContent describes a response with status code 204, with default header values.

CircuitsProvidersDeleteNoContent circuits providers delete no content
*/
type CircuitsProvidersDeleteNoContent struct {
}

func (o *CircuitsProvidersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/{id}/][%d] circuitsProvidersDeleteNoContent ", 204)
}

func (o *CircuitsProvidersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCircuitsProvidersDeleteDefault creates a CircuitsProvidersDeleteDefault with default headers values
func NewCircuitsProvidersDeleteDefault(code int) *CircuitsProvidersDeleteDefault {
	return &CircuitsProvidersDeleteDefault{
		_statusCode: code,
	}
}

/* CircuitsProvidersDeleteDefault describes a response with status code -1, with default header values.

CircuitsProvidersDeleteDefault circuits providers delete default
*/
type CircuitsProvidersDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers delete default response
func (o *CircuitsProvidersDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProvidersDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/{id}/][%d] circuits_providers_delete default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsProvidersDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
