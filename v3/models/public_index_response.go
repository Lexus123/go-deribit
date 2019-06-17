// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicIndexResponse public index response
// swagger:model public_index_response
type PublicIndexResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *PublicIndexResponseAO1Result `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PublicIndexResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *PublicIndexResponseAO1Result `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PublicIndexResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *PublicIndexResponseAO1Result `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this public index response
func (m *PublicIndexResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIndexResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIndexResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIndexResponse) UnmarshalBinary(b []byte) error {
	var res PublicIndexResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicIndexResponseAO1Result public index response a o1 result
// swagger:model PublicIndexResponseAO1Result
type PublicIndexResponseAO1Result struct {

	// The current index price for BTC-USD (only for selected currency == BTC)
	// Required: true
	BTC *float64 `json:"BTC"`

	// The current index price for ETH-USD (only for selected currency == ETH)
	ETH float64 `json:"ETH,omitempty"`

	// Estimated delivery price for the currency. For more details, see Documentation > General > Expiration Price
	// Required: true
	Edp *float64 `json:"edp"`
}

// Validate validates this public index response a o1 result
func (m *PublicIndexResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBTC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIndexResponseAO1Result) validateBTC(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"BTC", "body", m.BTC); err != nil {
		return err
	}

	return nil
}

func (m *PublicIndexResponseAO1Result) validateEdp(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"edp", "body", m.Edp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIndexResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIndexResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res PublicIndexResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
