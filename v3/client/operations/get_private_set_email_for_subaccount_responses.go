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

// GetPrivateSetEmailForSubaccountReader is a Reader for the GetPrivateSetEmailForSubaccount structure.
type GetPrivateSetEmailForSubaccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSetEmailForSubaccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSetEmailForSubaccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSetEmailForSubaccountOK creates a GetPrivateSetEmailForSubaccountOK with default headers values
func NewGetPrivateSetEmailForSubaccountOK() *GetPrivateSetEmailForSubaccountOK {
	return &GetPrivateSetEmailForSubaccountOK{}
}

/*GetPrivateSetEmailForSubaccountOK handles this case with default header values.

ok response
*/
type GetPrivateSetEmailForSubaccountOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateSetEmailForSubaccountOK) Error() string {
	return fmt.Sprintf("[GET /private/set_email_for_subaccount][%d] getPrivateSetEmailForSubaccountOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSetEmailForSubaccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
