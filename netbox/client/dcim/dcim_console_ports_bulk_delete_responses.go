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

// DcimConsolePortsBulkDeleteReader is a Reader for the DcimConsolePortsBulkDelete structure.
type DcimConsolePortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsBulkDeleteNoContent creates a DcimConsolePortsBulkDeleteNoContent with default headers values
func NewDcimConsolePortsBulkDeleteNoContent() *DcimConsolePortsBulkDeleteNoContent {
	return &DcimConsolePortsBulkDeleteNoContent{}
}

/* DcimConsolePortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortsBulkDeleteNoContent dcim console ports bulk delete no content
*/
type DcimConsolePortsBulkDeleteNoContent struct {
}

func (o *DcimConsolePortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/][%d] dcimConsolePortsBulkDeleteNoContent ", 204)
}

func (o *DcimConsolePortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsolePortsBulkDeleteDefault creates a DcimConsolePortsBulkDeleteDefault with default headers values
func NewDcimConsolePortsBulkDeleteDefault(code int) *DcimConsolePortsBulkDeleteDefault {
	return &DcimConsolePortsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimConsolePortsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimConsolePortsBulkDeleteDefault dcim console ports bulk delete default
*/
type DcimConsolePortsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports bulk delete default response
func (o *DcimConsolePortsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/][%d] dcim_console-ports_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsolePortsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
