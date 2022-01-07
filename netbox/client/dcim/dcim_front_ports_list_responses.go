// Code generated by go-swagger; DO NOT EDIT.

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimFrontPortsListReader is a Reader for the DcimFrontPortsList structure.
type DcimFrontPortsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsListOK creates a DcimFrontPortsListOK with default headers values
func NewDcimFrontPortsListOK() *DcimFrontPortsListOK {
	return &DcimFrontPortsListOK{}
}

/* DcimFrontPortsListOK describes a response with status code 200, with default header values.

DcimFrontPortsListOK dcim front ports list o k
*/
type DcimFrontPortsListOK struct {
	Payload *DcimFrontPortsListOKBody
}

func (o *DcimFrontPortsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/][%d] dcimFrontPortsListOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsListOK) GetPayload() *DcimFrontPortsListOKBody {
	return o.Payload
}

func (o *DcimFrontPortsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimFrontPortsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsListDefault creates a DcimFrontPortsListDefault with default headers values
func NewDcimFrontPortsListDefault(code int) *DcimFrontPortsListDefault {
	return &DcimFrontPortsListDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortsListDefault describes a response with status code -1, with default header values.

DcimFrontPortsListDefault dcim front ports list default
*/
type DcimFrontPortsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports list default response
func (o *DcimFrontPortsListDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/][%d] dcim_front-ports_list default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DcimFrontPortsListOKBody dcim front ports list o k body
swagger:model DcimFrontPortsListOKBody
*/
type DcimFrontPortsListOKBody struct {

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
	Results []*models.FrontPort `json:"results"`
}

// Validate validates this dcim front ports list o k body
func (o *DcimFrontPortsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimFrontPortsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimFrontPortsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimFrontPortsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimFrontPortsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimFrontPortsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimFrontPortsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimFrontPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimFrontPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim front ports list o k body based on the context it is used
func (o *DcimFrontPortsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimFrontPortsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimFrontPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimFrontPortsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimFrontPortsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimFrontPortsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimFrontPortsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
