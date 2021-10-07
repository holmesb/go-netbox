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

// DcimDeviceRolesReadReader is a Reader for the DcimDeviceRolesRead structure.
type DcimDeviceRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesReadOK creates a DcimDeviceRolesReadOK with default headers values
func NewDcimDeviceRolesReadOK() *DcimDeviceRolesReadOK {
	return &DcimDeviceRolesReadOK{}
}

/* DcimDeviceRolesReadOK describes a response with status code 200, with default header values.

DcimDeviceRolesReadOK dcim device roles read o k
*/
type DcimDeviceRolesReadOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcimDeviceRolesReadOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesReadOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesReadDefault creates a DcimDeviceRolesReadDefault with default headers values
func NewDcimDeviceRolesReadDefault(code int) *DcimDeviceRolesReadDefault {
	return &DcimDeviceRolesReadDefault{
		_statusCode: code,
	}
}

/* DcimDeviceRolesReadDefault describes a response with status code -1, with default header values.

DcimDeviceRolesReadDefault dcim device roles read default
*/
type DcimDeviceRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles read default response
func (o *DcimDeviceRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/{id}/][%d] dcim_device-roles_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDeviceRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
