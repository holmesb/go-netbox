// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomLinksBulkDeleteReader is a Reader for the ExtrasCustomLinksBulkDelete structure.
type ExtrasCustomLinksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomLinksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksBulkDeleteNoContent creates a ExtrasCustomLinksBulkDeleteNoContent with default headers values
func NewExtrasCustomLinksBulkDeleteNoContent() *ExtrasCustomLinksBulkDeleteNoContent {
	return &ExtrasCustomLinksBulkDeleteNoContent{}
}

/* ExtrasCustomLinksBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomLinksBulkDeleteNoContent extras custom links bulk delete no content
*/
type ExtrasCustomLinksBulkDeleteNoContent struct {
}

func (o *ExtrasCustomLinksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-links/][%d] extrasCustomLinksBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomLinksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasCustomLinksBulkDeleteDefault creates a ExtrasCustomLinksBulkDeleteDefault with default headers values
func NewExtrasCustomLinksBulkDeleteDefault(code int) *ExtrasCustomLinksBulkDeleteDefault {
	return &ExtrasCustomLinksBulkDeleteDefault{
		_statusCode: code,
	}
}

/* ExtrasCustomLinksBulkDeleteDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksBulkDeleteDefault extras custom links bulk delete default
*/
type ExtrasCustomLinksBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links bulk delete default response
func (o *ExtrasCustomLinksBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomLinksBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-links/][%d] extras_custom-links_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasCustomLinksBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
