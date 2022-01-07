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

// ExtrasJournalEntriesBulkUpdateReader is a Reader for the ExtrasJournalEntriesBulkUpdate structure.
type ExtrasJournalEntriesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesBulkUpdateOK creates a ExtrasJournalEntriesBulkUpdateOK with default headers values
func NewExtrasJournalEntriesBulkUpdateOK() *ExtrasJournalEntriesBulkUpdateOK {
	return &ExtrasJournalEntriesBulkUpdateOK{}
}

/* ExtrasJournalEntriesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesBulkUpdateOK extras journal entries bulk update o k
*/
type ExtrasJournalEntriesBulkUpdateOK struct {
	Payload *models.JournalEntry
}

func (o *ExtrasJournalEntriesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasJournalEntriesBulkUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesBulkUpdateDefault creates a ExtrasJournalEntriesBulkUpdateDefault with default headers values
func NewExtrasJournalEntriesBulkUpdateDefault(code int) *ExtrasJournalEntriesBulkUpdateDefault {
	return &ExtrasJournalEntriesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasJournalEntriesBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesBulkUpdateDefault extras journal entries bulk update default
*/
type ExtrasJournalEntriesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras journal entries bulk update default response
func (o *ExtrasJournalEntriesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extras_journal-entries_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasJournalEntriesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
