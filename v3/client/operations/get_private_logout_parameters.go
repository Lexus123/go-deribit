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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateLogoutParams creates a new GetPrivateLogoutParams object
// with the default values initialized.
func NewGetPrivateLogoutParams() *GetPrivateLogoutParams {

	return &GetPrivateLogoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateLogoutParamsWithTimeout creates a new GetPrivateLogoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateLogoutParamsWithTimeout(timeout time.Duration) *GetPrivateLogoutParams {

	return &GetPrivateLogoutParams{

		timeout: timeout,
	}
}

// NewGetPrivateLogoutParamsWithContext creates a new GetPrivateLogoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateLogoutParamsWithContext(ctx context.Context) *GetPrivateLogoutParams {

	return &GetPrivateLogoutParams{

		Context: ctx,
	}
}

// NewGetPrivateLogoutParamsWithHTTPClient creates a new GetPrivateLogoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateLogoutParamsWithHTTPClient(client *http.Client) *GetPrivateLogoutParams {

	return &GetPrivateLogoutParams{
		HTTPClient: client,
	}
}

/*GetPrivateLogoutParams contains all the parameters to send to the API endpoint
for the get private logout operation typically these are written to a http.Request
*/
type GetPrivateLogoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private logout params
func (o *GetPrivateLogoutParams) WithTimeout(timeout time.Duration) *GetPrivateLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private logout params
func (o *GetPrivateLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private logout params
func (o *GetPrivateLogoutParams) WithContext(ctx context.Context) *GetPrivateLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private logout params
func (o *GetPrivateLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private logout params
func (o *GetPrivateLogoutParams) WithHTTPClient(client *http.Client) *GetPrivateLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private logout params
func (o *GetPrivateLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
