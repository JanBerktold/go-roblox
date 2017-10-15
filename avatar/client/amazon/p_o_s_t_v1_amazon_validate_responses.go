// Code generated by go-swagger; DO NOT EDIT.

package amazon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/avatar/models"
)

// POSTV1AmazonValidateReader is a Reader for the POSTV1AmazonValidate structure.
type POSTV1AmazonValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1AmazonValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1AmazonValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPOSTV1AmazonValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPOSTV1AmazonValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPOSTV1AmazonValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPOSTV1AmazonValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPOSTV1AmazonValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPOSTV1AmazonValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1AmazonValidateOK creates a POSTV1AmazonValidateOK with default headers values
func NewPOSTV1AmazonValidateOK() *POSTV1AmazonValidateOK {
	return &POSTV1AmazonValidateOK{}
}

/*POSTV1AmazonValidateOK handles this case with default header values.

POSTV1AmazonValidateOK p o s t v1 amazon validate o k
*/
type POSTV1AmazonValidateOK struct {
	Payload models.APIEmptyResponseModel
}

func (o *POSTV1AmazonValidateOK) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateOK  %+v", 200, o.Payload)
}

func (o *POSTV1AmazonValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1AmazonValidateBadRequest creates a POSTV1AmazonValidateBadRequest with default headers values
func NewPOSTV1AmazonValidateBadRequest() *POSTV1AmazonValidateBadRequest {
	return &POSTV1AmazonValidateBadRequest{}
}

/*POSTV1AmazonValidateBadRequest handles this case with default header values.

Service has thrown an uknown exception.
*/
type POSTV1AmazonValidateBadRequest struct {
}

func (o *POSTV1AmazonValidateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateBadRequest ", 400)
}

func (o *POSTV1AmazonValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonValidateUnauthorized creates a POSTV1AmazonValidateUnauthorized with default headers values
func NewPOSTV1AmazonValidateUnauthorized() *POSTV1AmazonValidateUnauthorized {
	return &POSTV1AmazonValidateUnauthorized{}
}

/*POSTV1AmazonValidateUnauthorized handles this case with default header values.

Unauthorized to make request. Auth cookie is needed
*/
type POSTV1AmazonValidateUnauthorized struct {
}

func (o *POSTV1AmazonValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateUnauthorized ", 401)
}

func (o *POSTV1AmazonValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonValidateForbidden creates a POSTV1AmazonValidateForbidden with default headers values
func NewPOSTV1AmazonValidateForbidden() *POSTV1AmazonValidateForbidden {
	return &POSTV1AmazonValidateForbidden{}
}

/*POSTV1AmazonValidateForbidden handles this case with default header values.

Unauthorized to make request. XSRF token is needed.
*/
type POSTV1AmazonValidateForbidden struct {
}

func (o *POSTV1AmazonValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateForbidden ", 403)
}

func (o *POSTV1AmazonValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonValidateNotFound creates a POSTV1AmazonValidateNotFound with default headers values
func NewPOSTV1AmazonValidateNotFound() *POSTV1AmazonValidateNotFound {
	return &POSTV1AmazonValidateNotFound{}
}

/*POSTV1AmazonValidateNotFound handles this case with default header values.

{validateModel} is invalid.
*/
type POSTV1AmazonValidateNotFound struct {
}

func (o *POSTV1AmazonValidateNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateNotFound ", 404)
}

func (o *POSTV1AmazonValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonValidateTooManyRequests creates a POSTV1AmazonValidateTooManyRequests with default headers values
func NewPOSTV1AmazonValidateTooManyRequests() *POSTV1AmazonValidateTooManyRequests {
	return &POSTV1AmazonValidateTooManyRequests{}
}

/*POSTV1AmazonValidateTooManyRequests handles this case with default header values.

Service has been rate limited to user.
*/
type POSTV1AmazonValidateTooManyRequests struct {
}

func (o *POSTV1AmazonValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateTooManyRequests ", 429)
}

func (o *POSTV1AmazonValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonValidateServiceUnavailable creates a POSTV1AmazonValidateServiceUnavailable with default headers values
func NewPOSTV1AmazonValidateServiceUnavailable() *POSTV1AmazonValidateServiceUnavailable {
	return &POSTV1AmazonValidateServiceUnavailable{}
}

/*POSTV1AmazonValidateServiceUnavailable handles this case with default header values.

Service has been turned off. Header response will include Retry-After in seconds.
*/
type POSTV1AmazonValidateServiceUnavailable struct {
}

func (o *POSTV1AmazonValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/validate][%d] pOSTV1AmazonValidateServiceUnavailable ", 503)
}

func (o *POSTV1AmazonValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
