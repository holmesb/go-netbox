// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// DcimInventoryItemsUpdateReader is a Reader for the DcimInventoryItemsUpdate structure.
type DcimInventoryItemsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInventoryItemsUpdateOK creates a DcimInventoryItemsUpdateOK with default headers values
func NewDcimInventoryItemsUpdateOK() *DcimInventoryItemsUpdateOK {
	return &DcimInventoryItemsUpdateOK{}
}

/*DcimInventoryItemsUpdateOK handles this case with default header values.

DcimInventoryItemsUpdateOK dcim inventory items update o k
*/
type DcimInventoryItemsUpdateOK struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcimInventoryItemsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemsUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
