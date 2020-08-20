// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateTradehistoryResponse private tradehistory response
//
// swagger:model private_tradehistory_response
type PrivateTradehistoryResponse struct {

	// result
	// Required: true
	Result *PrivateTradehistoryResponseResult `json:"result"`
}

// Validate validates this private tradehistory response
func (m *PrivateTradehistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateTradehistoryResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateTradehistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateTradehistoryResponse) UnmarshalBinary(b []byte) error {
	var res PrivateTradehistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateTradehistoryResponseResult private tradehistory response result
//
// swagger:model PrivateTradehistoryResponseResult
type PrivateTradehistoryResponseResult struct {

	// Trade direction of the taker
	// Required: true
	Direction *string `json:"direction"`

	// Index price at trade
	// Required: true
	IndexPrice *string `json:"indexPrice"`

	// The name of the instrument
	// Required: true
	Instrument *string `json:"instrument"`

	// option implied volatility for the price (Options only)
	// Required: true
	Iv *string `json:"iv"`

	// The price of the trade
	// Required: true
	Price *string `json:"price"`

	// The quantity traded
	// Required: true
	Quantity *string `json:"quantity"`

	// Direction of the "tick".
	// Required: true
	TickDirection *string `json:"tickDirection"`

	// The timestamp of the trade in ms
	// Required: true
	TimeStamp *int64 `json:"timeStamp"`

	// The ID for the trade
	// Required: true
	TradeID *string `json:"tradeId"`

	// The trade sequence number
	// Required: true
	TradeSeq *string `json:"tradeSeq"`
}

// Validate validates this private tradehistory response result
func (m *PrivateTradehistoryResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTickDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeSeq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateTradehistoryResponseResult) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateIndexPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"indexPrice", "body", m.IndexPrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateInstrument(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"instrument", "body", m.Instrument); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateIv(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"iv", "body", m.Iv); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateTickDirection(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"tickDirection", "body", m.TickDirection); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateTimeStamp(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"timeStamp", "body", m.TimeStamp); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateTradeID(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"tradeId", "body", m.TradeID); err != nil {
		return err
	}

	return nil
}

func (m *PrivateTradehistoryResponseResult) validateTradeSeq(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"tradeSeq", "body", m.TradeSeq); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateTradehistoryResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateTradehistoryResponseResult) UnmarshalBinary(b []byte) error {
	var res PrivateTradehistoryResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
