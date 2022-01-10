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

// IpamIPRangesBulkDeleteReader is a Reader for the IpamIPRangesBulkDelete structure.
type IpamIPRangesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPRangesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesBulkDeleteNoContent creates a IpamIPRangesBulkDeleteNoContent with default headers values
func NewIpamIPRangesBulkDeleteNoContent() *IpamIPRangesBulkDeleteNoContent {
	return &IpamIPRangesBulkDeleteNoContent{}
}

/* IpamIPRangesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamIPRangesBulkDeleteNoContent ipam Ip ranges bulk delete no content
*/
type IpamIPRangesBulkDeleteNoContent struct {
}

func (o *IpamIPRangesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/][%d] ipamIpRangesBulkDeleteNoContent ", 204)
}

func (o *IpamIPRangesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamIPRangesBulkDeleteDefault creates a IpamIPRangesBulkDeleteDefault with default headers values
func NewIpamIPRangesBulkDeleteDefault(code int) *IpamIPRangesBulkDeleteDefault {
	return &IpamIPRangesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamIPRangesBulkDeleteDefault describes a response with status code -1, with default header values.

IpamIPRangesBulkDeleteDefault ipam ip ranges bulk delete default
*/
type IpamIPRangesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges bulk delete default response
func (o *IpamIPRangesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPRangesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/][%d] ipam_ip-ranges_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamIPRangesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
