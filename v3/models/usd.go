// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// Usd Option price in USD (Only if `advanced="usd"`)
// swagger:model usd
type Usd float64

// Validate validates this usd
func (m Usd) Validate(formats strfmt.Registry) error {
	return nil
}
