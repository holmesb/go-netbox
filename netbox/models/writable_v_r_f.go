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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableVRF writable v r f
//
// swagger:model WritableVRF
type WritableVRF struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Enforce unique space
	//
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique bool `json:"enforce_unique,omitempty"`

	// export targets
	// Unique: true
	ExportTargets []int64 `json:"export_targets"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// import targets
	// Unique: true
	ImportTargets []int64 `json:"import_targets"`

	// Ipaddress count
	// Read Only: true
	IpaddressCount int64 `json:"ipaddress_count,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Prefix count
	// Read Only: true
	PrefixCount int64 `json:"prefix_count,omitempty"`

	// Route distinguisher
	//
	// Unique route distinguisher (as defined in RFC 4364)
	// Max Length: 21
	Rd *string `json:"rd,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable v r f
func (m *WritableVRF) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVRF) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateExportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("export_targets", "body", m.ExportTargets); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateImportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("import_targets", "body", m.ImportTargets); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateRd(formats strfmt.Registry) error {
	if swag.IsZero(m.Rd) { // not required
		return nil
	}

	if err := validate.MaxLength("rd", "body", *m.Rd, 21); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableVRF) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable v r f based on the context it is used
func (m *WritableVRF) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpaddressCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefixCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVRF) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidateIpaddressCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ipaddress_count", "body", int64(m.IpaddressCount)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidatePrefixCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prefix_count", "body", int64(m.PrefixCount)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVRF) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableVRF) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVRF) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVRF) UnmarshalBinary(b []byte) error {
	var res WritableVRF
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
