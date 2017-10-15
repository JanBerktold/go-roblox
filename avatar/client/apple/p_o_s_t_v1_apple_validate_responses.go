// Code generated by go-swagger; DO NOT EDIT.

package apple

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/avatar/models"
)

// POSTV1AppleValidateReader is a Reader for the POSTV1AppleValidate structure.
type POSTV1AppleValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1AppleValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1AppleValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPOSTV1AppleValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPOSTV1AppleValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPOSTV1AppleValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPOSTV1AppleValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPOSTV1AppleValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPOSTV1AppleValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1AppleValidateOK creates a POSTV1AppleValidateOK with default headers values
func NewPOSTV1AppleValidateOK() *POSTV1AppleValidateOK {
	return &POSTV1AppleValidateOK{}
}

/*POSTV1AppleValidateOK handles this case with default header values.

POSTV1AppleValidateOK p o s t v1 apple validate o k
*/
type POSTV1AppleValidateOK struct {
	Payload models.APIEmptyResponseModel
}

func (o *POSTV1AppleValidateOK) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateOK  %+v", 200, o.Payload)
}

func (o *POSTV1AppleValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1AppleValidateBadRequest creates a POSTV1AppleValidateBadRequest with default headers values
func NewPOSTV1AppleValidateBadRequest() *POSTV1AppleValidateBadRequest {
	return &POSTV1AppleValidateBadRequest{}
}

/*POSTV1AppleValidateBadRequest handles this case with default header values.

Service has thrown an uknown exception.
*/
type POSTV1AppleValidateBadRequest struct {
}

func (o *POSTV1AppleValidateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateBadRequest ", 400)
}

func (o *POSTV1AppleValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AppleValidateUnauthorized creates a POSTV1AppleValidateUnauthorized with default headers values
func NewPOSTV1AppleValidateUnauthorized() *POSTV1AppleValidateUnauthorized {
	return &POSTV1AppleValidateUnauthorized{}
}

/*POSTV1AppleValidateUnauthorized handles this case with default header values.

Unauthorized to make request. Auth cookie is needed
*/
type POSTV1AppleValidateUnauthorized struct {
}

func (o *POSTV1AppleValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateUnauthorized ", 401)
}

func (o *POSTV1AppleValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AppleValidateForbidden creates a POSTV1AppleValidateForbidden with default headers values
func NewPOSTV1AppleValidateForbidden() *POSTV1AppleValidateForbidden {
	return &POSTV1AppleValidateForbidden{}
}

/*POSTV1AppleValidateForbidden handles this case with default header values.

Unauthorized to make request. XSRF token is needed.
*/
type POSTV1AppleValidateForbidden struct {
}

func (o *POSTV1AppleValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateForbidden ", 403)
}

func (o *POSTV1AppleValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AppleValidateNotFound creates a POSTV1AppleValidateNotFound with default headers values
func NewPOSTV1AppleValidateNotFound() *POSTV1AppleValidateNotFound {
	return &POSTV1AppleValidateNotFound{}
}

/*POSTV1AppleValidateNotFound handles this case with default header values.

{validateModel} is invalid.
*/
type POSTV1AppleValidateNotFound struct {
}

func (o *POSTV1AppleValidateNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateNotFound ", 404)
}

func (o *POSTV1AppleValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AppleValidateTooManyRequests creates a POSTV1AppleValidateTooManyRequests with default headers values
func NewPOSTV1AppleValidateTooManyRequests() *POSTV1AppleValidateTooManyRequests {
	return &POSTV1AppleValidateTooManyRequests{}
}

/*POSTV1AppleValidateTooManyRequests handles this case with default header values.

Service has been rate limited to user.
*/
type POSTV1AppleValidateTooManyRequests struct {
}

func (o *POSTV1AppleValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateTooManyRequests ", 429)
}

func (o *POSTV1AppleValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AppleValidateServiceUnavailable creates a POSTV1AppleValidateServiceUnavailable with default headers values
func NewPOSTV1AppleValidateServiceUnavailable() *POSTV1AppleValidateServiceUnavailable {
	return &POSTV1AppleValidateServiceUnavailable{}
}

/*POSTV1AppleValidateServiceUnavailable handles this case with default header values.

Service has been turned off. Header response will include Retry-After in seconds.
*/
type POSTV1AppleValidateServiceUnavailable struct {
}

func (o *POSTV1AppleValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/apple/validate][%d] pOSTV1AppleValidateServiceUnavailable ", 503)
}

func (o *POSTV1AppleValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}