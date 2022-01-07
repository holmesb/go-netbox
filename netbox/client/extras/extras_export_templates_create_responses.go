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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/holmesb/go-netbox/netbox/models"
)

// ExtrasExportTemplatesCreateReader is a Reader for the ExtrasExportTemplatesCreate structure.
type ExtrasExportTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasExportTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesCreateCreated creates a ExtrasExportTemplatesCreateCreated with default headers values
func NewExtrasExportTemplatesCreateCreated() *ExtrasExportTemplatesCreateCreated {
	return &ExtrasExportTemplatesCreateCreated{}
}

/* ExtrasExportTemplatesCreateCreated describes a response with status code 201, with default header values.

ExtrasExportTemplatesCreateCreated extras export templates create created
*/
type ExtrasExportTemplatesCreateCreated struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extrasExportTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasExportTemplatesCreateCreated) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesCreateDefault creates a ExtrasExportTemplatesCreateDefault with default headers values
func NewExtrasExportTemplatesCreateDefault(code int) *ExtrasExportTemplatesCreateDefault {
	return &ExtrasExportTemplatesCreateDefault{
		_statusCode: code,
	}
}

/* ExtrasExportTemplatesCreateDefault describes a response with status code -1, with default header values.

ExtrasExportTemplatesCreateDefault extras export templates create default
*/
type ExtrasExportTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates create default response
func (o *ExtrasExportTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extras_export-templates_create default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasExportTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
