// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Lexus123/go-deribit/models"
)

// GetPrivateGetAddressBookReader is a Reader for the GetPrivateGetAddressBook structure.
type GetPrivateGetAddressBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetAddressBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetAddressBookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetAddressBookOK creates a GetPrivateGetAddressBookOK with default headers values
func NewGetPrivateGetAddressBookOK() *GetPrivateGetAddressBookOK {
	return &GetPrivateGetAddressBookOK{}
}

/*GetPrivateGetAddressBookOK handles this case with default header values.

GetPrivateGetAddressBookOK get private get address book o k
*/
type GetPrivateGetAddressBookOK struct {
	Payload *models.PrivateAddressBookResponse
}

func (o *GetPrivateGetAddressBookOK) Error() string {
	return fmt.Sprintf("[GET /private/get_address_book][%d] getPrivateGetAddressBookOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetAddressBookOK) GetPayload() *models.PrivateAddressBookResponse {
	return o.Payload
}

func (o *GetPrivateGetAddressBookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateAddressBookResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}