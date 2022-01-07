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

// IpamVlansPartialUpdateReader is a Reader for the IpamVlansPartialUpdate structure.
type IpamVlansPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansPartialUpdateOK creates a IpamVlansPartialUpdateOK with default headers values
func NewIpamVlansPartialUpdateOK() *IpamVlansPartialUpdateOK {
	return &IpamVlansPartialUpdateOK{}
}

/* IpamVlansPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlansPartialUpdateOK ipam vlans partial update o k
*/
type IpamVlansPartialUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlans/{id}/][%d] ipamVlansPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlansPartialUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlansPartialUpdateDefault creates a IpamVlansPartialUpdateDefault with default headers values
func NewIpamVlansPartialUpdateDefault(code int) *IpamVlansPartialUpdateDefault {
	return &IpamVlansPartialUpdateDefault{
		_statusCode: code,
	}
}

/* IpamVlansPartialUpdateDefault describes a response with status code -1, with default header values.

IpamVlansPartialUpdateDefault ipam vlans partial update default
*/
type IpamVlansPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlans partial update default response
func (o *IpamVlansPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlans/{id}/][%d] ipam_vlans_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlansPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
