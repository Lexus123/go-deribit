// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// Quantity The number of contracts to be traded.
//
// swagger:model quantity
type Quantity float64

// Validate validates this quantity
func (m Quantity) Validate(formats strfmt.Registry) error {
	return nil
}