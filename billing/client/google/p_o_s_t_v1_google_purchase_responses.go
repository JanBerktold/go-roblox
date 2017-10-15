// Code generated by go-swagger; DO NOT EDIT.

package google

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/billing/models"
)

// POSTV1GooglePurchaseReader is a Reader for the POSTV1GooglePurchase structure.
type POSTV1GooglePurchaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1GooglePurchaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1GooglePurchaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPOSTV1GooglePurchaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPOSTV1GooglePurchaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1GooglePurchaseOK creates a POSTV1GooglePurchaseOK with default headers values
func NewPOSTV1GooglePurchaseOK() *POSTV1GooglePurchaseOK {
	return &POSTV1GooglePurchaseOK{}
}

/*POSTV1GooglePurchaseOK handles this case with default header values.

POSTV1GooglePurchaseOK p o s t v1 google purchase o k
*/
type POSTV1GooglePurchaseOK struct {
	Payload models.APIEmptyResponseModel
}

func (o *POSTV1GooglePurchaseOK) Error() string {
	return fmt.Sprintf("[POST /v1/google/purchase][%d] pOSTV1GooglePurchaseOK  %+v", 200, o.Payload)
}

func (o *POSTV1GooglePurchaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1GooglePurchaseBadRequest creates a POSTV1GooglePurchaseBadRequest with default headers values
func NewPOSTV1GooglePurchaseBadRequest() *POSTV1GooglePurchaseBadRequest {
	return &POSTV1GooglePurchaseBadRequest{}
}

/*POSTV1GooglePurchaseBadRequest handles this case with default header values.

Service was unable to validate receiptId.
*/
type POSTV1GooglePurchaseBadRequest struct {
}

func (o *POSTV1GooglePurchaseBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/google/purchase][%d] pOSTV1GooglePurchaseBadRequest ", 400)
}

func (o *POSTV1GooglePurchaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1GooglePurchaseUnauthorized creates a POSTV1GooglePurchaseUnauthorized with default headers values
func NewPOSTV1GooglePurchaseUnauthorized() *POSTV1GooglePurchaseUnauthorized {
	return &POSTV1GooglePurchaseUnauthorized{}
}

/*POSTV1GooglePurchaseUnauthorized handles this case with default header values.

Unauthorized to make request.
*/
type POSTV1GooglePurchaseUnauthorized struct {
}

func (o *POSTV1GooglePurchaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/google/purchase][%d] pOSTV1GooglePurchaseUnauthorized ", 401)
}

func (o *POSTV1GooglePurchaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
