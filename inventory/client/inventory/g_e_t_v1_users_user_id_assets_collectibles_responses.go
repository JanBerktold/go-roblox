// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/inventory/models"
)

// GETV1UsersUserIDAssetsCollectiblesReader is a Reader for the GETV1UsersUserIDAssetsCollectibles structure.
type GETV1UsersUserIDAssetsCollectiblesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETV1UsersUserIDAssetsCollectiblesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETV1UsersUserIDAssetsCollectiblesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGETV1UsersUserIDAssetsCollectiblesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETV1UsersUserIDAssetsCollectiblesOK creates a GETV1UsersUserIDAssetsCollectiblesOK with default headers values
func NewGETV1UsersUserIDAssetsCollectiblesOK() *GETV1UsersUserIDAssetsCollectiblesOK {
	return &GETV1UsersUserIDAssetsCollectiblesOK{}
}

/*GETV1UsersUserIDAssetsCollectiblesOK handles this case with default header values.

Collectibles owned by user
*/
type GETV1UsersUserIDAssetsCollectiblesOK struct {
	Payload *models.APIPageResponseCollectibleUserAssetModel
}

func (o *GETV1UsersUserIDAssetsCollectiblesOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{userId}/assets/collectibles][%d] gETV1UsersUserIdAssetsCollectiblesOK  %+v", 200, o.Payload)
}

func (o *GETV1UsersUserIDAssetsCollectiblesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPageResponseCollectibleUserAssetModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETV1UsersUserIDAssetsCollectiblesBadRequest creates a GETV1UsersUserIDAssetsCollectiblesBadRequest with default headers values
func NewGETV1UsersUserIDAssetsCollectiblesBadRequest() *GETV1UsersUserIDAssetsCollectiblesBadRequest {
	return &GETV1UsersUserIDAssetsCollectiblesBadRequest{}
}

/*GETV1UsersUserIDAssetsCollectiblesBadRequest handles this case with default header values.

The specified asset type(s) are invalid.
*/
type GETV1UsersUserIDAssetsCollectiblesBadRequest struct {
}

func (o *GETV1UsersUserIDAssetsCollectiblesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users/{userId}/assets/collectibles][%d] gETV1UsersUserIdAssetsCollectiblesBadRequest ", 400)
}

func (o *GETV1UsersUserIDAssetsCollectiblesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}