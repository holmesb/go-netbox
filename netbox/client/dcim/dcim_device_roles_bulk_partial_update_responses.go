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

	"github.com/holmesb/go-netbox/netbox/models"
)

// DcimDeviceRolesBulkPartialUpdateReader is a Reader for the DcimDeviceRolesBulkPartialUpdate structure.
type DcimDeviceRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesBulkPartialUpdateOK creates a DcimDeviceRolesBulkPartialUpdateOK with default headers values
func NewDcimDeviceRolesBulkPartialUpdateOK() *DcimDeviceRolesBulkPartialUpdateOK {
	return &DcimDeviceRolesBulkPartialUpdateOK{}
}

/* DcimDeviceRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesBulkPartialUpdateOK dcim device roles bulk partial update o k
*/
type DcimDeviceRolesBulkPartialUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/][%d] dcimDeviceRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesBulkPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesBulkPartialUpdateDefault creates a DcimDeviceRolesBulkPartialUpdateDefault with default headers values
func NewDcimDeviceRolesBulkPartialUpdateDefault(code int) *DcimDeviceRolesBulkPartialUpdateDefault {
	return &DcimDeviceRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDeviceRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceRolesBulkPartialUpdateDefault dcim device roles bulk partial update default
*/
type DcimDeviceRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles bulk partial update default response
func (o *DcimDeviceRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/][%d] dcim_device-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDeviceRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
