// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DepositState Deposit state, allowed values : `pending`, `completed`, `rejected`, `replaced`
//
// swagger:model deposit_state
type DepositState string

const (

	// DepositStatePending captures enum value "pending"
	DepositStatePending DepositState = "pending"

	// DepositStateCompleted captures enum value "completed"
	DepositStateCompleted DepositState = "completed"

	// DepositStateRejected captures enum value "rejected"
	DepositStateRejected DepositState = "rejected"

	// DepositStateReplaced captures enum value "replaced"
	DepositStateReplaced DepositState = "replaced"
)

// for schema
var depositStateEnum []interface{}

func init() {
	var res []DepositState
	if err := json.Unmarshal([]byte(`["pending","completed","rejected","replaced"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		depositStateEnum = append(depositStateEnum, v)
	}
}

func (m DepositState) validateDepositStateEnum(path, location string, value DepositState) error {
	if err := validate.EnumCase(path, location, value, depositStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this deposit state
func (m DepositState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDepositStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
