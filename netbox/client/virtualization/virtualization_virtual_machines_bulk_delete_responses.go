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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationVirtualMachinesBulkDeleteReader is a Reader for the VirtualizationVirtualMachinesBulkDelete structure.
type VirtualizationVirtualMachinesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationVirtualMachinesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesBulkDeleteNoContent creates a VirtualizationVirtualMachinesBulkDeleteNoContent with default headers values
func NewVirtualizationVirtualMachinesBulkDeleteNoContent() *VirtualizationVirtualMachinesBulkDeleteNoContent {
	return &VirtualizationVirtualMachinesBulkDeleteNoContent{}
}

/* VirtualizationVirtualMachinesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationVirtualMachinesBulkDeleteNoContent virtualization virtual machines bulk delete no content
*/
type VirtualizationVirtualMachinesBulkDeleteNoContent struct {
}

func (o *VirtualizationVirtualMachinesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationVirtualMachinesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationVirtualMachinesBulkDeleteDefault creates a VirtualizationVirtualMachinesBulkDeleteDefault with default headers values
func NewVirtualizationVirtualMachinesBulkDeleteDefault(code int) *VirtualizationVirtualMachinesBulkDeleteDefault {
	return &VirtualizationVirtualMachinesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* VirtualizationVirtualMachinesBulkDeleteDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesBulkDeleteDefault virtualization virtual machines bulk delete default
*/
type VirtualizationVirtualMachinesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines bulk delete default response
func (o *VirtualizationVirtualMachinesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/][%d] virtualization_virtual-machines_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationVirtualMachinesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
