// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeribitPriceRankingNotification deribit price ranking notification
//
// swagger:model deribit_price_ranking_notification
type DeribitPriceRankingNotification []*DeribitPriceRankingNotificationItems0

// Validate validates this deribit price ranking notification
func (m DeribitPriceRankingNotification) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// DeribitPriceRankingNotificationItems0 deribit price ranking notification items0
//
// swagger:model DeribitPriceRankingNotificationItems0
type DeribitPriceRankingNotificationItems0 struct {

	// Stock exchange status
	Enabled bool `json:"enabled,omitempty"`

	// Stock exchange identifier
	Identifier string `json:"identifier,omitempty"`

	// Stock exchange index price
	Price float64 `json:"price,omitempty"`

	// The timestamp of the last update from stock exchange
	Timestamp int64 `json:"timestamp,omitempty"`

	// The weight of the ranking given in percent
	Weight float64 `json:"weight,omitempty"`
}

// Validate validates this deribit price ranking notification items0
func (m *DeribitPriceRankingNotificationItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeribitPriceRankingNotificationItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeribitPriceRankingNotificationItems0) UnmarshalBinary(b []byte) error {
	var res DeribitPriceRankingNotificationItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
