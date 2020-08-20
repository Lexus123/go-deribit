// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Deposit deposit
//
// swagger:model deposit
type Deposit struct {

	// address
	// Required: true
	Address CurrencyAddress `json:"address"`

	// amount
	// Required: true
	Amount CurrencyAmount `json:"amount"`

	// currency
	// Required: true
	Currency Currency `json:"currency"`

	// received timestamp
	// Required: true
	ReceivedTimestamp Timestamp `json:"received_timestamp"`

	// state
	// Required: true
	State DepositState `json:"state"`

	// transaction id
	// Required: true
	TransactionID CurrencyTransactionID `json:"transaction_id"`

	// updated timestamp
	// Required: true
	UpdatedTimestamp Timestamp `json:"updated_timestamp"`
}

// Validate validates this deposit
func (m *Deposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deposit) validateAddress(formats strfmt.Registry) error {

	if err := m.Address.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("address")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateCurrency(formats strfmt.Registry) error {

	if err := m.Currency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("currency")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateReceivedTimestamp(formats strfmt.Registry) error {

	if err := m.ReceivedTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("received_timestamp")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateTransactionID(formats strfmt.Registry) error {

	if err := m.TransactionID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transaction_id")
		}
		return err
	}

	return nil
}

func (m *Deposit) validateUpdatedTimestamp(formats strfmt.Registry) error {

	if err := m.UpdatedTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("updated_timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deposit) UnmarshalBinary(b []byte) error {
	var res Deposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}