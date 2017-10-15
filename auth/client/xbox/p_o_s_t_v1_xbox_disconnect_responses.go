// Code generated by go-swagger; DO NOT EDIT.

package xbox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/auth/models"
)

// POSTV1XboxDisconnectReader is a Reader for the POSTV1XboxDisconnect structure.
type POSTV1XboxDisconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1XboxDisconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1XboxDisconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPOSTV1XboxDisconnectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1XboxDisconnectOK creates a POSTV1XboxDisconnectOK with default headers values
func NewPOSTV1XboxDisconnectOK() *POSTV1XboxDisconnectOK {
	return &POSTV1XboxDisconnectOK{}
}

/*POSTV1XboxDisconnectOK handles this case with default header values.

OK
*/
type POSTV1XboxDisconnectOK struct {
	Payload *models.APISuccessResponse
}

func (o *POSTV1XboxDisconnectOK) Error() string {
	return fmt.Sprintf("[POST /v1/xbox/disconnect][%d] pOSTV1XboxDisconnectOK  %+v", 200, o.Payload)
}

func (o *POSTV1XboxDisconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1XboxDisconnectForbidden creates a POSTV1XboxDisconnectForbidden with default headers values
func NewPOSTV1XboxDisconnectForbidden() *POSTV1XboxDisconnectForbidden {
	return &POSTV1XboxDisconnectForbidden{}
}

/*POSTV1XboxDisconnectForbidden handles this case with default header values.

Forbidden
*/
type POSTV1XboxDisconnectForbidden struct {
}

func (o *POSTV1XboxDisconnectForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/xbox/disconnect][%d] pOSTV1XboxDisconnectForbidden ", 403)
}

func (o *POSTV1XboxDisconnectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}