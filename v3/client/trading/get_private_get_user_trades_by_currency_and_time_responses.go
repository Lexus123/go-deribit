// Code generated by go-swagger; DO NOT EDIT.

package trading

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Lexus123/go-deribit/v3/models"
)

// GetPrivateGetUserTradesByCurrencyAndTimeReader is a Reader for the GetPrivateGetUserTradesByCurrencyAndTime structure.
type GetPrivateGetUserTradesByCurrencyAndTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetUserTradesByCurrencyAndTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetUserTradesByCurrencyAndTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetUserTradesByCurrencyAndTimeOK creates a GetPrivateGetUserTradesByCurrencyAndTimeOK with default headers values
func NewGetPrivateGetUserTradesByCurrencyAndTimeOK() *GetPrivateGetUserTradesByCurrencyAndTimeOK {
	return &GetPrivateGetUserTradesByCurrencyAndTimeOK{}
}

/*GetPrivateGetUserTradesByCurrencyAndTimeOK handles this case with default header values.

GetPrivateGetUserTradesByCurrencyAndTimeOK get private get user trades by currency and time o k
*/
type GetPrivateGetUserTradesByCurrencyAndTimeOK struct {
	Payload *models.PrivateGetUserTradesHistoryResponse
}

func (o *GetPrivateGetUserTradesByCurrencyAndTimeOK) Error() string {
	return fmt.Sprintf("[GET /private/get_user_trades_by_currency_and_time][%d] getPrivateGetUserTradesByCurrencyAndTimeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetUserTradesByCurrencyAndTimeOK) GetPayload() *models.PrivateGetUserTradesHistoryResponse {
	return o.Payload
}

func (o *GetPrivateGetUserTradesByCurrencyAndTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetUserTradesHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
