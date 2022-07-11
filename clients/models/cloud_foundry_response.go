// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudFoundryResponse cloud foundry response
// swagger:model CloudFoundryResponse

type CloudFoundryResponse struct {

	// pagination
	Pagination CloudFoundryPagination `json:"pagination"`

	// resources
	Resources CloudFoundryResponseResources `json:"resources"`
}

/* polymorph CloudFoundryResponse pagination false */

/* polymorph CloudFoundryResponse resources false */

// Validate validates this cloud foundry response
func (m *CloudFoundryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CloudFoundryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudFoundryResponse) UnmarshalBinary(b []byte) error {
	var res CloudFoundryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
