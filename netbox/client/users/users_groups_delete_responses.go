// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersGroupsDeleteReader is a Reader for the UsersGroupsDelete structure.
type UsersGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsDeleteNoContent creates a UsersGroupsDeleteNoContent with default headers values
func NewUsersGroupsDeleteNoContent() *UsersGroupsDeleteNoContent {
	return &UsersGroupsDeleteNoContent{}
}

/* UsersGroupsDeleteNoContent describes a response with status code 204, with default header values.

UsersGroupsDeleteNoContent users groups delete no content
*/
type UsersGroupsDeleteNoContent struct {
}

func (o *UsersGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/{id}/][%d] usersGroupsDeleteNoContent ", 204)
}

func (o *UsersGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersGroupsDeleteDefault creates a UsersGroupsDeleteDefault with default headers values
func NewUsersGroupsDeleteDefault(code int) *UsersGroupsDeleteDefault {
	return &UsersGroupsDeleteDefault{
		_statusCode: code,
	}
}

/* UsersGroupsDeleteDefault describes a response with status code -1, with default header values.

UsersGroupsDeleteDefault users groups delete default
*/
type UsersGroupsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups delete default response
func (o *UsersGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UsersGroupsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/{id}/][%d] users_groups_delete default  %+v", o._statusCode, o.Payload)
}
func (o *UsersGroupsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
