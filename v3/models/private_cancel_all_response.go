// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateCancelAllResponse private cancel all response
//
// swagger:model private_cancel_all_response
type PrivateCancelAllResponse struct {
	BaseMessage

	PrivateCancelAllResponseAllOf1

	// Total number of successfully cancelled orders
	// Required: true
	Result *float64 `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrivateCancelAllResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var aO1 PrivateCancelAllResponseAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PrivateCancelAllResponseAllOf1 = aO1

	// AO2
	var dataAO2 struct {
		Result *float64 `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Result = dataAO2.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrivateCancelAllResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PrivateCancelAllResponseAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		Result *float64 `json:"result"`
	}

	dataAO2.Result = m.Result

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this private cancel all response
func (m *PrivateCancelAllResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PrivateCancelAllResponseAllOf1

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateCancelAllResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateCancelAllResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateCancelAllResponse) UnmarshalBinary(b []byte) error {
	var res PrivateCancelAllResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateCancelAllResponseAllOf1 private cancel all response all of1
//
// swagger:model PrivateCancelAllResponseAllOf1
type PrivateCancelAllResponseAllOf1 interface{}
