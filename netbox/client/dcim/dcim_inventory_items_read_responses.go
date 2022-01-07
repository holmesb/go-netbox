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

// DcimInventoryItemsReadReader is a Reader for the DcimInventoryItemsRead structure.
type DcimInventoryItemsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsReadOK creates a DcimInventoryItemsReadOK with default headers values
func NewDcimInventoryItemsReadOK() *DcimInventoryItemsReadOK {
	return &DcimInventoryItemsReadOK{}
}

/* DcimInventoryItemsReadOK describes a response with status code 200, with default header values.

DcimInventoryItemsReadOK dcim inventory items read o k
*/
type DcimInventoryItemsReadOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/{id}/][%d] dcimInventoryItemsReadOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsReadOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsReadDefault creates a DcimInventoryItemsReadDefault with default headers values
func NewDcimInventoryItemsReadDefault(code int) *DcimInventoryItemsReadDefault {
	return &DcimInventoryItemsReadDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemsReadDefault describes a response with status code -1, with default header values.

DcimInventoryItemsReadDefault dcim inventory items read default
*/
type DcimInventoryItemsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items read default response
func (o *DcimInventoryItemsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/{id}/][%d] dcim_inventory-items_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
