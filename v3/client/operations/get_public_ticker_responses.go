// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/v3/models"
)

// GetPublicTickerReader is a Reader for the GetPublicTicker structure.
type GetPublicTickerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicTickerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicTickerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicTickerOK creates a GetPublicTickerOK with default headers values
func NewGetPublicTickerOK() *GetPublicTickerOK {
	return &GetPublicTickerOK{}
}

/*GetPublicTickerOK handles this case with default header values.

ok response
*/
type GetPublicTickerOK struct {
	Payload *models.PublicTickerResponse
}

func (o *GetPublicTickerOK) Error() string {
	return fmt.Sprintf("[GET /public/ticker][%d] getPublicTickerOK  %+v", 200, o.Payload)
}

func (o *GetPublicTickerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicTickerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
