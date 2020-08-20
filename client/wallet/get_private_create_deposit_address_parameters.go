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

// NewGetPrivateCreateDepositAddressParams creates a new GetPrivateCreateDepositAddressParams object
// with the default values initialized.
func NewGetPrivateCreateDepositAddressParams() *GetPrivateCreateDepositAddressParams {
	var ()
	return &GetPrivateCreateDepositAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateCreateDepositAddressParamsWithTimeout creates a new GetPrivateCreateDepositAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateCreateDepositAddressParamsWithTimeout(timeout time.Duration) *GetPrivateCreateDepositAddressParams {
	var ()
	return &GetPrivateCreateDepositAddressParams{

		timeout: timeout,
	}
}

// NewGetPrivateCreateDepositAddressParamsWithContext creates a new GetPrivateCreateDepositAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateCreateDepositAddressParamsWithContext(ctx context.Context) *GetPrivateCreateDepositAddressParams {
	var ()
	return &GetPrivateCreateDepositAddressParams{

		Context: ctx,
	}
}

// NewGetPrivateCreateDepositAddressParamsWithHTTPClient creates a new GetPrivateCreateDepositAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateCreateDepositAddressParamsWithHTTPClient(client *http.Client) *GetPrivateCreateDepositAddressParams {
	var ()
	return &GetPrivateCreateDepositAddressParams{
		HTTPClient: client,
	}
}

/*GetPrivateCreateDepositAddressParams contains all the parameters to send to the API endpoint
for the get private create deposit address operation typically these are written to a http.Request
*/
type GetPrivateCreateDepositAddressParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) WithTimeout(timeout time.Duration) *GetPrivateCreateDepositAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) WithContext(ctx context.Context) *GetPrivateCreateDepositAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) WithHTTPClient(client *http.Client) *GetPrivateCreateDepositAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) WithCurrency(currency string) *GetPrivateCreateDepositAddressParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private create deposit address params
func (o *GetPrivateCreateDepositAddressParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateCreateDepositAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}