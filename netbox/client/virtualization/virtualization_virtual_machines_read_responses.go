// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationVirtualMachinesReadReader is a Reader for the VirtualizationVirtualMachinesRead structure.
type VirtualizationVirtualMachinesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesReadOK creates a VirtualizationVirtualMachinesReadOK with default headers values
func NewVirtualizationVirtualMachinesReadOK() *VirtualizationVirtualMachinesReadOK {
	return &VirtualizationVirtualMachinesReadOK{}
}

/* VirtualizationVirtualMachinesReadOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesReadOK virtualization virtual machines read o k
*/
type VirtualizationVirtualMachinesReadOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesReadOK  %+v", 200, o.Payload)
}
func (o *VirtualizationVirtualMachinesReadOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesReadDefault creates a VirtualizationVirtualMachinesReadDefault with default headers values
func NewVirtualizationVirtualMachinesReadDefault(code int) *VirtualizationVirtualMachinesReadDefault {
	return &VirtualizationVirtualMachinesReadDefault{
		_statusCode: code,
	}
}

/* VirtualizationVirtualMachinesReadDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesReadDefault virtualization virtual machines read default
*/
type VirtualizationVirtualMachinesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines read default response
func (o *VirtualizationVirtualMachinesReadDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesReadDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_read default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationVirtualMachinesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
