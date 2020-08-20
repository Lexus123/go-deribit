// Code generated by go-swagger; DO NOT EDIT.

package trading

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

// NewGetPrivateCancelAllByCurrencyParams creates a new GetPrivateCancelAllByCurrencyParams object
// with the default values initialized.
func NewGetPrivateCancelAllByCurrencyParams() *GetPrivateCancelAllByCurrencyParams {
	var ()
	return &GetPrivateCancelAllByCurrencyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateCancelAllByCurrencyParamsWithTimeout creates a new GetPrivateCancelAllByCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateCancelAllByCurrencyParamsWithTimeout(timeout time.Duration) *GetPrivateCancelAllByCurrencyParams {
	var ()
	return &GetPrivateCancelAllByCurrencyParams{

		timeout: timeout,
	}
}

// NewGetPrivateCancelAllByCurrencyParamsWithContext creates a new GetPrivateCancelAllByCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateCancelAllByCurrencyParamsWithContext(ctx context.Context) *GetPrivateCancelAllByCurrencyParams {
	var ()
	return &GetPrivateCancelAllByCurrencyParams{

		Context: ctx,
	}
}

// NewGetPrivateCancelAllByCurrencyParamsWithHTTPClient creates a new GetPrivateCancelAllByCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateCancelAllByCurrencyParamsWithHTTPClient(client *http.Client) *GetPrivateCancelAllByCurrencyParams {
	var ()
	return &GetPrivateCancelAllByCurrencyParams{
		HTTPClient: client,
	}
}

/*GetPrivateCancelAllByCurrencyParams contains all the parameters to send to the API endpoint
for the get private cancel all by currency operation typically these are written to a http.Request
*/
type GetPrivateCancelAllByCurrencyParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Kind
	  Instrument kind, if not provided instruments of all kinds are considered

	*/
	Kind *string
	/*Type
	  Order type - limit, stop or all, default - `all`

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithTimeout(timeout time.Duration) *GetPrivateCancelAllByCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithContext(ctx context.Context) *GetPrivateCancelAllByCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithHTTPClient(client *http.Client) *GetPrivateCancelAllByCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithCurrency(currency string) *GetPrivateCancelAllByCurrencyParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithKind adds the kind to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithKind(kind *string) *GetPrivateCancelAllByCurrencyParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithType adds the typeVar to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) WithType(typeVar *string) *GetPrivateCancelAllByCurrencyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private cancel all by currency params
func (o *GetPrivateCancelAllByCurrencyParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateCancelAllByCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}