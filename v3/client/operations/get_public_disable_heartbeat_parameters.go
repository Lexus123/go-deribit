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

// NewGetPublicDisableHeartbeatParams creates a new GetPublicDisableHeartbeatParams object
// with the default values initialized.
func NewGetPublicDisableHeartbeatParams() *GetPublicDisableHeartbeatParams {

	return &GetPublicDisableHeartbeatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicDisableHeartbeatParamsWithTimeout creates a new GetPublicDisableHeartbeatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicDisableHeartbeatParamsWithTimeout(timeout time.Duration) *GetPublicDisableHeartbeatParams {

	return &GetPublicDisableHeartbeatParams{

		timeout: timeout,
	}
}

// NewGetPublicDisableHeartbeatParamsWithContext creates a new GetPublicDisableHeartbeatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicDisableHeartbeatParamsWithContext(ctx context.Context) *GetPublicDisableHeartbeatParams {

	return &GetPublicDisableHeartbeatParams{

		Context: ctx,
	}
}

// NewGetPublicDisableHeartbeatParamsWithHTTPClient creates a new GetPublicDisableHeartbeatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicDisableHeartbeatParamsWithHTTPClient(client *http.Client) *GetPublicDisableHeartbeatParams {

	return &GetPublicDisableHeartbeatParams{
		HTTPClient: client,
	}
}

/*GetPublicDisableHeartbeatParams contains all the parameters to send to the API endpoint
for the get public disable heartbeat operation typically these are written to a http.Request
*/
type GetPublicDisableHeartbeatParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) WithTimeout(timeout time.Duration) *GetPublicDisableHeartbeatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) WithContext(ctx context.Context) *GetPublicDisableHeartbeatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) WithHTTPClient(client *http.Client) *GetPublicDisableHeartbeatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public disable heartbeat params
func (o *GetPublicDisableHeartbeatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicDisableHeartbeatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
