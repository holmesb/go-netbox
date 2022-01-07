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

// DcimPlatformsReadReader is a Reader for the DcimPlatformsRead structure.
type DcimPlatformsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsReadOK creates a DcimPlatformsReadOK with default headers values
func NewDcimPlatformsReadOK() *DcimPlatformsReadOK {
	return &DcimPlatformsReadOK{}
}

/* DcimPlatformsReadOK describes a response with status code 200, with default header values.

DcimPlatformsReadOK dcim platforms read o k
*/
type DcimPlatformsReadOK struct {
	Payload *models.Platform
}

func (o *DcimPlatformsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/platforms/{id}/][%d] dcimPlatformsReadOK  %+v", 200, o.Payload)
}
func (o *DcimPlatformsReadOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsReadDefault creates a DcimPlatformsReadDefault with default headers values
func NewDcimPlatformsReadDefault(code int) *DcimPlatformsReadDefault {
	return &DcimPlatformsReadDefault{
		_statusCode: code,
	}
}

/* DcimPlatformsReadDefault describes a response with status code -1, with default header values.

DcimPlatformsReadDefault dcim platforms read default
*/
type DcimPlatformsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms read default response
func (o *DcimPlatformsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimPlatformsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/platforms/{id}/][%d] dcim_platforms_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPlatformsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
