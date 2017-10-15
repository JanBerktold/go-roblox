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

// POSTV1AmazonPurchaseReader is a Reader for the POSTV1AmazonPurchase structure.
type POSTV1AmazonPurchaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTV1AmazonPurchaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPOSTV1AmazonPurchaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPOSTV1AmazonPurchaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPOSTV1AmazonPurchaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPOSTV1AmazonPurchaseOK creates a POSTV1AmazonPurchaseOK with default headers values
func NewPOSTV1AmazonPurchaseOK() *POSTV1AmazonPurchaseOK {
	return &POSTV1AmazonPurchaseOK{}
}

/*POSTV1AmazonPurchaseOK handles this case with default header values.

POSTV1AmazonPurchaseOK p o s t v1 amazon purchase o k
*/
type POSTV1AmazonPurchaseOK struct {
	Payload models.APIEmptyResponseModel
}

func (o *POSTV1AmazonPurchaseOK) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/purchase][%d] pOSTV1AmazonPurchaseOK  %+v", 200, o.Payload)
}

func (o *POSTV1AmazonPurchaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTV1AmazonPurchaseBadRequest creates a POSTV1AmazonPurchaseBadRequest with default headers values
func NewPOSTV1AmazonPurchaseBadRequest() *POSTV1AmazonPurchaseBadRequest {
	return &POSTV1AmazonPurchaseBadRequest{}
}

/*POSTV1AmazonPurchaseBadRequest handles this case with default header values.

Service was unable to validate receiptId.
*/
type POSTV1AmazonPurchaseBadRequest struct {
}

func (o *POSTV1AmazonPurchaseBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/purchase][%d] pOSTV1AmazonPurchaseBadRequest ", 400)
}

func (o *POSTV1AmazonPurchaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPOSTV1AmazonPurchaseUnauthorized creates a POSTV1AmazonPurchaseUnauthorized with default headers values
func NewPOSTV1AmazonPurchaseUnauthorized() *POSTV1AmazonPurchaseUnauthorized {
	return &POSTV1AmazonPurchaseUnauthorized{}
}

/*POSTV1AmazonPurchaseUnauthorized handles this case with default header values.

Unauthorized to make request.
*/
type POSTV1AmazonPurchaseUnauthorized struct {
}

func (o *POSTV1AmazonPurchaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/amazon/purchase][%d] pOSTV1AmazonPurchaseUnauthorized ", 401)
}

func (o *POSTV1AmazonPurchaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}