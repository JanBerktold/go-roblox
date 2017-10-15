// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GETV1UserStudiodataReader is a Reader for the GETV1UserStudiodata structure.
type GETV1UserStudiodataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETV1UserStudiodataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETV1UserStudiodataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGETV1UserStudiodataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETV1UserStudiodataOK creates a GETV1UserStudiodataOK with default headers values
func NewGETV1UserStudiodataOK() *GETV1UserStudiodataOK {
	return &GETV1UserStudiodataOK{}
}

/*GETV1UserStudiodataOK handles this case with default header values.

GETV1UserStudiodataOK g e t v1 user studiodata o k
*/
type GETV1UserStudiodataOK struct {
	Payload GETV1UserStudiodataOKBody
}

func (o *GETV1UserStudiodataOK) Error() string {
	return fmt.Sprintf("[GET /v1/user/studiodata][%d] gETV1UserStudiodataOK  %+v", 200, o.Payload)
}

func (o *GETV1UserStudiodataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETV1UserStudiodataNoContent creates a GETV1UserStudiodataNoContent with default headers values
func NewGETV1UserStudiodataNoContent() *GETV1UserStudiodataNoContent {
	return &GETV1UserStudiodataNoContent{}
}

/*GETV1UserStudiodataNoContent handles this case with default header values.

No data was found.
*/
type GETV1UserStudiodataNoContent struct {
	Payload GETV1UserStudiodataNoContentBody
}

func (o *GETV1UserStudiodataNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/user/studiodata][%d] gETV1UserStudiodataNoContent  %+v", 204, o.Payload)
}

func (o *GETV1UserStudiodataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GETV1UserStudiodataNoContentBody g e t v1 user studiodata no content body
swagger:model GETV1UserStudiodataNoContentBody
*/

type GETV1UserStudiodataNoContentBody interface{}

/*GETV1UserStudiodataOKBody g e t v1 user studiodata o k body
swagger:model GETV1UserStudiodataOKBody
*/

type GETV1UserStudiodataOKBody interface{}
