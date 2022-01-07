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

// DcimPowerOutletTemplatesUpdateReader is a Reader for the DcimPowerOutletTemplatesUpdate structure.
type DcimPowerOutletTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesUpdateOK creates a DcimPowerOutletTemplatesUpdateOK with default headers values
func NewDcimPowerOutletTemplatesUpdateOK() *DcimPowerOutletTemplatesUpdateOK {
	return &DcimPowerOutletTemplatesUpdateOK{}
}

/* DcimPowerOutletTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesUpdateOK dcim power outlet templates update o k
*/
type DcimPowerOutletTemplatesUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesUpdateDefault creates a DcimPowerOutletTemplatesUpdateDefault with default headers values
func NewDcimPowerOutletTemplatesUpdateDefault(code int) *DcimPowerOutletTemplatesUpdateDefault {
	return &DcimPowerOutletTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPowerOutletTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesUpdateDefault dcim power outlet templates update default
*/
type DcimPowerOutletTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlet templates update default response
func (o *DcimPowerOutletTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerOutletTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
