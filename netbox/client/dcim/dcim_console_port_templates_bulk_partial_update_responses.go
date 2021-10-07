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

// DcimConsolePortTemplatesBulkPartialUpdateReader is a Reader for the DcimConsolePortTemplatesBulkPartialUpdate structure.
type DcimConsolePortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesBulkPartialUpdateOK creates a DcimConsolePortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimConsolePortTemplatesBulkPartialUpdateOK() *DcimConsolePortTemplatesBulkPartialUpdateOK {
	return &DcimConsolePortTemplatesBulkPartialUpdateOK{}
}

/* DcimConsolePortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesBulkPartialUpdateOK dcim console port templates bulk partial update o k
*/
type DcimConsolePortTemplatesBulkPartialUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesBulkPartialUpdateDefault creates a DcimConsolePortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimConsolePortTemplatesBulkPartialUpdateDefault(code int) *DcimConsolePortTemplatesBulkPartialUpdateDefault {
	return &DcimConsolePortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsolePortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortTemplatesBulkPartialUpdateDefault dcim console port templates bulk partial update default
*/
type DcimConsolePortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates bulk partial update default response
func (o *DcimConsolePortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/][%d] dcim_console-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsolePortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
