// Code generated by go-swagger; DO NOT EDIT.

package internal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/v3/models"
)

// GetPublicGetOptionMarkPricesReader is a Reader for the GetPublicGetOptionMarkPrices structure.
type GetPublicGetOptionMarkPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetOptionMarkPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicGetOptionMarkPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetOptionMarkPricesOK creates a GetPublicGetOptionMarkPricesOK with default headers values
func NewGetPublicGetOptionMarkPricesOK() *GetPublicGetOptionMarkPricesOK {
	return &GetPublicGetOptionMarkPricesOK{}
}

/*GetPublicGetOptionMarkPricesOK handles this case with default header values.

ok response
*/
type GetPublicGetOptionMarkPricesOK struct {
	Payload models.PublicGetOptionMarkPricesResponse
}

func (o *GetPublicGetOptionMarkPricesOK) Error() string {
	return fmt.Sprintf("[GET /public/get_option_mark_prices][%d] getPublicGetOptionMarkPricesOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetOptionMarkPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
