/* 
 * Roblox.Api.Authentication V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// A model containing information to reset a password with a ticket.
type PasswordResetModel struct {

	// The ticket.
	Ticket string `json:"ticket,omitempty"`

	// The new password.
	Password string `json:"password,omitempty"`
}
