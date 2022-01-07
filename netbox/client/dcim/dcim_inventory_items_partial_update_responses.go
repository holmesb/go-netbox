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

// DcimInventoryItemsPartialUpdateReader is a Reader for the DcimInventoryItemsPartialUpdate structure.
type DcimInventoryItemsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsPartialUpdateOK creates a DcimInventoryItemsPartialUpdateOK with default headers values
func NewDcimInventoryItemsPartialUpdateOK() *DcimInventoryItemsPartialUpdateOK {
	return &DcimInventoryItemsPartialUpdateOK{}
}

/* DcimInventoryItemsPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsPartialUpdateOK dcim inventory items partial update o k
*/
type DcimInventoryItemsPartialUpdateOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcimInventoryItemsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemsPartialUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsPartialUpdateDefault creates a DcimInventoryItemsPartialUpdateDefault with default headers values
func NewDcimInventoryItemsPartialUpdateDefault(code int) *DcimInventoryItemsPartialUpdateDefault {
	return &DcimInventoryItemsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemsPartialUpdateDefault dcim inventory items partial update default
*/
type DcimInventoryItemsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items partial update default response
func (o *DcimInventoryItemsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcim_inventory-items_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
