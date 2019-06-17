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

// TransferType Type of transfer: `user` - sent to user, `subaccount` - sent to subaccount
// swagger:model transfer_type
type TransferType string

const (

	// TransferTypeUser captures enum value "user"
	TransferTypeUser TransferType = "user"

	// TransferTypeSubaccount captures enum value "subaccount"
	TransferTypeSubaccount TransferType = "subaccount"
)

// for schema
var transferTypeEnum []interface{}

func init() {
	var res []TransferType
	if err := json.Unmarshal([]byte(`["user","subaccount"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transferTypeEnum = append(transferTypeEnum, v)
	}
}

func (m TransferType) validateTransferTypeEnum(path, location string, value TransferType) error {
	if err := validate.Enum(path, location, value, transferTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this transfer type
func (m TransferType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransferTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
