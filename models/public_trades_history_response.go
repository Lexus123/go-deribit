// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicTradesHistoryResponse public trades history response
//
// swagger:model public_trades_history_response
type PublicTradesHistoryResponse struct {

	// result
	// Required: true
	Result *PublicTradesHistoryResponseResult `json:"result"`
}

// Validate validates this public trades history response
func (m *PublicTradesHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicTradesHistoryResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PublicTradesHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTradesHistoryResponse) UnmarshalBinary(b []byte) error {
	var res PublicTradesHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicTradesHistoryResponseResult public trades history response result
//
// swagger:model PublicTradesHistoryResponseResult
type PublicTradesHistoryResponseResult struct {

	// has more
	// Required: true
	HasMore *bool `json:"has_more"`

	// trades
	// Required: true
	Trades []*PublicTrade `json:"trades"`
}

// Validate validates this public trades history response result
func (m *PublicTradesHistoryResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasMore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrades(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicTradesHistoryResponseResult) validateHasMore(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"has_more", "body", m.HasMore); err != nil {
		return err
	}

	return nil
}

func (m *PublicTradesHistoryResponseResult) validateTrades(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"trades", "body", m.Trades); err != nil {
		return err
	}

	for i := 0; i < len(m.Trades); i++ {
		if swag.IsZero(m.Trades[i]) { // not required
			continue
		}

		if m.Trades[i] != nil {
			if err := m.Trades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + "trades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicTradesHistoryResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTradesHistoryResponseResult) UnmarshalBinary(b []byte) error {
	var res PublicTradesHistoryResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
