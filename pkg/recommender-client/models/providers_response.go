// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProvidersResponse ProvidersResponse ProvidersResponse is the response used for the supported providers
// swagger:model ProvidersResponse
type ProvidersResponse struct {

	// providers
	Providers []*Provider `json:"providers"`
}

// Validate validates this providers response
func (m *ProvidersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersResponse) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersResponse) UnmarshalBinary(b []byte) error {
	var res ProvidersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
