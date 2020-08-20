// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPrivateAddToAddressBookParams creates a new GetPrivateAddToAddressBookParams object
// with the default values initialized.
func NewGetPrivateAddToAddressBookParams() *GetPrivateAddToAddressBookParams {
	var ()
	return &GetPrivateAddToAddressBookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateAddToAddressBookParamsWithTimeout creates a new GetPrivateAddToAddressBookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateAddToAddressBookParamsWithTimeout(timeout time.Duration) *GetPrivateAddToAddressBookParams {
	var ()
	return &GetPrivateAddToAddressBookParams{

		timeout: timeout,
	}
}

// NewGetPrivateAddToAddressBookParamsWithContext creates a new GetPrivateAddToAddressBookParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateAddToAddressBookParamsWithContext(ctx context.Context) *GetPrivateAddToAddressBookParams {
	var ()
	return &GetPrivateAddToAddressBookParams{

		Context: ctx,
	}
}

// NewGetPrivateAddToAddressBookParamsWithHTTPClient creates a new GetPrivateAddToAddressBookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateAddToAddressBookParamsWithHTTPClient(client *http.Client) *GetPrivateAddToAddressBookParams {
	var ()
	return &GetPrivateAddToAddressBookParams{
		HTTPClient: client,
	}
}

/*GetPrivateAddToAddressBookParams contains all the parameters to send to the API endpoint
for the get private add to address book operation typically these are written to a http.Request
*/
type GetPrivateAddToAddressBookParams struct {

	/*Address
	  Address in currency format, it must be in address book

	*/
	Address string
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Name
	  Name of address book entry

	*/
	Name string
	/*Tfa
	  TFA code, required when TFA is enabled for current account

	*/
	Tfa *string
	/*Type
	  Address book type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithTimeout(timeout time.Duration) *GetPrivateAddToAddressBookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithContext(ctx context.Context) *GetPrivateAddToAddressBookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithHTTPClient(client *http.Client) *GetPrivateAddToAddressBookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithAddress(address string) *GetPrivateAddToAddressBookParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetAddress(address string) {
	o.Address = address
}

// WithCurrency adds the currency to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithCurrency(currency string) *GetPrivateAddToAddressBookParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithName adds the name to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithName(name string) *GetPrivateAddToAddressBookParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetName(name string) {
	o.Name = name
}

// WithTfa adds the tfa to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithTfa(tfa *string) *GetPrivateAddToAddressBookParams {
	o.SetTfa(tfa)
	return o
}

// SetTfa adds the tfa to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetTfa(tfa *string) {
	o.Tfa = tfa
}

// WithType adds the typeVar to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) WithType(typeVar string) *GetPrivateAddToAddressBookParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private add to address book params
func (o *GetPrivateAddToAddressBookParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateAddToAddressBookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param address
	qrAddress := o.Address
	qAddress := qrAddress
	if qAddress != "" {
		if err := r.SetQueryParam("address", qAddress); err != nil {
			return err
		}
	}

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if o.Tfa != nil {

		// query param tfa
		var qrTfa string
		if o.Tfa != nil {
			qrTfa = *o.Tfa
		}
		qTfa := qrTfa
		if qTfa != "" {
			if err := r.SetQueryParam("tfa", qTfa); err != nil {
				return err
			}
		}

	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}