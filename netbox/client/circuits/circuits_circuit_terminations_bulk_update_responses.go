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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/holmesb/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsBulkUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkUpdate structure.
type CircuitsCircuitTerminationsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateOK creates a CircuitsCircuitTerminationsBulkUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkUpdateOK() *CircuitsCircuitTerminationsBulkUpdateOK {
	return &CircuitsCircuitTerminationsBulkUpdateOK{}
}

/* CircuitsCircuitTerminationsBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkUpdateOK circuits circuit terminations bulk update o k
*/
type CircuitsCircuitTerminationsBulkUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsBulkUpdateDefault creates a CircuitsCircuitTerminationsBulkUpdateDefault with default headers values
func NewCircuitsCircuitTerminationsBulkUpdateDefault(code int) *CircuitsCircuitTerminationsBulkUpdateDefault {
	return &CircuitsCircuitTerminationsBulkUpdateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitTerminationsBulkUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsBulkUpdateDefault circuits circuit terminations bulk update default
*/
type CircuitsCircuitTerminationsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations bulk update default response
func (o *CircuitsCircuitTerminationsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/][%d] circuits_circuit-terminations_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
