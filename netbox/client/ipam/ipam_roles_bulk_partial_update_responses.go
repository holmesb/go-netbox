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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/holmesb/go-netbox/netbox/models"
)

// IpamRolesBulkPartialUpdateReader is a Reader for the IpamRolesBulkPartialUpdate structure.
type IpamRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesBulkPartialUpdateOK creates a IpamRolesBulkPartialUpdateOK with default headers values
func NewIpamRolesBulkPartialUpdateOK() *IpamRolesBulkPartialUpdateOK {
	return &IpamRolesBulkPartialUpdateOK{}
}

/* IpamRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamRolesBulkPartialUpdateOK ipam roles bulk partial update o k
*/
type IpamRolesBulkPartialUpdateOK struct {
	Payload *models.Role
}

func (o *IpamRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/][%d] ipamRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamRolesBulkPartialUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesBulkPartialUpdateDefault creates a IpamRolesBulkPartialUpdateDefault with default headers values
func NewIpamRolesBulkPartialUpdateDefault(code int) *IpamRolesBulkPartialUpdateDefault {
	return &IpamRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* IpamRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamRolesBulkPartialUpdateDefault ipam roles bulk partial update default
*/
type IpamRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles bulk partial update default response
func (o *IpamRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/][%d] ipam_roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
