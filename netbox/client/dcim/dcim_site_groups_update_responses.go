// Code generated by go-swagger; DO NOT EDIT.

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

// DcimSiteGroupsUpdateReader is a Reader for the DcimSiteGroupsUpdate structure.
type DcimSiteGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSiteGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsUpdateOK creates a DcimSiteGroupsUpdateOK with default headers values
func NewDcimSiteGroupsUpdateOK() *DcimSiteGroupsUpdateOK {
	return &DcimSiteGroupsUpdateOK{}
}

/* DcimSiteGroupsUpdateOK describes a response with status code 200, with default header values.

DcimSiteGroupsUpdateOK dcim site groups update o k
*/
type DcimSiteGroupsUpdateOK struct {
	Payload *models.SiteGroup
}

func (o *DcimSiteGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/{id}/][%d] dcimSiteGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSiteGroupsUpdateOK) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsUpdateDefault creates a DcimSiteGroupsUpdateDefault with default headers values
func NewDcimSiteGroupsUpdateDefault(code int) *DcimSiteGroupsUpdateDefault {
	return &DcimSiteGroupsUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSiteGroupsUpdateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsUpdateDefault dcim site groups update default
*/
type DcimSiteGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups update default response
func (o *DcimSiteGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/{id}/][%d] dcim_site-groups_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSiteGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
