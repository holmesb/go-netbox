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
)

// DcimVirtualChassisBulkDeleteReader is a Reader for the DcimVirtualChassisBulkDelete structure.
type DcimVirtualChassisBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimVirtualChassisBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisBulkDeleteNoContent creates a DcimVirtualChassisBulkDeleteNoContent with default headers values
func NewDcimVirtualChassisBulkDeleteNoContent() *DcimVirtualChassisBulkDeleteNoContent {
	return &DcimVirtualChassisBulkDeleteNoContent{}
}

/* DcimVirtualChassisBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimVirtualChassisBulkDeleteNoContent dcim virtual chassis bulk delete no content
*/
type DcimVirtualChassisBulkDeleteNoContent struct {
}

func (o *DcimVirtualChassisBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkDeleteNoContent ", 204)
}

func (o *DcimVirtualChassisBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimVirtualChassisBulkDeleteDefault creates a DcimVirtualChassisBulkDeleteDefault with default headers values
func NewDcimVirtualChassisBulkDeleteDefault(code int) *DcimVirtualChassisBulkDeleteDefault {
	return &DcimVirtualChassisBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimVirtualChassisBulkDeleteDefault describes a response with status code -1, with default header values.

DcimVirtualChassisBulkDeleteDefault dcim virtual chassis bulk delete default
*/
type DcimVirtualChassisBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis bulk delete default response
func (o *DcimVirtualChassisBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcim_virtual-chassis_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimVirtualChassisBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
