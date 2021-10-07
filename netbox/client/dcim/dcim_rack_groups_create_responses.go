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

// DcimRackGroupsCreateReader is a Reader for the DcimRackGroupsCreate structure.
type DcimRackGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackGroupsCreateCreated creates a DcimRackGroupsCreateCreated with default headers values
func NewDcimRackGroupsCreateCreated() *DcimRackGroupsCreateCreated {
	return &DcimRackGroupsCreateCreated{}
}

/* DcimRackGroupsCreateCreated describes a response with status code 201, with default header values.

DcimRackGroupsCreateCreated dcim rack groups create created
*/
type DcimRackGroupsCreateCreated struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-groups/][%d] dcimRackGroupsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRackGroupsCreateCreated) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackGroupsCreateDefault creates a DcimRackGroupsCreateDefault with default headers values
func NewDcimRackGroupsCreateDefault(code int) *DcimRackGroupsCreateDefault {
	return &DcimRackGroupsCreateDefault{
		_statusCode: code,
	}
}

/* DcimRackGroupsCreateDefault describes a response with status code -1, with default header values.

DcimRackGroupsCreateDefault dcim rack groups create default
*/
type DcimRackGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack groups create default response
func (o *DcimRackGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-groups/][%d] dcim_rack-groups_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRackGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
