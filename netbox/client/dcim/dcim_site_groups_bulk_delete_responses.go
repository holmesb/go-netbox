// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimSiteGroupsBulkDeleteReader is a Reader for the DcimSiteGroupsBulkDelete structure.
type DcimSiteGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSiteGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsBulkDeleteNoContent creates a DcimSiteGroupsBulkDeleteNoContent with default headers values
func NewDcimSiteGroupsBulkDeleteNoContent() *DcimSiteGroupsBulkDeleteNoContent {
	return &DcimSiteGroupsBulkDeleteNoContent{}
}

/* DcimSiteGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimSiteGroupsBulkDeleteNoContent dcim site groups bulk delete no content
*/
type DcimSiteGroupsBulkDeleteNoContent struct {
}

func (o *DcimSiteGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/site-groups/][%d] dcimSiteGroupsBulkDeleteNoContent ", 204)
}

func (o *DcimSiteGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimSiteGroupsBulkDeleteDefault creates a DcimSiteGroupsBulkDeleteDefault with default headers values
func NewDcimSiteGroupsBulkDeleteDefault(code int) *DcimSiteGroupsBulkDeleteDefault {
	return &DcimSiteGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimSiteGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimSiteGroupsBulkDeleteDefault dcim site groups bulk delete default
*/
type DcimSiteGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups bulk delete default response
func (o *DcimSiteGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/site-groups/][%d] dcim_site-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSiteGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
