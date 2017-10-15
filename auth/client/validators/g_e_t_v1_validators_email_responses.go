// Code generated by go-swagger; DO NOT EDIT.

package validators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/auth/models"
)

// GETV1ValidatorsEmailReader is a Reader for the GETV1ValidatorsEmail structure.
type GETV1ValidatorsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETV1ValidatorsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETV1ValidatorsEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETV1ValidatorsEmailOK creates a GETV1ValidatorsEmailOK with default headers values
func NewGETV1ValidatorsEmailOK() *GETV1ValidatorsEmailOK {
	return &GETV1ValidatorsEmailOK{}
}

/*GETV1ValidatorsEmailOK handles this case with default header values.

OK
*/
type GETV1ValidatorsEmailOK struct {
	Payload *models.EmailValidationResponseModel
}

func (o *GETV1ValidatorsEmailOK) Error() string {
	return fmt.Sprintf("[GET /v1/validators/email][%d] gETV1ValidatorsEmailOK  %+v", 200, o.Payload)
}

func (o *GETV1ValidatorsEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailValidationResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}