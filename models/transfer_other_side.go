// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// TransferOtherSide For transfer from/to subaccount returns this subaccount name, for transfer to other account returns address, for transfer from other account returns that accounts username.
//
// swagger:model transfer_other_side
type TransferOtherSide string

// Validate validates this transfer other side
func (m TransferOtherSide) Validate(formats strfmt.Registry) error {
	return nil
}
