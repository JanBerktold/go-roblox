// Code generated by go-swagger; DO NOT EDIT.

package passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/auth/models"
)

// POSTV1UserPasswordsChangeReader is a Reader for the POSTV1UserPasswordsChange structure.
type POSTV1UserPasswordsChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1UserPasswordsChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1UserPasswordsChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPOSTV1UserPasswordsChangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPOSTV1UserPasswordsChangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPOSTV1UserPasswordsChangeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1UserPasswordsChangeOK creates a POSTV1UserPasswordsChangeOK with default headers values
func NewPOSTV1UserPasswordsChangeOK() *POSTV1UserPasswordsChangeOK {
	return &POSTV1UserPasswordsChangeOK{}
}

/*POSTV1UserPasswordsChangeOK handles this case with default header values.

Password successfully changed.
*/
type POSTV1UserPasswordsChangeOK struct {
	Payload models.APIEmptyResponseModel
}

func (o *POSTV1UserPasswordsChangeOK) Error() string {
	return fmt.Sprintf("[POST /v1/user/passwords/change][%d] pOSTV1UserPasswordsChangeOK  %+v", 200, o.Payload)
}

func (o *POSTV1UserPasswordsChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1UserPasswordsChangeBadRequest creates a POSTV1UserPasswordsChangeBadRequest with default headers values
func NewPOSTV1UserPasswordsChangeBadRequest() *POSTV1UserPasswordsChangeBadRequest {
	return &POSTV1UserPasswordsChangeBadRequest{}
}

/*POSTV1UserPasswordsChangeBadRequest handles this case with default header values.

{Roblox.Api.Authentication.ResponseEnums.PasswordErrors.InvalidCurrentPassword}
            OR
            {Roblox.Api.Authentication.ResponseEnums.PasswordErrors.InvalidPassword}
*/
type POSTV1UserPasswordsChangeBadRequest struct {
}

func (o *POSTV1UserPasswordsChangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/user/passwords/change][%d] pOSTV1UserPasswordsChangeBadRequest ", 400)
}

func (o *POSTV1UserPasswordsChangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1UserPasswordsChangeForbidden creates a POSTV1UserPasswordsChangeForbidden with default headers values
func NewPOSTV1UserPasswordsChangeForbidden() *POSTV1UserPasswordsChangeForbidden {
	return &POSTV1UserPasswordsChangeForbidden{}
}

/*POSTV1UserPasswordsChangeForbidden handles this case with default header values.

{Roblox.Api.Authentication.ResponseEnums.PasswordErrors.PinLocked}
*/
type POSTV1UserPasswordsChangeForbidden struct {
}

func (o *POSTV1UserPasswordsChangeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/user/passwords/change][%d] pOSTV1UserPasswordsChangeForbidden ", 403)
}

func (o *POSTV1UserPasswordsChangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1UserPasswordsChangeTooManyRequests creates a POSTV1UserPasswordsChangeTooManyRequests with default headers values
func NewPOSTV1UserPasswordsChangeTooManyRequests() *POSTV1UserPasswordsChangeTooManyRequests {
	return &POSTV1UserPasswordsChangeTooManyRequests{}
}

/*POSTV1UserPasswordsChangeTooManyRequests handles this case with default header values.

{Roblox.Api.Authentication.ResponseEnums.PasswordErrors.Flooded}
*/
type POSTV1UserPasswordsChangeTooManyRequests struct {
}

func (o *POSTV1UserPasswordsChangeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /v1/user/passwords/change][%d] pOSTV1UserPasswordsChangeTooManyRequests ", 429)
}

func (o *POSTV1UserPasswordsChangeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
