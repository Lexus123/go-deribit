// Code generated by go-swagger; DO NOT EDIT.

package websocket_only

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

// NewGetPublicUnsubscribeParams creates a new GetPublicUnsubscribeParams object
// with the default values initialized.
func NewGetPublicUnsubscribeParams() *GetPublicUnsubscribeParams {
	var ()
	return &GetPublicUnsubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicUnsubscribeParamsWithTimeout creates a new GetPublicUnsubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicUnsubscribeParamsWithTimeout(timeout time.Duration) *GetPublicUnsubscribeParams {
	var ()
	return &GetPublicUnsubscribeParams{

		timeout: timeout,
	}
}

// NewGetPublicUnsubscribeParamsWithContext creates a new GetPublicUnsubscribeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicUnsubscribeParamsWithContext(ctx context.Context) *GetPublicUnsubscribeParams {
	var ()
	return &GetPublicUnsubscribeParams{

		Context: ctx,
	}
}

// NewGetPublicUnsubscribeParamsWithHTTPClient creates a new GetPublicUnsubscribeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicUnsubscribeParamsWithHTTPClient(client *http.Client) *GetPublicUnsubscribeParams {
	var ()
	return &GetPublicUnsubscribeParams{
		HTTPClient: client,
	}
}

/*GetPublicUnsubscribeParams contains all the parameters to send to the API endpoint
for the get public unsubscribe operation typically these are written to a http.Request
*/
type GetPublicUnsubscribeParams struct {

	/*Channels
	  A list of channels to unsubscribe from.

	*/
	Channels []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) WithTimeout(timeout time.Duration) *GetPublicUnsubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) WithContext(ctx context.Context) *GetPublicUnsubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) WithHTTPClient(client *http.Client) *GetPublicUnsubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannels adds the channels to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) WithChannels(channels []string) *GetPublicUnsubscribeParams {
	o.SetChannels(channels)
	return o
}

// SetChannels adds the channels to the get public unsubscribe params
func (o *GetPublicUnsubscribeParams) SetChannels(channels []string) {
	o.Channels = channels
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicUnsubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesChannels := o.Channels

	joinedChannels := swag.JoinByFormat(valuesChannels, "multi")
	// query array param channels
	if err := r.SetQueryParam("channels", joinedChannels...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}