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

// ExtrasTagsPartialUpdateReader is a Reader for the ExtrasTagsPartialUpdate structure.
type ExtrasTagsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsPartialUpdateOK creates a ExtrasTagsPartialUpdateOK with default headers values
func NewExtrasTagsPartialUpdateOK() *ExtrasTagsPartialUpdateOK {
	return &ExtrasTagsPartialUpdateOK{}
}

/* ExtrasTagsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsPartialUpdateOK extras tags partial update o k
*/
type ExtrasTagsPartialUpdateOK struct {
	Payload *models.Tag
}

func (o *ExtrasTagsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/tags/{id}/][%d] extrasTagsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasTagsPartialUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsPartialUpdateDefault creates a ExtrasTagsPartialUpdateDefault with default headers values
func NewExtrasTagsPartialUpdateDefault(code int) *ExtrasTagsPartialUpdateDefault {
	return &ExtrasTagsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasTagsPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasTagsPartialUpdateDefault extras tags partial update default
*/
type ExtrasTagsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags partial update default response
func (o *ExtrasTagsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasTagsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/tags/{id}/][%d] extras_tags_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasTagsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
