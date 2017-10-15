// Code generated by go-swagger; DO NOT EDIT.

package membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janberktold/go-roblox/groups/models"
)

// GETV1GroupsGroupIDRolesRoleSetIDUsersReader is a Reader for the GETV1GroupsGroupIDRolesRoleSetIDUsers structure.
type GETV1GroupsGroupIDRolesRoleSetIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETV1GroupsGroupIDRolesRoleSetIDUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETV1GroupsGroupIDRolesRoleSetIDUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersOK creates a GETV1GroupsGroupIDRolesRoleSetIDUsersOK with default headers values
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersOK() *GETV1GroupsGroupIDRolesRoleSetIDUsersOK {
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersOK{}
}

/*GETV1GroupsGroupIDRolesRoleSetIDUsersOK handles this case with default header values.

GETV1GroupsGroupIDRolesRoleSetIDUsersOK g e t v1 groups group Id roles role set Id users o k
*/
type GETV1GroupsGroupIDRolesRoleSetIDUsersOK struct {
	Payload *models.APIPageResponseUserModel
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/groups/{groupId}/roles/{roleSetId}/users][%d] gETV1GroupsGroupIdRolesRoleSetIdUsersOK  %+v", 200, o.Payload)
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPageResponseUserModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest creates a GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest with default headers values
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest() *GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest {
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest{}
}

/*GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest handles this case with default header values.

{Roblox.Groups.Api.ResponseEnums.MembershipErrors.InvalidGroupId}
*/
type GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest struct {
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/groups/{groupId}/roles/{roleSetId}/users][%d] gETV1GroupsGroupIdRolesRoleSetIdUsersBadRequest ", 400)
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETV1GroupsGroupIDRolesRoleSetIDUsersForbidden creates a GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden with default headers values
func NewGETV1GroupsGroupIDRolesRoleSetIDUsersForbidden() *GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden {
	return &GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden{}
}

/*GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden handles this case with default header values.

{Roblox.Groups.Api.ResponseEnums.MembershipErrors.InvalidRoleSetId}
*/
type GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden struct {
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/groups/{groupId}/roles/{roleSetId}/users][%d] gETV1GroupsGroupIdRolesRoleSetIdUsersForbidden ", 403)
}

func (o *GETV1GroupsGroupIDRolesRoleSetIDUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
