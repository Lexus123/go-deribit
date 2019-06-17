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

// AddressBookType Address book type
// swagger:model address_book_type
type AddressBookType string

const (

	// AddressBookTypeTransfer captures enum value "transfer"
	AddressBookTypeTransfer AddressBookType = "transfer"

	// AddressBookTypeWithdrawal captures enum value "withdrawal"
	AddressBookTypeWithdrawal AddressBookType = "withdrawal"
)

// for schema
var addressBookTypeEnum []interface{}

func init() {
	var res []AddressBookType
	if err := json.Unmarshal([]byte(`["transfer","withdrawal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addressBookTypeEnum = append(addressBookTypeEnum, v)
	}
}

func (m AddressBookType) validateAddressBookTypeEnum(path, location string, value AddressBookType) error {
	if err := validate.Enum(path, location, value, addressBookTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this address book type
func (m AddressBookType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAddressBookTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
