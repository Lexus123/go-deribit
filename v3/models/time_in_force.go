// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TimeInForce Order time in force: `"good_til_cancelled"`, `"fill_or_kill"`, `"immediate_or_cancel"`
// swagger:model time_in_force
type TimeInForce string

const (

	// TimeInForceGoodTilCancelled captures enum value "good_til_cancelled"
	TimeInForceGoodTilCancelled TimeInForce = "good_til_cancelled"

	// TimeInForceFillOrKill captures enum value "fill_or_kill"
	TimeInForceFillOrKill TimeInForce = "fill_or_kill"

	// TimeInForceImmediateOrCancel captures enum value "immediate_or_cancel"
	TimeInForceImmediateOrCancel TimeInForce = "immediate_or_cancel"
)

// for schema
var timeInForceEnum []interface{}

func init() {
	var res []TimeInForce
	if err := json.Unmarshal([]byte(`["good_til_cancelled","fill_or_kill","immediate_or_cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeInForceEnum = append(timeInForceEnum, v)
	}
}

func (m TimeInForce) validateTimeInForceEnum(path, location string, value TimeInForce) error {
	if err := validate.Enum(path, location, value, timeInForceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this time in force
func (m TimeInForce) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTimeInForceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
