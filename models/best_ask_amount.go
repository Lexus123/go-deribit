// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// BestAskAmount It represents the requested order size of all best asks
//
// swagger:model best_ask_amount
type BestAskAmount float64

// Validate validates this best ask amount
func (m BestAskAmount) Validate(formats strfmt.Registry) error {
	return nil
}