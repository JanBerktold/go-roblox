// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/auth/models"
)

// GETV1SamlMetadataReader is a Reader for the GETV1SamlMetadata structure.
type GETV1SamlMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETV1SamlMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETV1SamlMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETV1SamlMetadataOK creates a GETV1SamlMetadataOK with default headers values
func NewGETV1SamlMetadataOK() *GETV1SamlMetadataOK {
	return &GETV1SamlMetadataOK{}
}

/*GETV1SamlMetadataOK handles this case with default header values.

OK
*/
type GETV1SamlMetadataOK struct {
	Payload models.SamlMetadataResult
}

func (o *GETV1SamlMetadataOK) Error() string {
	return fmt.Sprintf("[GET /v1/saml/metadata][%d] gETV1SamlMetadataOK  %+v", 200, o.Payload)
}

func (o *GETV1SamlMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}