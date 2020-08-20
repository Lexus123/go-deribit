// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Lexus123/go-deribit/models"
)

// GetPrivateLogoutReader is a Reader for the GetPrivateLogout structure.
type GetPrivateLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrivateLogoutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateLogoutOK creates a GetPrivateLogoutOK with default headers values
func NewGetPrivateLogoutOK() *GetPrivateLogoutOK {
	return &GetPrivateLogoutOK{}
}

/*GetPrivateLogoutOK handles this case with default header values.

foo
*/
type GetPrivateLogoutOK struct {
}

func (o *GetPrivateLogoutOK) Error() string {
	return fmt.Sprintf("[GET /private/logout][%d] getPrivateLogoutOK ", 200)
}

func (o *GetPrivateLogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrivateLogoutBadRequest creates a GetPrivateLogoutBadRequest with default headers values
func NewGetPrivateLogoutBadRequest() *GetPrivateLogoutBadRequest {
	return &GetPrivateLogoutBadRequest{}
}

/*GetPrivateLogoutBadRequest handles this case with default header values.

result when used via rest/HTTP
*/
type GetPrivateLogoutBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateLogoutBadRequest) Error() string {
	return fmt.Sprintf("[GET /private/logout][%d] getPrivateLogoutBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateLogoutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
