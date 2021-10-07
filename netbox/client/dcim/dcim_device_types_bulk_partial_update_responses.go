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

// DcimDeviceTypesBulkPartialUpdateReader is a Reader for the DcimDeviceTypesBulkPartialUpdate structure.
type DcimDeviceTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceTypesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesBulkPartialUpdateOK creates a DcimDeviceTypesBulkPartialUpdateOK with default headers values
func NewDcimDeviceTypesBulkPartialUpdateOK() *DcimDeviceTypesBulkPartialUpdateOK {
	return &DcimDeviceTypesBulkPartialUpdateOK{}
}

/* DcimDeviceTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkPartialUpdateOK dcim device types bulk partial update o k
*/
type DcimDeviceTypesBulkPartialUpdateOK struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceTypesBulkPartialUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesBulkPartialUpdateDefault creates a DcimDeviceTypesBulkPartialUpdateDefault with default headers values
func NewDcimDeviceTypesBulkPartialUpdateDefault(code int) *DcimDeviceTypesBulkPartialUpdateDefault {
	return &DcimDeviceTypesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDeviceTypesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceTypesBulkPartialUpdateDefault dcim device types bulk partial update default
*/
type DcimDeviceTypesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types bulk partial update default response
func (o *DcimDeviceTypesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcim_device-types_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDeviceTypesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
