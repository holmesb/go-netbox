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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/holmesb/go-netbox/netbox/models"
)

// TenancyTenantGroupsBulkPartialUpdateReader is a Reader for the TenancyTenantGroupsBulkPartialUpdate structure.
type TenancyTenantGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsBulkPartialUpdateOK creates a TenancyTenantGroupsBulkPartialUpdateOK with default headers values
func NewTenancyTenantGroupsBulkPartialUpdateOK() *TenancyTenantGroupsBulkPartialUpdateOK {
	return &TenancyTenantGroupsBulkPartialUpdateOK{}
}

/* TenancyTenantGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsBulkPartialUpdateOK tenancy tenant groups bulk partial update o k
*/
type TenancyTenantGroupsBulkPartialUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsBulkPartialUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantGroupsBulkPartialUpdateDefault creates a TenancyTenantGroupsBulkPartialUpdateDefault with default headers values
func NewTenancyTenantGroupsBulkPartialUpdateDefault(code int) *TenancyTenantGroupsBulkPartialUpdateDefault {
	return &TenancyTenantGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyTenantGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantGroupsBulkPartialUpdateDefault tenancy tenant groups bulk partial update default
*/
type TenancyTenantGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenant groups bulk partial update default response
func (o *TenancyTenantGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/][%d] tenancy_tenant-groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyTenantGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
