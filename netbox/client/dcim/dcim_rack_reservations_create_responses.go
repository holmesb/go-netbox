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

// DcimRackReservationsCreateReader is a Reader for the DcimRackReservationsCreate structure.
type DcimRackReservationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackReservationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackReservationsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackReservationsCreateCreated creates a DcimRackReservationsCreateCreated with default headers values
func NewDcimRackReservationsCreateCreated() *DcimRackReservationsCreateCreated {
	return &DcimRackReservationsCreateCreated{}
}

/* DcimRackReservationsCreateCreated describes a response with status code 201, with default header values.

DcimRackReservationsCreateCreated dcim rack reservations create created
*/
type DcimRackReservationsCreateCreated struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-reservations/][%d] dcimRackReservationsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRackReservationsCreateCreated) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackReservationsCreateDefault creates a DcimRackReservationsCreateDefault with default headers values
func NewDcimRackReservationsCreateDefault(code int) *DcimRackReservationsCreateDefault {
	return &DcimRackReservationsCreateDefault{
		_statusCode: code,
	}
}

/* DcimRackReservationsCreateDefault describes a response with status code -1, with default header values.

DcimRackReservationsCreateDefault dcim rack reservations create default
*/
type DcimRackReservationsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack reservations create default response
func (o *DcimRackReservationsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackReservationsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-reservations/][%d] dcim_rack-reservations_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRackReservationsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
