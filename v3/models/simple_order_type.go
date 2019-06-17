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

// SimpleOrderType Order type, `"all"`, `"limit"`, `"stop"`
// swagger:model simple_order_type
type SimpleOrderType string

const (

	// SimpleOrderTypeAll captures enum value "all"
	SimpleOrderTypeAll SimpleOrderType = "all"

	// SimpleOrderTypeLimit captures enum value "limit"
	SimpleOrderTypeLimit SimpleOrderType = "limit"

	// SimpleOrderTypeStop captures enum value "stop"
	SimpleOrderTypeStop SimpleOrderType = "stop"
)

// for schema
var simpleOrderTypeEnum []interface{}

func init() {
	var res []SimpleOrderType
	if err := json.Unmarshal([]byte(`["all","limit","stop"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simpleOrderTypeEnum = append(simpleOrderTypeEnum, v)
	}
}

func (m SimpleOrderType) validateSimpleOrderTypeEnum(path, location string, value SimpleOrderType) error {
	if err := validate.Enum(path, location, value, simpleOrderTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this simple order type
func (m SimpleOrderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSimpleOrderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
