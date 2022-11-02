// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdmissionPlugin AdmissionPlugin represents an admission plugin
//
// swagger:model AdmissionPlugin
type AdmissionPlugin struct {

	// name
	Name string `json:"name,omitempty"`

	// plugin
	Plugin string `json:"plugin,omitempty"`

	// from version
	FromVersion Semver `json:"fromVersion,omitempty"`
}

// Validate validates this admission plugin
func (m *AdmissionPlugin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdmissionPlugin) validateFromVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.FromVersion) { // not required
		return nil
	}

	if err := m.FromVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fromVersion")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fromVersion")
		}
		return err
	}

	return nil
}

// ContextValidate validate this admission plugin based on the context it is used
func (m *AdmissionPlugin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFromVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdmissionPlugin) contextValidateFromVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FromVersion.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fromVersion")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fromVersion")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdmissionPlugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdmissionPlugin) UnmarshalBinary(b []byte) error {
	var res AdmissionPlugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}