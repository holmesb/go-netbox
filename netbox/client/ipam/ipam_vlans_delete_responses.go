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
)

// IpamVlansDeleteReader is a Reader for the IpamVlansDelete structure.
type IpamVlansDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlansDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansDeleteNoContent creates a IpamVlansDeleteNoContent with default headers values
func NewIpamVlansDeleteNoContent() *IpamVlansDeleteNoContent {
	return &IpamVlansDeleteNoContent{}
}

/* IpamVlansDeleteNoContent describes a response with status code 204, with default header values.

IpamVlansDeleteNoContent ipam vlans delete no content
*/
type IpamVlansDeleteNoContent struct {
}

func (o *IpamVlansDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/{id}/][%d] ipamVlansDeleteNoContent ", 204)
}

func (o *IpamVlansDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVlansDeleteDefault creates a IpamVlansDeleteDefault with default headers values
func NewIpamVlansDeleteDefault(code int) *IpamVlansDeleteDefault {
	return &IpamVlansDeleteDefault{
		_statusCode: code,
	}
}

/* IpamVlansDeleteDefault describes a response with status code -1, with default header values.

IpamVlansDeleteDefault ipam vlans delete default
*/
type IpamVlansDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlans delete default response
func (o *IpamVlansDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/{id}/][%d] ipam_vlans_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlansDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
