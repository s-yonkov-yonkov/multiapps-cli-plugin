// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudFoundryResourceMetadata cloud foundry resource metadata
// swagger:model CloudFoundryResource_metadata

type CloudFoundryResourceMetadata struct {

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

/* polymorph CloudFoundryResource_metadata labels false */

/* polymorph CloudFoundryResource_metadata annotations false */

// Validate validates this cloud foundry resource metadata
func (m *CloudFoundryResourceMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CloudFoundryResourceMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudFoundryResourceMetadata) UnmarshalBinary(b []byte) error {
	var res CloudFoundryResourceMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
