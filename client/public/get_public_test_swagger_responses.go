// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Lexus123/go-deribit/models"
)

// GetPublicTestReader is a Reader for the GetPublicTest structure.
type GetPublicTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPublicTestOK creates a GetPublicTestOK with default headers values
func NewGetPublicTestOK() *GetPublicTestOK {
	return &GetPublicTestOK{}
}

/*GetPublicTestOK handles this case with default header values.

GetPublicTestOK get public test o k
*/
type GetPublicTestOK struct {
	Payload *models.PublicTestResponse
}

func (o *GetPublicTestOK) Error() string {
	return fmt.Sprintf("[GET /public/test][%d] getPublicTestOK  %+v", 200, o.Payload)
}

func (o *GetPublicTestOK) GetPayload() *models.PublicTestResponse {
	return o.Payload
}

func (o *GetPublicTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicTestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}