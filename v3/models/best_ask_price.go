// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// BestAskPrice The current best ask price, `null` if there aren't any asks
// swagger:model best_ask_price
type BestAskPrice float64

// Validate validates this best ask price
func (m BestAskPrice) Validate(formats strfmt.Registry) error {
	return nil
}
