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
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/holmesb/go-netbox/netbox/models"
)

// DcimPowerPortsListReader is a Reader for the DcimPowerPortsList structure.
type DcimPowerPortsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsListOK creates a DcimPowerPortsListOK with default headers values
func NewDcimPowerPortsListOK() *DcimPowerPortsListOK {
	return &DcimPowerPortsListOK{}
}

/* DcimPowerPortsListOK describes a response with status code 200, with default header values.

DcimPowerPortsListOK dcim power ports list o k
*/
type DcimPowerPortsListOK struct {
	Payload *DcimPowerPortsListOKBody
}

func (o *DcimPowerPortsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/][%d] dcimPowerPortsListOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortsListOK) GetPayload() *DcimPowerPortsListOKBody {
	return o.Payload
}

func (o *DcimPowerPortsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimPowerPortsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsListDefault creates a DcimPowerPortsListDefault with default headers values
func NewDcimPowerPortsListDefault(code int) *DcimPowerPortsListDefault {
	return &DcimPowerPortsListDefault{
		_statusCode: code,
	}
}

/* DcimPowerPortsListDefault describes a response with status code -1, with default header values.

DcimPowerPortsListDefault dcim power ports list default
*/
type DcimPowerPortsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power ports list default response
func (o *DcimPowerPortsListDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/][%d] dcim_power-ports_list default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerPortsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DcimPowerPortsListOKBody dcim power ports list o k body
swagger:model DcimPowerPortsListOKBody
*/
type DcimPowerPortsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.PowerPort `json:"results"`
}

// Validate validates this dcim power ports list o k body
func (o *DcimPowerPortsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimPowerPortsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPortsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPortsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPortsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPortsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim power ports list o k body based on the context it is used
func (o *DcimPowerPortsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimPowerPortsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimPowerPortsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimPowerPortsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimPowerPortsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
