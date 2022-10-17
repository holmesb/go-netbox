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

// DcimModuleBaysReadReader is a Reader for the DcimModuleBaysRead structure.
type DcimModuleBaysReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysReadOK creates a DcimModuleBaysReadOK with default headers values
func NewDcimModuleBaysReadOK() *DcimModuleBaysReadOK {
	return &DcimModuleBaysReadOK{}
}

/*
DcimModuleBaysReadOK describes a response with status code 200, with default header values.

DcimModuleBaysReadOK dcim module bays read o k
*/
type DcimModuleBaysReadOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays read o k response has a 2xx status code
func (o *DcimModuleBaysReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays read o k response has a 3xx status code
func (o *DcimModuleBaysReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays read o k response has a 4xx status code
func (o *DcimModuleBaysReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays read o k response has a 5xx status code
func (o *DcimModuleBaysReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays read o k response a status code equal to that given
func (o *DcimModuleBaysReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBaysReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/module-bays/{id}/][%d] dcimModuleBaysReadOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/module-bays/{id}/][%d] dcimModuleBaysReadOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysReadOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysReadDefault creates a DcimModuleBaysReadDefault with default headers values
func NewDcimModuleBaysReadDefault(code int) *DcimModuleBaysReadDefault {
	return &DcimModuleBaysReadDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysReadDefault describes a response with status code -1, with default header values.

DcimModuleBaysReadDefault dcim module bays read default
*/
type DcimModuleBaysReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays read default response
func (o *DcimModuleBaysReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays read default response has a 2xx status code
func (o *DcimModuleBaysReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays read default response has a 3xx status code
func (o *DcimModuleBaysReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays read default response has a 4xx status code
func (o *DcimModuleBaysReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays read default response has a 5xx status code
func (o *DcimModuleBaysReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays read default response a status code equal to that given
func (o *DcimModuleBaysReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/module-bays/{id}/][%d] dcim_module-bays_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/module-bays/{id}/][%d] dcim_module-bays_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}