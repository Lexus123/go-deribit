// Code generated by go-swagger; DO NOT EDIT.

package private

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Lexus123/go-deribit/models"
)

// GetPrivateToggleSubaccountLoginReader is a Reader for the GetPrivateToggleSubaccountLogin structure.
type GetPrivateToggleSubaccountLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateToggleSubaccountLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateToggleSubaccountLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateToggleSubaccountLoginOK creates a GetPrivateToggleSubaccountLoginOK with default headers values
func NewGetPrivateToggleSubaccountLoginOK() *GetPrivateToggleSubaccountLoginOK {
	return &GetPrivateToggleSubaccountLoginOK{}
}

/*GetPrivateToggleSubaccountLoginOK handles this case with default header values.

ok response
*/
type GetPrivateToggleSubaccountLoginOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateToggleSubaccountLoginOK) Error() string {
	return fmt.Sprintf("[GET /private/toggle_subaccount_login][%d] getPrivateToggleSubaccountLoginOK  %+v", 200, o.Payload)
}

func (o *GetPrivateToggleSubaccountLoginOK) GetPayload() *models.OkResponse {
	return o.Payload
}

func (o *GetPrivateToggleSubaccountLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}