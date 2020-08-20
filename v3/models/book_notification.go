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

// BookNotification book notification
//
// swagger:model book_notification
type BookNotification struct {

	// asks
	// Required: true
	Asks Asks `json:"asks"`

	// bids
	// Required: true
	Bids Bids `json:"bids"`

	// id of the notification
	// Required: true
	ChangeID *int64 `json:"change_id"`

	// instrument name
	// Required: true
	InstrumentName InstrumentName `json:"instrument_name"`

	// timestamp
	Timestamp TimestampForBookNotifications `json:"timestamp,omitempty"`
}

// Validate validates this book notification
func (m *BookNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BookNotification) validateAsks(formats strfmt.Registry) error {

	if err := validate.Required("asks", "body", m.Asks); err != nil {
		return err
	}

	if err := m.Asks.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("asks")
		}
		return err
	}

	return nil
}

func (m *BookNotification) validateBids(formats strfmt.Registry) error {

	if err := validate.Required("bids", "body", m.Bids); err != nil {
		return err
	}

	if err := m.Bids.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bids")
		}
		return err
	}

	return nil
}

func (m *BookNotification) validateChangeID(formats strfmt.Registry) error {

	if err := validate.Required("change_id", "body", m.ChangeID); err != nil {
		return err
	}

	return nil
}

func (m *BookNotification) validateInstrumentName(formats strfmt.Registry) error {

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

func (m *BookNotification) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := m.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BookNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookNotification) UnmarshalBinary(b []byte) error {
	var res BookNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
