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

// GetPrivateGetOrderStateReader is a Reader for the GetPrivateGetOrderState structure.
type GetPrivateGetOrderStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetOrderStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetOrderStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrivateGetOrderStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetOrderStateOK creates a GetPrivateGetOrderStateOK with default headers values
func NewGetPrivateGetOrderStateOK() *GetPrivateGetOrderStateOK {
	return &GetPrivateGetOrderStateOK{}
}

/*GetPrivateGetOrderStateOK handles this case with default header values.

foo
*/
type GetPrivateGetOrderStateOK struct {
	Payload *models.PrivateGetOrderStateResponse
}

func (o *GetPrivateGetOrderStateOK) Error() string {
	return fmt.Sprintf("[GET /private/get_order_state][%d] getPrivateGetOrderStateOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetOrderStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetOrderStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateGetOrderStateBadRequest creates a GetPrivateGetOrderStateBadRequest with default headers values
func NewGetPrivateGetOrderStateBadRequest() *GetPrivateGetOrderStateBadRequest {
	return &GetPrivateGetOrderStateBadRequest{}
}

/*GetPrivateGetOrderStateBadRequest handles this case with default header values.

result when used via rest/HTTP
*/
type GetPrivateGetOrderStateBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateGetOrderStateBadRequest) Error() string {
	return fmt.Sprintf("[GET /private/get_order_state][%d] getPrivateGetOrderStateBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateGetOrderStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
