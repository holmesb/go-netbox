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

package ipam

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

// IpamRouteTargetsListReader is a Reader for the IpamRouteTargetsList structure.
type IpamRouteTargetsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsListOK creates a IpamRouteTargetsListOK with default headers values
func NewIpamRouteTargetsListOK() *IpamRouteTargetsListOK {
	return &IpamRouteTargetsListOK{}
}

/* IpamRouteTargetsListOK describes a response with status code 200, with default header values.

IpamRouteTargetsListOK ipam route targets list o k
*/
type IpamRouteTargetsListOK struct {
	Payload *IpamRouteTargetsListOKBody
}

func (o *IpamRouteTargetsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/route-targets/][%d] ipamRouteTargetsListOK  %+v", 200, o.Payload)
}
func (o *IpamRouteTargetsListOK) GetPayload() *IpamRouteTargetsListOKBody {
	return o.Payload
}

func (o *IpamRouteTargetsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamRouteTargetsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsListDefault creates a IpamRouteTargetsListDefault with default headers values
func NewIpamRouteTargetsListDefault(code int) *IpamRouteTargetsListDefault {
	return &IpamRouteTargetsListDefault{
		_statusCode: code,
	}
}

/* IpamRouteTargetsListDefault describes a response with status code -1, with default header values.

IpamRouteTargetsListDefault ipam route targets list default
*/
type IpamRouteTargetsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets list default response
func (o *IpamRouteTargetsListDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/route-targets/][%d] ipam_route-targets_list default  %+v", o._statusCode, o.Payload)
}
func (o *IpamRouteTargetsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IpamRouteTargetsListOKBody ipam route targets list o k body
swagger:model IpamRouteTargetsListOKBody
*/
type IpamRouteTargetsListOKBody struct {

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
	Results []*models.RouteTarget `json:"results"`
}

// Validate validates this ipam route targets list o k body
func (o *IpamRouteTargetsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *IpamRouteTargetsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamRouteTargetsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamRouteTargetsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamRouteTargetsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamRouteTargetsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamRouteTargetsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamRouteTargetsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamRouteTargetsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamRouteTargetsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamRouteTargetsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam route targets list o k body based on the context it is used
func (o *IpamRouteTargetsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamRouteTargetsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamRouteTargetsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamRouteTargetsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamRouteTargetsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamRouteTargetsListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamRouteTargetsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
