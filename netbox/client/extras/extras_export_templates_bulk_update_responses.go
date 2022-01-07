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

// ExtrasExportTemplatesBulkUpdateReader is a Reader for the ExtrasExportTemplatesBulkUpdate structure.
type ExtrasExportTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesBulkUpdateOK creates a ExtrasExportTemplatesBulkUpdateOK with default headers values
func NewExtrasExportTemplatesBulkUpdateOK() *ExtrasExportTemplatesBulkUpdateOK {
	return &ExtrasExportTemplatesBulkUpdateOK{}
}

/* ExtrasExportTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesBulkUpdateOK extras export templates bulk update o k
*/
type ExtrasExportTemplatesBulkUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/][%d] extrasExportTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasExportTemplatesBulkUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesBulkUpdateDefault creates a ExtrasExportTemplatesBulkUpdateDefault with default headers values
func NewExtrasExportTemplatesBulkUpdateDefault(code int) *ExtrasExportTemplatesBulkUpdateDefault {
	return &ExtrasExportTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasExportTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasExportTemplatesBulkUpdateDefault extras export templates bulk update default
*/
type ExtrasExportTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates bulk update default response
func (o *ExtrasExportTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/][%d] extras_export-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasExportTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
