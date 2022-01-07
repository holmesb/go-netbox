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

	"github.com/holmesb/go-netbox/netbox/models"
)

// DcimSitesBulkUpdateReader is a Reader for the DcimSitesBulkUpdate structure.
type DcimSitesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesBulkUpdateOK creates a DcimSitesBulkUpdateOK with default headers values
func NewDcimSitesBulkUpdateOK() *DcimSitesBulkUpdateOK {
	return &DcimSitesBulkUpdateOK{}
}

/* DcimSitesBulkUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkUpdateOK dcim sites bulk update o k
*/
type DcimSitesBulkUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/sites/][%d] dcimSitesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesBulkUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesBulkUpdateDefault creates a DcimSitesBulkUpdateDefault with default headers values
func NewDcimSitesBulkUpdateDefault(code int) *DcimSitesBulkUpdateDefault {
	return &DcimSitesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSitesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimSitesBulkUpdateDefault dcim sites bulk update default
*/
type DcimSitesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites bulk update default response
func (o *DcimSitesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/sites/][%d] dcim_sites_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSitesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
