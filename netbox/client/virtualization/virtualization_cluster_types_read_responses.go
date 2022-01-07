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

// VirtualizationClusterTypesReadReader is a Reader for the VirtualizationClusterTypesRead structure.
type VirtualizationClusterTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesReadOK creates a VirtualizationClusterTypesReadOK with default headers values
func NewVirtualizationClusterTypesReadOK() *VirtualizationClusterTypesReadOK {
	return &VirtualizationClusterTypesReadOK{}
}

/* VirtualizationClusterTypesReadOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesReadOK virtualization cluster types read o k
*/
type VirtualizationClusterTypesReadOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesReadOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterTypesReadOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesReadDefault creates a VirtualizationClusterTypesReadDefault with default headers values
func NewVirtualizationClusterTypesReadDefault(code int) *VirtualizationClusterTypesReadDefault {
	return &VirtualizationClusterTypesReadDefault{
		_statusCode: code,
	}
}

/* VirtualizationClusterTypesReadDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesReadDefault virtualization cluster types read default
*/
type VirtualizationClusterTypesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types read default response
func (o *VirtualizationClusterTypesReadDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesReadDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-types/{id}/][%d] virtualization_cluster-types_read default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationClusterTypesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
