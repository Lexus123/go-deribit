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

// NewGetPrivateEditParams creates a new GetPrivateEditParams object
// with the default values initialized.
func NewGetPrivateEditParams() *GetPrivateEditParams {
	var (
		postOnlyDefault = bool(true)
	)
	return &GetPrivateEditParams{
		PostOnly: &postOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateEditParamsWithTimeout creates a new GetPrivateEditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateEditParamsWithTimeout(timeout time.Duration) *GetPrivateEditParams {
	var (
		postOnlyDefault = bool(true)
	)
	return &GetPrivateEditParams{
		PostOnly: &postOnlyDefault,

		timeout: timeout,
	}
}

// NewGetPrivateEditParamsWithContext creates a new GetPrivateEditParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateEditParamsWithContext(ctx context.Context) *GetPrivateEditParams {
	var (
		postOnlyDefault = bool(true)
	)
	return &GetPrivateEditParams{
		PostOnly: &postOnlyDefault,

		Context: ctx,
	}
}

// NewGetPrivateEditParamsWithHTTPClient creates a new GetPrivateEditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateEditParamsWithHTTPClient(client *http.Client) *GetPrivateEditParams {
	var (
		postOnlyDefault = bool(true)
	)
	return &GetPrivateEditParams{
		PostOnly:   &postOnlyDefault,
		HTTPClient: client,
	}
}

/*GetPrivateEditParams contains all the parameters to send to the API endpoint
for the get private edit operation typically these are written to a http.Request
*/
type GetPrivateEditParams struct {

	/*Advanced
	  Advanced option order type. If you have posted an advanced option order, it is necessary to re-supply this parameter when editing it (Only for options)

	*/
	Advanced *string
	/*Amount
	  It represents the requested order size. For perpetual and futures the amount is in USD units, for options it is amount of corresponding cryptocurrency contracts, e.g., BTC or ETH

	*/
	Amount float64
	/*OrderID
	  The order id

	*/
	OrderID string
	/*PostOnly
	  <p>If true, the order is considered post-only. If the new price would cause the order to be filled immediately (as taker), the price will be changed to be just below the bid.</p> <p>Only valid in combination with time_in_force=`"good_til_cancelled"`</p>

	*/
	PostOnly *bool
	/*Price
	  <p>The order price in base currency.</p> <p>When editing an option order with advanced=usd, the field price should be the option price value in USD.</p> <p>When editing an option order with advanced=implv, the field price should be a value of implied volatility in percentages. For example,  price=100, means implied volatility of 100%</p>

	*/
	Price float64
	/*StopPrice
	  Stop price, required for stop limit orders (Only for stop orders)

	*/
	StopPrice *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private edit params
func (o *GetPrivateEditParams) WithTimeout(timeout time.Duration) *GetPrivateEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private edit params
func (o *GetPrivateEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private edit params
func (o *GetPrivateEditParams) WithContext(ctx context.Context) *GetPrivateEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private edit params
func (o *GetPrivateEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private edit params
func (o *GetPrivateEditParams) WithHTTPClient(client *http.Client) *GetPrivateEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private edit params
func (o *GetPrivateEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvanced adds the advanced to the get private edit params
func (o *GetPrivateEditParams) WithAdvanced(advanced *string) *GetPrivateEditParams {
	o.SetAdvanced(advanced)
	return o
}

// SetAdvanced adds the advanced to the get private edit params
func (o *GetPrivateEditParams) SetAdvanced(advanced *string) {
	o.Advanced = advanced
}

// WithAmount adds the amount to the get private edit params
func (o *GetPrivateEditParams) WithAmount(amount float64) *GetPrivateEditParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get private edit params
func (o *GetPrivateEditParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithOrderID adds the orderID to the get private edit params
func (o *GetPrivateEditParams) WithOrderID(orderID string) *GetPrivateEditParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get private edit params
func (o *GetPrivateEditParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WithPostOnly adds the postOnly to the get private edit params
func (o *GetPrivateEditParams) WithPostOnly(postOnly *bool) *GetPrivateEditParams {
	o.SetPostOnly(postOnly)
	return o
}

// SetPostOnly adds the postOnly to the get private edit params
func (o *GetPrivateEditParams) SetPostOnly(postOnly *bool) {
	o.PostOnly = postOnly
}

// WithPrice adds the price to the get private edit params
func (o *GetPrivateEditParams) WithPrice(price float64) *GetPrivateEditParams {
	o.SetPrice(price)
	return o
}

// SetPrice adds the price to the get private edit params
func (o *GetPrivateEditParams) SetPrice(price float64) {
	o.Price = price
}

// WithStopPrice adds the stopPrice to the get private edit params
func (o *GetPrivateEditParams) WithStopPrice(stopPrice *float64) *GetPrivateEditParams {
	o.SetStopPrice(stopPrice)
	return o
}

// SetStopPrice adds the stopPrice to the get private edit params
func (o *GetPrivateEditParams) SetStopPrice(stopPrice *float64) {
	o.StopPrice = stopPrice
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Advanced != nil {

		// query param advanced
		var qrAdvanced string
		if o.Advanced != nil {
			qrAdvanced = *o.Advanced
		}
		qAdvanced := qrAdvanced
		if qAdvanced != "" {
			if err := r.SetQueryParam("advanced", qAdvanced); err != nil {
				return err
			}
		}

	}

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatFloat64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// query param order_id
	qrOrderID := o.OrderID
	qOrderID := qrOrderID
	if qOrderID != "" {
		if err := r.SetQueryParam("order_id", qOrderID); err != nil {
			return err
		}
	}

	if o.PostOnly != nil {

		// query param post_only
		var qrPostOnly bool
		if o.PostOnly != nil {
			qrPostOnly = *o.PostOnly
		}
		qPostOnly := swag.FormatBool(qrPostOnly)
		if qPostOnly != "" {
			if err := r.SetQueryParam("post_only", qPostOnly); err != nil {
				return err
			}
		}

	}

	// query param price
	qrPrice := o.Price
	qPrice := swag.FormatFloat64(qrPrice)
	if qPrice != "" {
		if err := r.SetQueryParam("price", qPrice); err != nil {
			return err
		}
	}

	if o.StopPrice != nil {

		// query param stop_price
		var qrStopPrice float64
		if o.StopPrice != nil {
			qrStopPrice = *o.StopPrice
		}
		qStopPrice := swag.FormatFloat64(qrStopPrice)
		if qStopPrice != "" {
			if err := r.SetQueryParam("stop_price", qStopPrice); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
