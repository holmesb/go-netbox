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

// DcimConsoleServerPortTemplatesBulkUpdateReader is a Reader for the DcimConsoleServerPortTemplatesBulkUpdate structure.
type DcimConsoleServerPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesBulkUpdateOK creates a DcimConsoleServerPortTemplatesBulkUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesBulkUpdateOK() *DcimConsoleServerPortTemplatesBulkUpdateOK {
	return &DcimConsoleServerPortTemplatesBulkUpdateOK{}
}

/* DcimConsoleServerPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesBulkUpdateOK dcim console server port templates bulk update o k
*/
type DcimConsoleServerPortTemplatesBulkUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesBulkUpdateDefault creates a DcimConsoleServerPortTemplatesBulkUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesBulkUpdateDefault(code int) *DcimConsoleServerPortTemplatesBulkUpdateDefault {
	return &DcimConsoleServerPortTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesBulkUpdateDefault dcim console server port templates bulk update default
*/
type DcimConsoleServerPortTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates bulk update default response
func (o *DcimConsoleServerPortTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/][%d] dcim_console-server-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
