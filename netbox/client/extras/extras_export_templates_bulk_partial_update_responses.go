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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasExportTemplatesBulkPartialUpdateReader is a Reader for the ExtrasExportTemplatesBulkPartialUpdate structure.
type ExtrasExportTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesBulkPartialUpdateOK creates a ExtrasExportTemplatesBulkPartialUpdateOK with default headers values
func NewExtrasExportTemplatesBulkPartialUpdateOK() *ExtrasExportTemplatesBulkPartialUpdateOK {
	return &ExtrasExportTemplatesBulkPartialUpdateOK{}
}

/* ExtrasExportTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesBulkPartialUpdateOK extras export templates bulk partial update o k
*/
type ExtrasExportTemplatesBulkPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/][%d] extrasExportTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesBulkPartialUpdateDefault creates a ExtrasExportTemplatesBulkPartialUpdateDefault with default headers values
func NewExtrasExportTemplatesBulkPartialUpdateDefault(code int) *ExtrasExportTemplatesBulkPartialUpdateDefault {
	return &ExtrasExportTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasExportTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasExportTemplatesBulkPartialUpdateDefault extras export templates bulk partial update default
*/
type ExtrasExportTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates bulk partial update default response
func (o *ExtrasExportTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/][%d] extras_export-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasExportTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
