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

// ExtrasCustomLinksPartialUpdateReader is a Reader for the ExtrasCustomLinksPartialUpdate structure.
type ExtrasCustomLinksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksPartialUpdateOK creates a ExtrasCustomLinksPartialUpdateOK with default headers values
func NewExtrasCustomLinksPartialUpdateOK() *ExtrasCustomLinksPartialUpdateOK {
	return &ExtrasCustomLinksPartialUpdateOK{}
}

/* ExtrasCustomLinksPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksPartialUpdateOK extras custom links partial update o k
*/
type ExtrasCustomLinksPartialUpdateOK struct {
	Payload *models.CustomLink
}

func (o *ExtrasCustomLinksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-links/{id}/][%d] extrasCustomLinksPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomLinksPartialUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksPartialUpdateDefault creates a ExtrasCustomLinksPartialUpdateDefault with default headers values
func NewExtrasCustomLinksPartialUpdateDefault(code int) *ExtrasCustomLinksPartialUpdateDefault {
	return &ExtrasCustomLinksPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasCustomLinksPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksPartialUpdateDefault extras custom links partial update default
*/
type ExtrasCustomLinksPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links partial update default response
func (o *ExtrasCustomLinksPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomLinksPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-links/{id}/][%d] extras_custom-links_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasCustomLinksPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
