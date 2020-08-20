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

// GetPrivateGetPositionReader is a Reader for the GetPrivateGetPosition structure.
type GetPrivateGetPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrivateGetPositionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetPositionOK creates a GetPrivateGetPositionOK with default headers values
func NewGetPrivateGetPositionOK() *GetPrivateGetPositionOK {
	return &GetPrivateGetPositionOK{}
}

/*GetPrivateGetPositionOK handles this case with default header values.

foo
*/
type GetPrivateGetPositionOK struct {
	Payload *models.PrivateGetPositionResponse
}

func (o *GetPrivateGetPositionOK) Error() string {
	return fmt.Sprintf("[GET /private/get_position][%d] getPrivateGetPositionOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetPositionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateGetPositionBadRequest creates a GetPrivateGetPositionBadRequest with default headers values
func NewGetPrivateGetPositionBadRequest() *GetPrivateGetPositionBadRequest {
	return &GetPrivateGetPositionBadRequest{}
}

/*GetPrivateGetPositionBadRequest handles this case with default header values.

When some error occurs
*/
type GetPrivateGetPositionBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateGetPositionBadRequest) Error() string {
	return fmt.Sprintf("[GET /private/get_position][%d] getPrivateGetPositionBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateGetPositionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
