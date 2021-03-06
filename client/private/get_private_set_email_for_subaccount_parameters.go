// Code generated by go-swagger; DO NOT EDIT.

package private

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
	"github.com/go-openapi/swag"
)

// NewGetPrivateSetEmailForSubaccountParams creates a new GetPrivateSetEmailForSubaccountParams object
// with the default values initialized.
func NewGetPrivateSetEmailForSubaccountParams() *GetPrivateSetEmailForSubaccountParams {
	var ()
	return &GetPrivateSetEmailForSubaccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateSetEmailForSubaccountParamsWithTimeout creates a new GetPrivateSetEmailForSubaccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateSetEmailForSubaccountParamsWithTimeout(timeout time.Duration) *GetPrivateSetEmailForSubaccountParams {
	var ()
	return &GetPrivateSetEmailForSubaccountParams{

		timeout: timeout,
	}
}

// NewGetPrivateSetEmailForSubaccountParamsWithContext creates a new GetPrivateSetEmailForSubaccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateSetEmailForSubaccountParamsWithContext(ctx context.Context) *GetPrivateSetEmailForSubaccountParams {
	var ()
	return &GetPrivateSetEmailForSubaccountParams{

		Context: ctx,
	}
}

// NewGetPrivateSetEmailForSubaccountParamsWithHTTPClient creates a new GetPrivateSetEmailForSubaccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateSetEmailForSubaccountParamsWithHTTPClient(client *http.Client) *GetPrivateSetEmailForSubaccountParams {
	var ()
	return &GetPrivateSetEmailForSubaccountParams{
		HTTPClient: client,
	}
}

/*GetPrivateSetEmailForSubaccountParams contains all the parameters to send to the API endpoint
for the get private set email for subaccount operation typically these are written to a http.Request
*/
type GetPrivateSetEmailForSubaccountParams struct {

	/*Email
	  The email address for the subaccount

	*/
	Email string
	/*Sid
	  The user id for the subaccount

	*/
	Sid int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) WithTimeout(timeout time.Duration) *GetPrivateSetEmailForSubaccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) WithContext(ctx context.Context) *GetPrivateSetEmailForSubaccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) WithHTTPClient(client *http.Client) *GetPrivateSetEmailForSubaccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) WithEmail(email string) *GetPrivateSetEmailForSubaccountParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) SetEmail(email string) {
	o.Email = email
}

// WithSid adds the sid to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) WithSid(sid int64) *GetPrivateSetEmailForSubaccountParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the get private set email for subaccount params
func (o *GetPrivateSetEmailForSubaccountParams) SetSid(sid int64) {
	o.Sid = sid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSetEmailForSubaccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param email
	qrEmail := o.Email
	qEmail := qrEmail
	if qEmail != "" {
		if err := r.SetQueryParam("email", qEmail); err != nil {
			return err
		}
	}

	// query param sid
	qrSid := o.Sid
	qSid := swag.FormatInt64(qrSid)
	if qSid != "" {
		if err := r.SetQueryParam("sid", qSid); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
