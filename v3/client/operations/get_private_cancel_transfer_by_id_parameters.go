// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateCancelTransferByIDParams creates a new GetPrivateCancelTransferByIDParams object
// with the default values initialized.
func NewGetPrivateCancelTransferByIDParams() *GetPrivateCancelTransferByIDParams {
	var ()
	return &GetPrivateCancelTransferByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateCancelTransferByIDParamsWithTimeout creates a new GetPrivateCancelTransferByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateCancelTransferByIDParamsWithTimeout(timeout time.Duration) *GetPrivateCancelTransferByIDParams {
	var ()
	return &GetPrivateCancelTransferByIDParams{

		timeout: timeout,
	}
}

// NewGetPrivateCancelTransferByIDParamsWithContext creates a new GetPrivateCancelTransferByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateCancelTransferByIDParamsWithContext(ctx context.Context) *GetPrivateCancelTransferByIDParams {
	var ()
	return &GetPrivateCancelTransferByIDParams{

		Context: ctx,
	}
}

// NewGetPrivateCancelTransferByIDParamsWithHTTPClient creates a new GetPrivateCancelTransferByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateCancelTransferByIDParamsWithHTTPClient(client *http.Client) *GetPrivateCancelTransferByIDParams {
	var ()
	return &GetPrivateCancelTransferByIDParams{
		HTTPClient: client,
	}
}

/*GetPrivateCancelTransferByIDParams contains all the parameters to send to the API endpoint
for the get private cancel transfer by ID operation typically these are written to a http.Request
*/
type GetPrivateCancelTransferByIDParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string
	/*ID
	  Id of transfer

	*/
	ID int64
	/*Tfa
	  TFA code, required when TFA is enabled for current account

	*/
	Tfa *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithTimeout(timeout time.Duration) *GetPrivateCancelTransferByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithContext(ctx context.Context) *GetPrivateCancelTransferByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithHTTPClient(client *http.Client) *GetPrivateCancelTransferByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithCurrency(currency string) *GetPrivateCancelTransferByIDParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithID adds the id to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithID(id int64) *GetPrivateCancelTransferByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetID(id int64) {
	o.ID = id
}

// WithTfa adds the tfa to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) WithTfa(tfa *string) *GetPrivateCancelTransferByIDParams {
	o.SetTfa(tfa)
	return o
}

// SetTfa adds the tfa to the get private cancel transfer by ID params
func (o *GetPrivateCancelTransferByIDParams) SetTfa(tfa *string) {
	o.Tfa = tfa
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateCancelTransferByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
