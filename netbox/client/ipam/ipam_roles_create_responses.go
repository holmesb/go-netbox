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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamRolesCreateReader is a Reader for the IpamRolesCreate structure.
type IpamRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesCreateCreated creates a IpamRolesCreateCreated with default headers values
func NewIpamRolesCreateCreated() *IpamRolesCreateCreated {
	return &IpamRolesCreateCreated{}
}

/* IpamRolesCreateCreated describes a response with status code 201, with default header values.

IpamRolesCreateCreated ipam roles create created
*/
type IpamRolesCreateCreated struct {
	Payload *models.Role
}

func (o *IpamRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/roles/][%d] ipamRolesCreateCreated  %+v", 201, o.Payload)
}
func (o *IpamRolesCreateCreated) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesCreateDefault creates a IpamRolesCreateDefault with default headers values
func NewIpamRolesCreateDefault(code int) *IpamRolesCreateDefault {
	return &IpamRolesCreateDefault{
		_statusCode: code,
	}
}

/* IpamRolesCreateDefault describes a response with status code -1, with default header values.

IpamRolesCreateDefault ipam roles create default
*/
type IpamRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles create default response
func (o *IpamRolesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/roles/][%d] ipam_roles_create default  %+v", o._statusCode, o.Payload)
}
func (o *IpamRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
